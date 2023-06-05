package base

type NASARsp struct {
	Code           int    `json:"code"`
	Msg            string `json:"msg"`
	ServiceVersion string `json:"service_version"`

	HttpError    string `json:"http_error"`
	ErrorMessage string `json:"error_message"`
	Request      string `json:"request"`

	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
