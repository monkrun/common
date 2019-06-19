package request

type RequestBody struct {
	ActionType int         //请求类型
	ActionId   string      //请求ID
	Body       interface{} //请求内容主体
}
