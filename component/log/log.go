package log

//Logger logger 对象
var Logger LoggerInterface

func init() {
	Logger = &DefaultLogger{}
}

//InitLogger 初始化logger对象
func InitLogger(ILogger LoggerInterface) {
	Logger = ILogger
}

//LoggerInterface 日志接口
type LoggerInterface interface {
	//Debugf(format string, params ...interface{})
	//
	//Infof(format string, params ...interface{})
	//
	//Warnf(format string, params ...interface{}) error
	//
	//Errorf(format string, params ...interface{}) error
	//
	//Debug(v ...interface{})
	//
	//Info(v ...interface{})
	//
	//Warn(v ...interface{}) error
	//
	//Error(v ...interface{}) error

	Debug(arg0 interface{}, args ...interface{})

	Info(arg0 interface{}, args ...interface{})

	Warn(arg0 interface{}, args ...interface{}) error

	Error(arg0 interface{}, args ...interface{}) error
}

//Debugf debug 格式化
func Debugf(format string, params ...interface{}) {
	Logger.Debug(format, params)
}

//Infof 打印info
func Infof(format string, params ...interface{}) {
	Logger.Info(format, params)
}

//Warnf warn格式化
func Warnf(format string, params ...interface{}) error {
	return Logger.Warn(format, params)
}

//Errorf error格式化
func Errorf(format string, params ...interface{}) error {
	return Logger.Error(format, params)
}

//Debug 打印debug
func Debug(v ...interface{}) {
	if len(v) > 1 {
		Logger.Debug(v[0], v[1:]...)
	} else {
		Logger.Debug(v)
	}
}

//Info 打印Info
func Info(v ...interface{}) {
	if len(v) > 1 {
		Logger.Info(v[0], v[1:]...)
	} else {
		Logger.Info(v)
	}
}

//Warn 打印Warn
func Warn(v ...interface{}) {
	if len(v) > 1 {
		Logger.Warn(v[0], v[1:]...)
	} else {
		Logger.Warn(v)
	}
}

//Error 打印Error
func Error(v ...interface{}) {
	if len(v) > 1 {
		Logger.Error(v[0], v[1:]...)
	} else {
		Logger.Error(v)
	}
}

//DefaultLogger 默认日志实现
type DefaultLogger struct {
}

//Debugf debug 格式化
func (d *DefaultLogger) Debugf(format string, params ...interface{}) {

}

//Infof 打印info
func (d *DefaultLogger) Infof(format string, params ...interface{}) {

}

//Warnf warn格式化
func (d *DefaultLogger) Warnf(format string, params ...interface{}) error {
	return nil
}

//Errorf error格式化
func (d *DefaultLogger) Errorf(format string, params ...interface{}) error {
	return nil
}

//Debug 打印debug
func (d *DefaultLogger) Debug(arg0 interface{}, args ...interface{}) {

}

//Info 打印Info
func (d *DefaultLogger) Info(arg0 interface{}, args ...interface{}) {

}

//Warn 打印Warn
func (d *DefaultLogger) Warn(arg0 interface{}, args ...interface{}) error {
	return nil
}

//Error 打印Error
func (d *DefaultLogger) Error(arg0 interface{}, args ...interface{}) error {
	return nil
}
