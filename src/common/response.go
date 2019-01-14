// auth: kunlun
// date: 2019-01-14
// description:
package common

import "time"

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Body      interface{} `json:"body"`
	TimeStamp time.Time   `json:"time_stamp"`
}

const (
	SUCECSS    = 0
	FAIL       = 1
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

// SUCCESS
func (res Response) Success(data interface{}) Response {
	return Response{SUCECSS, "success", data, time.Now()}
}

// FAIL
func (res *Response) Fail(message string) *Response {
	return &Response{FAIL, message, nil, time.Now()}
}
