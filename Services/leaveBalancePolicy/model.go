package Services

import (
	"EmployeeService/Controller/Dto"
	leaveBalancePolicy "EmployeeService/Repository/leaveBalancePolicy"
)

var (
	LeaveBalancePolicy leavebalancePolicyInterface = &leaveBalancePolicyStruct{}
)

type leavebalancePolicyInterface interface {
	SelectLeaveBalancePolicy(params Dto.LeaveBalancePolicyDTO) ([]leaveBalancePolicy.LeaveBalancePolicy, error)
	SaveLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) error
	UpdateLeaveBalancePolicy(leaveBalanceBonusList string, params Dto.LeaveBalancePolicyDTO) error
	ValidationDuplicateCompany(companyId int64) (exists bool, err error)
	CheckValidationCompanyName(companyName string) (leaveBalancePolicy.LeaveBalancePolicy, error)
}

type leaveBalancePolicyStruct struct{}
