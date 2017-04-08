package environment

import (
	"os"
	"errors"

	"github.com/sirupsen/logrus"
)

func (env *Env) initLogger() error {
	defer func() { env.Log.Info("Log initialized") }()

	env.Log = logrus.New()
	env.Log.Formatter = &logrus.JSONFormatter{}

	logL, err := logrus.ParseLevel(*logLevel)
	if err != nil {
		return err
	}
	env.Log.Level = logL

	if *logFilePath == "" {
		return nil
	}

	f, err := os.OpenFile(*logFilePath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return errors.New("Unable to open log file - " + *logFilePath)
	}
	env.Log.Out = f

	return nil
}
