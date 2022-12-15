package Dto

import "time"

type EmployeeParams struct {
	EmployeeDataDTO
	SortColumn string `json:"sortColumn"`
	SortOrder  string `json:"sortOrder"`
	PageNumber int64  `json:"pageNumber"`
	PageSize   int64  `json:"pageSize"`
}

type EmployeeBrowseParams struct {
	EmployeeBrowseDTO
	SortColumn string `json:"sortColumn"`
	SortOrder  string `json:"sortOrder"`
	PageNumber int64  `json:"pageNumber"`
	PageSize   int64  `json:"pageSize"`
}

type EmployeeSearchDTO struct {
	Search string `json:"search"`
}

type GetHeadParams struct {
	Transaction string `json:"transaction"`
}

type EmployeSubcribeDTO struct {
	ID   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type EmployeeBrowseDTO struct {
	Code                string `json:"code" validate:"max=25" bson:"code"`
	Name                string `json:"name" bson:"name"`
	MachineId           *int64 `json:"machineId" bson:"machine_id"`
	CompanyName         string `json:"companyName" bson:"company_name" validate:"required"`
	CompanyLocationCode string `json:"companyLocationCode" bson:"company_location_code" validate:"required"`
	DepartmentName      string `json:"departmentName" bson:"department_name" validate:"required"`
	WorkStatus          string `json:"workStatus" bson:"work_status"`
}

type EmployeeDataDTO struct {
	ID                     int64      `json:"id" bson:"id"`
	Code                   string     `json:"code" validate:"max=25" bson:"code"`
	Type                   string     `json:"type" bson:"type"`
	FingerPrintId          string     `json:"fingerPrintId" bson:"finger_print_id"`
	FaceId                 string     `json:"faceId" bson:"face_id"`
	MachineId              int64      `json:"machineId" bson:"machine_id"`
	MachineName            string     `json:"machineName" bson:"machine_name"`
	DepartmentId           int64      `json:"departmentId" bson:"department_id" validate:"required"`
	DepartmentName         string     `json:"departmentName" bson:"department_name" validate:"required"`
	SectionId              int64      `json:"sectionId" bson:"section_id" validate:"required"`
	SectionName            string     `json:"sectionName" bson:"section_name" validate:"required"`
	PositionId             int64      `json:"positionId" bson:"position_id" validate:"required"`
	PositionName           string     `json:"positionName" bson:"position_name" validate:"required"`
	CompanyId              int64      `json:"companyId" bson:"company_id" validate:"required"`
	CompanyName            string     `json:"companyName" bson:"company_name" validate:"required"`
	LocationId             int64      `json:"locationId" bson:"location_id" validate:"required"`
	LocationName           string     `json:"locationName" bson:"location_name" validate:"required"`
	CompanyLocationId      int64      `json:"companyLocationId" bson:"company_location_id" validate:"required"`
	CompanyLocationCode    string     `json:"companyLocationCode" bson:"company_location_code" validate:"required"`
	CompanyLocationAlias   string     `json:"companyLocationAlias" bson:"company_location_alias" validate:"required"`
	Parent                 int64      `json:"parent" bson:"parent" `
	ParentName             string     `json:"parentName" bson:"parent_name" `
	IdentityNo             string     `json:"identityNo" bson:"identity_no" `
	DrivingLicenseNo       string     `json:"drivingLicenseNo" bson:"driving_license_no" `
	LicenseType            string     `json:"licenseType" bson:"license_type" `
	NpwpNo                 string     `json:"npwpNo" bson:"npwp_no" `
	Name                   string     `json:"name" bson:"name" `
	PlaceOfBirth           string     `json:"placeOfBirth" bson:"place_of_birth" `
	DateOfBirth            *time.Time `json:"dateOfBirth" bson:"date_of_birth" `
	Email                  string     `json:"email" bson:"email" `
	Address                string     `json:"address" bson:"address" `
	TemporaryAddress       string     `json:"temporaryAddress" bson:"temporary_address" `
	NeighbourHoodWardNo    string     `json:"neighbourHoodWardNo" bson:"neighbour_hood_ward_no" `
	UrbanName              string     `json:"urbanName" bson:"urban_name" `
	SubDistrictName        string     `json:"subDistrictName" bson:"sub_district_name" `
	Religion               string     `json:"religion" bson:"religion" `
	MaritalStatus          string     `json:"maritalStatus" bson:"marital_status" `
	Citizen                string     `json:"citizen" bson:"citizen" `
	Gender                 string     `json:"gender" bson:"gender" `
	Ethnic                 string     `json:"ethnic" bson:"ethnic" `
	MobilePhoneNo          string     `json:"mobilePhoneNo" bson:"mobile_phone_no" `
	PhoneNo                string     `json:"phoneNo" bson:"phone_no" `
	ShirtSize              string     `json:"shirtSize" bson:"shirt_size" `
	PantSize               int64      `json:"pantSize" bson:"pant_size" `
	ShoeSize               int64      `json:"shoeSize" bson:"shoe_size" `
	JoinDate               *time.Time `json:"joinDate" bson:"join_date" `
	ResignDate             *time.Time `json:"resignDate" bson:"resign_date" `
	ResignReason           string     `json:"resignReason" bson:"resign_reason" `
	Bank                   string     `json:"bank" bson:"bank" `
	BankAccountNo          string     `json:"bankAccountNo" bson:"bank_account_no" `
	BankAccountName        string     `json:"bankAccountName" bson:"bank_account_name" `
	FamilyMobilePhoneNo    string     `json:"familyMobilePhoneNo" bson:"family_mobile_phone_no" `
	WorkStatus             string     `json:"workStatus" bson:"work_status" `
	ProfilePhoto           string     `json:"profilePhoto" bson:"profile_photo" `
	ContractStart          *time.Time `json:"contractStart" bson:"contract_start" `
	ContractEnd            *time.Time `json:"contractEnd" bson:"contract_end" `
	BpjsNo                 string     `json:"bpjsNo" bson:"bpjs_no" `
	JamsostekNo            string     `json:"jamsostekNo" bson:"jamsostek_no" `
	JamsostekType          string     `json:"jamsostekType" bson:"jamsostek_type" `
	JamsostekBalance       int64      `json:"jamsostekBalance" bson:"jamsostek_balance" `
	FamilyCardNo           string     `json:"familyCardNo" bson:"family_card_no" `
	ActiveStatus           bool       `json:"activeStatus" bson:"active_status" `
	Deleted                bool       `json:"deleted" bson:"deleted" `
	CreatedBy              string     `json:"createdBy" bson:"created_by" `
	CreatedTime            time.Time  `json:"createdTime" bson:"created_time" `
	ModifiedBy             *string    `json:"modifiedBy" bson:"modified_by" `
	ModifiedTime           *time.Time `json:"modifiedTime" bson:"modified_time" `
	DeletedBy              *string    `json:"deletedBy" bson:"deleted_by" `
	DeletedTime            *time.Time `json:"deletedTime" bson:"deleted_time" `
	Extension              string     `json:"extension" bson:"extension" `
	PostalCode             string     `json:"postalCode" bson:"postal_code" `
	WorkStatusChangeReason string     `json:"workStatusChangeReason" bson:"work_status_change_reason" `
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate" bson:"work_status_change_date" `
	Whatsapp               string     `json:"whatsapp" bson:"whatsapp"`
	Instagram              string     `json:"instagram" bson:"instagram"`
	Twitter                string     `json:"twitter" bson:"twitter"`
	Gmail                  string     `json:"gmail" bson:"gmail"`
	Facebook               string     `json:"facebook" bson:"facebook"`
}

type EmployeeUpdateDTO struct {
	ID                  int64      `json:"id" bson:"id"  `
	Parent              int64      `json:"parent" bson:"parent"  `
	IdentityNo          string     `json:"identityNo" bson:"identity_no"`
	DrivingLicenseNo    string     `json:"drivingLicenseNo" bson:"driving_license_no"`
	LicenseType         string     `json:"licenseType" bson:"license_type" `
	NpwpNo              string     `json:"npwpNo" bson:"npwp_no"`
	Name                string     `json:"name" bson:"name"`
	PlaceOfBirth        string     `json:"placeOfBirth" bson:"place_of_birth"`
	DateOfBirth         *time.Time `json:"dateOfBirth" bson:"date_of_birth"`
	Email               string     `json:"email" bson:"email"`
	Address             string     `json:"address" bson:"address"`
	TemporaryAddress    string     `json:"temporaryAddress" bson:"temporary_address"`
	NeighbourHoodWardNo string     `json:"neighbourHoodWardNo" bson:"neighbour_hood_ward_no"`
	UrbanName           string     `json:"urbanName" bson:"urban_name"`
	SubDistrictName     string     `json:"subDistrictName" bson:"sub_district_name"`
	Religion            string     `json:"religion" bson:"religion"`
	MaritalStatus       string     `json:"maritalStatus" bson:"marital_status"`
	Citizen             string     `json:"citizen" bson:"citizen"`
	Gender              string     `json:"gender" bson:"gender"`
	Ethnic              string     `json:"ethnic" bson:"ethnic"`
	MobilePhoneNo       string     `json:"mobilePhoneNo" bson:"mobile_phone_no"`
	PhoneNo             string     `json:"phoneNo" bson:"phone_no"`
	ShirtSize           string     `json:"shirtSize" bson:"shirt_size"`
	PantSize            int64      `json:"pantSize" bson:"pant_size"`
	ShoeSize            int64      `json:"shoeSize" bson:"shoe_size"`
	Bank                string     `json:"bank" bson:"bank"`
	BankAccountNo       string     `json:"bankAccountNo" bson:"bank_account_no"`
	BankAccountName     string     `json:"bankAccountName" bson:"bank_account_name"`
	FamilyMobilePhoneNo string     `json:"familyMobilePhoneNo" bson:"family_mobile_phone_no"`
	BpjsNo              string     `json:"bpjsNo" bson:"bpjs_no"`
	JamsostekNo         string     `json:"jamsostekNo" bson:"jamsostek_no"`
	JamsostekType       string     `json:"jamsostekType" bson:"jamsostek_type"`
	JamsostekBalance    int64      `json:"jamsostekBalance" bson:"jamsostek_balance"`
	FamilyCardNo        string     `json:"familyCardNo" bson:"family_card_no"`
	Extension           string     `json:"extension" bson:"extension" `
	ModifiedBy          *string    `json:"modifiedBy" bson:"modified_by"`
	ModifiedTime        *time.Time `json:"modifiedTime" bson:"modified_time"`
	Whatsapp            string     `json:"whatsapp" bson:"whatsapp"`
	Instagram           string     `json:"instagram" bson:"instagram"`
	Twitter             string     `json:"twitter" bson:"twitter"`
	Gmail               string     `json:"gmail" bson:"gmail"`
	Facebook            string     `json:"facebook" bson:"facebook"`
}

type SubsEmployee struct {
	Key               string                 `json:"key"`
	StateStore        string                 `json:"stateStore"`
	NotificationData  NotificationInfoDTO    `json:"notificationData"`
	Data              EmployeeDataDTO        `json:"data"`
	JWTDocodedPayload map[string]interface{} `json:"jwtDocodedPayload"`
}
type SubsEmployeeRollback struct {
	Key               string                 `json:"key"`
	StateStore        string                 `json:"stateStore"`
	NotificationData  NotificationInfoDTO    `json:"notificationData"`
	Data              EmployeeUpdateDTO      `json:"data"`
	JWTDocodedPayload map[string]interface{} `json:"jwtDocodedPayload"`
}

type SubsEmployeeUpdateWorkStatus struct {
	Key               string                   `json:"key"`
	StateStore        string                   `json:"stateStore"`
	NotificationData  NotificationInfoDTO      `json:"notificationData"`
	Data              UpdateWorkStatusStateDTO `json:"data"`
	JWTDocodedPayload map[string]interface{}   `json:"jwtDocodedPayload"`
}

type SaveWorkStatusDTO struct {
	WorkStatus    string     `json:"workStatus" validate:"required"`
	EmployeeId    int64      `json:"employeeId" validate:"required"`
	ContractStart *time.Time `json:"contractStart" validate:"required"`
	ContractEnd   *time.Time `json:"contractEnd" validate:"required"`
	Deleted       bool       `json:"deleted" bson:"deleted" `
	CreatedBy     string     `json:"createdBy" bson:"created_by" `
	CreatedTime   time.Time  `json:"createdTime" bson:"created_time" `
}

type WorkStatusUpdateDTO struct {
	ID                     int64      `json:"id" validate:"required"`
	WorkStatus             string     `json:"workStatus" validate:"required"`
	ContractStart          *time.Time `json:"contractStart" `
	ContractEnd            *time.Time `json:"contractEnd" `
	WorkStatusChangeReason string     `json:"workStatusChangeReason" validate:"required"`
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate" validate:"required"`
	CreatedBy              string     `json:"createdBy" bson:"created_by" `
	CreatedTime            time.Time  `json:"createdTime" bson:"created_time" `
	ModifiedBy             *string    `json:"modifiedBy" bson:"modified_by" `
	ModifiedTime           *time.Time `json:"modifiedTime" bson:"modified_time" `
}

type ResignStatusUpdateDTO struct {
	ID           int64      `json:"id" validate:"required"`
	WorkStatus   string     `json:"workStatus" validate:"required"`
	JoinDate     *time.Time `json:"joinDate"`
	ResignDate   *time.Time `json:"resignDate"`
	ResignReason string     `json:"resignReason"`
}

type UpdateResignStatusStateDTO struct {
	ID           int64      `json:"id" validate:"required"`
	WorkStatus   string     `json:"workStatus" validate:"required"`
	JoinDate     *time.Time `json:"joinDate"`
	ResignDate   *time.Time `json:"resignDate"`
	ResignReason string     `json:"resignReason"`
	WorkStatusId int64      `json:"workStatusId"`
}

type FingerprintUpdateDTO struct {
	ID            int64      `json:"id" bson:"id"`
	FingerPrintId string     `json:"fingerPrintId" validate:"required"`
	FaceId        string     `json:"faceId"`
	MachineId     int64      `json:"machineId"`
	ModifiedBy    *string    `json:"modifiedBy" bson:"modified_by" `
	ModifiedTime  *time.Time `json:"modifiedTime" bson:"modified_time" `
}

type MachineIdUpdateDTO struct {
	ID        int64 `json:"id" bson:"id"`
	MachineId int64 `json:"machineId"`
}

type FaceIdUpdateDTO struct {
	ID        int64  `json:"id" bson:"id"`
	FaceId    string `json:"faceId"`
	MachineId int64  `json:"machineId"`
}

type UpdateWorkStatusStateDTO struct {
	EmployeeId             int64      `json:"employeeId"`
	WorkStatus             string     `json:"workStatus"`
	ContractStart          *time.Time `json:"contractStart"`
	ContractEnd            *time.Time `json:"contractEnd"`
	WorkStatusChangeReason string     `json:"workStatusChangeReason"`
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate"`
	WorkStatusId           int64      `json:"workStatusId"`
}

type SubscribeEmployeePsgl struct {
	EmployeeId             int64  `json:"employeeId"`
	NewEmployeeCode        string `json:"newEmployeeCode"`
	NewIdCompany           int64  `json:"newIdCompany"`
	NewCompanyName         string `json:"newCompanyName"`
	NewIdLocation          int64  `json:"newIdLocation"`
	NewLocationName        string `json:"newLocationName"`
	NewIdDepartment        int64  `json:"newIdDepartment"`
	NewDepartmentName      string `json:"newDepartmentName"`
	NewIdSection           int64  `json:"newIdSection"`
	NewSectionName         string `json:"newSectionName"`
	NewIdPosition          int64  `json:"newIdPosition"`
	NewPositionName        string `json:"newPositionName"`
	NewCompanyLocationCode string `json:"newCompanyLocationCode"`
}

type SubscribeCompanyLocationCodeUpdateDTO struct {
	CompanyLocationCode    string `json:"companyLocationCode" bson:"company_location_code" `
	OldCompanyLocationCode string `json:"oldCompanyLocationCode" bson:"old_company_location_code" `
}

type SubscribeCompanyUpdateDTO struct {
	CompanyId      int64  `json:"id"`
	OldCompanyName string `json:"oldName"`
	CompanyName    string `json:"name"`
}

type SubscribeDepartmentUpdateDTO struct {
	DepartmentId      int64  `json:"id"`
	OldDepartmentName string `json:"oldName"`
	DepartmentName    string `json:"name"`
}

type SubscribeLocationUpdateDTO struct {
	LocationId      int64  `json:"id"`
	OldLocationName string `json:"oldName"`
	LocationName    string `json:"name"`
}

type SubscribeSectionUpdateDTO struct {
	SectionId      int64  `json:"id"`
	OldSectionName string `json:"oldName"`
	SectionName    string `json:"name"`
}

type SubscribePositionUpdateDTO struct {
	PositionId     int64  `json:"id"`
	OldPostionName string `json:"oldName"`
	PositionName   string `json:"name"`
}
