package transfer

import "time"

type NewTransferRepository struct{}

type NewSubscribeTransferRepository struct{}

type TransfersData struct {
	RecordsTotal int         `json:"recordsTotal"`
	HasReachMax  bool        `json:"hasReachMax"`
	Data         []Transfers `json:"data"`
}

type Transfers struct {
	ID                     int64      `json:"id"  `
	TransferCode           string     `json:"transferCode"`
	TransferDate           time.Time  `json:"transferDate"`
	EmployeeId             int64      `json:"employeeId"`
	EmployeeCode           string     `json:"employeeCode"`
	EmployeeName           string     `json:"employeeName"`
	MachineId              int64      `json:"machineId"`
	MachineName            string     `json:"machineName"`
	FingerPrintId          string     `json:"fingerPrintId"`
	OldEmployeeCode        string     `json:"oldEmployeeCode"`
	OldIdCompany           int64      `json:"oldIdCompany"`
	OldCompanyName         string     `json:"oldCompanyName"`
	OldIdlocation          int64      `json:"oldIdlocation"`
	OldLocationName        string     `json:"oldLocationName"`
	OldIdDepartment        int64      `json:"oldIdDepartment"`
	OldDepartmentName      string     `json:"oldDepartmentName"`
	OldIdSection           int64      `json:"oldIdSection"`
	OldSectionName         string     `json:"oldSectionName"`
	OldIdPosition          int64      `json:"oldIdPosition"`
	OldPositionName        string     `json:"oldPositionName"`
	OldCompanyLocationCode string     `json:"oldCompanyLocationCode"`
	NewEmployeeCode        string     `json:"newEmployeeCode"`
	NewIdCompany           int64      `json:"newIdCompany"`
	NewCompanyName         string     `json:"newCompanyName"`
	NewIdLocation          int64      `json:"newIdLocation"`
	NewLocationName        string     `json:"newLocationName"`
	NewIdDepartment        int64      `json:"newIdDepartment"`
	NewDepartmentName      string     `json:"newDepartmentName"`
	NewIdSection           int64      `json:"newIdSection"`
	NewSectionName         string     `json:"newSectionName"`
	NewIdPosition          int64      `json:"newIdPosition"`
	NewPositionName        string     `json:"newPositionName"`
	NewCompanyLocationCode string     `json:"newCompanyLocationCode"`
	Reason                 string     `json:"reason"`
	Deleted                bool       `json:"deleted"`
	CreatedBy              string     `json:"createdBy"`
	CreatedTime            time.Time  `json:"createdTime"`
	ModifiedBy             *string    `json:"modifiedBy"`
	ModifiedTime           *time.Time `json:"modifiedTime"`
	DeletedBy              *string    `json:"deletedBy"`
	DeletedTime            *time.Time `json:"deletedTime"`
}

var allowedFields = map[string]string{
	"transferCode":           "transfer_code",
	"transferDate":           "transfer_date",
	"employeeId":             "employee_id",
	"employeeCode":           "employee_code",
	"oldEmployeeCode":        "old_employee_code",
	"oldCompanyName":         "old_company_name",
	"oldLocationName":        "old_location_name",
	"oldDepartmentName":      "old_department_name",
	"oldSectionName":         "old_section_name",
	"oldPositionName":        "old_position_name",
	"oldCompanyLocationCode": "old_company_location_code",
	"newEmployeeCode":        "new_employee_code",
	"newCompanyName":         "new_company_name",
	"newLocationName":        "new_location_name",
	"newDepartmentName":      "new_department_name",
	"newSectionName":         "new_section_name",
	"newPositionName":        "new_position_name",
	"newCompanyLocationCode": "new_company_location_code",
	"reason":                 "reason",
	"deleted":                "deleted",
	"createdBy":              "created_by",
	"createdTime":            "created_time",
	"modifiedBy":             "modified_by",
	"modifiedTime":           "modified_time",
	"deletedBy":              "deleted_by",
	"deletedTime":            "deleted_time",
}
