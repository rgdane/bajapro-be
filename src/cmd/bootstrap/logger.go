package bootstrap

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := strings.ToUpper(entry.Level.String())
	message := fmt.Sprintf("[%s] [%s] - %s\n", timestamp, level, entry.Message)
	return []byte(message), nil
}

var Log = logrus.New()

var once sync.Once

func InitLogger() {
	once.Do(func() {
		Log.SetFormatter(&CustomFormatter{})
		Log.SetOutput(os.Stdout)
		Log.SetLevel(logrus.InfoLevel)
	})
}
