package Services

import (
	"EmployeeService/Controller/Dto"
	leavebalance "EmployeeService/Repository/leaveBalance"
	"database/sql"
)

func (c leaveBalanceStruct) UpdateLeaveBalance(params Dto.LeaveBalanceDataDTO) error {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.UpdateLeaveBalance(params)
}

func (c leaveBalanceStruct) SelectEmployeeLeaveBalance(params Dto.LeaveBalanceDataDTO) (leavebalance.LeaveBalance, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.SelectEmployeeLeaveBalance(params)
}

func (c leaveBalanceStruct) SaveLeaveBalanceAdjusment(params Dto.LeaveBalanceAdjustmentDTO) (leavebalance.LeaveBalanceAdjustment, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.SaveLeaveBalanceAdjusment(params)
}

func (c leaveBalanceStruct) DataLeaveBalance(params Dto.LeaveBalanceParams) (leavebalance.LeaveBalanceData, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.DataLeaveBalance(params)
}

func (c leaveBalanceStruct) DetailLeaveBalance(EmployeeId int64) (leavebalance.DetailLeaveBalance, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.DetailLeaveBalance(EmployeeId)
}

func (c leaveBalanceStruct) DetailLeaveBalanceAdjustment(Tahun, EmployeeId int64) (leavebalance.LeaveBalanceAdjustmentData, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.DetailLeaveBalanceAdjustment(Tahun, EmployeeId)
}

func (c leaveBalanceStruct) ValidationDuplicateDataLeaveBalance(EmployeeId int64) (bool, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.ValidationDuplicateDataLeaveBalance(EmployeeId)
}

func (c leaveBalanceStruct) HardDeleteLeaveBalance(EmployeeId int64) (bool, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.HardDeleteLeaveBalance(EmployeeId)
}

func (c leaveBalanceStruct) SubscriberLeaveBalanceUpdate(params Dto.SubscribeTransferDTO) (*sql.Tx, error) {
	repo := &leavebalance.NewSubscriberLeaveBalanceRepository{}
	return repo.LeaveBalanceUpdate(params)
}

func (c leaveBalanceStruct) SubscriberLeaveBalanceApproved(params Dto.LeavesConfirmation) (*sql.Tx, error) {
	repo := &leavebalance.NewSubscriberLeaveBalanceRepository{}
	return repo.LeaveBalanceApproved(params)
}

func (c leaveBalanceStruct) SelectLeaveBalanceActive() ([]leavebalance.ResetLeaveBalance, error) {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.SelectLeaveBalanceActive()
}

func (c leaveBalanceStruct) UpdateLeaveBalanceActive(ID int64) error {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.UpdateLeaveBalanceActive(ID)
}

// Dieksekusi setelah EmployeeAdd dijalankan pada service employee
func (c leaveBalanceStruct) SaveLeaveBalance(params Dto.LeaveBalanceDataDTO) error {
	repo := &leavebalance.NewLeaveBalanceRepository{}
	return repo.SaveLeaveBalance(params)
}
