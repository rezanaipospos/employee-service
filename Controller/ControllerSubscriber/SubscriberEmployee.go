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

type EmployeeInterface interface {
	SubsEmployeeAdded(ctx context.Context, m *pubsub.Message)
	SubsEmployeePersonalInfoUpdated(ctx context.Context, m *pubsub.Message)
	SubsEmployeeDeleted(ctx context.Context, m *pubsub.Message)

	SubsEmployeeWorkStatusUpdated(ctx context.Context, m *pubsub.Message)
	SubsRollbackEmployeeWorkStatusUpdated(ctx context.Context, m *pubsub.Message)

	SubsEmployeeResigned(ctx context.Context, m *pubsub.Message)
	SubsRollbackEmployeeResigned(ctx context.Context, m *pubsub.Message)

	SubsEmployeeFingerUpdated(ctx context.Context, m *pubsub.Message)
	SubsEmployeeMachineIdUpdated(ctx context.Context, m *pubsub.Message)
	SubsCompanyLocationUpdated(ctx context.Context, m *pubsub.Message)
	SubsCompanyUpdated(ctx context.Context, m *pubsub.Message)
	SubsDepartmentUpdate(ctx context.Context, m *pubsub.Message)
	SubsLocationUpdate(ctx context.Context, m *pubsub.Message)
	SubsSectionUpdate(ctx context.Context, m *pubsub.Message)
	SubsPositionUpdate(ctx context.Context, m *pubsub.Message)
	SubsEmployeeFaceIdUpdated(ctx context.Context, m *pubsub.Message)
}

// Keterangan : function subscriber ini digunakan untuk menambahkan data kaywayawan ke dalam database mongo
func (c Subscriber) SubsEmployeeAdded(ctx context.Context, m *pubsub.Message) {
	log.Info().Msgf("SubsEmployeeAdded , data %s", string(m.Data))
	var param Dto.SubsEmployee
	err := json.Unmarshal(m.Data, &param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
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
	err = Services.Employee.SubscriberEmployeeSave(param.Data)

	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeeAdded->SubscriberEmployeeSave(EmployeeID : %d)", param.Data.ID),
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
	employeeData, err := Services.Employee.CheckExistEmployeeID(EmployeeID)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.CheckExistEmployeeID(Dto.EmployeeDataDTO{ID: %d})", EmployeeID),
			"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
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
		Title:          "Employee Added",
		Message:        "Success",
		Transaction:    "EmployeeAdded",
		TransactionId:  param.Data.Code,
		UrlPath:        "",
	}
	err = Notification.PushNotification(notif)
	if err != nil {
		log.Info().Msgf(" Notification.PushNotification(%v), error: "+err.Error(), notif)
		Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Notification.PushNotification(%v)", notif), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
	}
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengupdate data karyawan di database mongo sesuai dengan ID karyawan
func (c Subscriber) SubsEmployeePersonalInfoUpdated(ctx context.Context, m *pubsub.Message) {
	log.Info().Msgf("SubsEmployeePersonalInfoUpdated , data %s", string(m.Data))
	var param Dto.SubsEmployee
	err := json.Unmarshal(m.Data, &param)
	log.Info().Msgf("param is  %+v", param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string) //Employee-EmployeePersonalInfoUpdated
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberEmployeePersonalInfoUpdated(param.Data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeePersonalInfoUpdated->SubscriberEmployeePersonalInfoUpdated(EmployeeID : %d)", param.Data.ID),
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
	var stateValue Dapr.StateValue
	bytes, err := Dapr.StateGet(param.StateStore, param.Key)
	if err != nil {
		log.Info().Msgf(" Dapr.StateGet(%s, %s), error: "+err.Error(), param.StateStore, param.Key)
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	log.Info().Msgf("bytes is  %s", string(bytes))
	err = json.Unmarshal(bytes, &stateValue)
	if err != nil {
		log.Info().Msgf(" json.Unmarshal(%s, %v), error: "+err.Error(), string(bytes), stateValue)
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	stateValue.Services = Dapr.SteteUpdateService(subscriptionID, "ok", stateValue.Services)
	stateCfg := Dapr.GetConfigState(nil, param.StateStore, param.Key, stateValue.State, stateValue.Services...)
	fmt.Printf("stateCfg %+v \n", stateCfg)
	err = Dapr.StatePublish(stateCfg)
	if err != nil {
		log.Info().Msgf(" Dapr.StatePublish(%v), error: "+err.Error(), stateCfg)
		Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Dapr.StatePublish(%v)", stateCfg), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	m.Ack()

}

// Keterangan : function subscriber ini digunakan untuk mengahapus data karyawan di database mongo sesuai dengan ID karyawan
func (c Subscriber) SubsEmployeeDeleted(ctx context.Context, m *pubsub.Message) {
	var data Dto.EmployeeDataDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberEmployeeDeleted(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeeDeleted->SubscriberEmployeeDeleted(EmployeeID : %d)", data.ID),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengupdate data status pekerja dari seorang karyawan (Mis. Karyawan Tetap,Kontrak,Trainning)
func (c Subscriber) SubsEmployeeWorkStatusUpdated(ctx context.Context, m *pubsub.Message) {
	var param Dto.SubsEmployeeUpdateWorkStatus
	// var data Dto.EmployeeDataDTO
	err := json.Unmarshal(m.Data, &param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
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

	workStatusUpdateData := Dto.WorkStatusUpdateDTO{
		WorkStatus:             param.Data.WorkStatus,
		ContractStart:          param.Data.ContractStart,
		ContractEnd:            param.Data.ContractEnd,
		WorkStatusChangeReason: param.Data.WorkStatusChangeReason,
		WorkStatusChangeDate:   param.Data.WorkStatusChangeDate,
	}
	err = Services.Employee.SubscriberEmployeeWorkStatusUpdated(workStatusUpdateData)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.SubscriberEmployeeWorkStatusUpdated(%+v)", workStatusUpdateData),
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

	var stateValue Dapr.StateValue
	bytes, err := Dapr.StateGet(param.StateStore, param.Key)
	if err != nil {
		log.Info().Msgf(" Dapr.StateGet(%s, %s), error: "+err.Error(), param.StateStore, param.Key)
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	log.Info().Msgf("bytes is  %s", string(bytes))
	err = json.Unmarshal(bytes, &stateValue)
	if err != nil {
		log.Info().Msgf(" json.Unmarshal(%s, %v), error: "+err.Error(), string(bytes), stateValue)
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	stateValue.Services = Dapr.SteteUpdateService(subscriptionID, "ok", stateValue.Services)
	stateCfg := Dapr.GetConfigState(nil, param.StateStore, param.Key, stateValue.State, stateValue.Services...)
	fmt.Printf("stateCfg %+v \n", stateCfg)
	err = Dapr.StatePublish(stateCfg)
	if err != nil {
		log.Info().Msgf(" Dapr.StatePublish(%v), error: "+err.Error(), stateCfg)
		Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Dapr.StatePublish(%v)", stateCfg), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

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
		TransactionId:  fmt.Sprintf("%d", param.Data.EmployeeId),
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

// Keterangan : function subscriber ini digunakan untuk rollback data EmployeeWorkStatus
func (c Subscriber) SubsRollbackEmployeeWorkStatusUpdated(ctx context.Context, m *pubsub.Message) {
	var param Dto.SubsEmployeeUpdateWorkStatus

	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)

	err := json.Unmarshal(m.Data, &param)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		// }
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
	bytes, err := Dapr.StateGet(subscriptionID, param.Key)
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
	var statePayload Dto.UpdateWorkStatusStateDTO
	err = json.Unmarshal(stateBytes, &statePayload)
	if err != nil {
		log.Info().Msg(" json.Unmarshal(stateBytes, &statePayload), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}
	// if Dapr.GetService(subscriptionID, stateValue.Services) {
	//ROLLBACK WITH STATE, THIS OLD DATA
	tx, _, err := Services.Employee.UpdateEmployeeWorkStatus(Dto.WorkStatusUpdateDTO{
		WorkStatus:             statePayload.WorkStatus,
		ContractStart:          statePayload.ContractStart,
		ContractEnd:            statePayload.ContractEnd,
		WorkStatusChangeReason: statePayload.WorkStatusChangeReason,
		WorkStatusChangeDate:   statePayload.WorkStatusChangeDate,
	})
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.SubsUpdateEmployeeNameInShiftPatternEmployees(statePayload : %v)", statePayload),
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

	err = Services.Employee.HardDeleteWorkStatus(statePayload.WorkStatusId)
	if err != nil {
		tx.Rollback()
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.HardDeleteWorkStatus(%+v)", statePayload.WorkStatusId),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}
	if err = tx.Commit(); err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       "tx.Commit()",
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
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
		Title:          "Work Status Update",
		Message:        "Failed",
		Transaction:    "EmployeeWorkStatusUpdated",
		TransactionId:  fmt.Sprintf("%d", param.Data.EmployeeId),
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

// Keterangan : function subscriber ini digunakan untuk mengupdate data fingerprint pegawai
func (c Subscriber) SubsEmployeeFingerUpdated(ctx context.Context, m *pubsub.Message) {
	var data Dto.FingerprintUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberEmployeeFingerUpdated(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeeFingerUpdated->SubscriberEmployeeFingerUpdated(EmployeeID : %d)", data.ID),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengupdate data kode mesin pegawai
func (c Subscriber) SubsEmployeeMachineIdUpdated(ctx context.Context, m *pubsub.Message) {
	var data Dto.MachineIdUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberEmployeeMachineIdUpdated(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {}
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.SubscriberEmployeeMachineIdUpdated(MachineId : %d)", data.MachineId),
			"subscriptionID": subscriptionID,
			"errors":         err, "messageID": m.ID}, nil)
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			m.Ack()
			return
		}
		m.Nack()
		return
	}
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengupdate data face id pegawai
func (c Subscriber) SubsEmployeeFaceIdUpdated(ctx context.Context, m *pubsub.Message) {
	var data Dto.FaceIdUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberEmployeeFaceIdUpdated(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeeFaceIdUpdated->SubscriberEmployeeFaceIdUpdated(EmployeeID : %d)", data.ID),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengupdate status pegawai menjadi resign dan alasan resign
func (c Subscriber) SubsEmployeeResigned(ctx context.Context, m *pubsub.Message) {
	log.Info().Msgf("SubsEmployeeResigned , data %s", string(m.Data))
	var param Dto.SubsEmployee
	err := json.Unmarshal(m.Data, &param)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
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

	err = Services.Employee.SubscriberEmployeeResigned(param.Data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsEmployeeResigned->SubscriberEmployeeResigned(EmployeeID : %d)", param.Data.ID),
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

	employeeData, err := Services.Employee.CheckExistEmployeeID(EmployeeID)
	if err != nil {
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.CheckExistEmployeeID(Dto.EmployeeDataDTO{ID: %d})", EmployeeID),
			"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

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
		Title:          "Employee Resign",
		Message:        "Success",
		Transaction:    "EmployeeResign",
		TransactionId:  param.Data.Code,
		UrlPath:        "",
	}
	err = Notification.PushNotification(notif)
	if err != nil {
		log.Info().Msgf(" Notification.PushNotification(%v), error: "+err.Error(), notif)
		Logging.LogError(map[string]interface{}{"function": fmt.Sprintf("Notification.PushNotification(%v)", notif), "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
	}
	m.Ack()
}

//Keterangan : function Rollback subscriber ini digunakan untuk mengontrol Topic EmployeePersonalInfoUpdated di PubSub
func (c Subscriber) SubsRollbackEmployeeResigned(ctx context.Context, m *pubsub.Message) {

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
	var statePayload Dto.UpdateResignStatusStateDTO
	err = json.Unmarshal(stateBytes, &statePayload)
	if err != nil {
		log.Info().Msg(" json.Unmarshal(stateBytes, &statePayload), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	// ROLLBACK WITH STATE, THIS OLD DATA
	// Update Data
	tx, _, err := Services.Employee.UpdateEmployeeResignStatus(Dto.ResignStatusUpdateDTO{
		WorkStatus:   statePayload.WorkStatus,
		JoinDate:     statePayload.JoinDate,
		ResignDate:   statePayload.ResignDate,
		ResignReason: statePayload.ResignReason,
	})
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

	// Delete Work Status
	err = Services.Employee.HardDeleteWorkStatus(statePayload.WorkStatusId)
	if err != nil {
		tx.Rollback()
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("Services.Employee.HardDeleteWorkStatus(%+v)", statePayload.WorkStatusId),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		m.Nack()
		return
	}

	if err := tx.Commit(); err != nil {
		log.Info().Msgf("tx.Commit(), error: " + err.Error())
		Logging.LogError(map[string]interface{}{"function": "tx.Commit()", "subscriptionID": subscriptionID, "errors": err.Error(), "messageID": m.ID}, nil)
		m.Nack()
	}

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
		Title:          "Employee Resign",
		Message:        "Failed",
		Transaction:    "EmployeeResign",
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

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Comloc apabila terjadi perubahan data pada master comloc(Service CompanyStructure)
func (c Subscriber) SubsCompanyLocationUpdated(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeCompanyLocationCodeUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := ctx.Value("subscriptionID").(string)

	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberCompanyLocationUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsCompanyLocationUpdated->SubscriberCompanyLocationUpdate(OldCompanyLocationCode  : %s)", data.OldCompanyLocationCode),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Company apabila terjadi perubahan data pada master Company(Service CompanyStructure)
func (c Subscriber) SubsCompanyUpdated(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeCompanyUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := ctx.Value("subscriptionID").(string)

	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberCompanyUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsCompanyUpdated->SubscriberCompanyUpdate(CompanyId : %d)", data.CompanyId),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Department apabila terjadi perubahan data pada master Department(Service CompanyStructure)
func (c Subscriber) SubsDepartmentUpdate(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeDepartmentUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)

	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberDepartmentUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsDepartmentUpdate->SubscriberDepartmentUpdate(DepartmentId : %d)", data.DepartmentId),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Location apabila terjadi perubahan data pada master Location(Service CompanyStructure)
func (c Subscriber) SubsLocationUpdate(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeLocationUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)

	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberLocationUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLocationUpdate->SubscriberLocationUpdate(LocationId : %d)", data.LocationId),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Bagian apabila terjadi perubahan data pada master Bagian(Service CompanyStructure)
func (c Subscriber) SubsSectionUpdate(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeSectionUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberSectionUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLocationUpdate->SubscriberLocationUpdate(SectionId : %d)", data.SectionId),
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
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah seluruh data Posisi/jabatan apabila terjadi perubahan data pada master Posisi/jabatan(Service CompanyStructure)
func (c Subscriber) SubsPositionUpdate(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribePositionUpdateDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	err = Services.Employee.SubscriberPositionUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsPositionUpdate->SubscriberPositionUpdate(PositionId : %d)", data.PositionId),
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
	m.Ack()
}
