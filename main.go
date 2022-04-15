package main

import (
	"UserAuth/common"
	"log"
	"time"

	"UserAuth/routers"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Port string `short:"p" long:"port" description:"set Tcp port to listen to"`
}

func main() {
	//Flags are used to read command line input while compiling the program.
	//Ex: go run main.go --port=8000
	//In the above example --port are flag with value 8000.
	//if flags are not provided the program executes the default one given in the program.
	//Here it is :8081
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalln("Error Parsing flags", err)
	}
	defer common.Db.Close()
	time.Sleep(time.Second * 3)
	common.Startup()
	e := routers.InitRoutes()

	//If flags are not empty means the server start at the given port
	//otherwise pick the default one.
	if opts.Port != "" {
		log.Panic(e.Start(":" + opts.Port))
	} else {
		log.Panic(e.Start(":8081"))
	}
}
