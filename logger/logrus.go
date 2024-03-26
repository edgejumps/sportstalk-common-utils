package logger

import (
	"fmt"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

func newLogrusLogger() *log.Logger {
	logger := log.New()
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	logger.Formatter = customFormatter
	return logger
}

func cInfo(logger *log.Logger, txt ...interface{}) {
	logger.Infof(color.GreenString(fmt.Sprint(txt...)))
}

func cInfof(logger *log.Logger, format string, args ...interface{}) {
	logger.Infof(color.GreenString(fmt.Sprint(format)), args...)
}

func cDebug(logger *log.Logger, txt ...interface{}) {
	logger.Debugf(color.YellowString(fmt.Sprint(txt...)))
}

func cDebugf(logger *log.Logger, format string, args ...interface{}) {
	logger.Debugf(color.YellowString(fmt.Sprint(format)), args...)
}

func cError(logger *log.Logger, txt ...interface{}) {
	logger.Errorf(color.RedString(fmt.Sprint(txt...)))
}

func cErrorf(logger *log.Logger, format string, args ...interface{}) {
	logger.Errorf(color.RedString(fmt.Sprint(format)), args...)
}

func cWarn(logger *log.Logger, txt ...interface{}) {
	logger.Warnf(color.BlueString(fmt.Sprint(txt...)))
}

func cWarnf(logger *log.Logger, format string, args ...interface{}) {
	logger.Warnf(color.BlueString(fmt.Sprint(format)), args...)
}
