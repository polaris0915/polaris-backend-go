package main

import (
	"errors"
	"github.com/backend/model"
	"github.com/backend/router"
	"github.com/jpillora/overseer"
	"log"
	"net/http"
)

func Program() func(state overseer.State) {
	return func(state overseer.State) {
		// 判断state.Listener是否为空，参考main.md
		if state.Listener == nil {
			return
		}
		ser := &http.Server{
			Addr: "127.0.0.1:8080",
			// gin.Engine实现了
			//type Handler interface {
			//	ServeHTTP(ResponseWriter, *Request)
			//}
			// 因此可以直接将gin.Engine写到这里
			Handler: router.InitRouter(),
		}
		// 关于这里为什么不用ListenAndServe而是用Serve，参考main.md
		if err := ser.Serve(state.Listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// TODO: 更换自己的日志系统
			log.Fatal(err)
		}
	}
}

func main() {
	//overseer.Run(overseer.Config{
	//	Program:          Program(),
	//	Address:          "127.0.0.1:8080",
	//	TerminateTimeout: 5 * time.Second,
	//})
	model.InitMysql()
	log.Fatal(router.InitRouter().Run())
}
