package Utilites

import (
	"webservice-pattern/Constants"
	"webservice-pattern/Models"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus Constants.ResponseStatus, data T) Models.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) Models.ApiResponse[T] {
	return Models.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
