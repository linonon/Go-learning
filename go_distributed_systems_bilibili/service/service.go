package service

import (
	"context"
	"distributed/registry"
	"fmt"
	"log"
	"net/http"
)

// Start will registry a service and start it.
func Start(
	ctx context.Context,
	reg registry.Registration,
	host, port string,
	registerHandlersFunc func(),
) (
	context.Context, error,
) {
	registerHandlersFunc()
	ctx = startService(ctx, reg.ServiceName, host, port)
	err := registry.RegisterService(reg)
	if err != nil {
		return ctx, err
	}
	return ctx, nil
}

func startService(
	ctx context.Context,
	serviceName registry.ServiceName,
	host, port string,
) context.Context {
	// 創建 context with server
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		cancel()
	}()

	// 創建一個 監聽 鍵盤的 goroutine , 點擊就關閉 context
	go func() {
		fmt.Printf("%v started. Press any key to shut down\n", serviceName)
		var s string
		fmt.Scanln(&s)
		err := registry.ShutdownService(fmt.Sprintf("http://%s:%s", host, port))
		if err != nil {
			log.Println(err)
		}
		srv.Shutdown(ctx)
		cancel()
	}()

	return ctx
}
