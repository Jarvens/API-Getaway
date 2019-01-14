// auth: kunlun
// date: 2019-01-14
// description:
package main

import (
	"fmt"
	"log"
	"net/http"
	"server/controller"
	"time"
)

func main() {

	// 网关操作
	http.HandleFunc("/GateWay/Add", controller.AddGateWay)
	http.HandleFunc("/GateWay/Edit", controller.EditGateWay)
	http.HandleFunc("GateWay/Delete", controller.DeleteGateWay)

	// 接口操作
	http.HandleFunc("/Api/Add", controller.AddApi)
	http.HandleFunc("/Api/Edit", controller.EditApi)
	http.HandleFunc("/Api/Delete", controller.DeleteApi)

	//策略操作
	http.HandleFunc("/Strategy/Add", controller.AddStrategy)
	http.HandleFunc("/Strategy/Edit", controller.EditStrategy)
	http.HandleFunc("/Strategy/Delete", controller.DeleteStrategy)

	err := http.ListenAndServe(":9900", nil)

	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("api-gateway starting success time: %v", time.Now())
}
