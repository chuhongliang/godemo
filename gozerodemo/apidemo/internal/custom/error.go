package custom

// 定义自定义错误类型
type CustomError struct {
	Message string
	Code    int
}

// 实现 Error 方法，满足 error 接口
func (e CustomError) Error() string {
	return e.Message
}
