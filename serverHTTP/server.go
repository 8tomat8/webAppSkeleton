package serverHTTP

import (
	"os"
	"os/signal"
	"syscall"
	"log"
	"fmt"
	"net/http"
	"context"
	"github.com/8tomat8/webAppSkeleton/environment"
)

type empty struct{}

type test struct {
	Host string
	Port string

	StopTimeout int `json:"stop_timeout" default:"60"`
}

func Run(env *environment.Env) {
	r := NewRouter(env)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", env.Conf.Server.Host, env.Conf.Server.Port),
		Handler: r,
	}

	// Graceful shutdown
	stop := make(chan empty)
	go func() {
		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

		log.Printf("Received %v signal\n", <-sigs)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("could not shutdown: %v", err)
		}
		close(stop)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	<-stop
}
