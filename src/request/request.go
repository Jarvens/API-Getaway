// auth: kunlun
// date: 2019-01-15
// description:
package request

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

var Version = "1.0"

// 定义接口
type Request interface {
	SetHeader(string, ...string) Request
	Headers() map[string][]string
	SetQueryParam(string, ...string) Request
	QueryParams() map[string][]string
	UrlPath() string
	SetJSON(string) Request
	SetRawBody([]byte) Request
	SetURL(string)
	SetFormParam(string, ...string) Request
	FormParams() map[string][]string

	// 文件操作
	AddFile(string, string, []byte) Request
	Send() (Response, err error)
}

// 定义表单文件上传结构体
type formFile struct {
	filename string
	data     []byte
}

// 定义请求结构体
type request struct {
	client      *http.Client
	method      string
	URL         string
	headers     map[string][]string
	cookies     map[string]string
	isJson      bool
	body        []byte
	formParams  map[string][]string
	queryParams map[string][]string
	files       map[string]*formFile
}

// 请求头设置
// value 不为空则 set 否则删除该headers中的key
func (request *request) SetHeader(key string, values ...string) Request {

	if len(values) > 0 {
		request.headers[key] = values[:]
	} else {
		delete(request.headers, key)
	}
	return request
}

// 获取所有头部信息
func (request *request) Headers() map[string][]string {
	headers := make(map[string][]string)
	for key, val := range request.headers {
		headers[key] = val[:]
	}
	return headers
}

// 设置查询参数
func (request *request) SetQueryParam(key string, values ...string) Request {
	if len(values) > 0 {
		request.queryParams[key] = values[:]
	} else {
		delete(request.queryParams, key)
	}
	return request
}

// 获取所有查询参数
func (request *request) QueryParams() map[string][]string {
	params := make(map[string][]string)
	for key, val := range request.queryParams {
		params[key] = val[:]
	}
	return params
}

func (request *request) UrlPath() string {
	if len(request.queryParams) > 0 {
		return request.URL + "?" + parseParams(request.queryParams).Encode()
	} else {
		return request.URL
	}
}

func (request *request) SetJSON(value string) Request {
	request.isJson = true
	request.body = []byte(value)
	return request
}

func (request *request) SetRawBody(body []byte) Request {
	request.isJson = false
	request.body = body
	return request
}

func (request *request) SetFormParam(key string, values ...string) Request {
	if len(values) > 0 {
		request.formParams[key] = values[:]
	} else {
		delete(request.formParams, key)
	}
	return request
}

func (this *request) AddFile(fieldname string, filename string, data []byte) Request {
	if fieldname != "" && filename != "" && data != nil {
		this.files[fieldname] = &formFile{filename: fieldname, data: data}
	}
	return this
}

func parseParams(params map[string][]string) url.Values {
	v := url.Values{}
	for key, values := range params {
		for _, value := range values {
			v.Add(key, value)
		}
	}
	return v
}

func (this *request) parseBody() (req *http.Request, err error) {
	if this.method == "GET" || this.method == "TRACE" {
		req, err = http.NewRequest(this.method, this.UrlPath(), nil)
	}

	if len(this.body) > 0 {
		if this.isJson {
			this.headers["Content-Type"] = []string{"application/json"}
			req, err = http.NewRequest(this.method, this.UrlPath(), strings.NewReader(string(this.body)))
		} else {
			var body *bytes.Buffer
			body = bytes.NewBuffer(this.body)
			req, err = http.NewRequest(this.method, this.UrlPath(), body)
		}
	} else if len(this.files) > 0 {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		var part io.Writer
		for fieldname, file := range this.files {
			part, err = writer.CreateFormFile(fieldname, file.filename)
			if err != nil {
				return
			}
			_, err = part.Write(file.data)
			if err != nil {
				return
			}
		}
		for fieldname, values := range this.formParams {
			temp := make(map[string][]string)
			temp[fieldname] = values

			value := parseParams(temp).Encode()
			err = writer.WriteField(fieldname, value)
			if err != nil {
				return
			}
		}
		err = writer.Close()
		if err != nil {
			return
		}
		this.headers["Content-Type"] = []string{writer.FormDataContentType()}
		req, err = http.NewRequest(this.method, this.UrlPath(), body)
	} else {
		this.headers["Content-Type"] = []string{"application/x-www-form-urlencoded"}
		req, err = http.NewRequest(this.method, this.UrlPath(),
			strings.NewReader(parseParams(this.formParams).Encode()))
	}
	return
}

func (this *request) Send() (res Response, err error) {
	req, err := this.parseBody()
	if err != nil {
		return
	}

}
