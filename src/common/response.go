// auth: kunlun
// date: 2019-01-14
// description:
package common

import "time"

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Body      interface{} `json:"body"`
	TimeStamp string      `json:"time_stamp"`
}

const (
	SUCCESS    = 0
	FAIL       = 1
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

// SUCCESS
func (res Response) Success(data interface{}) Response {
	return Response{SUCCESS, "success", data, time.Now().Format(TimeFormat)}
}

// FAIL
func (res *Response) Fail(message string) *Response {
	return &Response{FAIL, message, nil, time.Now().Format(TimeFormat)}
}
