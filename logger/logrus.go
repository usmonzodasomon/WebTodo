package logger

import (
	"io"
	"os"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()

	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})

	infoFile, err := os.OpenFile("./logs/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Не удалось создать файл info.log: ", err)
	}

	errorFile, err := os.OpenFile("./logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Не удалось создать файл error.log: ", err)
	}

	warnFile, err := os.OpenFile("./logs/warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Не удалось создать файл warn.log: ", err)
	}

	debugFile, err := os.OpenFile("./logs/debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Не удалось создать файл debug.log: ", err)
	}

	logger.AddHook(&fileHook{
		LevelsArr: []logrus.Level{
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
		Files: map[logrus.Level]*os.File{
			logrus.ErrorLevel: errorFile,
			logrus.WarnLevel:  warnFile,
			logrus.InfoLevel:  infoFile,
			logrus.DebugLevel: debugFile,
		},
	})
}

type fileHook struct {
	LevelsArr []logrus.Level
	Files     map[logrus.Level]*os.File
}

func (hook *fileHook) Fire(entry *logrus.Entry) error {
	for _, level := range hook.LevelsArr {
		if entry.Level <= level {
			entry.Logger.Out = io.MultiWriter(hook.Files[level], colorable.NewColorableStdout())
			break
		}
	}
	return nil
}

func (hook *fileHook) Levels() []logrus.Level {
	return hook.LevelsArr
}

func CloseFile() {
	fileHook, ok := logger.Hooks[logrus.ErrorLevel][0].(*fileHook)
	if ok {
		for _, file := range fileHook.Files {
			if err := file.Close(); err != nil {
				logger.Errorf("Failed to close log file: %s", err)
			}
		}
	}
}

func GetLogger() *logrus.Logger {
	return logger
}
