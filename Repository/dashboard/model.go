package dashboard

import "time"

type NewDashboardRepository struct{}

type NewConsumerDashboardRepository struct{}

type NewEmployeeData struct {
	EmployeeId     int64      `json:"employeeId" bson:"id"`
	EmployeeCode   string     `json:"employeeCode" bson:"code"`
	EmployeeName   string     `json:"employeeName" bson:"name"`
	DepartmentName string     `json:"departmentName" bson:"department_name"`
	CompanyName    string     `json:"companyName" bson:"company_name"`
	PositionName   string     `json:"positionName" bson:"position_name"`
	JoinDate       *time.Time `json:"joinDate" bson:"join_date"`
}

type TotalReligionSummary struct {
	Religion string `json:"religion"`
	Value    int64  `json:"value"`
}

type TotalWorkStatusSummary struct {
	WorkStatus string `json:"workStatus"`
	Value      int64  `json:"value"`
}

type TotalWillExpireEmployeeContract struct {
	Percentage    float64 `json:"percentage"`
	CountContract int64   `json:"countContract"`
	CountTraining int64   `json:"countTraining"`
}

type TotalEmployeeByLengthOfWork struct {
	TotalEmployees int64 `json:"totalEmployees"`
	TotalMale      int64 `json:"totalMale"`
	TotalFemale    int64 `json:"totalFemale"`
}
