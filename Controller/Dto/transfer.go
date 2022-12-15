package Dto

import "time"

type TransferParams struct {
	TransferDataDTO
	SortColumn string `json:"sortColumn"`
	SortOrder  string `json:"sortOrder"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
}

type TransferDataDTO struct {
	ID                     int64      `json:"id"  `
	TransferCode           string     `json:"transferCode" validate:"required"`
	TransferDate           time.Time  `json:"transferDate"`
	EmployeeId             int64      `json:"employeeId" validate:"required"`
	EmployeeCode           string     `json:"employeeCode"`
	EmployeeName           string     `json:"employeeName"`
	MachineId              int64      `json:"machineId" validate:"required"`
	MachineName            string     `json:"machineName" validate:"required"`
	FingerPrintId          string     `json:"fingerPrintId" validate:"required"`
	OldEmployeeCode        string     `json:"oldEmployeeCode" validate:"required"`
	OldIdCompany           int64      `json:"oldIdCompany" validate:"required"`
	OldCompanyName         string     `json:"oldCompanyName" validate:"required"`
	OldIdlocation          int64      `json:"oldIdlocation" validate:"required"`
	OldLocationName        string     `json:"oldLocationName" validate:"required"`
	OldIdDepartment        int64      `json:"oldIdDepartment" validate:"required"`
	OldDepartmentName      string     `json:"oldDepartmentName" validate:"required"`
	OldIdSection           int64      `json:"oldIdSection" validate:"required"`
	OldSectionName         string     `json:"oldSectionName" validate:"required"`
	OldIdPosition          int64      `json:"oldIdPosition" validate:"required"`
	OldPositionName        string     `json:"oldPositionName" validate:"required"`
	OldCompanyLocationCode string     `json:"oldCompanyLocationCode" validate:"required"`
	NewEmployeeCode        string     `json:"newEmployeeCode" validate:"required"`
	NewIdCompany           int64      `json:"newIdCompany" validate:"required"`
	NewCompanyName         string     `json:"newCompanyName" validate:"required"`
	NewIdLocation          int64      `json:"newIdLocation" validate:"required"`
	NewLocationName        string     `json:"newLocationName" validate:"required"`
	NewIdDepartment        int64      `json:"newIdDepartment" validate:"required"`
	NewDepartmentName      string     `json:"newDepartmentName" validate:"required"`
	NewIdSection           int64      `json:"newIdSection" validate:"required"`
	NewSectionName         string     `json:"newSectionName" validate:"required"`
	NewIdPosition          int64      `json:"newIdPosition" validate:"required"`
	NewPositionName        string     `json:"newPositionName" validate:"required"`
	NewCompanyLocationCode string     `json:"newCompanyLocationCode" validate:"required"`
	Reason                 string     `json:"reason"`
	Deleted                bool       `json:"deleted"`
	CreatedBy              string     `json:"createdBy"`
	CreatedTime            time.Time  `json:"createdTime"`
	ModifiedBy             *string    `json:"modifiedBy"`
	ModifiedTime           *time.Time `json:"modifiedTime"`
	DeletedBy              *string    `json:"deletedBy"`
	DeletedTime            *time.Time `json:"deletedTime"`
}
