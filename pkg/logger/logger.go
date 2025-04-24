package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/natefinch/lumberjack"
)

// Logger 定义日志结构体
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// GlobalLogger 全局日志实例
var GlobalLogger *Logger

func LogsSign() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// NewLogger 创建新的日志记录器
func NewLogger(handle *lumberjack.Logger) *Logger {
	// 创建多输出目标
	//infoWriter := io.MultiWriter(os.Stdout, handle)
	errorWriter := io.MultiWriter(os.Stderr, handle)
	debugWriter := io.MultiWriter(os.Stdout, handle)

	return &Logger{
		infoLogger:  log.New(handle, "[I] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(errorWriter, "[E] ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(debugWriter, "[D] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// InitLogger 初始化日志系统，并将其分配给全局 Logger
func InitLogger(logPath string) error {
	// 使用 lumberjack 库进行日志轮转
	logFile := &lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    10,      // 每10MB生成一个新文件
		MaxBackups: 3,       // 保留3个备份
		MaxAge:     28,      // 保留28天的日志
		Compress:   true,    // 是否压缩旧日志
	}

	// 创建新的 Logger 实例
	GlobalLogger = NewLogger(logFile)
	return nil
}

// Info 记录信息级别的日志
func Info(v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.infoLogger.Println(v...)
	}
}

// Infof 记录格式化信息级别的日志
func Infof(format string, v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.infoLogger.Printf(format, v...)
	}
}

// Error 记录错误级别的日志
func Error(v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.errorLogger.Println(v...)
		fmt.Println(v...)
	}
}

// Errorf 记录格式化错误级别的日志
func Errorf(format string, v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.errorLogger.Printf(format, v...)
		fmt.Printf(format, v...)
	}
}

// Debug 记录调试级别的日志
func Debug(v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.debugLogger.Println(v...)
	}
}

// Debugf 记录格式化调试级别的日志
func Debugf(format string, v ...interface{}) {
	if GlobalLogger != nil {
		GlobalLogger.debugLogger.Printf(format, v...)
	}
}
