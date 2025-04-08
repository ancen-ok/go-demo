package core

import (
	"gitee.com/molonglove/goboot/logger"
)

var Log *logger.ToolLogger

func InitLogger() {
	/*logPath, _ := homedir.Expand("~/.admin/logs")
	fmt.Println("日志保存路径：", logPath)*/
	Log = logger.DefaultLogger("./logs")
}
