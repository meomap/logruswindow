package logruswindow

import (
	"io/ioutil"
	"testing"

	"github.com/Sirupsen/logrus"
)

const (
	src = "mylog"
	msg = "Errors happened!"
)

func TestEventHook(t *testing.T) {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	hook, err := NewEventHook(src, []logrus.Level{
		logrus.ErrorLevel,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if closeErr := hook.Close(); closeErr != nil {
			t.Fatal(closeErr)
		}
	}()

	logger.Hooks.Add(hook)

	fields := logrus.Fields{
		"func":   "DoSomething",
		"server": "localhost",
		"tag":    "development",
	}

	logger.WithFields(fields).Error(msg)
	logger.WithFields(fields).Info(msg)
	logger.WithFields(fields).Debug(msg)
}
