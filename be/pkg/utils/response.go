package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
	Page    any    `json:"page,omitempty"`
	Limit   any    `json:"limit,omitempty"`
	Total   any    `json:"total,omitempty"`
	// Meta    any    `json:"meta,omitempty"`
}

type EmptyObj struct{}

func BuildResponseSuccess(message string, data any) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
	return res
}
func BuildResponseSuccessPaginate(message string, data any, page int, limit int, total int64) Response {
	res := Response{
		Status:  true,
		Message: message,
		Data:    data,
		Page:    page,
		Limit:   limit,
		Total:   total,
	}
	return res
}

func BuildResponseFailed(message string, err string, data any) Response {
	res := Response{
		Status:  false,
		Message: message,
		Error:   err,
		Data:    data,
	}
	return res
}
func BuildResponseFailedValidation(err any, data any) Response {
	res := Response{
		Status:  false,
		Message: "validation error",
		Error:   err,
		Data:    data,
	}
	return res
}
