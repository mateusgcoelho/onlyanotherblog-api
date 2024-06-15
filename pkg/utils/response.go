package utils

type defaultResponse struct {
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

type defaultErrorResponse struct {
	Message    *string     `json:"message"`
	StackTrace *string     `json:"stackTrace"`
	Data       interface{} `json:"data"`
}

func ReponseData(data interface{}) defaultResponse {
	return defaultResponse{
		Data: data,
	}
}

func ResponseErrorMessage(message string) defaultResponse {
	return defaultResponse{
		Message: &message,
	}
}

func ResponseErrorStackTrace(message string, stackStrace string) defaultErrorResponse {
	return defaultErrorResponse{
		Message:    &message,
		StackTrace: &stackStrace,
	}
}
