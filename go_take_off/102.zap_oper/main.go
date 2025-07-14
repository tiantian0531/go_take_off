package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// dev()
	// test()
	// prod()

	// SetLoggerLevel()
	// SetLoggerTimeFormat()
	// SetSugar()
	// SetStructLogger()

	//全局日志
	//initLogger()
	//globalLogger()

	//日志双写
	//LoggerWriteConsoleAndFile()

	// LoggerWriteConsoleAndFile_2()
	SplitLogger()
}

func dev() {

	logger, _ := zap.NewDevelopment()
	logger.Info("dev this is info")
	logger.Warn("dev this is warn")
	logger.Error("dev this is error")
	// logger.Panic("dev this is panic") 直接挂
}

func test() {
	logger := zap.NewExample()
	logger.Info("exam this is info")
	logger.Warn("exam this is warn")
	logger.Error("exam this is error")
}
func prod() {
	logger, _ := zap.NewProduction()
	logger.Info("prod this is info")
	logger.Warn("prod this is warn")
	logger.Error("prod this is error")
}

// 设置日志级别
func SetLoggerLevel() {

	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	logger, _ := cfg.Build()
	logger.Debug("spec level Debug")
	logger.Info("spec level Info")
	logger.Warn("spec level Warn")
	logger.Error("spec level Error")
}

// 日期格式化 必须  2006-01-02 15:04:05 设计约定
func SetLoggerTimeFormat() {

	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := cfg.Build()
	logger.Debug("spec level Debug")
	logger.Info("spec level Info")
	logger.Warn("spec level Warn")
	logger.Error("spec level Error")
}

func SetSugar() {

	logger, _ := zap.NewDevelopment()
	sl := logger.Sugar()
	sl.Info("dev this is info")
	sl.Warn("dev this is warn")
	sl.Error("dev this is error")
}

func SetStructLogger() {
	logger, _ := zap.NewDevelopment()
	logger.Info("this is info",
		zap.String("username", "admin"),
		zap.Int("user_id", 42),
		zap.Bool("active", true))

}

// 初始化全局日志
func initLogger() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建 Logger
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}

func globalLogger() {
	zap.L().Info("dev this is info")
	zap.L().Warn("dev this is warn")
	zap.L().Error("dev this is error")
	zap.S().Infof("dev this is info %s", "xxx")
	zap.S().Warnf("dev this is warn %s", "xxx")
	zap.S().Errorf("dev this is error %s", "xxx")
}

// 日志双写（写控制台、写文件）第一种方法
func LoggerWriteConsoleAndFile() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 输出console
	consoleWrite := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		os.Stdout,
		zap.InfoLevel,
	)

	// 输出文件
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend.Perm())
	fileWrite := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		file,
		zap.InfoLevel,
	)

	core := zapcore.NewTee(consoleWrite, fileWrite)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("xxxxx")
}

// 文件打开问题
func LoggerWriteConsoleAndFile_2() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 输出文件
	file, err := os.OpenFile(
		"./app.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644, // 关键修复：设置正确权限
	)
	if err != nil {
		panic(err) // 处理文件打开错误
	}

	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		writeSyncer,
		zap.DebugLevel,
	)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("xxxxx")

}

// 日志切片

type MyWrite struct {
	file        *os.File
	mutex       sync.Mutex
	currentTime string
}

func (w *MyWrite) Write(b []byte) (length int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	now := time.Now().Format("2001-01-02 15-04-05")
	if now == w.currentTime {
		return w.file.Write(b)
	}
	if w.file != nil {
		w.file.Close()
	}

	name := fmt.Sprintf("logs/%s.log", now)
	file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend.Perm())
	w.file = file
	w.currentTime = now
	return file.Write(b)
}

func SplitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&MyWrite{})),
		zap.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())

	for i := 0; i < 10; i++ {
		logger.Sugar().Infof("这是 %d log", i)
		time.Sleep(time.Second)
	}

}
