package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var (
	//port      = flag.Int("port", 50051, "The server port")
	//bootstrap = flag.String("bootstrap", "192.168.0.29:50051", "bootstrap address")
	//cmd       = flag.String("cmd", "", "commmand to issue [put/get]")
	//key       = flag.String("key", "", "key for command")
	//value     = flag.String("value", "", "value for command")
)

func main() {
	flag.Parse()

	log.Println("Starting . . .")

}


