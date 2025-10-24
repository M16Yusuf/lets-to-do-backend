package models

type Response struct {
	IsSuccess bool   `json:"is_success"`
	Code      int    `json:"code,omitempty"`
	Msg       string `json:"message,omitempty"`
}

type ErrorResponse struct {
	Response
	Err string `json:"error"`
}

type ResponseData struct {
	Response
	Data       any        `json:"data,omitempty"`
	Pagenation Pagenation `json:"pagenation"`
}

type Pagenation struct {
	Page      int `json:"page,omitempty"`
	Limit     int `json:"limit,omitempty"`
	Total     int `json:"total,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
}
