package main

import (
	"log"
	"github.com/8tomat8/webAppSkeleton/environment"
	"github.com/8tomat8/webAppSkeleton/serverHTTP"
)

func main() {
	var err error
	var env environment.Env

	err = env.Start()
	if err != nil {
		log.Fatal(err)
	}

	serverHTTP.Run(&env)

	env.Stop()
	env.Check(err)
}
