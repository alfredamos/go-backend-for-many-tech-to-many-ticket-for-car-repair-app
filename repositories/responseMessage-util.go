package repositories

type ResponseMessage struct {
	Message    string
	StatusCode int
	Status     string
}

func NewResponseMessage(message string, statusCode int, status string) ResponseMessage {
	return ResponseMessage{Message: message, StatusCode: statusCode, Status: status}
}
