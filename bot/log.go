package bot

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"go.uber.org/zap"
	"github.com/lnxjedi/robot"
)

// loggers of last resort, initialize early and update in start.go
func init() {
	botStdErrLogger = log.New(os.Stderr, "", log.LstdFlags)
	botStdOutLogger = log.New(os.Stdout, "", log.LstdFlags)
}

// initialized in start.go
var botStdErrLogger, botStdOutLogger *log.Logger

// Set by terminal connector
var terminalWriter io.Writer

var errorThreshold = robot.Warn

var (
	logger, _ = zap.NewProduction()
	sugar = logger.Sugar()
)

// Log logs messages whenever the connector log level is
// less than the given level
func Log(l robot.LogLevel, m string, v ...interface{}) bool {
	botLogger.Lock()
	currlevel := botLogger.level
	logger := botLogger.l
	botLogger.Unlock()
	prefix := logLevelToStr(l) + ":"
	msg := prefix + " " + m
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	if logger == nil && l >= currlevel {
		sugar.Info(msg)
		return true
	}
	if nullConn && l >= errorThreshold {
		sugar.Info(msg)
	}
	if l >= currlevel || l == robot.Audit {
		if l == robot.Fatal {
			sugar.Fatal(msg)
		} else {
			if localTerm && l >= errorThreshold {
				if terminalWriter != nil {
					terminalWriter.Write([]byte("LOG " + msg + "\n"))
				} else {
					botStdOutLogger.Print(msg)
				}
			}
			sugar.Info(msg)
			tsMsg := fmt.Sprintf("%s %s\n", time.Now().Format("Jan 2 15:04:05"), msg)
			botLogger.Lock()
			botLogger.buffer[botLogger.buffLine] = tsMsg
			botLogger.buffLine = (botLogger.buffLine + 1) % (buffLines - 1)
			botLogger.Unlock()
		}
		return true
	}
	return false
}
