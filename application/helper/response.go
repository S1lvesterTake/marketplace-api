package helper

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
