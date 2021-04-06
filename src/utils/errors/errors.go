package errors

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

// package errors
//
// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// )
//
// // type RestErr interface {
// // 	Message() string
// // 	Status() int
// // 	Error() string
// // 	Causes() []interface{}
// // }
//
// type RestErr struct {
// 	ErrMessage string        `json:"message"`
// 	ErrStatus  int           `json:"status"`
// 	ErrError   string        `json:"error"`
// 	ErrCauses  []interface{} `json:"causes"`
// }
//
// func (e RestErr) Message() string {
// 	return e.ErrMessage
// }
//
// func (e RestErr) Status() int {
// 	return e.ErrStatus
// }
//
// func (e RestErr) Error() string {
// 	return fmt.Sprintf("message: %s - status %d - error %s - causes [ %v ]",
// 		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
// }
//
// func (e RestErr) Causes() []interface{} {
// 	return e.ErrCauses
// }
//
// func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
// 	return RestErr{
// 		ErrMessage: message,
// 		ErrStatus:  status,
// 		ErrError:   err,
// 		ErrCauses:  causes,
// 	}
// }
//
// func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
// 	var apiErr RestErr
// 	if err := json.Unmarshal(bytes, &apiErr); err != nil {
// 		return RestErr{}, errors.New("invalid json")
// 	}
// 	return apiErr, nil
// }
//
// func NewBadRequestError(message string) RestErr {
// 	return RestErr{
// 		ErrMessage: message,
// 		ErrStatus:  http.StatusBadRequest,
// 		ErrError:   "bad_request",
// 	}
// }
//
// func NewNotFoundError(message string) RestErr {
// 	return RestErr{
// 		ErrMessage: message,
// 		ErrStatus:  http.StatusNotFound,
// 		ErrError:   "not_found",
// 	}
// }
//
// func NewUnauthorizedError(message string) RestErr {
// 	return RestErr{
// 		ErrMessage: message,
// 		ErrStatus:  http.StatusUnauthorized,
// 		ErrError:   "unauthorized",
// 	}
// }
//
// func NewInternalServerError(message string, err error) RestErr {
// 	result := RestErr{
// 		ErrMessage: message,
// 		ErrStatus:  http.StatusInternalServerError,
// 		ErrError:   "internal_server_error",
// 	}
// 	if err != nil {
// 		result.ErrCauses = append(result.ErrCauses, err.Error())
// 	}
// 	return result
// }
