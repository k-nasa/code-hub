package httputil

import "fmt"

type HTTPError struct {
	Message string `json:"message"`
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("message=%v", he.Message)
}
