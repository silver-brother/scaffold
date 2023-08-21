package logx

import (
	"fmt"
	"os"
	"scaffold/internal/config"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog(cnf config.Config) zerolog.Logger {

	// 配置 lumberjack 实例，用于日志文件分割
	logFile := &lumberjack.Logger{
		Filename:   cnf.Log.Dir + "/server.log", // 日志文件名称
		MaxSize:    10,                          // 每个日志文件最大尺寸 (MB)
		MaxBackups: 3,                           // 最多保留的旧日志文件数量
		MaxAge:     30,                          // 日志文件最多保留天数
		Compress:   true,                        // 是否压缩旧的日志文件
	}
	zerolog.TimestampFieldName = "ts"
	zerolog.LevelFieldName = "lv"
	zerolog.MessageFieldName = "msg"
	zerolog.TimeFieldFormat = time.DateTime
	zerolog.CallerMarshalFunc = func(_ uintptr, file string, line int) string {
		p, _ := os.Getwd()

		return fmt.Sprintf("%s:%d", strings.TrimPrefix(file, p+"/"), line)
	}
	l := zerolog.New(logFile).
		Hook(TracingHook{}).
		With()
	l = l.Str("service", cnf.Name)
	return l.Caller().Timestamp().Logger()
}
