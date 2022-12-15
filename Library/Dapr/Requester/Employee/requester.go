package Employee

import (
	"EmployeeService/Library/Dapr"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetEmployeeWithPhoneNo(phoneNo string) (data CheckEmployeeExistData, err error) {
	var config Dapr.ConfigRequest
	config.AppID = EmployeeService
	config.Method = http.MethodGet
	// config.Payload, _ = json.Marshal(phoneNo)
	config.Path = "/employee/check-mobile-phone-number/" + phoneNo
	result, err := Dapr.Request(config)
	if err != nil {
		return
	}
	resultByte, _ := json.Marshal(result.RestResult)
	err = json.Unmarshal(resultByte, &data)
	if err != nil {
		return
	}
	return
}
func GetEmployeeWithEmployeeId(UserId int64) (data CheckExistEmployeeID, err error) {
	var config Dapr.ConfigRequest
	config.AppID = EmployeeService
	config.Method = http.MethodGet
	// config.Payload, _ = json.Marshal(phoneNo)
	config.Path = "/employee/check-employee-id/" + fmt.Sprint(UserId)
	result, err := Dapr.Request(config)
	if err != nil {
		return
	}
	resultByte, _ := json.Marshal(result.RestResult)
	err = json.Unmarshal(resultByte, &data)
	if err != nil {
		return
	}
	return
}
