package ControllerSubscriber

import (
	"EmployeeService/Controller/Dto"
	utils "EmployeeService/Library/Helper/Utils"
	"EmployeeService/Library/Logging"
	ServicesEm "EmployeeService/Services/employee"
	Services "EmployeeService/Services/leaveBalance"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"cloud.google.com/go/pubsub"
)

type LeaveBalanceInterface interface {
	SubsLeaveBalanceUpdate(ctx context.Context, m *pubsub.Message)
	SubsLeaveBalanceApproved(ctx context.Context, m *pubsub.Message)
}

// Keterangan : function subscriber ini digunakan untuk mengubah data karyawan setelah di mutasi
func (c Subscriber) SubsLeaveBalanceUpdate(ctx context.Context, m *pubsub.Message) {
	var data Dto.SubscribeTransferDTO
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	tx, err := Services.LeaveBalance.SubscriberLeaveBalanceUpdate(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLeaveBalanceUpdate->SubscriberLeaveBalanceUpdate(EmployeeId : %d)", data.EmployeeId),
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

	//EMPLOYEE
	var dataEm Dto.SubscribeEmployeePsgl
	err = json.Unmarshal(m.Data, &dataEm)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	// Kemudian Data juga di ubah di table Employee, baik di DB psgl
	tx, err = ServicesEm.Employee.SubscriberEmployeeUpdatePsgl(tx, dataEm)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLeaveBalanceUpdate->SubscriberEmployeeUpdatePsgl(EmployeeId : %d)", dataEm.EmployeeId),
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

	// Kemudian Data juga di ubah di table Employee, baik di DB psgl
	err = ServicesEm.Employee.SubscriberEmployeeUpdateMngo(dataEm)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLeaveBalanceUpdate->SubscriberEmployeeUpdateMngo(EmployeeId : %d)", dataEm.EmployeeId),
			"subscriptionID": subscriptionID,
			"errors":         err,
			"messageID":      m.ID}, nil)
		tx.Rollback()
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			m.Ack()
			return
		}
		m.Nack()
		return
	}

	tx.Commit()
	m.Ack()
}

// Keterangan : function subscriber ini digunakan untuk mengubah saldo cuti pegawai setelah permohonan cuti disetujui oleh atasan
func (c Subscriber) SubsLeaveBalanceApproved(ctx context.Context, m *pubsub.Message) {
	var data Dto.LeavesConfirmation
	err := json.Unmarshal(m.Data, &data)
	subscriptionID := utils.GetValueOfContext("subscriptionID", ctx).(string)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{"subscriptionID": subscriptionID, "errors": err, "messageID": m.ID}, nil)
		m.Nack()
		return
	}

	totalCurrentBalance := data.LeaveDay - data.LeaveBonusDay
	if totalCurrentBalance < 1 {
		totalCurrentBalance = 0
	}

	data.CurrentBalance = totalCurrentBalance
	tx, err := Services.LeaveBalance.SubscriberLeaveBalanceApproved(data)
	if err != nil {
		// if m.DeliveryAttempt != nil {
		// }
		Logging.LogError(map[string]interface{}{
			"function":       fmt.Sprintf("SubsLeaveBalanceApproved->SubscriberLeaveBalanceApproved(EmployeeId : %d)", data.EmployeeId),
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

	tx.Commit()
	m.Ack()
}
