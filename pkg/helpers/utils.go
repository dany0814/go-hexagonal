package helpers

// Message api
func Message(code int, message string) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": message}
}

// Message api error
func MessageError(code int, err error) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": err.Error()}
}

// DataResponse api
func DataResponse(code int, message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": message, "data": data}
}
