package commondata

type JsonError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
