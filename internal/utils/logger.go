package utils

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//NewLogger 日志配置初始化器
func NewLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncooder()

	level := zap.AtomicLevel{}
	if err := level.UnmarshalText([]byte(viper.GetString("log.level"))); err != nil {
		level = zap.NewAtomicLevel() //默认用info级别
	}
	//core := zapcore.NewCore(encoder, writeSyncer, level)

	//根据app的模式把日志输出到不同的位置
	var core zapcore.Core
	if viper.GetString("app.mode") == gin.DebugMode {
		//consoleEncoder 往终端输出日志
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			//创建一个将Debug级别以上的日志输出到终端的配置信息
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger) //替换掉zap库里全局的logger
}

//日志格式
func getEncooder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

//日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   viper.GetString("log.filename"), //日志文件位置
		MaxSize:    viper.GetInt("olog.max_size"),
		MaxAge:     viper.GetInt("log.max_age"),
		MaxBackups: viper.GetInt("log.max_backups"),
		Compress:   viper.GetBool("log.compress"),
	}
	return zapcore.AddSync(lumberjackLogger)
}
