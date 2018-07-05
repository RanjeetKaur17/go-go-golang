package main

import (
	log "github.com/sirupsen/logrus"
	logConfig "go-go-golang/GoLogging/log"
)

func main() {

	logConfig.InitializeLogging("/tmp/test.log")

	log.Println("Something Happened!!")
	log.Fatalf("What Happened??")

}