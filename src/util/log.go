package util

import (
	"log"
	"os"
)

type logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var Logger logger

func InitLog() {
	//errFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalln("打开日志文件失败：", err)
	//}

	Logger.Info = log.New(os.Stdout, "Info:", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Warning = log.New(os.Stdout, "Warning:", log.Ldate|log.Ltime|log.Lshortfile)
	Logger.Error = log.New(os.Stdout, "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}
