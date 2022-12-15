package Services

import (
	"EmployeeService/Controller/Dto"
	leavebalance "EmployeeService/Repository/leaveBalance"
	"database/sql"
)

var (
	LeaveBalance leavebalanceInterface = &leaveBalanceStruct{}
)

type leavebalanceInterface interface {
	UpdateLeaveBalance(params Dto.LeaveBalanceDataDTO) error
	SelectEmployeeLeaveBalance(params Dto.LeaveBalanceDataDTO) (leavebalance.LeaveBalance, error)
	SaveLeaveBalanceAdjusment(params Dto.LeaveBalanceAdjustmentDTO) (leavebalance.LeaveBalanceAdjustment, error)
	DataLeaveBalance(params Dto.LeaveBalanceParams) (leavebalance.LeaveBalanceData, error)
	DetailLeaveBalance(EmployeeId int64) (leavebalance.DetailLeaveBalance, error)
	DetailLeaveBalanceAdjustment(Tahun, EmployeeId int64) (leavebalance.LeaveBalanceAdjustmentData, error)
	ValidationDuplicateDataLeaveBalance(EmployeeId int64) (bool, error)
	HardDeleteLeaveBalance(EmployeeId int64) (bool, error)

	SubscriberLeaveBalanceUpdate(params Dto.SubscribeTransferDTO) (*sql.Tx, error)
	SubscriberLeaveBalanceApproved(params Dto.LeavesConfirmation) (*sql.Tx, error)

	SelectLeaveBalanceActive() ([]leavebalance.ResetLeaveBalance, error)
	UpdateLeaveBalanceActive(ID int64) error
	// Dieksekusi setelah EmployeeAdd dijalankan pada service employee
	SaveLeaveBalance(params Dto.LeaveBalanceDataDTO) error
}

type leaveBalanceStruct struct {
}
