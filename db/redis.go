package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"pledge-backend-test/config"
	"pledge-backend-test/log"
	"time"
)

func InitReids() {
	log.Logger.Info("init redis")
	redisConfig := config.Config.Redis
	//建立连接池
	redisConn := &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,   //最大空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxActive:   redisConfig.MaxActive, //最大激活连接数，表示同时最多有N个连接   0 表示无穷大
		Wait:        true,                  // 如果连接数不足则阻塞等待
		IdleTimeout: 180 * time.Second,     //空闲连接的超时时间，表示超时后空闲连接会被清除
		// 连接初始化逻辑
		Dial: func() (redis.Conn, error) {
			//建立tcp连接
			dial, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port))
			if err != nil {
				return nil, err
			}
			//验证密码
			//_, err = dial.Do("auth", redisConfig.Password)
			//if err != nil {
			//	log.Logger.Error("redis auth err " + err.Error())
			//	//panic("redis auth err " + err.Error())
			//}
			//选择db
			_, err = dial.Do("select", redisConfig.Db)
			if err != nil {
				log.Logger.Error("redis select db err " + err.Error())
				//panic("redis select db err " + err.Error())
			}
			return dial, nil
		},
	}
	// 从池中获取连接
	err := redisConn.Get().Err()
	if err != nil {
		log.Logger.Error("redis init err " + err.Error())
		//panic("redis init err" + err.Error())
	}
	RedisConn = redisConn
}

// RedisSet 设置key、value、time
func RedisSet(key string, data interface{}, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if aliveSeconds > 0 {
		_, err = conn.Do("set", key, value, "EX", aliveSeconds)
	} else {
		_, err = conn.Do("set", key, value)
	}
	if err != nil {
		return err
	}
	return nil
}

// RedisSetString  设置key、value、time
func RedisSetString(key string, data string, aliveSeconds int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	var err error
	if aliveSeconds > 0 {
		_, err = redis.String(conn.Do("set", key, data, "EX", aliveSeconds))
	} else {
		_, err = redis.String(conn.Do("set", key, data))
	}
	if err != nil {
		return err
	}
	return nil
}

// RedisGet 获取Key
func RedisGet(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Bytes(conn.Do("get", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisGetString 获取Key
func RedisGetString(key string) (string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.String(conn.Do("get", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

// RedisSetInt64  set int64 value by key
func RedisSetInt64(key string, data int64, time int) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = redis.Int64(conn.Do("set", key, value))
	if err != nil {
		return err
	}
	if time != 0 {
		_, err = redis.Int64(conn.Do("expire", key, time))
		if err != nil {
			return err
		}
	}
	return nil
}

// RedisGetInt64 get int64 value by key
func RedisGetInt64(key string) (int64, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("get", key))
	if err != nil {
		return -1, err
	}
	return reply, nil
}

// RedisDelete 删除Key
func RedisDelete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	return redis.Bool(conn.Do("del", key))
}

// RedisFlushDB 清空当前DB
func RedisFlushDB() error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("flushdb")
	if err != nil {
		return err
	}
	return nil
}

// RedisGetHashOne 获取Heah其中一个值
func RedisGetHashOne(key, name string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := conn.Do("hgetall", key, name)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisSetHash 设置Hash
func RedisSetHash(key string, data map[string]string, time interface{}) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	for k, v := range data {
		err := conn.Send("hset", key, k, v)
		if err != nil {
			return err
		}
	}
	err := conn.Flush()
	if err != nil {
		return err
	}

	if time != nil {
		_, err = conn.Do("expire", key, time.(int))
		if err != nil {
			return err
		}
	}
	return nil
}

// RedisGetHash 获取Hash类型
func RedisGetHash(key string) (map[string]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.StringMap(conn.Do("hgetall", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// RedisDelHash 删除Hash
func RedisDelHash(key string) (bool, error) {

	return true, nil
}

// RedisExistsHash 检查Key是否存在
func RedisExistsHash(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("hexists", key))
	if err != nil {
		return false
	}
	return exists
}

// RedisExists 检查Key是否存在
func RedisExists(key string) bool {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	exists, err := redis.Bool(conn.Do("exists", key))
	if err != nil {
		return false
	}
	return exists
}

// RedisGetTTL 获取Key剩余时间
func RedisGetTTL(key string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Int64(conn.Do("ttl", key))
	if err != nil {
		return 0
	}
	return reply
}

// RedisSAdd set 集合
func RedisSAdd(k, v string) int64 {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := conn.Do("SAdd", k, v)
	if err != nil {
		return -1
	}
	return reply.(int64)
}

// RedisSmembers 获取集合元素
func RedisSmembers(k string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	reply, err := redis.Strings(conn.Do("smembers", k))
	if err != nil {
		return []string{}, errors.New("读取set错误")
	}
	return reply, err
}

type RedisEncryptionTask struct {
	RecordOrderFlowId int32  `json:"recordOrderFlow"` //密码转账表ID
	Encryption        string `json:"encryption"`      //密码串
	EndTime           int64  `json:"endTime"`         //失效截止时间
}

// RedisListRpush 列表右侧添加数据
func RedisListRpush(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("rpush", listName, encryption)
	return err
}

// RedisListLRange 取出列表中所有元素
func RedisListLRange(listName string) ([]string, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	res, err := redis.Strings(conn.Do("lrange", listName, 0, -1))
	return res, err
}

// RedisListLRem 删除列表中指定元素
func RedisListLRem(listName string, encryption string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("lrem", listName, 1, encryption)
	return err
}

// RedisListLength 列表长度
func RedisListLength(listName string) (interface{}, error) {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	len, err := conn.Do("llen", listName)
	return len, err
}

// RedisDelList list 删除整个列表
func RedisDelList(setName string) error {
	conn := RedisConn.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("del", setName)
	return err
}
