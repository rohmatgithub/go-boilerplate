package common

import "errors"

type ErrorModel struct {
	Code     int
	Error    error
	CausedBy error
	// ErrorParameter        []ErrorParameter
	AdditionalInformation interface{}
}

func GenerateErrorModel(code int, err string, causedBy error) ErrorModel {
	var errModel ErrorModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.CausedBy = causedBy
	return errModel
}

func GenerateInternalDBServerError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-APP-DBS-001", causedBy)
}

func GenerateUnknownError(causedBy error) ErrorModel {
	return GenerateErrorModel(500, "E-5-APP-SRV-001", causedBy)
}
