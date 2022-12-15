package Notification

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Dapr"
	"encoding/json"
	"net/http"
)

func PushNotification(param Dto.NotificationInfoDTO) (err error) {
	var config Dapr.ConfigRequest
	config.AppID = NotificationService
	config.Method = http.MethodPost
	config.Payload, _ = json.Marshal(param)
	config.Path = "/notification/notifications"
	_, err = Dapr.Request(config)
	if err != nil {
		return
	}
	return
}
