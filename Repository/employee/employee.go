package employee

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	validator "EmployeeService/Library/Helper/Validator"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultEmployeeSortColumn = `code`
	defaultEmployeeSortOrder  = `desc`
	defaultEmployeePageNumber = 0
	defaultEmployeePageSize   = 10

	queryInsertEmployee = `INSERT INTO public.employees 
								(code,
								type,
								finger_print_id,
								face_id,
								machine_id,
								department_id,
								section_id,
								position_id,
								company_id,
								location_id,
								company_location_id,
								company_location_code,
								parent,
								identity_no,
								driving_license_no,
								npwp_no,
								name,
								place_of_birth,
								date_of_birth,
								email,
								address,
								temporary_address,
								neighbour_hood_ward_no,
								urban_name,
								sub_district_name,
								religion,
								marital_status,
								citizen,
								gender,
								ethnic,
								mobile_phone_no,
								phone_no,
								shirt_size,
								pant_size,
								shoe_size,
								join_date,
								resign_date,
								resign_reason,
								bank,
								bank_account_no,
								bank_account_name,
								family_mobile_phone_no,
								work_status,
								profile_photo,
								contract_start,
								contract_end,
								bpjs_no,
								jamsostek_no,
								jamsostek_type,
								jamsostek_balance,
								family_card_no,
								active_status,
								deleted,
								created_by,
								created_time,
								extension,
								license_type,
								postal_code,
								whatsapp,
								instagram,
								twitter,
								gmail,
								facebook) 
							VALUES 
								($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63)  
							RETURNING id,code,type,finger_print_id,face_id,machine_id,department_id,section_id,position_id,company_id,location_id,company_location_id,company_location_code,parent,identity_no,driving_license_no,npwp_no,name,place_of_birth,date_of_birth,email,address,temporary_address,neighbour_hood_ward_no,urban_name,sub_district_name,religion,marital_status,citizen,gender,ethnic,mobile_phone_no,phone_no,shirt_size,pant_size,shoe_size,join_date,resign_date,resign_reason,bank,bank_account_no,bank_account_name,family_mobile_phone_no,work_status,profile_photo,contract_start,contract_end,bpjs_no,jamsostek_no,jamsostek_type,jamsostek_balance,family_card_no,active_status,deleted,created_by,created_time,modified_by,modified_time,deleted_by,deleted_time,extension,license_type,postal_code,whatsapp,instagram,twitter,gmail,facebook`

	queryUpdateEmployee = `UPDATE public.employees SET 
								identity_no=$1,
								driving_license_no=$2,
								npwp_no=$3,
								name=$4,
								place_of_birth=$5,
								date_of_birth=$6,
								email=$7,
								address=$8,
								temporary_address=$9,
								neighbour_hood_ward_no=$10,
								urban_name=$11,
								sub_district_name=$12,
								religion=$13,
								marital_status=$14,
								citizen=$15,
								gender=$16,
								ethnic=$17,
								mobile_phone_no=$18,
								phone_no=$19,
								shirt_size=$20,
								pant_size=$21,
								shoe_size=$22,
								bank=$23,
								bank_account_no=$24,
								bank_account_name=$25,
								family_mobile_phone_no=$26,
								bpjs_no=$27,
								jamsostek_no=$28,
								jamsostek_type=$29,
								jamsostek_balance=$30,
								family_card_no=$31,
								modified_by=$32, 
								modified_time=$33,
								parent=$34,
								extension=$35,
								license_type=$36,
								whatsapp=$37,
								instagram=$38,
								twitter=$39,
								gmail=$40,
								facebook=$41
							WHERE id=$42
							RETURNING id,parent,identity_no,driving_license_no,license_type,npwp_no,name,place_of_birth,date_of_birth,email,address,temporary_address,neighbour_hood_ward_no,urban_name,sub_district_name,religion,marital_status,citizen,gender,ethnic,mobile_phone_no,phone_no,shirt_size,pant_size,shoe_size,bank,bank_account_no,bank_account_name,family_mobile_phone_no,bpjs_no,jamsostek_no,jamsostek_type,jamsostek_balance,family_card_no,extension,whatsapp,instagram,twitter,gmail,facebook,modified_by,modified_time`

	querySoftDeleteEmployee = `UPDATE public.employees SET deleted=TRUE,deleted_by=$1,deleted_time=$2 WHERE id=$3`

	queryBrowseSubordinates = `WITH RECURSIVE subordinates AS (
	SELECT id, parent, code, name, phone_no, email, religion, gender, profile_photo, work_status
		FROM employees  WHERE deleted=FALSE AND parent = %d
	UNION SELECT e.id, e.parent, e.code, e.name, e.phone_no, e.email, e.religion, e.gender, e.profile_photo, e.work_status
		FROM employees e INNER JOIN subordinates s ON s.id = e.parent WHERE deleted=FALSE ) 
		SELECT * FROM subordinates WHERE code ILIKE '%%%s%%' AND name ILIKE '%%%s%%' AND work_status ILIKE '%%%s%%'`

	queryDetailSubordinates = `WITH RECURSIVE subordinates AS ( SELECT id, code, type, finger_print_id, face_id, machine_id, department_id, section_id, position_id, company_id, location_id, company_location_code, parent, identity_no, driving_license_no, npwp_no, name, place_of_birth, date_of_birth, email, address, temporary_address, neighbour_hood_ward_no, urban_name, sub_district_name, religion, marital_status, citizen, gender, ethnic, mobile_phone_no, phone_no, shirt_size, pant_size, shoe_size, join_date, resign_date, resign_reason, bank, bank_account_no, bank_account_name, family_mobile_phone_no, work_status, profile_photo, contract_start, contract_end, bpjs_no, jamsostek_no, jamsostek_type, jamsostek_balance, family_card_no, active_status, created_by, created_time, modified_by, modified_time ,extension,postal_code
	FROM employees WHERE deleted=FALSE AND parent = %d
	UNION  
	SELECT e.id, e.code, e.type, e.finger_print_id, e.face_id, e.machine_id, e.department_id, e.section_id, e.position_id, e.company_id, e.location_id, e.company_location_code, e.parent, e.identity_no, e.driving_license_no, e.npwp_no, e.name, e.place_of_birth, e.date_of_birth, e.email, e.address, e.temporary_address, e.neighbour_hood_ward_no, e.urban_name, e.sub_district_name, e.religion, e.marital_status, e.citizen, e.gender, e.ethnic, e.mobile_phone_no, e.phone_no, e.shirt_size, e.pant_size, e.shoe_size, e.join_date, e.resign_date, e.resign_reason, e.bank, e.bank_account_no, e.bank_account_name, e.family_mobile_phone_no, e.work_status, e.profile_photo, e.contract_start, e.contract_end, e.bpjs_no, e.jamsostek_no, e.jamsostek_type, e.jamsostek_balance, e.family_card_no, e.active_status, e.created_by, e.created_time, e.modified_by, e.modified_time ,e.extension,e.postal_code FROM employees e INNER JOIN subordinates s ON s.id = e.parent where deleted=FALSE)  
	SELECT*FROM subordinates where id = %d`

	queryGetEmployeeWorkStatus    = `SELECT id, work_status, contract_start, contract_end, work_status_change_reason, work_status_change_date FROM public.employees WHERE id = $1`
	queryUpdateEmployeeWorkStatus = `UPDATE public.employees SET 
										work_status = $1, 
										contract_start = $2,
										contract_end = $3,
										work_status_change_reason = $4,
										work_status_change_date = $5
									WHERE id=$6 and deleted=false RETURNING id,work_status,contract_start,contract_end,work_status_change_reason,work_status_change_date`

	queryUpdateEmployeeFingerprint = `UPDATE public.employees SET 
												finger_print_id = $1, 
												face_id = $2,
												machine_id = $3,
												modified_by=$4, 
												modified_time=$5 
											WHERE id=$6 and deleted=false RETURNING id,finger_print_id,face_id,machine_id`

	queryUpdateEmployeeMachineId = `UPDATE public.employees SET machine_id = $1 WHERE id=$2 and deleted=false RETURNING id,machine_id`

	queryUpdateEmployeeFaceId = `UPDATE public.employees SET face_id = $1,machine_id=$2  WHERE id=$3 and deleted=false RETURNING id,face_id,machine_id`

	queryUpdateEmployeeResignStatus = `UPDATE public.employees SET 
											work_status = $1, 
											resign_date = $2,
											resign_reason = $3
										WHERE id=$4 and deleted=false RETURNING id,work_status,join_date,resign_date,resign_reason`

	queryValidationEmployee = `SELECT EXISTS(SELECT * FROM public.employees WHERE code = $1 AND deleted=FALSE LIMIT 1 )`

	queryInsertWorkStatus     = `INSERT INTO public.work_status (work_status,employee_id,start_date,end_date,deleted,created_by,created_time) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`
	queryHardDeleteWorkStatus = `DELETE FROM public.work_status WHERE id = $1`

	queryValidateMobilePhone = `Select Exists(Select * From employees Where mobile_phone_no=$1 And deleted=False)`

	queryValidateEmail = `Select Exists(Select * From employees Where email=$1 And deleted=False)`

	queryCheckExistMobilePhoneNumber = `Select id, code, name, mobile_phone_no, email, active_status From employees Where mobile_phone_no=$1 And deleted=False limit 1`
	queryCheckExistEmail             = `Select id, code, name, mobile_phone_no, email, active_status From employees Where email=$1 And deleted=False limit 1`

	queryCountEmployees = `SELECT COUNT(*) FROM public.employees WHERE deleted = FALSE`
)

func (c NewEmployeeRepository) setDefaultParams(params Dto.EmployeeBrowseParams) Dto.EmployeeBrowseParams {
	if params.SortColumn == "" {
		params.SortColumn = defaultEmployeeSortColumn
	}
	if params.SortOrder == "" {
		params.SortOrder = defaultEmployeeSortOrder
	}
	if params.PageNumber == 0 {
		params.PageNumber = defaultEmployeePageNumber
	}
	if params.PageSize == 0 {
		params.PageSize = defaultEmployeePageSize
	}
	return params
}

// Mongo
func (c NewEmployeeRepository) BrowseEmployee(params Dto.EmployeeBrowseParams) (result EmployeesData, err error) {
	var order int64
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	params = c.setDefaultParams(params)
	params.SortOrder = validator.ValidateSortOrder(params.SortOrder, defaultEmployeeSortOrder)
	params.SortColumn = validator.ValidateSortColumn(allowedFields, params.SortColumn, defaultEmployeeSortColumn)

	switch {
	case strings.EqualFold(params.SortOrder, "DESC"):
		order = -1 //DESC
	case strings.EqualFold(params.SortOrder, "ASC"):
		order = 1 //ASC
	default:
		order = -1 //DESC
	}

	filter := bson.M{
		"code": bson.M{
			"$regex":   params.Code,
			"$options": "i",
		},
		"name": bson.M{
			"$regex":   params.Name,
			"$options": "i",
		},
		"company_name": bson.M{
			"$regex":   params.CompanyName,
			"$options": "i",
		},
		"company_location_code": bson.M{
			"$regex":   params.CompanyLocationCode,
			"$options": "i",
		},
		"department_name": bson.M{
			"$regex":   params.DepartmentName,
			"$options": "i",
		},
		"work_status": bson.M{
			"$regex":   params.WorkStatus,
			"$options": "i",
		},
	}

	if params.MachineId != nil {
		filter["machine_id"] = params.MachineId
	}

	projection := bson.M{
		"id":                    1,
		"code":                  1,
		"name":                  1,
		"company_name":          1,
		"machine_id":            1,
		"machine_name":          1,
		"company_location_code": 1,
		"location_name":         1,
		"department_name":       1,
		"position_name":         1,
		"finger_print_id":       1,
		"face_id":               1,
		"parent_name":           1,
		"work_status":           1,
		"profile_photo":         1,
	}

	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.D{{Key: params.SortColumn, Value: order}}).
		SetSkip(params.PageNumber * params.PageSize).
		SetLimit(params.PageSize)

	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err = cursor.All(context.TODO(), &result.Data); err != nil {
		return
	}
	defer cursor.Close(context.TODO())

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return
	}

	result.RecordsTotal = count

	if count == 0 || params.PageNumber >= count/params.PageSize {
		result.HasReachMax = true
	} else {
		result.HasReachMax = false
	}
	return
}

// Mongo
func (c NewEmployeeRepository) SearchEmployee(params Dto.EmployeeSearchDTO) (result []SearchEmployees, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{
		"name": bson.M{
			"$regex":   params.Search,
			"$options": "i",
		},
	}

	projection := bson.M{
		"id":                     1,
		"code":                   1,
		"name":                   1,
		"company_id":             1,
		"company_name":           1,
		"company_location_id":    1,
		"company_location_code":  1,
		"company_location_alias": 1,
		"department_id":          1,
		"department_name":        1,
		"finger_print_id":        1,
		"face_id":                1,
		"parent_name":            1,
		"work_status":            1,
		"profile_photo":          1,
	}

	opts := options.Find().
		SetProjection(projection).
		SetSort(bson.D{{Key: "name", Value: 1}})

	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err = cursor.All(context.TODO(), &result); err != nil {
		return
	}
	defer cursor.Close(context.TODO())
	return
}

// Mongo
func (r NewEmployeeRepository) GetHead(employeeId int64) (result GetHead, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{
		"id": employeeId,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return
	}
	return
}

// Mongo
func (c NewEmployeeRepository) BrowseEmployeeDetail(params Dto.EmployeeDataDTO) (result Employees, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{
		"id": params.ID,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return
	}

	return
}

// Mongo
func (c NewEmployeeRepository) CheckExistEmployeeID(employeeId int64) (result CheckExistEmployeeID, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{
		"id": employeeId,
	}

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return
	}

	return
}

func (r NewEmployeeRepository) SaveEmployee(params Dto.EmployeeDataDTO) (tx *sql.Tx, employee Employees, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryInsertEmployee,
		params.Code,
		params.Type,
		params.FingerPrintId,
		params.FaceId,
		params.MachineId,
		params.DepartmentId,
		params.SectionId,
		params.PositionId,
		params.CompanyId,
		params.LocationId,
		params.CompanyLocationId,
		params.CompanyLocationCode,
		params.Parent,
		params.IdentityNo,
		params.DrivingLicenseNo,
		params.NpwpNo,
		params.Name,
		params.PlaceOfBirth,
		params.DateOfBirth,
		params.Email,
		params.Address,
		params.TemporaryAddress,
		params.NeighbourHoodWardNo,
		params.UrbanName,
		params.SubDistrictName,
		params.Religion,
		params.MaritalStatus,
		params.Citizen,
		params.Gender,
		params.Ethnic,
		params.MobilePhoneNo,
		params.PhoneNo,
		params.ShirtSize,
		params.PantSize,
		params.ShoeSize,
		params.JoinDate,
		params.ResignDate,
		params.ResignReason,
		params.Bank,
		params.BankAccountNo,
		params.BankAccountName,
		params.FamilyMobilePhoneNo,
		params.WorkStatus,
		params.ProfilePhoto,
		params.ContractStart,
		params.ContractEnd,
		params.BpjsNo,
		params.JamsostekNo,
		params.JamsostekType,
		params.JamsostekBalance,
		params.FamilyCardNo,
		params.ActiveStatus,
		params.Deleted,
		params.CreatedBy,
		params.CreatedTime,
		params.Extension,
		params.LicenseType,
		params.PostalCode,
		params.Whatsapp,
		params.Instagram,
		params.Twitter,
		params.Gmail,
		params.Facebook)

	if err = row.Scan(
		&employee.ID,
		&employee.Code,
		&employee.Type,
		&employee.FingerPrintId,
		&employee.FaceId,
		&employee.MachineId,
		&employee.DepartmentId,
		&employee.SectionId,
		&employee.PositionId,
		&employee.CompanyId,
		&employee.LocationId,
		&employee.CompanyLocationId,
		&employee.CompanyLocationCode,
		&employee.Parent,
		&employee.IdentityNo,
		&employee.DrivingLicenseNo,
		&employee.NpwpNo,
		&employee.Name,
		&employee.PlaceOfBirth,
		&employee.DateOfBirth,
		&employee.Email,
		&employee.Address,
		&employee.TemporaryAddress,
		&employee.NeighbourHoodWardNo,
		&employee.UrbanName,
		&employee.SubDistrictName,
		&employee.Religion,
		&employee.MaritalStatus,
		&employee.Citizen,
		&employee.Gender,
		&employee.Ethnic,
		&employee.MobilePhoneNo,
		&employee.PhoneNo,
		&employee.ShirtSize,
		&employee.PantSize,
		&employee.ShoeSize,
		&employee.JoinDate,
		&employee.ResignDate,
		&employee.ResignReason,
		&employee.Bank,
		&employee.BankAccountNo,
		&employee.BankAccountName,
		&employee.FamilyMobilePhoneNo,
		&employee.WorkStatus,
		&employee.ProfilePhoto,
		&employee.ContractStart,
		&employee.ContractEnd,
		&employee.BpjsNo,
		&employee.JamsostekNo,
		&employee.JamsostekType,
		&employee.JamsostekBalance,
		&employee.FamilyCardNo,
		&employee.ActiveStatus,
		&employee.Deleted,
		&employee.CreatedBy,
		&employee.CreatedTime,
		&employee.ModifiedBy,
		&employee.ModifiedTime,
		&employee.DeletedBy,
		&employee.DeletedTime,
		&employee.Extension,
		&employee.LicenseType,
		&employee.PostalCode,
		&employee.Whatsapp,
		&employee.Instagram,
		&employee.Twitter,
		&employee.Gmail,
		&employee.Facebook); err != nil {
		tx.Rollback()
		return
	}
	return
}

func (r NewEmployeeRepository) UpdateEmployee(params Dto.EmployeeUpdateDTO) (tx *sql.Tx, employee EmployeesUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}

	row := tx.QueryRow(
		queryUpdateEmployee,
		params.IdentityNo,
		params.DrivingLicenseNo,
		params.NpwpNo,
		params.Name,
		params.PlaceOfBirth,
		params.DateOfBirth,
		params.Email,
		params.Address,
		params.TemporaryAddress,
		params.NeighbourHoodWardNo,
		params.UrbanName,
		params.SubDistrictName,
		params.Religion,
		params.MaritalStatus,
		params.Citizen,
		params.Gender,
		params.Ethnic,
		params.MobilePhoneNo,
		params.PhoneNo,
		params.ShirtSize,
		params.PantSize,
		params.ShoeSize,
		params.Bank,
		params.BankAccountNo,
		params.BankAccountName,
		params.FamilyMobilePhoneNo,
		params.BpjsNo,
		params.JamsostekNo,
		params.JamsostekType,
		params.JamsostekBalance,
		params.FamilyCardNo,
		params.ModifiedBy,
		params.ModifiedTime,
		params.Parent,
		params.Extension,
		params.LicenseType,
		params.Whatsapp,
		params.Instagram,
		params.Twitter,
		params.Gmail,
		params.Facebook,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.Parent,
		&employee.IdentityNo,
		&employee.DrivingLicenseNo,
		&employee.LicenseType,
		&employee.NpwpNo,
		&employee.Name,
		&employee.PlaceOfBirth,
		&employee.DateOfBirth,
		&employee.Email,
		&employee.Address,
		&employee.TemporaryAddress,
		&employee.NeighbourHoodWardNo,
		&employee.UrbanName,
		&employee.SubDistrictName,
		&employee.Religion,
		&employee.MaritalStatus,
		&employee.Citizen,
		&employee.Gender,
		&employee.Ethnic,
		&employee.MobilePhoneNo,
		&employee.PhoneNo,
		&employee.ShirtSize,
		&employee.PantSize,
		&employee.ShoeSize,
		&employee.Bank,
		&employee.BankAccountNo,
		&employee.BankAccountName,
		&employee.FamilyMobilePhoneNo,
		&employee.BpjsNo,
		&employee.JamsostekNo,
		&employee.JamsostekType,
		&employee.JamsostekBalance,
		&employee.FamilyCardNo,
		&employee.Extension,
		&employee.Whatsapp,
		&employee.Instagram,
		&employee.Twitter,
		&employee.Gmail,
		&employee.Facebook,
		&employee.ModifiedBy,
		&employee.ModifiedTime); err != nil {
		tx.Rollback()
		return
	}
	return
}

func (r NewEmployeeRepository) DeleteEmployee(params Dto.EmployeeDataDTO) (tx *sql.Tx, result bool, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	sqlResult, err := tx.Exec(
		querySoftDeleteEmployee,
		params.DeletedBy,
		params.DeletedTime,
		params.ID,
	)
	if err != nil {
		tx.Rollback()
		// tx.Commit()
		return
	}

	count, err := sqlResult.RowsAffected()
	if err != nil {
		tx.Rollback()
		// tx.Commit()
		return
	}

	// err = tx.Commit()
	if err != nil {
		tx.Rollback()
		// tx.Commit()
		return
	}

	if count > 0 {
		result = true
	}

	return
}

func (c NewEmployeeRepository) BrowseSubordinatesData(params Dto.EmployeeParams) (result []Subordinates, err error) {

	var dataSub Subordinates
	connection := Config.DATABASE_MAIN.Get()
	query := fmt.Sprintf(queryBrowseSubordinates, params.Parent, params.Code, params.Name, params.WorkStatus)
	fmt.Println(query)

	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		if err = rows.Scan(
			&dataSub.ID,
			&dataSub.Parent,
			&dataSub.Code,
			&dataSub.Name,
			&dataSub.PhoneNo,
			&dataSub.Email,
			&dataSub.Religion,
			&dataSub.Gender,
			&dataSub.ProfilePhoto,
			&dataSub.WorkStatus); err != nil {
			return
		}
		result = append(result, dataSub)
	}

	return
}

func (c NewEmployeeRepository) DetailSubordinatesData(Parent, ID int64) (result []DetailSubordinates, err error) {

	var dataDetail DetailSubordinates
	connection := Config.DATABASE_MAIN.Get()
	query := fmt.Sprintf(queryDetailSubordinates, Parent, ID)
	rows, err := connection.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	if rows.Err() != nil {
		err = rows.Err()
		return
	}
	for rows.Next() {
		if err = rows.Scan(
			&dataDetail.ID,
			&dataDetail.Code,
			&dataDetail.Type,
			&dataDetail.FingerPrintId,
			&dataDetail.FaceId,
			&dataDetail.MachineId,
			&dataDetail.DepartmentId,
			&dataDetail.SectionId,
			&dataDetail.PositionId,
			&dataDetail.CompanyId,
			&dataDetail.LocationId,
			&dataDetail.CompanyLocationCode,
			&dataDetail.Parent,
			&dataDetail.IdentityNo,
			&dataDetail.DrivingLicenseNo,
			&dataDetail.NpwpNo,
			&dataDetail.Name,
			&dataDetail.PlaceOfBirth,
			&dataDetail.DateOfBirth,
			&dataDetail.Email,
			&dataDetail.Address,
			&dataDetail.TemporaryAddress,
			&dataDetail.NeighbourHoodWardNo,
			&dataDetail.UrbanName,
			&dataDetail.SubDistrictName,
			&dataDetail.Religion,
			&dataDetail.MaritalStatus,
			&dataDetail.Citizen,
			&dataDetail.Gender,
			&dataDetail.Ethnic,
			&dataDetail.MobilePhoneNo,
			&dataDetail.PhoneNo,
			&dataDetail.ShirtSize,
			&dataDetail.PantSize,
			&dataDetail.ShoeSize,
			&dataDetail.JoinDate,
			&dataDetail.ResignDate,
			&dataDetail.ResignReason,
			&dataDetail.Bank,
			&dataDetail.BankAccountNo,
			&dataDetail.BankAccountName,
			&dataDetail.FamilyMobilePhoneNo,
			&dataDetail.WorkStatus,
			&dataDetail.ProfilePhoto,
			&dataDetail.ContractStart,
			&dataDetail.ContractEnd,
			&dataDetail.BpjsNo,
			&dataDetail.JamsostekNo,
			&dataDetail.JamsostekType,
			&dataDetail.JamsostekBalance,
			&dataDetail.FamilyCardNo,
			&dataDetail.ActiveStatus,
			&dataDetail.CreatedBy,
			&dataDetail.CreatedTime,
			&dataDetail.ModifiedBy,
			&dataDetail.ModifiedTime,
			&dataDetail.Extension,
			&dataDetail.PostalCode); err != nil {
			return
		}
		result = append(result, dataDetail)
	}

	return
}

func (r NewEmployeeRepository) UpdateEmployeeWorkStatus(params Dto.WorkStatusUpdateDTO) (tx *sql.Tx, employee WorkStatusUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(
		queryUpdateEmployeeWorkStatus,
		params.WorkStatus,
		params.ContractStart,
		params.ContractEnd,
		params.WorkStatusChangeReason,
		params.WorkStatusChangeDate,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.WorkStatus,
		&employee.ContractStart,
		&employee.ContractEnd,
		&employee.WorkStatusChangeReason,
		&employee.WorkStatusChangeDate); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r NewEmployeeRepository) UpdateEmployeeFingerprint(params Dto.FingerprintUpdateDTO) (tx *sql.Tx, employee FingerprintUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(
		queryUpdateEmployeeFingerprint,
		params.FingerPrintId,
		params.FaceId,
		params.MachineId,
		params.ModifiedBy,
		params.ModifiedTime,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.FingerPrintId,
		&employee.FaceId,
		&employee.MachineId); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r NewEmployeeRepository) UpdateEmployeeMachineId(params Dto.MachineIdUpdateDTO) (tx *sql.Tx, employee MachineIdUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(
		queryUpdateEmployeeMachineId,
		params.MachineId,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.MachineId); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r NewEmployeeRepository) UpdateEmployeeFaceId(params Dto.FaceIdUpdateDTO) (tx *sql.Tx, employee FaceIdUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(
		queryUpdateEmployeeFaceId,
		params.FaceId,
		params.MachineId,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.FaceId,
		&employee.MachineId); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (r NewEmployeeRepository) UpdateEmployeeResignStatus(params Dto.ResignStatusUpdateDTO) (tx *sql.Tx, employee ResignStatusUpdate, err error) {
	tx, err = Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	row := tx.QueryRow(
		queryUpdateEmployeeResignStatus,
		params.WorkStatus,
		params.ResignDate,
		params.ResignReason,
		params.ID,
	)

	if err = row.Scan(
		&employee.ID,
		&employee.WorkStatus,
		&employee.JoinDate,
		&employee.ResignDate,
		&employee.ResignReason); err != nil {
		tx.Rollback()
		return
	}

	return
}

func (c NewEmployeeRepository) ValidationDuplicateDataEmployee(Code string) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidationEmployee, Code).Scan(&exists)
	return
}

func (r NewEmployeeRepository) GetEmployeeWorkStatus(employeeId int64) (result GetEmployeeWorkStatus, err error) {
	connection := Config.DATABASE_MAIN.Get()

	row := connection.QueryRow(queryGetEmployeeWorkStatus, employeeId)

	if err = row.Scan(
		&result.ID,
		&result.WorkStatus,
		&result.ContractStart,
		&result.ContractEnd,
		&result.WorkStatusChangeReason,
		&result.WorkStatusChangeDate,
	); err != nil {
		return
	}

	return
}

func (r NewEmployeeRepository) SaveWorkStatus(tx *sql.Tx, params Dto.SaveWorkStatusDTO) (tempTx *sql.Tx, workStatusId int64, err error) {
	row := tx.QueryRow(
		queryInsertWorkStatus,
		params.WorkStatus,
		params.EmployeeId,
		params.ContractStart,
		params.ContractEnd,
		params.Deleted,
		params.CreatedBy,
		params.CreatedTime,
	)

	if err = row.Scan(&workStatusId); err != nil {
		tx.Rollback()
		return
	}

	tempTx = tx
	return
}

func (r NewEmployeeRepository) HardDeleteWorkStatus(workStatusId int64) (err error) {
	connection := Config.DATABASE_MAIN.Get()

	_, err = connection.Exec(queryHardDeleteWorkStatus, workStatusId)

	if err != nil {
		return
	}
	return
}

func (c NewEmployeeRepository) ValidationDuplicateMobilePhone(MobilePhoneNo string) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidateMobilePhone, MobilePhoneNo).Scan(&exists)
	return
}

func (c NewEmployeeRepository) ValidationDuplicateEmail(Email string) (exists bool, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryValidateEmail, Email).Scan(&exists)
	return
}

func (c NewEmployeeRepository) CheckExistMobilePhoneNumber(phoneNo string) (employee CheckExistData, err error) {
	connection := Config.DATABASE_MAIN.Get()

	row := connection.QueryRow(queryCheckExistMobilePhoneNumber, phoneNo)
	if err = row.Scan(
		&employee.ID,
		&employee.Code,
		&employee.Name,
		&employee.MobilePhoneNo,
		&employee.Email,
		&employee.ActiveStatus,
	); err != nil {
		return
	}
	return
}

func (c NewEmployeeRepository) CheckExistEmail(email string) (employee CheckExistData, err error) {
	connection := Config.DATABASE_MAIN.Get()

	row := connection.QueryRow(queryCheckExistEmail, email)
	if err = row.Scan(
		&employee.ID,
		&employee.Code,
		&employee.Name,
		&employee.MobilePhoneNo,
		&employee.Email,
		&employee.ActiveStatus,
	); err != nil {
		return
	}
	return
}

func (c NewEmployeeRepository) CountEmployees() (count int64, err error) {
	connection := Config.DATABASE_MAIN.Get()
	err = connection.QueryRow(queryCountEmployees).Scan(&count)
	return
}
