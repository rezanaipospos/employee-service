package employee

import "time"

type NewEmployeeRepository struct{}
type NewSubscriberEmployeeRepository struct{}

type EmployeesData struct {
	RecordsTotal int64          `json:"recordsTotal"`
	HasReachMax  bool           `json:"hasReachMax"`
	Data         []GetEmployees `json:"data"`
}

type Employees struct {
	ID                     int64      `json:"id" field:"required"`
	Code                   string     `json:"code" bson:"code"`
	Type                   string     `json:"type" bson:"type"`
	FingerPrintId          string     `json:"fingerPrintId" bson:"finger_print_id"`
	FaceId                 string     `json:"faceId" bson:"face_id"`
	MachineId              int64      `json:"machineId" bson:"machine_id"`
	MachineName            string     `json:"machineName" bson:"machine_name"`
	DepartmentId           int64      `json:"departmentId" bson:"department_id" `
	DepartmentName         string     `json:"departmentName" bson:"department_name" `
	SectionId              int64      `json:"sectionId" bson:"section_id" `
	SectionName            string     `json:"sectionName" bson:"section_name" `
	PositionId             int64      `json:"positionId" bson:"position_id" `
	PositionName           string     `json:"positionName" bson:"position_name" `
	CompanyId              int64      `json:"companyId" bson:"company_id" `
	CompanyName            string     `json:"companyName" bson:"company_name" `
	LocationId             int64      `json:"locationId" bson:"location_id" `
	LocationName           string     `json:"locationName" bson:"location_name" `
	CompanyLocationId      int64      `json:"companyLocationId" bson:"company_location_id" `
	CompanyLocationCode    string     `json:"companyLocationCode" bson:"company_location_code" `
	CompanyLocationAlias   string     `json:"companyLocationAlias" bson:"company_location_alias" `
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
	JoinDate               time.Time  `json:"joinDate" bson:"join_date" `
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
	PostalCode             string     `json:"postalCode" bson:"postal_code"`
	WorkStatusChangeReason string     `json:"workStatusChangeReason" bson:"work_status_change_reason"`
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate" bson:"work_status_change_date"`
	Whatsapp               string     `json:"whatsapp" bson:"whatsapp"`
	Instagram              string     `json:"instagram" bson:"instagram"`
	Twitter                string     `json:"twitter" bson:"twitter"`
	Gmail                  string     `json:"gmail" bson:"gmail"`
	Facebook               string     `json:"facebook" bson:"facebook"`
}

type EmployeesUpdate struct {
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

type GetEmployees struct {
	ID                  int64  `json:"id" bson:"id" field:"required"`
	Code                string `json:"code" bson:"code"`
	Name                string `json:"name" bson:"name"`
	CompanyName         string `json:"companyName" bson:"company_name"`
	MachineId           int64  `json:"machineId" bson:"machine_id"`
	MachineName         string `json:"machineName" bson:"machine_name"`
	CompanyLocationCode string `json:"companyLocationCode" bson:"company_location_code"`
	LocationName        string `json:"locationName" bson:"location_name" `
	DepartmentName      string `json:"departmentName" bson:"department_name"`
	PositionName        string `json:"positionName" bson:"position_name" `
	FingerPrintId       string `json:"fingerPrintId" bson:"finger_print_id"`
	FaceId              string `json:"faceId" bson:"face_id"`
	ParentName          string `json:"parentName" bson:"parent_name"`
	WorkStatus          string `json:"workStatus" bson:"work_status"`
	ProfilePhoto        string `json:"profilePhoto" bson:"profile_photo"`
}

type SearchEmployees struct {
	ID                   int64  `json:"id" bson:"id" field:"required"`
	Code                 string `json:"code" bson:"code"`
	Name                 string `json:"name" bson:"name"`
	CompanyId            int64  `json:"companyId" bson:"company_id"`
	CompanyName          string `json:"companyName" bson:"company_name"`
	CompanyLocationId    int64  `json:"companyLocationId" bson:"company_location_id"`
	CompanyLocationCode  string `json:"companyLocationCode" bson:"company_location_code"`
	CompanyLocationAlias string `json:"companyLocationAlias" bson:"company_location_alias"`
	DepartmentId         int64  `json:"departmentId" bson:"department_id"`
	DepartmentName       string `json:"departmentName" bson:"department_name"`
	FingerPrintId        string `json:"fingerPrintId" bson:"finger_print_id"`
	FaceId               string `json:"faceId" bson:"face_id"`
	ParentName           string `json:"parentName" bson:"parent_name"`
	WorkStatus           string `json:"workStatus" bson:"work_status"`
	ProfilePhoto         string `json:"profilePhoto" bson:"profile_photo"`
}

type Subordinates struct {
	ID           int64  `json:"id"  `
	Parent       int64  `json:"parent"  `
	Code         string `json:"code"`
	Name         string `json:"name"`
	PhoneNo      string `json:"phoneNo"`
	Email        string `json:"email"`
	Religion     string `json:"religion"`
	Gender       string `json:"gender"`
	ProfilePhoto string `json:"profilePhoto"`
	WorkStatus   string `json:"workStatus"`
}

type DetailSubordinates struct {
	ID                  int64      `json:"id" `
	Code                string     `json:"code"`
	Type                string     `json:"type"`
	FingerPrintId       string     `json:"fingerPrintId"`
	FaceId              string     `json:"faceId"`
	MachineId           int64      `json:"machineId"`
	DepartmentId        int64      `json:"departmentId"`
	SectionId           int64      `json:"sectionId"`
	PositionId          int64      `json:"positionId"`
	CompanyId           int64      `json:"companyId"`
	LocationId          int64      `json:"locationId"`
	CompanyLocationCode string     `json:"companyLocationCode"`
	Parent              int64      `json:"parent"`
	IdentityNo          string     `json:"identityNo"`
	DrivingLicenseNo    string     `json:"drivingLicenseNo"`
	LicenseType         string     `json:"licenseType"`
	NpwpNo              string     `json:"npwpNo"`
	Name                string     `json:"name"`
	PlaceOfBirth        string     `json:"placeOfBirth"`
	DateOfBirth         *time.Time `json:"dateOfBirth"`
	Email               string     `json:"email"`
	Address             string     `json:"address"`
	TemporaryAddress    string     `json:"temporaryAddress"`
	NeighbourHoodWardNo string     `json:"neighbourHoodWardNo"`
	UrbanName           string     `json:"urbanName"`
	SubDistrictName     string     `json:"subDistrictName"`
	Religion            string     `json:"religion"`
	MaritalStatus       string     `json:"maritalStatus"`
	Citizen             string     `json:"citizen"`
	Gender              string     `json:"gender"`
	Ethnic              string     `json:"ethnic"`
	MobilePhoneNo       string     `json:"mobilePhoneNo"`
	PhoneNo             string     `json:"phoneNo"`
	ShirtSize           string     `json:"shirtSize"`
	PantSize            int64      `json:"pantSize"`
	ShoeSize            int64      `json:"shoeSize"`
	JoinDate            *time.Time `json:"joinDate"`
	ResignDate          *time.Time `json:"resignDate"`
	ResignReason        string     `json:"resignReason"`
	Bank                string     `json:"bank"`
	BankAccountNo       string     `json:"bankAccountNo"`
	BankAccountName     string     `json:"bankAccountName"`
	FamilyMobilePhoneNo string     `json:"familyMobilePhoneNo"`
	WorkStatus          string     `json:"workStatus"`
	ProfilePhoto        string     `json:"profilePhoto"`
	ContractStart       *time.Time `json:"contractStart"`
	ContractEnd         *time.Time `json:"contractEnd"`
	BpjsNo              string     `json:"bpjsNo"`
	JamsostekNo         string     `json:"jamsostekNo"`
	JamsostekType       string     `json:"jamsostekType"`
	JamsostekBalance    int64      `json:"jamsostekBalance"`
	FamilyCardNo        string     `json:"familyCardNo"`
	ActiveStatus        bool       `json:"activeStatus"`
	CreatedBy           string     `json:"createdBy"`
	CreatedTime         time.Time  `json:"createdTime"`
	ModifiedBy          *string    `json:"modifiedBy"`
	ModifiedTime        *time.Time `json:"modifiedTime"`
	Extension           string     `json:"extension"`
	PostalCode          string     `json:"postalCode"`
}

type GetEmployeeWorkStatus struct {
	ID                     int64      `json:"id"`
	WorkStatus             string     `json:"workStatus"`
	ContractStart          *time.Time `json:"contractStart"`
	ContractEnd            *time.Time `json:"contractEnd"`
	WorkStatusChangeReason string     `json:"workStatusChangeReason"`
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate"`
}

type WorkStatusUpdate struct {
	ID                     int64      `json:"id"`
	WorkStatus             string     `json:"workStatus"`
	ContractStart          *time.Time `json:"contractStart"`
	ContractEnd            *time.Time `json:"contractEnd"`
	WorkStatusChangeReason string     `json:"workStatusChangeReason"`
	WorkStatusChangeDate   *time.Time `json:"workStatusChangeDate"`
}

type ResignStatusUpdate struct {
	ID           int64      `json:"id"`
	WorkStatus   string     `json:"workStatus"`
	JoinDate     *time.Time `json:"joinDate"`
	ResignDate   *time.Time `json:"resignDate"`
	ResignReason string     `json:"resignReason"`
}

type FingerprintUpdate struct {
	ID            int64  `json:"id"`
	FingerPrintId string `json:"fingerPrintId"`
	FaceId        string `json:"faceId"`
	MachineId     int64  `json:"machineId"`
}

type MachineIdUpdate struct {
	ID        int64 `json:"id"`
	MachineId int64 `json:"machineId"`
}

type FaceIdUpdate struct {
	ID        int64  `json:"id"`
	FaceId    string `json:"faceId"`
	MachineId int64  `json:"machineId"`
}

type GetHead struct {
	ID           int64  `json:"id" bson:"id" field:"required"`
	Code         string `json:"code"  bson:"code"`
	Name         string `json:"name" bson:"name" `
	Parent       int64  `json:"parent" bson:"parent" `
	PositionName string `json:"positionName" bson:"position_name"`
}

type CheckExistData struct {
	ID            int64  `json:"id"`
	Code          string `json:"code" `
	Name          string `json:"name" `
	MobilePhoneNo string `json:"mobilePhoneNo"`
	Email         string `json:"email"`
	ActiveStatus  bool   `json:"activeStatus" `
}

type CheckExistEmployeeID struct {
	ID            int64  `json:"id" bson:"id"`
	Code          string `json:"code" bson:"code"`
	Name          string `json:"name" bson:"name"`
	MobilePhoneNo string `json:"mobilePhoneNo" bson:"mobile_phone_no"`
	Email         string `json:"email" bson:"email"`
	ActiveStatus  bool   `json:"activeStatus" bson:"active_status"`
	MachineId     int64  `json:"machineId" bson:"machine_id"`
	MachineName   string `json:"machineName" bson:"machine_name"`
	ProfilePhoto  string `json:"profilePhoto"`
}

var allowedFields = map[string]string{
	"id":                  "id",
	"code":                "code",
	"name":                "name",
	"companyName":         "company_name",
	"machineName":         "machine_name",
	"companyLocationCode": "company_location_code",
	"locationName":        "location_name",
	"departmentName":      "department_name",
	"positionName":        "position_name",
	"fingerPrintId":       "finger_print_id",
	"faceId":              "face_id",
	"parentName":          "parent_name",
	"workStatus":          "work_status",
	"profilePhoto":        "profile_photo",
}
