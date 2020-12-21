package errors

import (
	"encoding/json"
	"fmt"
)

var (
	// 全部请求需要实现 200 正常请求
	// 默认实现200
	//c.JSON(http.StatusOK,....)

	// 负状态码，跟请求的时候的各种参数有关以及一些程序验证问题逻辑问题
	// GET问题 -100 到 -199
	ApiErrGet = ApiStandardError{
		ErrorCode: -112,
		ErrorMsg:  "参数异常，请检查错误信息，当前请求GET",
	}

	// POST问题 -200 到 -299
	ApiErrPostForm = ApiStandardError{
		ErrorCode: -212,
		ErrorMsg:  "参数不足或参数格式不对",
	}

	// DELETE -500 到 -599
	ApiErrWriteDelete = ApiStandardError{
		ErrorCode: -512,
		ErrorMsg:  "删除资源错误，具体错误未知",
	}
	//服务器错误处理函数以及结构体等数据的错误运行响应
	ApiErrFunc = ApiStandardError{
		ErrorCode: -500,
		ErrorMsg:  "服务器错误处理函数以及结构体等数据的错误运行响应",
	}
	//文件处理类错误
	ApiErrFile = ApiStandardError{
		ErrorCode: -11500,
		ErrorMsg:  "文件处理类错误",
	}
	//七牛云存储处理类错误
	ApiErrFileQiNiu = ApiStandardError{
		ErrorCode: -11501,
		ErrorMsg:  "七牛云存储处理类错误",
	}

	//微信类错误，不管是什么错误
	ApiErrWXChat = ApiStandardError{
		ErrorCode: -12501,
		ErrorMsg:  "微信处理类错误",
	}

	//服务器正常响应
	// code = 0
	ApiErrSuccess = ApiStandardError{
		ErrorCode: 0,
		ErrorMsg:  "正常!请安心食用!",
	}

	//举例
	//ApiErrSuccess.SetDate("任何数据类型")

	// 正状态码
	// 服务器过载问题拒绝，一般常见于自动503返回，一般不需要手动去调用，由HTTP—ROUTER自动完成  503
	// 服务器因为权限没有登录的问题拒绝 401
	ApiErrAuth = ApiStandardError{
		ErrorCode: 401,
		ErrorMsg:  "服务器因为权限没有登录的问题拒绝",
	}
	// 服务器认为你在非法操作拒绝 403
	ApiErrRefuse = ApiStandardError{
		ErrorCode: 403,
		ErrorMsg:  "服务器认为你在非法操作拒绝",
	}
	//请求头错误，安全漏洞
	ApiErrHeader = ApiStandardError{
		ErrorCode: 412,
		ErrorMsg:  "请求头错误，安全漏洞",
	}
	// 服务器MPG调度失败 600 ,一般多用于go协程使用过程中，channel 过载等处理
	ApiErrChannel = ApiStandardError{
		ErrorCode: 600,
		ErrorMsg:  "服务器MPG调度失败",
	}
	// 服务器GC失败 1001
	ApiErrGC = ApiStandardError{
		ErrorCode: 1001,
		ErrorMsg:  "服务器GC失败",
	}
	// 服务器处理超时任务失败 1004 超时
	ApiErrTimeout = ApiStandardError{
		ErrorCode: 1004,
		ErrorMsg:  "服务器处理超时任务失败",
	}
	//数据库处理错误
	ApiErrSql =  ApiStandardError{
		ErrorCode: 1005,
		ErrorMsg:  "任务失败，请检查具体错误信息",
	}

	ApiPending = ApiStandardError{
		ErrorCode: 2020,
		ErrorMsg:  "持续开发中",
	}
)

//设置返回给前端的内容，这个是一个常用接口
func (p ApiStandardError) SetDate(body interface{}) *ApiStandardError {
	p.Data = body
	return &p
}

//重新设置错误的信息，这个是一个常用接口
func (p ApiStandardError) SetErrMsg(body string) *ApiStandardError {
	p.ErrorMsg = body
	return &p
}

// 标准错误类结构体
type ApiStandardError struct {
	ErrorCode int         `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	Data      interface{} `json:"data,omitempty"` // 如果没有数据，那么不返回这个字段，而不是返回一个null
}

//自行声明一个不了解的错误,不了解的错误是默认400，其实也就是数据库的错误
func NewUnknownErr(e error) ApiStandardError {
	return ApiStandardError{400, e.Error(), nil}
}
func NewErr(errorCode int, errorMsg string, date interface{}) ApiStandardError {
	return ApiStandardError{errorCode, errorMsg, date}
}
func (p ApiStandardError) Error() string {
	bytes, e := json.Marshal(p)
	if e != nil {
		fmt.Println(e)
		return e.Error()
	}
	return string(bytes)
}
