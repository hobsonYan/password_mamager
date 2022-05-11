package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 200

	SelectFailureCode ResponseCode = 500

	SelectSuccessMsg ResponseMsg = "查询成功"

	SelectFailureMsg ResponseMsg = "查询失败"
)
