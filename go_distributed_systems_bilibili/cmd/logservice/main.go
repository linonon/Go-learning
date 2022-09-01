// log service entry point
package main

import (
	"context"
	"distributed/log"
	"distributed/registry"
	"distributed/service"
	"fmt"
	stlog "log"
)

func main() {

	// Define service attributes
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	// Init Registration Infomation
	r := registry.Registration{
		ServiceName:      registry.LogService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{},
		ServiceUpdateURL: serviceAddress + "/services",
	}

	// Start Service
	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatalln(err)
	}

	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
