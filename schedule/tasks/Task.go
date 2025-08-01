package tasks

import (
	"github.com/jasonlvhit/gocron"
	"pledge-backend-test/db"
	"pledge-backend-test/schedule/common"
	"pledge-backend-test/schedule/service"
	"time"
)

func Task() {
	//获取环境变量
	common.GetEnv()
	//刷新redis
	err := db.RedisFlushDB()
	if err != nil {
		panic("clear redis err" + err.Error())
	}
	//init task
	service.NewPoolService().UpdateAllPoolInfo()
	service.NewTokenPriceService().UpdateContractPrice()
	service.NewTokenSymbolService().UpdateTokenSymbol()
	service.NewTokenLogoService().UpdateTokenLogo()
	service.NewBanlanceMonitor().Monitor()

	//开启定时任务
	scheduler := gocron.NewScheduler() //开启调度器
	scheduler.ChangeLoc(time.UTC)      //设置时区
	//添加定时任务
	_ = scheduler.Every(2).Minutes().From(gocron.NextTick()).Do(service.NewPoolService().UpdateAllPoolInfo)
	_ = scheduler.Every(1).Minutes().From(gocron.NextTick()).Do(service.NewTokenPriceService().UpdateContractPrice)
	_ = scheduler.Every(2).Hour().From(gocron.NextTick()).Do(service.NewTokenSymbolService().UpdateTokenSymbol)
	_ = scheduler.Every(2).Hour().From(gocron.NextTick()).Do(service.NewTokenLogoService().UpdateTokenLogo)
	_ = scheduler.Every(30).Minutes().From(gocron.NextTick()).Do(service.NewBanlanceMonitor().Monitor)

	<-scheduler.Start() //调度器启动
}
