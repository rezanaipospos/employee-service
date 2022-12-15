package Services

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Repository/employee"
	"database/sql"
)

func (c employeeStruct) BrowseEmployee(params Dto.EmployeeBrowseParams) (employee.EmployeesData, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.BrowseEmployee(params)
}

func (c employeeStruct) SearchEmployee(params Dto.EmployeeSearchDTO) ([]employee.SearchEmployees, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.SearchEmployee(params)
}

func (c employeeStruct) GetHead(parent int64) (employee.GetHead, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.GetHead(parent)
}

func (c employeeStruct) BrowseEmployeeDetail(params Dto.EmployeeDataDTO) (employee.Employees, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.BrowseEmployeeDetail(params)
}

func (c employeeStruct) CheckExistEmployeeID(employeeId int64) (employee.CheckExistEmployeeID, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.CheckExistEmployeeID(employeeId)
}

func (c employeeStruct) CheckExistMobilePhoneNumber(phoneNo string) (employee.CheckExistData, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.CheckExistMobilePhoneNumber(phoneNo)
}

func (c employeeStruct) CheckExistEmail(email string) (employee.CheckExistData, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.CheckExistEmail(email)
}

func (c employeeStruct) SaveEmployee(params Dto.EmployeeDataDTO) (*sql.Tx, employee.Employees, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.SaveEmployee(params)
}

func (c employeeStruct) UpdateEmployee(params Dto.EmployeeUpdateDTO) (*sql.Tx, employee.EmployeesUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployee(params)
}

func (c employeeStruct) DeleteEmployee(params Dto.EmployeeDataDTO) (*sql.Tx, bool, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.DeleteEmployee(params)
}

func (c employeeStruct) BrowseSubordinatesData(params Dto.EmployeeParams) ([]employee.Subordinates, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.BrowseSubordinatesData(params)
}

func (c employeeStruct) DetailSubordinatesData(Parent, ID int64) ([]employee.DetailSubordinates, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.DetailSubordinatesData(Parent, ID)
}

func (c employeeStruct) GetEmployeeWorkStatus(employeeId int64) (employee.GetEmployeeWorkStatus, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.GetEmployeeWorkStatus(employeeId)
}

func (c employeeStruct) SaveWorkStatus(tx *sql.Tx, params Dto.SaveWorkStatusDTO) (*sql.Tx, int64, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.SaveWorkStatus(tx, params)
}

func (c employeeStruct) HardDeleteWorkStatus(workStatusId int64) error {
	repo := &employee.NewEmployeeRepository{}
	return repo.HardDeleteWorkStatus(workStatusId)
}

func (c employeeStruct) UpdateEmployeeWorkStatus(params Dto.WorkStatusUpdateDTO) (*sql.Tx, employee.WorkStatusUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployeeWorkStatus(params)
}

func (c employeeStruct) UpdateEmployeeFingerprint(params Dto.FingerprintUpdateDTO) (*sql.Tx, employee.FingerprintUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployeeFingerprint(params)
}

func (c employeeStruct) UpdateEmployeeMachineId(params Dto.MachineIdUpdateDTO) (*sql.Tx, employee.MachineIdUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployeeMachineId(params)
}

func (c employeeStruct) UpdateEmployeeFaceId(params Dto.FaceIdUpdateDTO) (*sql.Tx, employee.FaceIdUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployeeFaceId(params)
}

func (c employeeStruct) UpdateEmployeeResignStatus(params Dto.ResignStatusUpdateDTO) (*sql.Tx, employee.ResignStatusUpdate, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.UpdateEmployeeResignStatus(params)
}

func (c employeeStruct) ValidationDuplicateDataEmployee(Code string) (bool, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.ValidationDuplicateDataEmployee(Code)
}

func (c employeeStruct) ValidationDuplicateMobilePhone(MobilePhoneNo string) (bool, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.ValidationDuplicateMobilePhone(MobilePhoneNo)
}

func (c employeeStruct) ValidationDuplicateEmail(Email string) (bool, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.ValidationDuplicateEmail(Email)
}

func (c employeeStruct) CountEmployees() (int64, error) {
	repo := &employee.NewEmployeeRepository{}
	return repo.CountEmployees()
}

func (c employeeStruct) SubscriberEmployeeSave(params Dto.EmployeeDataDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeSave(params)
}

func (c employeeStruct) SubscriberEmployeeUpdatePsgl(tx *sql.Tx, params Dto.SubscribeEmployeePsgl) (*sql.Tx, error) {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeUpdatePsgl(tx, params)
}

func (c employeeStruct) SubscriberEmployeeUpdateMngo(params Dto.SubscribeEmployeePsgl) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeUpdateMngo(params)
}

func (c employeeStruct) SubscriberEmployeePersonalInfoUpdated(params Dto.EmployeeDataDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeePersonalInfoUpdated(params)
}

func (c employeeStruct) SubscriberEmployeeDeleted(params Dto.EmployeeDataDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeDeleted(params)
}

func (c employeeStruct) SubscriberEmployeeWorkStatusUpdated(params Dto.WorkStatusUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeWorkStatusUpdated(params)
}

func (c employeeStruct) SubscriberEmployeeFingerUpdated(params Dto.FingerprintUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeFingerUpdated(params)
}

func (c employeeStruct) SubscriberEmployeeMachineIdUpdated(params Dto.MachineIdUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeMachineIdUpdated(params)
}

func (c employeeStruct) SubscriberEmployeeFaceIdUpdated(params Dto.FaceIdUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeFaceIdUpdated(params)
}

func (c employeeStruct) SubscriberEmployeeResigned(params Dto.EmployeeDataDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.EmployeeResigned(params)
}

func (c employeeStruct) SubscriberCompanyLocationUpdate(params Dto.SubscribeCompanyLocationCodeUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.CompanyLocationUpdate(params)
}

func (c employeeStruct) SubscriberCompanyUpdate(params Dto.SubscribeCompanyUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.CompanyUpdate(params)
}

func (c employeeStruct) SubscriberDepartmentUpdate(params Dto.SubscribeDepartmentUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.DepartmentUpdate(params)
}

func (c employeeStruct) SubscriberLocationUpdate(params Dto.SubscribeLocationUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.LocationUpdate(params)
}

func (c employeeStruct) SubscriberSectionUpdate(params Dto.SubscribeSectionUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.SectionUpdate(params)
}

func (c employeeStruct) SubscriberPositionUpdate(params Dto.SubscribePositionUpdateDTO) error {
	repo := &employee.NewSubscriberEmployeeRepository{}
	return repo.PositionUpdate(params)
}
