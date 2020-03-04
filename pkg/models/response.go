package models

type Response struct {
	Error       bool              `json:"error"`
	ErrorText   string            `json:"error_text"`
	Data        *Data             `json:"data"`
	CustomError map[string]string `json:"custom_error"`
}

type Data struct{
	Res bool
}
