// auth: kunlun
// date: 2019-01-15
// description:
package middle

import (
	"encoding/json"
	"net/http"
	"time"
	"yihuo"
)

func GetVisitCount(res http.ResponseWriter, req *http.Request, context *yihuo.Context) {
	visitCount, _ := json.Marshal(map[string]interface{}{
		"gatewaySuccessCount": context.VisitCount.SuccessCount.GetCount(),
		"gatewayFailureCount": context.VisitCount.FailureCount.GetCount(),
		"gatewayDayCount":     context.VisitCount.TotalCount.GetCount(),
		"lastUpdateTime":      time.Now().Format("15:04:05"),
	})

	res.Write(visitCount)
	return
}
