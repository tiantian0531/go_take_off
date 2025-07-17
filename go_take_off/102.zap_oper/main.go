package main

import (
	"os"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//dev()
	//test()
	//prod()

	//SetLoggerLevel()
	// SetLoggerTimeFormat()
	//SetSugar()
	// SetStructLogger()

	//全局日志
	//initLogger()
	//globalLogger()

	//日志双写
	//WriteLoggerToConsoleAndFile()

	// LoggerWriteConsoleAndFile_2()
	SplitLogger()
}

func dev() {

	logger, _ := zap.NewDevelopment()
	logger.Info("dev this is info", zap.String("info", "dev this is info"))
	logger.Warn("dev this is warn", zap.Int("warn", 110))
	logger.Error("dev this is error", zap.Duration("duration", time.Second))
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

	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := cfg.Build()
	logger.Debug("这是[Debug]日志")
	logger.Info("这是[Info]日志")
	logger.Warn("这是[Warn]日志",
		zap.String("Key1", "Value1"),
		zap.Int("Key2", 10086),
		zap.Duration("Duration", time.Second*100))
	logger.Error("这是[Error]日志")
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
	defer logger.Sync() // 在程序退出时刷新缓冲区
	sl.Infof("开发输出Info日志是: %d \n", 9999)
	sl.Warnln("开发输出Warn日志是:", "测试")
}

func SetStructLogger() {
	logger, _ := zap.NewDevelopment()
	logger.Info("this is info",
		zap.String("username", "admin"),
		zap.Int("user_id", 42),
		zap.Bool("active", true))

}

// 初始化全局日志
func initGlobalLogger() {
	cfg := zap.NewDevelopmentConfig()
	// 替换时间格式化方式
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}

func globalLogger() {
	zap.L().Debug("全局标准日志")
	zap.S().Debug("Sugar日志")
}

// 日志双写（写控制台、写文件）第一种方法
func WriteLoggerToConsoleAndFile() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 输出console
	conWrite := zapcore.NewCore(
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

	// 合并
	core := zapcore.NewTee(conWrite, fileWrite)
	logger := zap.New(core, zap.AddCaller())
	logger.Info("这个日志将会输出在控制台和写入文件中")
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

type CustomizeWrite struct {
	file        *os.File
	mutex       sync.Mutex
	currentTime string
}

func (w *CustomizeWrite) Write(b []byte) (length int, err error) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	// 获取当前时间并且对比写入时时间是否相同,如果相同直接写入
	now := time.Now().Format("2001-01-02 15-04-05")
	if now == w.currentTime && w.file != nil {
		return w.file.Write(b)
	}

	if w.file != nil {
		w.file.Close()
		w.file = nil
	}

	// 创建目录
	if err := os.MkdirAll("./logs", 0755); err != nil {
		return 0, err
	}

	name := filepath.Join("logs", now+".log")
	file, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend.Perm())
	if err != nil {
		return 0, err
	}
	w.file = file
	w.currentTime = now
	return file.Write(b)
}

func SplitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&CustomizeWrite{})),
		zap.InfoLevel,
	)
	logger := zap.New(core, zap.AddCaller())

	for i := 0; i < 10; i++ {
		logger.Sugar().Infof("这是 %d log", i)
		time.Sleep(time.Second)
	}
}
