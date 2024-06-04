package main

import (
	"context"
	"distributed-learn/log"
	"distributed-learn/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done() // 如果服务shutdown，ctx的Done管道会传过来，然后继续执行
	fmt.Println("Shutting down log service.")
}
