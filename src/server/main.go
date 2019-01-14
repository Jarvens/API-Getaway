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

	err := http.ListenAndServe(":9900", nil)

	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("api-gateway starting success time: %v", time.Now())
}
