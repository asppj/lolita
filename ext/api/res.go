package api

// Res api response结构
type Res struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

// DefaultRes 获取默认api response
func DefaultRes() *Res {
	return &Res{true, SUCCESS, struct{}{}}
}

// NewRes 新建api response
func NewRes(success bool, msg string, data interface{}) (res Res) {
	res = Res{}
	res.Success = success
	res.Msg = msg
	res.Data = data
	return
}

// FailedRes 错误返回结构（just for swagger）
type FailedRes struct {
	Success bool        `json:"success" binding:"required" example:"false"`          // 请求结果，失败:false，成功:true
	Msg     string      `json:"msg" binding:"required" example:"req_data_val_error"` // 请求结果的message
	Data    interface{} `json:"data" binding:"required"`                             // 返回的数据
}

// SuccessRes 正确返回结构（just for swagger）
type SuccessRes struct {
	Success bool        `json:"success" binding:"required" example:"true"` // 请求结果，失败:false，成功:true
	Msg     string      `json:"msg" binding:"required" example:"ok"`       // 请求结果的message
	Data    interface{} `json:"data" binding:"required"`                   // 返回的数据
}
