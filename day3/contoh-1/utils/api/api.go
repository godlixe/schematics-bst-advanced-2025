package apix

type HTTPResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
