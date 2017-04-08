package environment

import (
	"github.com/jinzhu/gorm"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"

	"flag"
	"github.com/sirupsen/logrus"
)

var configFilePath = flag.String("config", "config.json", "Path to configuration file.")
var logFilePath = flag.String("logfile", "", "Path to log file.")
var logLevel = flag.String("loglevel", "info", "Log level.")

type empty struct{}

type Env struct {
	Log       *logrus.Logger
	DB        *gorm.DB
	Conf      *config
	RedisPool *redis.Pool

	wg   *sync.WaitGroup
	stop chan empty
}

func (env *Env) Start() (error) {
	env.wg = &sync.WaitGroup{}
	env.stop = make(chan empty)
	flag.Parse()

	err := env.initLogger()
	if err != nil {
		return err
	}

	err = env.parseConfig()
	if env.Check(err) {
		return err
	}

	err = env.dbInit()
	if env.Check(err) {
		return err
	}

	err = env.redisInit()
	if env.Check(err) {
		return err
	}

	// Add here your custom initializer

	return nil
}

func (env *Env) Stop() {
	close(env.stop)

	stop := make(chan empty)
	go func() {
		defer close(stop)
		env.wg.Wait()
	}()

	if env.Conf.Server.StopTimeout == 0 {
		env.Conf.Server.StopTimeout = 60
	}

	for {
		select {
		case <-stop:
			env.Log.Info("Environment was successfully shut down.")
			return
		case <-time.NewTimer(time.Second * time.Duration(env.Conf.Server.StopTimeout)).C:
			env.Log.Warn("Application was stoped by timeout!")
			return
		}
	}

}
