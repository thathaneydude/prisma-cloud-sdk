package client

import (
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	msg string
}

type InternalServerError struct {
	msg string
}

type TooManyRequestError struct {
	msg string
}

func (m *UnauthorizedError) Error() string {
	return fmt.Sprintf("%v unauthorized: %v", http.StatusUnauthorized, m.msg)
}

func (m *InternalServerError) Error() string {
	return fmt.Sprintf("%v internal server error: %v", http.StatusInternalServerError, m.msg)
}

func (m *TooManyRequestError) Error() string {
	return fmt.Sprintf("%v too many tequests: %v", http.StatusTooManyRequests, m.msg)
}
