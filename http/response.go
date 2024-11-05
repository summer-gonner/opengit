package http

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
	Status           string `json:"status"`
}
