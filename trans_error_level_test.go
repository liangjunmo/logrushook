package logrushook_test

import (
	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook"
)

var (
	ErrorCode gocode.Code = "error"
	WarnCode  gocode.Code = "warn"
)

func ExampleTransErrorLevelLogrusHook() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logrus.AddHook(
		logrushook.NewTransErrorLevelLogrusHook(
			logrus.WarnLevel,
			[]gocode.Code{ErrorCode},
			true,
		),
	)

	logrus.WithError(WarnCode).Error("warn")
}