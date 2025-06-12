package schedule

import (
	"pledge-backend-test/db"
	"pledge-backend-test/schedule/tasks"
)

func main() {
	db.InitMysql()
	db.InitReids()
	tasks.Task()
}
