package environment

import (
	"github.com/garyburd/redigo/redis"
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

func (env *Env) dbInit() (err error) {
	switch env.Conf.DB.Type {
	case "mysql":
		env.DB, err = env.initMysql()
		if env.DB.Error != nil {
			return env.DB.Error
		}
	default:
		return errors.New("There is no initializer for " + env.Conf.DB.Type)
	}

	// SQL queries tracing
	env.DB.LogMode(env.Conf.DB.Debug)

	env.wg.Add(1)
	go func() {
		<-env.stop
		env.DB.Close()
		env.wg.Done()
	}()
	return nil
}

func (env *Env) initMysql() (db *gorm.DB,err error){
	db, err = gorm.Open(env.Conf.DB.Type,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			env.Conf.DB.User,
			env.Conf.DB.Pass,
			env.Conf.DB.Host,
			env.Conf.DB.Port,
			env.Conf.DB.DB))

	if err != nil {
		return
	}
	return
}

func (env *Env) redisInit() (err error) {
	env.RedisPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%s", env.Conf.Redis.Host, env.Conf.Redis.Port))
		},
	}

	env.wg.Add(1)
	go func() {
		<-env.stop
		env.RedisPool.Close()
		env.wg.Done()
	}()
	return nil
}