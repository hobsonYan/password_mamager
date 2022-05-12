package constant

type ResponseCode int
type ResponseMsg string

const (
	SelectSuccessCode ResponseCode = 200

	SelectFailureCode ResponseCode = 500

	SelectSuccessMsg ResponseMsg = "操作成功"

	SelectFailureMsg ResponseMsg = "操作成功"
)
