package employee

import (
	"EmployeeService/Config"
	"EmployeeService/Controller/Dto"
	"context"
	"database/sql"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	queryCompanyUpdate         = `UPDATE public.leave_balances SET company_name = $1 WHERE company_name = $2`
	queryDepartmentUpdate      = `UPDATE public.leave_balances SET department_name = $1 WHERE department_name = $2`
	queryLocationUpdate        = `UPDATE public.leave_balances SET location_name = $1 WHERE location_name = $2`
	queryCompanyLocationUpdate = `UPDATE public.employees SET company_location_code = $1 WHERE company_location_code = $2`
)

//EmployeeAdded
func (r NewSubscriberEmployeeRepository) EmployeeSave(params Dto.EmployeeDataDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	_, err = collection.InsertOne(context.TODO(), params)
	if err != nil {
		return
	}
	return
}

func (c NewSubscriberEmployeeRepository) ValidationDuplicateData(params Dto.EmployeeDataDTO) (exists bool, err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	find := bson.M{
		"code": params.Code,
	}
	result := collection.FindOne(context.TODO(), find)

	var employee Employees
	if err = result.Decode(&employee); err != nil {
		if strings.Contains(err.Error(), "mongo: no documents in result") {
			exists = false
			err = nil
			return
		}
		return
	}

	exists = true
	return
}

// EmployeeTransfered
func (r NewSubscriberEmployeeRepository) EmployeeUpdatePsgl(tx *sql.Tx, params Dto.SubscribeEmployeePsgl) (tempTx *sql.Tx, err error) {

	queryConsume := `UPDATE public.employees SET 	
						department_id=$1, 
						section_id=$2,
						position_id=$3, 
						company_id=$4, 
						location_id=$5, 
						company_location_code=$6, 
						code=$7 
					where id = $8 `
	row := tx.QueryRow(
		queryConsume,
		params.NewIdDepartment,
		params.NewIdSection,
		params.NewIdPosition,
		params.NewIdCompany,
		params.NewIdLocation,
		params.NewCompanyLocationCode,
		params.NewEmployeeCode,
		params.EmployeeId,
	)

	if err = row.Err(); err != nil {
		tx.Rollback()
		return
	}
	tempTx = tx
	return
}

func (r NewSubscriberEmployeeRepository) EmployeeUpdateMngo(params Dto.SubscribeEmployeePsgl) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{"id": params.EmployeeId}
	update := bson.M{
		"$set": bson.M{
			"code":                  params.NewEmployeeCode,
			"department_id":         params.NewIdDepartment,
			"department_name":       params.NewDepartmentName,
			"section_id":            params.NewIdSection,
			"section_name":          params.NewSectionName,
			"position_id":           params.NewIdPosition,
			"position_name":         params.NewPositionName,
			"company_id":            params.NewIdCompany,
			"company_name":          params.NewCompanyName,
			"location_id":           params.NewIdLocation,
			"location_name":         params.NewLocationName,
			"company_location_code": params.NewCompanyLocationCode,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeePersonalInfoUpdated
func (r NewSubscriberEmployeeRepository) EmployeePersonalInfoUpdated(params Dto.EmployeeDataDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"identity_no":            params.IdentityNo,
			"driving_license_no":     params.DrivingLicenseNo,
			"license_type":           params.LicenseType,
			"npwp_no":                params.NpwpNo,
			"name":                   params.Name,
			"place_of_birth":         params.PlaceOfBirth,
			"date_of_birth":          params.DateOfBirth,
			"email":                  params.Email,
			"address":                params.Address,
			"temporary_address":      params.TemporaryAddress,
			"neighbour_hood_ward_no": params.NeighbourHoodWardNo,
			"urban_name":             params.UrbanName,
			"sub_district_name":      params.SubDistrictName,
			"religion":               params.Religion,
			"marital_status":         params.MaritalStatus,
			"citizen":                params.Citizen,
			"gender":                 params.Gender,
			"ethnic":                 params.Ethnic,
			"mobile_phone_no":        params.MobilePhoneNo,
			"phone_no":               params.PhoneNo,
			"shirt_size":             params.ShirtSize,
			"pant_size":              params.PantSize,
			"shoe_size":              params.ShoeSize,
			"bank":                   params.Bank,
			"bank_account_no":        params.BankAccountNo,
			"bank_account_name":      params.BankAccountName,
			"family_mobile_phone_no": params.FamilyMobilePhoneNo,
			"bpjs_no":                params.BpjsNo,
			"jamsostek_no":           params.JamsostekNo,
			"jamsostek_type":         params.JamsostekType,
			"jamsostek_balance":      params.JamsostekBalance,
			"family_card_no":         params.FamilyCardNo,
			"modified_by":            params.ModifiedBy,
			"modified_time":          params.ModifiedTime,
			"parent":                 params.Parent,
			"whatsapp":               params.Whatsapp,
			"instagram":              params.Instagram,
			"twitter":                params.Twitter,
			"gmail":                  params.Gmail,
			"facebook":               params.Facebook,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeeDeleted
func (r NewSubscriberEmployeeRepository) EmployeeDeleted(params Dto.EmployeeDataDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{"id": params.ID}
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return
	}
	return
}

// EmployeeUpdateWorkStatus
func (r NewSubscriberEmployeeRepository) EmployeeWorkStatusUpdated(params Dto.WorkStatusUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"work_status":               params.WorkStatus,
			"contract_start":            params.ContractStart,
			"contract_end":              params.ContractEnd,
			"work_status_change_reason": params.WorkStatusChangeReason,
			"work_status_change_date":   params.WorkStatusChangeDate,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeeFingerUpdated
func (r NewSubscriberEmployeeRepository) EmployeeFingerUpdated(params Dto.FingerprintUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"finger_print_id": params.FingerPrintId,
			"face_id":         params.FaceId,
			"machine_id":      params.MachineId,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeeMachineIDUpdated
func (r NewSubscriberEmployeeRepository) EmployeeMachineIdUpdated(params Dto.MachineIdUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"machine_id": params.MachineId,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeeFaceIdUpdated
func (r NewSubscriberEmployeeRepository) EmployeeFaceIdUpdated(params Dto.FaceIdUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"face_id":    params.FaceId,
			"machine_id": params.MachineId,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

// EmployeeResigned
func (r NewSubscriberEmployeeRepository) EmployeeResigned(params Dto.EmployeeDataDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.M{"id": params.ID}
	update := bson.M{
		"$set": bson.M{
			"work_status":   params.WorkStatus,
			"resign_date":   params.ResignDate,
			"resign_reason": params.ResignReason,
		},
	}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

func (r NewSubscriberEmployeeRepository) CompanyUpdate(params Dto.SubscribeCompanyUpdateDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		queryCompanyUpdate,
		params.CompanyName,
		params.OldCompanyName,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	// Mongo
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{{Key: "company_id", Value: params.CompanyId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "company_name", Value: params.CompanyName}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}
	tx.Commit()
	return
}

func (r NewSubscriberEmployeeRepository) CompanyLocationUpdate(params Dto.SubscribeCompanyLocationCodeUpdateDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		queryCompanyLocationUpdate,
		params.CompanyLocationCode,
		params.OldCompanyLocationCode,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	//Mongo
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{{Key: "company_location_code", Value: params.CompanyLocationCode}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "company_location_code", Value: params.CompanyLocationCode}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}

	tx.Commit()
	return
}

func (r NewSubscriberEmployeeRepository) DepartmentUpdate(params Dto.SubscribeDepartmentUpdateDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		queryDepartmentUpdate,
		params.DepartmentName,
		params.OldDepartmentName,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")

	filter := bson.D{{Key: "department_id", Value: params.DepartmentId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "department_name", Value: params.DepartmentName}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}
	tx.Commit()
	return
}

func (r NewSubscriberEmployeeRepository) LocationUpdate(params Dto.SubscribeLocationUpdateDTO) (err error) {
	tx, err := Config.DATABASE_MAIN.Get().Begin()
	if err != nil {
		return
	}
	_, err = tx.Exec(
		queryLocationUpdate,
		params.LocationName,
		params.OldLocationName,
	)
	if err != nil {
		tx.Rollback()
		return
	}

	// Mongo
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{{Key: "location_id", Value: params.LocationId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "location_name", Value: params.LocationName}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}

	tx.Commit()
	return
}

func (r NewSubscriberEmployeeRepository) SectionUpdate(params Dto.SubscribeSectionUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{{Key: "section_id", Value: params.SectionId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "section_name", Value: params.SectionName}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}
	return
}

func (r NewSubscriberEmployeeRepository) PositionUpdate(params Dto.SubscribePositionUpdateDTO) (err error) {
	clientMongo := Config.DATABASE_MONGO.Get()
	collection := clientMongo.Database("employee").Collection("employees")
	filter := bson.D{{Key: "position_id", Value: params.PositionId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "position_name", Value: params.PositionName}},
		},
	}

	_, err = collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return
	}

	return
}
