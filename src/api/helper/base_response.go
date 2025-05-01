package helper

import "tahaNGU/api/validation"

type BaseHttpResponse struct {
	Result          any                           `json:"result"`
	Success         bool                          `json:"success"`
	ResultCode      int                           `json:"resultCode"`
	ValidationError *[]validation.ValidationError `json:"validationErrors"`
	Error           any                           `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}

func GenerateBaseResponseWithValidationError(result any, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:          result,
		ResultCode:      resultCode,
		ValidationError: validation.GetValidationError(err),
	}
}
