package logger

import (
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
)

var (
	useLogurs    = false
	zapLogger    *zap.SugaredLogger
	logrusLogger *log.Logger
)

func init() {
	zapLogger = newZapLogger()
}

// ToggleColorfulLogger toggles between logrus and zap logger
// By default, it will use zap logger.
func ToggleColorfulLogger() {
	useLogurs = !useLogurs

	if useLogurs && logrusLogger == nil {
		logrusLogger = newLogrusLogger()
	} else if !useLogurs {
		logrusLogger = nil
	}
}

func Info(msg ...interface{}) {
	if useLogurs {
		cInfo(logrusLogger, msg...)
	} else {
		zInfo(zapLogger, msg...)
	}
}

func Error(msg ...interface{}) {
	if useLogurs {
		cError(logrusLogger, msg...)
	} else {
		zError(zapLogger, msg...)
	}
}

func Debug(msg ...interface{}) {
	if useLogurs {
		cDebug(logrusLogger, msg...)
	} else {
		zDebug(zapLogger, msg...)
	}
}

func Warn(msg ...interface{}) {
	if useLogurs {
		cWarn(logrusLogger, msg...)
	} else {
		zWarn(zapLogger, msg...)
	}
}

func Infof(format string, args ...interface{}) {
	if useLogurs {
		cInfof(logrusLogger, format, args...)
	} else {
		zInfof(zapLogger, format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if useLogurs {
		cErrorf(logrusLogger, format, args...)
	} else {
		zErrorf(zapLogger, format, args...)
	}
}

func Debugf(format string, args ...interface{}) {
	if useLogurs {
		cDebugf(logrusLogger, format, args...)
	} else {
		zDebugf(zapLogger, format, args...)
	}
}

func Warnf(format string, args ...interface{}) {

	if useLogurs {
		cWarnf(logrusLogger, format, args...)
	} else {
		zWarnf(zapLogger, format, args...)
	}
}
