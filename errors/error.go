package errors

import (
	"github.com/zzihyeon/go-graphql-server/graph/model"
	"github.com/zzihyeon/go-graphql-server/types"
)

func getErrorCode(code string) types.JHErrorCode {
	errorCode := make(map[string]types.JHErrorCode)
	/**
	controller
	**/
	/**
	service
	**/
	/**
	domain
	**/
	errorCode["Stock_Create_Failed"] = types.JHErrorCode{Code: 10301, Msg: "Stock create failed"}
	errorCode["Stock_Create_Panic"] = types.JHErrorCode{Code: 10302, Msg: "Stock create Panic fail"}
	errorCode["Stock_Delete_Failed"] = types.JHErrorCode{Code: 10303, Msg: "Stock delete failed"}
	errorCode["Stock_Delete_Panic"] = types.JHErrorCode{Code: 10304, Msg: "Stock delete Panic fail"}
	return errorCode[code]
}
func GetErrorResponse(code string, data interface{}) model.StandardResponse {
	defaultMsg := ""
	if code == "" {
		return model.StandardResponse{
			Code: 0,
			Payload: &model.StandardPayload{
				Msg:  &defaultMsg,
				Data: data,
			},
		}
	}
	errorData := getErrorCode(code)
	return model.StandardResponse{
		Code: errorData.Code,
		Payload: &model.StandardPayload{
			Msg:  &errorData.Msg,
			Data: nil,
		},
	}
}
