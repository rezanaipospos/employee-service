package leaveBalance

import "time"

type NewLeaveBalanceRepository struct{}

type NewSubscriberLeaveBalanceRepository struct{}

type LeaveBalanceData struct {
	RecordsTotal int            `json:"recordsTotal"`
	HasReachMax  bool           `json:"hasReachMax"`
	Data         []LeaveBalance `json:"data"`
}

type LeaveBalance struct {
	ID                int64      `json:"id" field:"required"`
	EmployeeId        int64      `json:"employeeId" `
	EmployeeCode      string     `json:"employeeCode" `
	CompanyName       string     `json:"companyName" `
	LocationName      string     `json:"locationName" `
	DepartmentName    string     `json:"departmentName" `
	StartBalances     int64      `json:"startBalances" `
	IncreaseBalance   int64      `json:"increaseBalance" `
	DecreaseBalance   int64      `json:"decreaseBalance" `
	LastPeriodBalance int64      `json:"lastPeriodBalance" `
	CurrentBalance    int64      `json:"currentBalance" `
	IsActive          bool       `json:"isActive" `
	Period            string     `json:"period" `
	JoinDate          *time.Time `json:"joinDate"`
	ExpiredDate       *time.Time `json:"expiredDate" `
	Deleted           bool       `json:"deleted"`
	CreatedBy         string     `json:"createdBy"`
	CreatedTime       time.Time  `json:"createdTime"`
	ModifiedBy        *string    `json:"modifiedBy"`
	ModifiedTime      *time.Time `json:"modifiedTime"`
	DeletedBy         *string    `json:"deletedBy"`
	DeletedTime       *time.Time `json:"deletedTime"`
}

type DetailLeaveBalance struct {
	ID                int64      `json:"id" field:"required"`
	EmployeeId        int64      `json:"employeeId" `
	EmployeeCode      string     `json:"employeeCode" `
	EmployeeName      string     `json:"employeeName" `
	CompanyName       string     `json:"companyName" `
	LocationName      string     `json:"locationName" `
	DepartmentName    string     `json:"departmentName" `
	StartBalances     int64      `json:"startBalances" `
	IncreaseBalance   int64      `json:"increaseBalance" `
	DecreaseBalance   int64      `json:"decreaseBalance" `
	LastPeriodBalance int64      `json:"lastPeriodBalance" `
	CurrentBalance    int64      `json:"currentBalance" `
	IsActive          bool       `json:"isActive" `
	Period            string     `json:"period" `
	JoinDate          *time.Time `json:"joinDate"`
	ExpiredDate       *time.Time `json:"expiredDate" `
}

type LeaveBalanceAdjustmentData struct {
	Data []LeaveBalanceAdjustment `json:"data"`
}

type LeaveBalanceAdjustment struct {
	Id           int64     `json:"id" `
	EmployeeId   int64     `json:"employeeId" `
	EmployeeCode string    `json:"employeeCode" `
	StartDate    time.Time `json:"startDate" `
	EndDate      time.Time `json:"endDate" `
	Type         string    `json:"type" `
	Quantity     int64     `json:"quantity" `
	Reason       string    `json:"reason" `
	Deleted      bool      `json:"deleted" `
	CreatedBy    string    `json:"createdBy" `
	CreatedTime  time.Time `json:"createdTime" `
	LeaveId      int64     `json:"leaveId" `
}

type LeaveBalanceDetailAdjustment struct {
	ID           int64     `json:"id" `
	EmployeeId   int64     `json:"employeeId" `
	EmployeeCode string    `json:"employeeCode" `
	Name         string    `json:"name" `
	StartDate    time.Time `json:"startDate" `
	EndDate      time.Time `json:"endDate" `
	Type         string    `json:"type" `
	Quantity     int64     `json:"quantity" `
	Reason       string    `json:"reason" `
	LeaveId      int64     `json:"leaveId" `
}

type ResetLeaveBalance struct {
	ID                int64      `json:"id"  `
	EmployeeId        int64      `json:"employeeId" `
	EmployeeCode      string     `json:"employeeCode" `
	CompanyName       string     `json:"companyName" `
	LocationName      string     `json:"locationName" `
	DepartmentName    string     `json:"departmentName" `
	StartBalances     int64      `json:"startBalances" `
	IncreaseBalance   int64      `json:"increaseBalance" `
	DecreaseBalance   int64      `json:"decreaseBalance" `
	LastPeriodBalance int64      `json:"lastPeriodBalance" `
	CurrentBalance    int64      `json:"currentBalance" `
	IsActive          bool       `json:"isActive" `
	Period            string     `json:"period" `
	JoinDate          *time.Time `json:"joinDate"`
	ExpiredDate       time.Time  `json:"expiredDate" `
}

var allowedFields = map[string]string{
	"employeeId":     "employee_id",
	"employeeCode":   "employee_code",
	"companyName":    "company_name",
	"locationName":   "location_name",
	"departmentName": "department_name",
	"period":         "period",
	"joinDate":       "join_date",
	"expiredDate":    "expired_date",
	"deleted":        "deleted",
	"createdBy":      "created_by",
	"createdTime":    "created_time",
	"modifiedBy":     "modified_by",
	"modifiedTime":   "name",
	"deletedBy":      "deleted_by",
	"deletedTime":    "deleted_time",
}
