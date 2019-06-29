package logging

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

var Log *zap.Logger

type ConfStruct struct {

	APPName string `default:"app name"`

	Log struct{
		LogDir  string `yaml:"log_dir"`
		LogFile string `yaml:"log_file"`
		LogDays int `yaml:"days"`
		Rotationtime string `yaml:"rotationtime"`
	} `yaml:"logger"`
}



func Load(file string, c *ConfStruct) error {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf, &c)
}

func GetConfig(file string) ConfStruct {
	c := ConfStruct{}
	if err := Load(file, &c); err!=nil {
		panic(err)
	}
	return c
}


func initLogger(filename, seg string) *zap.Logger {

	var rtime  time.Duration
	var fm string
	if seg == "hour" {
		fm = ".%Y-%m-%d-%H"
		rtime = time.Hour
	} else if seg == "minute" {
		fm = ".%Y-%m-%d-%H-%M"
		rtime = time.Minute
	} else{
		fm = ".%Y-%m-%d"
		rtime = time.Hour * 24
	}

	hook, _ := rotatelogs.New(
		filename + fm,
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(rtime),
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.InfoLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development, zap.AddCallerSkip(1))

	//logger.Info("log 初始化成功")
	return logger
}

type KV map[string]interface{}

func kvToZapField(kv KV) []zap.Field {
	zf := make([]zap.Field,0, len(kv))
	for k, v := range kv {
		zf = append(zf, zap.Reflect(k, v))
	}
	return zf
}

func Debugln(kv KV, msg string)            { Log.Debug(msg, kvToZapField(kv)...) }
func Infoln(kv KV, msg string)             { Log.Info(msg, kvToZapField(kv)...) }
func Warnln(kv KV, msg string)             {Log.Warn(msg, kvToZapField(kv)...)}
func Errorln(err error, kv KV, msg string) { Log.Error(msg, kvToZapField(kv)...) }
func Fatalln(kv KV, msg string)            { Log.Fatal(msg, kvToZapField(kv)...) }



func InitLogger(filename, seg string)  {
	Log = initLogger(filename, seg )
}