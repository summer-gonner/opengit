package http

import (
	"io"
	"k8s.io/apimachinery/pkg/util/json"
	"log"
)

func IoToErrorResponse(origin io.ReadCloser) (*ErrorResponse, error) {
	body, err := io.ReadAll(origin)
	if err != nil {
		log.Printf("错误1%+v\n", err)
		return nil, err
	}
	var errorResponse *ErrorResponse
	err = json.Unmarshal(body, &errorResponse)
	if err != nil {
		log.Printf("错误2%+v\n", err)

		return nil, err
	}
	return errorResponse, nil
}
