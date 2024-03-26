package logger

import "go.uber.org/zap"

func newZapLogger() *zap.SugaredLogger {
	l, _ := zap.NewDevelopment()
	defer l.Sync()

	return l.Sugar()
}

func zInfo(logger *zap.SugaredLogger, msg ...interface{}) {
	logger.Info(msg...)
}

func zError(logger *zap.SugaredLogger, msg ...interface{}) {
	logger.Error(msg...)
}

func zDebug(logger *zap.SugaredLogger, msg ...interface{}) {
	logger.Debug(msg...)
}

func zWarn(logger *zap.SugaredLogger, msg ...interface{}) {
	logger.Warn(msg...)
}

func zInfof(logger *zap.SugaredLogger, format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func zErrorf(logger *zap.SugaredLogger, format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func zDebugf(logger *zap.SugaredLogger, format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func zWarnf(logger *zap.SugaredLogger, format string, args ...interface{}) {
	logger.Warnf(format, args...)
}
