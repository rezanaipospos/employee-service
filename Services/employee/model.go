package Services

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Repository/employee"
	"database/sql"
)

var (
	Employee employeeInterface = &employeeStruct{}
)

type employeeInterface interface {
	BrowseEmployee(params Dto.EmployeeBrowseParams) (employee.EmployeesData, error)
	BrowseEmployeeDetail(params Dto.EmployeeDataDTO) (employee.Employees, error)
	SearchEmployee(params Dto.EmployeeSearchDTO) ([]employee.SearchEmployees, error)
	GetHead(parent int64) (employee.GetHead, error)

	SaveEmployee(params Dto.EmployeeDataDTO) (*sql.Tx, employee.Employees, error)
	ValidationDuplicateDataEmployee(Code string) (bool, error)
	ValidationDuplicateMobilePhone(MobilePhoneNo string) (bool, error)
	ValidationDuplicateEmail(Email string) (bool, error)

	UpdateEmployee(params Dto.EmployeeUpdateDTO) (*sql.Tx, employee.EmployeesUpdate, error)
	DeleteEmployee(params Dto.EmployeeDataDTO) (*sql.Tx, bool, error)
	BrowseSubordinatesData(params Dto.EmployeeParams) ([]employee.Subordinates, error)
	DetailSubordinatesData(Parent, ID int64) ([]employee.DetailSubordinates, error)

	UpdateEmployeeWorkStatus(params Dto.WorkStatusUpdateDTO) (*sql.Tx, employee.WorkStatusUpdate, error)
	UpdateEmployeeFingerprint(params Dto.FingerprintUpdateDTO) (*sql.Tx, employee.FingerprintUpdate, error)
	UpdateEmployeeMachineId(params Dto.MachineIdUpdateDTO) (*sql.Tx, employee.MachineIdUpdate, error)
	UpdateEmployeeFaceId(params Dto.FaceIdUpdateDTO) (*sql.Tx, employee.FaceIdUpdate, error)
	UpdateEmployeeResignStatus(params Dto.ResignStatusUpdateDTO) (*sql.Tx, employee.ResignStatusUpdate, error)
	GetEmployeeWorkStatus(employeeId int64) (employee.GetEmployeeWorkStatus, error)
	SaveWorkStatus(tx *sql.Tx, params Dto.SaveWorkStatusDTO) (*sql.Tx, int64, error)
	HardDeleteWorkStatus(workStatusId int64) error

	CheckExistMobilePhoneNumber(phoneNo string) (employee.CheckExistData, error)
	CheckExistEmail(email string) (employee.CheckExistData, error)
	CheckExistEmployeeID(employeeId int64) (employee.CheckExistEmployeeID, error)
	CountEmployees() (int64, error)

	//EmployeAdded
	SubscriberEmployeeSave(params Dto.EmployeeDataDTO) error
	// SubscriberValidationDuplicateData(params Dto.EmployeeDataDTO) (bool, error)
	//EmployeeTransfered
	SubscriberEmployeeUpdatePsgl(tx *sql.Tx, params Dto.SubscribeEmployeePsgl) (*sql.Tx, error)
	SubscriberEmployeeUpdateMngo(params Dto.SubscribeEmployeePsgl) error
	//EmployeePersonalInfoUpdated
	SubscriberEmployeePersonalInfoUpdated(params Dto.EmployeeDataDTO) error
	//EmployeeDeleted
	SubscriberEmployeeDeleted(params Dto.EmployeeDataDTO) error
	//EmployeeWorkStatusUpdated
	SubscriberEmployeeWorkStatusUpdated(params Dto.WorkStatusUpdateDTO) error
	//EmployeeFingerUpdated
	SubscriberEmployeeFingerUpdated(params Dto.FingerprintUpdateDTO) error
	//EmployeeResigned
	SubscriberEmployeeResigned(params Dto.EmployeeDataDTO) error
	//CompanyLocationUpdate
	SubscriberCompanyLocationUpdate(params Dto.SubscribeCompanyLocationCodeUpdateDTO) error
	//CompanyUpdate
	SubscriberCompanyUpdate(params Dto.SubscribeCompanyUpdateDTO) error
	//DepartmentUpdate
	SubscriberDepartmentUpdate(params Dto.SubscribeDepartmentUpdateDTO) error
	//LocationUpdate
	SubscriberLocationUpdate(params Dto.SubscribeLocationUpdateDTO) error
	//SectionUpdate
	SubscriberSectionUpdate(params Dto.SubscribeSectionUpdateDTO) error
	//PositionUpdate
	SubscriberPositionUpdate(params Dto.SubscribePositionUpdateDTO) error
	//EmployeeMachineIdUpdated
	SubscriberEmployeeMachineIdUpdated(params Dto.MachineIdUpdateDTO) error
	//EmployeeMachineIdUpdated
	SubscriberEmployeeFaceIdUpdated(params Dto.FaceIdUpdateDTO) error
}

type employeeStruct struct{}
