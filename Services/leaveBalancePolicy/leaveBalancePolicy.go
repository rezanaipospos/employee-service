package Services

import (
	"EmployeeService/Controller/Dto"
	leaveBalancePolicy "EmployeeService/Repository/leaveBalancePolicy"
)

func (c leaveBalancePolicyStruct) SelectLeaveBalancePolicy(params Dto.LeaveBalancePolicyDTO) ([]leaveBalancePolicy.LeaveBalancePolicy, error) {
	repo := &leaveBalancePolicy.NewLeaveBalancePolicyRepository{}
	return repo.SelectLeaveBalancePolicy(params)
}

func (c leaveBalancePolicyStruct) SaveLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) error {
	repo := &leaveBalancePolicy.NewLeaveBalancePolicyRepository{}
	return repo.SaveLeaveBalancePolicy(leaveBalanceBonusList, params)
}

func (c leaveBalancePolicyStruct) UpdateLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) error {
	repo := &leaveBalancePolicy.NewLeaveBalancePolicyRepository{}
	return repo.UpdateLeaveBalancePolicy(leaveBalanceBonusList, params)
}

func (c leaveBalancePolicyStruct) ValidationDuplicateCompany(companyId int64) (exists bool, err error) {
	repo := &leaveBalancePolicy.NewLeaveBalancePolicyRepository{}
	return repo.ValidationDuplicateCompany(companyId)
}

func (c leaveBalancePolicyStruct) CheckValidationCompanyName(companyName string) (leaveBalancePolicy.LeaveBalancePolicy, error) {
	repo := &leaveBalancePolicy.NewLeaveBalancePolicyRepository{}
	return repo.CheckValidationCompanyName(companyName)
}
