package Response

import (
	"EmployeeService/Constant"
	"encoding/json"
	"net/http"
)

type RespErrorStruct struct {
	RestTitle     string        `json:"title" example:"bad_request"`
	RestStatus    int           `json:"status" example:"400"`
	RestMessage   string        `json:"message"`
	RestResultErr []interface{} `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

type RespResultStruct struct {
	RestTitle   string      `json:"title"`
	RestStatus  int         `json:"status"`
	RestMessage string      `json:"message"`
	RestResult  interface{} `json:"result" swaggertype:"object,string" example:"key:value,key2:value2"`
}

func ResponseJson(w http.ResponseWriter, body interface{}, internalCode Constant.InternalSuccess) {
	result := RespResultStruct{
		RestTitle:   internalCode.Info().Title,
		RestMessage: internalCode.Info().Description,
		RestStatus:  int(internalCode),
		RestResult:  body,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(internalCode.Info().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseError(w http.ResponseWriter, err error, internalCode Constant.InternalError) {
	result := RespErrorStruct{
		RestTitle:   internalCode.Info().Title,
		RestMessage: internalCode.Info().Description,
		RestStatus:  int(internalCode),
	}
	if err != nil {
		result.RestResultErr = append(result.RestResultErr, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(internalCode.Info().HttpCode)
	json.NewEncoder(w).Encode(result)
}

func ResponseFile(w http.ResponseWriter, file []byte) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(file)
}
