package Services

import (
	"EmployeeService/Repository/dashboard"
)

func (c dashboardStruct) NewEmployeeData() ([]dashboard.NewEmployeeData, error) {
	repo := &dashboard.NewDashboardRepository{}
	return repo.NewEmployeeData()
}

func (c dashboardStruct) TotalReligionSummary() ([]dashboard.TotalReligionSummary, error) {
	repo := &dashboard.NewDashboardRepository{}
	return repo.TotalReligionSummary()
}

func (c dashboardStruct) TotalWorkStatusSummary() ([]dashboard.TotalWorkStatusSummary, error) {
	repo := &dashboard.NewDashboardRepository{}
	return repo.TotalWorkStatusSummary()
}

func (c dashboardStruct) TotalWillExpireEmployeeContract() (dashboard.TotalWillExpireEmployeeContract, error) {
	repo := &dashboard.NewDashboardRepository{}
	return repo.TotalWillExpireEmployeeContract()
}

func (c dashboardStruct) TotalEmployeeByLengthOfWork(numberOfYear int64) (dashboard.TotalEmployeeByLengthOfWork, error) {
	repo := &dashboard.NewDashboardRepository{}
	return repo.TotalEmployeeByLengthOfWork(numberOfYear)
}
