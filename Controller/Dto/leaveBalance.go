package Dto

import "time"

type LeaveBalanceDataDTO struct {
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
	Deleted           bool       `json:"deleted"`
	CreatedBy         string     `json:"createdBy"`
	CreatedTime       time.Time  `json:"createdTime"`
	ModifiedBy        *string    `json:"modifiedBy"`
	ModifiedTime      *time.Time `json:"modifiedTime"`
	DeletedBy         *string    `json:"deletedBy"`
	DeletedTime       *time.Time `json:"deletedTime"`
}

type LeaveBalanceAdjustmentDTO struct {
	Id           int64     `json:"id" `
	Tahun        int64     `json:"tahun"`
	EmployeeId   int64     `json:"employeeId" `
	EmployeeCode string    `json:"employeeCode" validate:"required,max=50"`
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

type LeaveBalanceParams struct {
	LeaveBalanceDataDTO
	SortColumn string `json:"sortColumn"`
	SortOrder  string `json:"sortOrder"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
}

type SubscribeTransferDTO struct {
	EmployeeId             int64  `json:"employeeId"`
	NewEmployeeCode        string `json:"newEmployeeCode"`
	NewCompanyName         string `json:"newCompanyName"`
	NewLocationName        string `json:"newLocationName"`
	NewDepartmentName      string `json:"newDepartmentName"`
	NewSectionName         string `json:"newSectionName"`
	NewPositionName        string `json:"newPositionName"`
	NewCompanyLocationCode string `json:"newCompanyLocationCode"`
}

// Leave Confirmation subs
type LeavesConfirmation struct {
	ID               int64     `json:"id"  `
	LeavesSettingsId int64     `json:"leavesSettingsId"`
	Name             string    `json:"name" validate:"required,max=200"`
	EmployeeId       int64     `json:"employeeId"`
	EmployeeCode     string    `json:"employeeCode"`
	EmployeeName     string    `json:"employeeName"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	Reason           string    `json:"reason"`
	LeaveDay         int64     `json:"leaveDay"`
	LeaveBonusDay    int64     `json:"leaveBonusDay"`
	Files            string    `json:"files"`
	Status           string    `json:"status"`
	CurrentBalance   int64     `json:"currentBalance"`
}
