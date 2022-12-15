package ControllerSubscriber

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Dapr"
	"EmployeeService/Library/Dapr/Requester/Notification"
	utils "EmployeeService/Library/Helper/Utils"
	"EmployeeService/Library/Logging"
	Services "EmployeeService/Services/employee"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
)

type OrchestratorInterface interface {
	SubsOrchestratorEmployeePersonalInfoUpdated(ctx context.Context, m *pubsub.Message)
	SubsRollbackPersonalInfoUpdated(ctx context.Context, m *pubsub.Message)
}

// Keterangan : function Orchestrator subscriber ini digunakan untuk mengontrol Topic EmployeePersonalInfoUpdated di PubSub
func (c Subscriber) SubsOrchestratorEmployeePersonalInfoUpdated(ctx context.Context, m *pubsub.Message) {
	log.Info().Msgf("SubsOrchestratorEmployeePersonalInfoUpdated , data %s", string(m.Data))
	var param Dto.SubsEmployee
	err := json.Unmarshal(m.Data, &param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		log.Info().Msg("json.Unmarshal(m.Data, &data), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	UserID, err := strconv.ParseInt(fmt.Sprintf("%v", param.JWTDocodedPayload["UserID"]), 10, 64)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("strconv.ParseInt(%v", param.JWTDocodedPayload["UserID"]),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}
	EmployeeID, err := strconv.ParseInt(fmt.Sprintf("%v", param.JWTDocodedPayload["EmployeeID"]), 10, 64)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("strconv.ParseInt(%v", param.JWTDocodedPayload["EmployeeID"]),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}
	employeeData, err := Services.Employee.CheckExistEmployeeID(EmployeeID)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.CheckExistEmployeeID(Dto.EmployeeDataDTO{ID: %d})", EmployeeID),
			"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	var state Dapr.StateValue
	bytes, err := Dapr.StateGet(param.StateStore, param.Key)
	if err != nil {
		log.Info().Msgf(" Dapr.StateGet(%s, %s), error: "+err.Error(), param.StateStore, param.Key)
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		log.Info().Msg(" json.Unmarshal(bytes, &state), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	log.Info().Msgf("Dapr.CheckServices(subscriptionID:%%d, state.Services: %+v)", subscriptionID, state.Services)
	if Dapr.CheckServices(subscriptionID, state.Services) {
		notif := Dto.NotificationInfoDTO{
			SenderUserId:   UserID,
			SenderId:       employeeData.ID,
			SenderCode:     employeeData.Code,
			SenderName:     employeeData.Name,
			SenderPhoto:    employeeData.ProfilePhoto,
			ReceiverUserId: UserID,
			ReceiverId:     employeeData.ID,
			ReceiverCode:   employeeData.Code,
			ReceiverName:   employeeData.Name,
			ReceiverPhoto:  employeeData.ProfilePhoto,
			Application:    "EmployeeService",
			Title:          "Personal Info Update",
			Message:        "Success",
			Transaction:    "EmployeePersonalInfoUpdated",
			TransactionId:  param.Data.Code,
			UrlPath:        "",
		}
		err = Notification.PushNotification(notif)
		if err != nil {
			log.Info().Msgf(" Notification.PushNotification(%v), error: "+err.Error(), notif)
			Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Notification.PushNotification(%v)", notif), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
			// m.Nack()
			// return
		}
		// Dapr.StateDelete(param.StateStore, param.Key)
		m.Ack()
		return
	}
	m.Nack()
}

//Keterangan : function Rollback subscriber ini digunakan untuk mengontrol Topic EmployeePersonalInfoUpdated di PubSub
func (c Subscriber) SubsRollbackPersonalInfoUpdated(ctx context.Context, m *pubsub.Message) {

	var param Dto.SubsEmployeeRollback
	err := json.Unmarshal(m.Data, &param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		log.Info().Msg("json.Unmarshal(m.Data, &data), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	UserID, err := strconv.ParseInt(fmt.Sprintf("%v", param.JWTDocodedPayload["UserID"]), 10, 64)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("strconv.ParseInt(%v", param.JWTDocodedPayload["UserID"]),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}
	EmployeeID, err := strconv.ParseInt(fmt.Sprintf("%v", param.JWTDocodedPayload["EmployeeID"]), 10, 64)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("strconv.ParseInt(%v", param.JWTDocodedPayload["EmployeeID"]),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}
	employeeData, err := Services.Employee.CheckExistEmployeeID(EmployeeID)
	if err != nil {
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	var stateValue Dapr.StateValue
	bytes, err := Dapr.StateGet(param.StateStore, param.Key)
	if err != nil {
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	err = json.Unmarshal(bytes, &stateValue)
	if err != nil {
		log.Info().Msg(" json.Unmarshal(bytes, &state), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	stateBytes, _ := json.Marshal(stateValue.State)
	var statePayload Dto.EmployeeUpdateDTO
	err = json.Unmarshal(stateBytes, &statePayload)
	if err != nil {
		log.Info().Msg(" json.Unmarshal(stateBytes, &statePayload), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	// if Dapr.GetService(subscriptionID, stateValue.Services) {
	//ROLLBACK WITH STATE, THIS OLD DATA
	tx, _, err := Services.Employee.UpdateEmployee(statePayload)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.SubscriberEmployeePersonalInfoUpdated(statePayload : %v)", statePayload),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			m.Ack()
			return
		}
		m.Nack()
		return
	}
	if err := tx.Commit(); err != nil {
		log.Info().Msgf("tx.Commit(), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"function": "tx.Commit()", "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
	}
	// }
	notif := Dto.NotificationInfoDTO{
		SenderUserId:   UserID,
		SenderId:       employeeData.ID,
		SenderCode:     employeeData.Code,
		SenderName:     employeeData.Name,
		SenderPhoto:    employeeData.ProfilePhoto,
		ReceiverUserId: UserID,
		ReceiverId:     employeeData.ID,
		ReceiverCode:   employeeData.Code,
		ReceiverName:   employeeData.Name,
		ReceiverPhoto:  employeeData.ProfilePhoto,
		Application:    "EmployeeService",
		Title:          "Personal Info Update",
		Message:        "Failed",
		Transaction:    "EmployeePersonalInfoUpdated",
		TransactionId:  fmt.Sprintf("%d", param.Data.ID),
		UrlPath:        "",
	}
	err = Notification.PushNotification(notif)
	if err != nil {
		log.Info().Msgf(" Notification.PushNotification(%v), error: "+err.Error(), notif)
		Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Notification.PushNotification(%v)", notif), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	m.Ack()
}
