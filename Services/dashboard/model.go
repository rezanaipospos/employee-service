package Services

import (
	"EmployeeService/Repository/dashboard"
)

var (
	Dashboard dashboardInterface = &dashboardStruct{}
)

type dashboardInterface interface {
	NewEmployeeData() ([]dashboard.NewEmployeeData, error)
	TotalReligionSummary() ([]dashboard.TotalReligionSummary, error)
	TotalWorkStatusSummary() ([]dashboard.TotalWorkStatusSummary, error)
	TotalWillExpireEmployeeContract() (dashboard.TotalWillExpireEmployeeContract, error)
	TotalEmployeeByLengthOfWork(numberOfYear int64) (dashboard.TotalEmployeeByLengthOfWork, error)
}

type dashboardStruct struct {
}
