package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Dapr"
	"EmployeeService/Library/Helper/Response"
	validator "EmployeeService/Library/Helper/Validator"
	"EmployeeService/Library/Logging"
	"EmployeeService/Library/PubSub/Publisher"
	"EmployeeService/Library/Storage"
	"EmployeeService/Repository/employee"
	Services "EmployeeService/Services/employee"
	LeaveBalanceServices "EmployeeService/Services/leaveBalance"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/mongo"
)

type Employee interface {
	EmployeeBrowse(w http.ResponseWriter, r *http.Request)
	EmployeeBrowseDetail(w http.ResponseWriter, r *http.Request)
	EmployeeSearch(w http.ResponseWriter, r *http.Request)
	EmployeeGetHead(w http.ResponseWriter, r *http.Request)
	EmployeeAdded(w http.ResponseWriter, r *http.Request)
	EmployeePersonalInfoUpdated(w http.ResponseWriter, r *http.Request)
	EmployeeDeleted(w http.ResponseWriter, r *http.Request)
	SubordinatesData(w http.ResponseWriter, r *http.Request)
	DetailSubordinatesData(w http.ResponseWriter, r *http.Request)
	EmployeeWorkStatusUpdated(w http.ResponseWriter, r *http.Request)
	DeleteWorkStatus(w http.ResponseWriter, r *http.Request)
	EmployeeResigned(w http.ResponseWriter, r *http.Request)
	EmployeeFingerUpdated(w http.ResponseWriter, r *http.Request)
	EmployeeMachineIdUpdated(w http.ResponseWriter, r *http.Request)
	EmployeeFaceIdUpdated(w http.ResponseWriter, r *http.Request)

	EmployeeUploadPhoto(w http.ResponseWriter, r *http.Request)
	EmployeeReadPhoto(w http.ResponseWriter, r *http.Request)

	//
	CheckExistMobilePhoneNumber(w http.ResponseWriter, r *http.Request)
	CheckExistEmail(w http.ResponseWriter, r *http.Request)
	CheckExistEmployeeID(w http.ResponseWriter, r *http.Request)

	CountEmployees(w http.ResponseWriter, r *http.Request)
}

// @Tags         Employee
// @Summary      Employee Upload Photo
// @Description  Sample Parameter: `3`
// @Accept       json
// @Produce      json
// @Param        id    path      int64                        true  "Id"
// @Param        file  formData  file                         true  "file Photo"
// @Success      200   {object}  Response.RespResultStruct{}  "OK"
// @Failure      500   {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/upload-photo [post]
// Keterangan : function ini digunakan untuk mengunggah photo karyawan dengan name kode karyawan yang di encrypt mengunakan md5 ke cloud storage
func (c Controller) EmployeeUploadPhoto(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeDataDTO
	var err error
	idEmployee := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 200*1024) // 200 Kb

	uploadedFile, handler, err := r.FormFile("file")
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}
	defer uploadedFile.Close()

	if err := validator.ValidateExtension(filepath.Ext(handler.Filename)); err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.BrowseEmployeeDetail(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	if err = Storage.UploadFile(result.Code, uploadedFile); err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	Response.ResponseJson(w, true, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Employee Upload Photo
// @Description  Sample Parameter: `3`
// @Accept       json
// @Produce      json
// @Param        id   path      int64                        true  "Id"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/read-photo [get]
// Keterangan : function ini digunakan untuk mengunduh photo karyawan dengan name kode karyawan yang di encrypt mengunakan md5 ke cloud storage
func (c Controller) EmployeeReadPhoto(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeDataDTO
	var err error
	idEmployee := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.BrowseEmployeeDetail(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}
	file, err := Storage.ReadFile(result.Code)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}
	Response.ResponseFile(w, file)
}

// @Tags         Employee
// @Summary      EmployeeBrowse
// @Description  Sample Parameter: `?sortColumn=code&sortOrder=desc&pageSize=10&pageNumber=1&code=111222333&name=staff`
// @Accept       json
// @Produce      json
// @Param        sortColumn           query     string                       false  "sortColumn"
// @Param        sortOrder            query     string                       false  "sortOrder"
// @Param        pageSize             query     int64                        false  "pageSize"
// @Param        pageNumber           query     int64                        false  "pageNumber"
// @Param        code                 query     string                       false  "code"
// @Param        name                 query     string                       false  "name"
// @Param        machineId            query     int64                        false  "machineId"
// @Param        companyName          query     string                       false  "companyName"
// @Param        companyLocationCode  query     string                       false  "companyLocationCode"
// @Param        departmentName       query     string                       false  "departmentName"
// @Param        workStatus           query     string                       false  "workStatus"
// @Success      200                  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500                  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/ [get]
// Keterangan : function ini digunakan untuk menampilkan data karyawan dengan filter seperti nama,id karyawan dll -> Data di tampilkan dari Mongo
func (c Controller) EmployeeBrowse(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeBrowseParams
	var decoder = schema.NewDecoder()
	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}
	result, err := Services.Employee.BrowseEmployee(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeBrowseDetail
// @Description  Sample Parameter: `3`
// @Accept       json
// @Produce      json
// @Param        id   path      int64                        true  "Id"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id} [get]
// Keterangan : function ini digunakan untuk menampilkan spesifik data karyawan dengan filter id karyawan
func (c Controller) EmployeeBrowseDetail(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeDataDTO
	var err error

	idEmployee := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.BrowseEmployeeDetail(params)
	if err != nil {
		if strings.Contains(err.Error(), mongo.ErrNoDocuments.Error()) {
			Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeSearch
// @Description  Sample Parameter: `?search=wijaya`
// @Accept       json
// @Produce      json
// @Param        search                 query     string                       false  "search"
// @Success      200                  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500                  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employee [get]
// Keterangan : function ini digunakan untuk menampilkan data karyawan dengan filter nama untuk keperluan combobox -> Data di tampilkan dari Mongo
func (c Controller) EmployeeSearch(w http.ResponseWriter, r *http.Request) {
	var params Dto.EmployeeSearchDTO
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.SearchEmployee(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeGetHead
// @Description  Sample Parameter: `?transaction=leave`
// @Accept       json
// @Produce      json
// @Param        id           path      int64                        true  "Id"
// @Param        transaction  query     string                       true  "transaction"  Enums(permission, leave, business-trip,change-schedulue, job-vacancy-request, overtimes)
// @Success      200          {object}  Response.RespResultStruct{}  "OK"
// @Failure      500          {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/head [get]
// Keterangan : function ini digunakan untuk menampilkan Atasan seorang karyawan sesuai dengan jenis transaski yang di pilih
func (c Controller) EmployeeGetHead(w http.ResponseWriter, r *http.Request) {

	var err error
	var transaction Dto.GetHeadParams
	var decoder = schema.NewDecoder()

	err = decoder.Decode(&transaction, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	// Jumlah data atasan yang muncul sesuai dengan jenis transaksi
	jumlahAtasan := 0
	if transaction.Transaction == "permission" {
		jumlahAtasan = 1
	} else if transaction.Transaction == "leave" {
		jumlahAtasan = 2
	} else if transaction.Transaction == "business-trip" {
		jumlahAtasan = 2
	} else if transaction.Transaction == "change-schedulue" {
		jumlahAtasan = 1
	} else if transaction.Transaction == "job-vacancy-request" {
		jumlahAtasan = 1
	} else if transaction.Transaction == "overtimes" {
		jumlahAtasan = 1
	} else {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	var employeeID int64
	idEmployee := chi.URLParam(r, "id")
	employeeID, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.GetHead(employeeID)
	if err != nil {
		if strings.Contains(err.Error(), mongo.ErrNoDocuments.Error()) {
			Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	idParent := result.Parent
	parents := make([]employee.GetHead, 0)
	i := 1
	for i <= jumlahAtasan {
		if idParent != 0 {
			getHeadList, err := Services.Employee.GetHead(idParent)
			if err != nil {
				if strings.Contains(err.Error(), mongo.ErrNoDocuments.Error()) {
					Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
					return
				}
				Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
				return
			}
			parents = append(parents, getHeadList)
			idParent = getHeadList.Parent
			i++
			continue
		}
		break
	}

	Response.ResponseJson(w, parents, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeAdded
// @Description  Sample Payload: `{"code":"12345.67890","type":"-","fingerPrintId":"100200300","faceId":"50264525","machineId":11884632,"MachineName":"Mesin","departmentId":196347512,"departmentName":"deptName","sectionId":1,"sectionName":"sectName","positionId":1,"positionName":"potiName","companyId":1,"companyName":"CompnName","locationId":1,"locationName":"LocName","companyLocationId":1,"companyLocationCode":"SSSS","parent":61,"parentName":"Juliana","identityNo":"1271000524163","drivingLicenseNo":"150404020603","npwpNo":"0","name":"Staff","placeOfBirth":"Medan","dateOfBirth":"1990-02-10T07:48:48.243Z","email":"test@ssss.com","address":"Jalan Veteran","temporaryAddress":"Jalan Veteran","neighbourHoodWardNo":"1020","urbanName":"Medan","subDistrictName":"Medan Timur","religion":"Buddha","maritalStatus":"Divorced","citizen":"Indonesia","gender":"Male","ethnic":"-","mobilePhoneNo":"0821789456","phoneNo":"0821789456","shirtSize":"XL","pantSize":30,"shoeSize":40,"joinDate":"2022-02-10T07:48:48.243Z","resignDate":null,"resignReason":null,"bank":"Mandiri","bankAccountNo":"1006001890","bankAccountName":"Test","familyMobilePhoneNo":"555666999","workStatus":"Kontrak","profilePhoto":"/Img/test","contractStart":"2022-02-10T07:48:48.243Z","contractEnd":"2022-05-10T07:48:48.243Z","bpjsNo":"0","jamsostekNo":"0","jamsostekType":"-","jamsostekBalance":0,"familyCardNo":"0","activeStatus":true,"extension":"123456","whatsapp":"082189890","instagram":"@instagram","twitter":"@twitter","gmail":"@gmail","facebook":"@facebook"}`
// @Accept       json
// @Produce      json
// @Param        "payload" body Dto.EmployeeDataDTO true "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees [post]
// Keterangan : function ini digunakan untuk menambahkan data karyawan
func (c Controller) EmployeeAdded(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeDataDTO
	var paramsSC Dto.LeaveBalanceDataDTO

	err := json.NewDecoder(r.Body).Decode(&params)
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, nil)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	timeStamp := time.Now().In(location)
	year := time.Now().Year()
	params.CreatedBy = EmployeeName
	params.CreatedTime = timeStamp

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	exists, err := Services.Employee.ValidationDuplicateDataEmployee(params.Code)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if exists {
		Response.ResponseError(w, err, Constant.StatusBadRequestAlreadyExists)
		return
	}

	// Check nomor hp jika duplikat
	mobilePhoneExist, err := Services.Employee.ValidationDuplicateMobilePhone(params.MobilePhoneNo)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if mobilePhoneExist {
		Response.ResponseError(w, err, Constant.StatusBadRequestMobilePhoneAlreadyExists)
		return
	}

	// Check Email jika duplikat
	emailExist, err := Services.Employee.ValidationDuplicateEmail(params.Email)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if emailExist {
		Response.ResponseError(w, err, Constant.StatusBadRequestEmailAlreadyExists)
		return
	}

	// Simpan Karyawan
	tx, result, err := Services.Employee.SaveEmployee(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	result.CompanyLocationAlias = params.CompanyLocationAlias
	result.DepartmentName = params.DepartmentName
	result.SectionName = params.SectionName
	result.PositionName = params.PositionName
	result.CompanyName = params.CompanyName
	result.LocationName = params.LocationName
	result.MachineName = params.MachineName
	result.ParentName = params.ParentName

	// Simpan Data Work Status
	params.ID = result.ID
	// err = Services.Employee.SaveWorkStatus(params)
	// if err != nil {
	// 	Logging.LogError(err, r)
	// 	Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
	// 	return
	// }

	paramsSC.EmployeeId = result.ID
	paramsSC.EmployeeCode = result.Code
	paramsSC.CompanyName = params.CompanyName
	paramsSC.LocationName = params.LocationName
	paramsSC.DepartmentName = params.DepartmentName
	paramsSC.StartBalances = 0
	paramsSC.IncreaseBalance = 0
	paramsSC.DecreaseBalance = 0
	paramsSC.LastPeriodBalance = 0
	paramsSC.CurrentBalance = 0
	paramsSC.IsActive = true
	paramsSC.Period = strconv.FormatInt(int64(year), 10)
	paramsSC.JoinDate = &result.JoinDate
	paramsSC.ExpiredDate = result.JoinDate.AddDate(1, 0, 0)
	paramsSC.CreatedBy = EmployeeName
	paramsSC.CreatedTime = timeStamp

	// Buat data saldo cuti pertama kali ketika karyawan berhasil di tambahkan dengan saldo cuti 0
	err2 := LeaveBalanceServices.LeaveBalance.SaveLeaveBalance(paramsSC)
	if err2 != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err2, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEADDED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	_, err = client.Publish(Publisher.Publisher{
		Data:              result,
		JWTDocodedPayload: contextMap,
	})

	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeePersonalInfoUpdated
// @Description  Sample Payload: `{"parent":2,"identityNo":"15246352","drivingLicenseNo":"584967222","npwpNo":"78995525","name":"Cris Jhon","placeOfBirth":"Jakarta","dateOfBirth":"2000-03-10T07:48:48.243Z","email":"crhis@email.com","address":"Jalan Buku","temporaryAddress":"Jalan Buku","neighbourHoodWardNo":"12345","urbanName":"Medan","subDistrictName":"Marelan","religion":"Atheis","maritalStatus":"Married","citizen":"Medan","gender":"Pria","ethnic":"ethnic","mobilePhoneNo":"082134566","phoneNo":"098765465","shirtSize":"L","pantSize":10,"shoeSize":10,"bank":"BCA","bankAccountNo":"150505002","bankAccountName":"Chris Jhon","familyMobilePhoneNo":"082134567673","bpjsNo":"123455","jamsostekNo":"1234567890","jamsostekType":"Kelas 1","jamsostekBalance":20000,"familyCardNo":"241231","extension":"-","licenseType":"-","whatsapp":"0821898901","instagram":"@instagramm","twitter":"@twitterr","gmail":"@gmaill","facebook":"@facebookk"}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.EmployeeUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id} [put]
// Keterangan : function ini digunakan untuk update data karyawan
func (c Controller) EmployeePersonalInfoUpdated(w http.ResponseWriter, r *http.Request) {
	var params Dto.EmployeeUpdateDTO
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	employeeStr := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(employeeStr, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	timeNow := time.Now().In(location)
	params.ModifiedBy = &EmployeeName
	params.ModifiedTime = &timeNow

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	tx, result, err := Services.Employee.UpdateEmployee(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEPERSONALINFOUPDATED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	subs, err := client.Subscriptions()
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	// statePayload get old data, publisher get new data
	oldData, err := Services.Employee.BrowseEmployeeDetail(Dto.EmployeeDataDTO{ID: params.ID})
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	state := Dapr.GetConfigState(subs, "", "", oldData)
	fmt.Printf("%+v \n", state)
	err = Dapr.StatePublish(state)
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	_, err = client.Publish(Publisher.Publisher{
		Key:               state.StateFmt.Key,
		StateStore:        state.StateStore,
		JWTDocodedPayload: contextMap,
		Data:              result,
	})

	if err != nil {
		Logging.LogError(err.Error(), r)
		Dapr.StateDelete(state.StateStore, state.StateFmt.Key)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags      Employee
// @Summary   EmployeeDeleted
// @Accept    json
// @Produce   json
// @Param     id   path      int64                        true  "Id"
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /employees/{id} [delete]
// Keterangan : function ini digunakan untuk menghapus data karyawan
func (c Controller) EmployeeDeleted(w http.ResponseWriter, r *http.Request) {

	var params Dto.EmployeeDataDTO
	var err error

	employeeStr := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(employeeStr, 10, 64)

	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	timeNow := time.Now().In(location)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	params.Deleted = true
	params.CreatedBy = EmployeeName
	params.CreatedTime = timeNow
	params.DeletedBy = &EmployeeName
	params.DeletedTime = &timeNow

	tx, delete, err := Services.Employee.DeleteEmployee(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	params.ContractStart = params.DeletedTime
	params.ContractEnd = params.DeletedTime

	// err = Services.Employee.SaveWorkStatus(params)
	// if err != nil {
	// 	Logging.LogError(err, r)
	// 	Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
	// 	return
	// }
	result := employee.Employees{ID: params.ID}

	client, err := Publisher.TOPIC_EMPLOYEEDELETED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	_, err = client.Publish(result)
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, delete, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      SubordinatesData
// @Description  Sample Parameter: `?name=asis&workStatus=Active&code=10.0003`
// @Accept       json
// @Produce      json
// @Param        id          path      int64                        true   "Id"
// @Param        name        query     string                       false  "name"
// @Param        workStatus  query     string                       false  "workStatus"
// @Param        code        query     string                       false  "code"
// @Success      200         {object}  Response.RespResultStruct{}  "OK"
// @Failure      500         {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/subordinates [get]
// Keterangan : function ini digunakan untuk menampilkan data bawahan dari seorang karyawan
func (c Controller) SubordinatesData(w http.ResponseWriter, r *http.Request) {
	var params Dto.EmployeeParams
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	idEmployee := chi.URLParam(r, "id")
	params.Parent, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.BrowseSubordinatesData(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      DetailSubordinatesData
// @Description  Sample Parameter: `/2/subordinates/9`
// @Accept       json
// @Produce      json
// @Param        id_parent  path      int64                        true  "ParentID"
// @Param        id         path      int64                        true  "Id"
// @Success      200        {object}  Response.RespResultStruct{}  "OK"
// @Failure      500        {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id_parent}/subordinates/{id} [get]
// Keterangan : function ini digunakan untuk menampilkan data detail bawahan dari seorang karyawan
func (c Controller) DetailSubordinatesData(w http.ResponseWriter, r *http.Request) {
	var Parent, ID int64
	var err error

	idParent := chi.URLParam(r, "id_parent")
	idEmployee := chi.URLParam(r, "id")

	Parent, err = strconv.ParseInt(idParent, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	ID, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.DetailSubordinatesData(Parent, ID)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Employee Work Status Updated (Training, Kontrak, Karyawan Tetap, Mitra Kerja, Resign)
// @Description  Sample Payload: `{"WorkStatus" : "Karyawan Tetap","ContractStart" : "2022-05-11T07:48:48.243Z","ContractEnd" : "2022-12-10T07:48:48.243Z","WorkStatusChangeReason" : "alasan","WorkStatusChangeDate" : "2022-12-10T07:48:48.243Z"}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.WorkStatusUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/workstatus [put]
// Keterangan : function ini digunakan untuk mengupdate data status pekerja dari seorang karyawan (Mis. Karyawan Tetap,Kontrak,Trainning)
func (c Controller) EmployeeWorkStatusUpdated(w http.ResponseWriter, r *http.Request) {

	var params Dto.WorkStatusUpdateDTO
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	employeeStr := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(employeeStr, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	// Check Status apabila karyawan sudah resign, maka tidak bisa di ubah lagi
	checkStatus, err := Services.Employee.BrowseEmployeeDetail(Dto.EmployeeDataDTO{ID: params.ID})
	if err != nil {
		Logging.LogError(err.Error(), r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if checkStatus.WorkStatus == "Resign" {
		Response.ResponseError(w, nil, Constant.StatusBadRequestEmployeeAlreadyResign)
		return
	}

	timeNow := time.Now().In(location)
	params.CreatedBy = EmployeeName
	params.CreatedTime = time.Now().In(location)
	params.ModifiedBy = &EmployeeName
	params.ModifiedTime = &timeNow

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	if params.WorkStatus == "Training" || params.WorkStatus == "Kontrak" {
		if params.ContractStart == nil {
			Response.ResponseError(w, err, Constant.StatusBadRequestContractStartNull)
			return
		}

		if params.ContractEnd == nil {
			Response.ResponseError(w, err, Constant.StatusBadRequestContractEndNull)
			return
		}
	}

	tx, result, err := Services.Employee.UpdateEmployeeWorkStatus(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	paramSaveWorkStatus := Dto.SaveWorkStatusDTO{
		WorkStatus:    params.WorkStatus,
		EmployeeId:    params.ID,
		ContractStart: params.ContractStart,
		ContractEnd:   params.ContractEnd,
		CreatedBy:     params.CreatedBy,
		CreatedTime:   params.CreatedTime,
	}
	tx, workStatusId, err := Services.Employee.SaveWorkStatus(tx, paramSaveWorkStatus)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEWORKSTATUSUPDATED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}

	subs, err := client.Subscriptions()
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	// statePayload get old data, publisher get new data
	oldData, err := Services.Employee.GetEmployeeWorkStatus(params.ID)
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	paramsUpdateWorkStatus := Dto.UpdateWorkStatusStateDTO{
		EmployeeId:             oldData.ID,
		WorkStatus:             oldData.WorkStatus,
		ContractStart:          oldData.ContractStart,
		ContractEnd:            oldData.ContractEnd,
		WorkStatusChangeReason: oldData.WorkStatusChangeReason,
		WorkStatusChangeDate:   oldData.WorkStatusChangeDate,
		WorkStatusId:           workStatusId,
	}

	state := Dapr.GetConfigState(subs, "", "", paramsUpdateWorkStatus)
	fmt.Printf("%+v \n", state)
	err = Dapr.StatePublish(state)
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	_, err = client.Publish(Publisher.Publisher{
		Key:               state.StateFmt.Key,
		StateStore:        state.StateStore,
		JWTDocodedPayload: contextMap,
		Data:              result,
	})
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}

	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeFingerUpdated
// @Description  Sample Payload: `{"FingerPrintId":"Update FingerPrintId","FaceId":"Update FaceId","MachineId":10}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.FingerprintUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/fingerprint [put]
// Keterangan : function ini digunakan untuk mengupdate data fingerprint pegawai
func (c Controller) EmployeeFingerUpdated(w http.ResponseWriter, r *http.Request) {

	var params Dto.FingerprintUpdateDTO
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	idData := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idData, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	timeNow := time.Now().In(location)
	params.ModifiedBy = &EmployeeName
	params.ModifiedTime = &timeNow

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	tx, result, err := Services.Employee.UpdateEmployeeFingerprint(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEFINGERUPDATED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	_, err = client.Publish(result)
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Update Machine Id
// @Description  Sample Payload: `{"MachineId":666}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.MachineIdUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/machineid [put]
// Keterangan : function ini digunakan untuk mengupdate data kode mesin pegawai
func (c Controller) EmployeeMachineIdUpdated(w http.ResponseWriter, r *http.Request) {

	var params Dto.MachineIdUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	idData := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idData, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	tx, result, err := Services.Employee.UpdateEmployeeMachineId(params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEMACHINEIDUPDATED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	_, err = client.Publish(result)
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Update Face Id
// @Description  Sample Payload: `{"faceId":"3211","machineId":3212}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.FaceIdUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/faceid [put]
// Keterangan : function ini digunakan untuk mengupdate data Face Id pegawai
func (c Controller) EmployeeFaceIdUpdated(w http.ResponseWriter, r *http.Request) {

	var params Dto.FaceIdUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	idData := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idData, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	tx, result, err := Services.Employee.UpdateEmployeeFaceId(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEEFACEIDUPDATED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	_, err = client.Publish(result)
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}
	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      EmployeeResigned
// @Description  Sample Payload: `{"WorkStatus":"Resign","ResignDate":"2022-03-12T07:48:48.243Z","ResignReason":"Pindah Kota"}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.ResignStatusUpdateDTO true  "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /employees/{id}/resign [put]
// Keterangan : function ini digunakan untuk mengupdate status pegawai menjadi resign dan alasan resign
func (c Controller) EmployeeResigned(w http.ResponseWriter, r *http.Request) {

	var params Dto.ResignStatusUpdateDTO
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	idData := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idData, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	location, err := time.LoadLocation(Constant.TimeLocation)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerError)
		return
	}

	tx, result, err := Services.Employee.UpdateEmployeeResignStatus(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	// Status resign di insert ke dalam table work status -> Sebagai history
	paramSaveWorkStatus := Dto.SaveWorkStatusDTO{
		WorkStatus:  params.WorkStatus,
		EmployeeId:  params.ID,
		CreatedBy:   EmployeeName,
		CreatedTime: time.Now().In(location),
	}
	tx, _, err = Services.Employee.SaveWorkStatus(tx, paramSaveWorkStatus)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	client, err := Publisher.TOPIC_EMPLOYEERESIGNED.Get()
	if err != nil {
		Logging.LogError(err, r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}

	// Get Subscriptions from Topic
	subs, err := client.Subscriptions()
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}

	// Get Old Data
	oldData, err := Services.Employee.BrowseEmployeeDetail(Dto.EmployeeDataDTO{ID: params.ID})
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	state := Dapr.GetConfigState(subs, "", "", oldData)
	fmt.Printf("%+v \n", state)
	err = Dapr.StatePublish(state)
	if err != nil {
		Logging.LogError(err.Error(), r)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	_, err = client.Publish(Publisher.Publisher{
		Key:               state.StateFmt.Key,
		StateStore:        state.StateStore,
		JWTDocodedPayload: contextMap,
		Data:              result,
	})

	if err != nil {
		Logging.LogError(err, r)
		Dapr.StateDelete(state.StateStore, state.StateFmt.Key)
		tx.Rollback()
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorPubSub)
		return
	}

	err = tx.Commit()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Employee Check Mobile Phone Number
// @Description  Sample Parameter: `082134567890`
// @Accept       json
// @Produce      json
// @Param        mobilePhoneNo  path      string                       true  "mobilePhoneNo"
// @Success      200            {object}  Response.RespResultStruct{}  "OK"
// @Failure      500            {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /check-mobile-phone-number/{mobilePhoneNumber} [get]
func (c Controller) CheckExistMobilePhoneNumber(w http.ResponseWriter, r *http.Request) {
	mobilePhoneNumber := chi.URLParam(r, "mobilePhoneNumber")
	result, err := Services.Employee.CheckExistMobilePhoneNumber(mobilePhoneNumber)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Employee Check Email
// @Description  Sample Parameter: `putra@ssss.com`
// @Accept       json
// @Produce      json
// @Param        email  path      string                       true  "email"
// @Success      200    {object}  Response.RespResultStruct{}  "OK"
// @Failure      500    {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /check-email/{email} [get]
func (c Controller) CheckExistEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")
	result, err := Services.Employee.CheckExistEmail(email)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Employee
// @Summary      Employee Check ID
// @Description  Sample Parameter: `1`
// @Accept       json
// @Produce      json
// @Param        id   path      int                          true  "id"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /check-employee-id/{id} [get]
func (c Controller) CheckExistEmployeeID(w http.ResponseWriter, r *http.Request) {

	var employeeId int64
	var err error

	idEmployee := chi.URLParam(r, "id")
	employeeId, err = strconv.ParseInt(idEmployee, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Employee.CheckExistEmployeeID(employeeId)
	if err != nil {
		if strings.Contains(err.Error(), mongo.ErrNoDocuments.Error()) {
			Response.ResponseError(w, nil, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

func (c Controller) CountEmployees(w http.ResponseWriter, r *http.Request) {
	count, err := Services.Employee.CountEmployees()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, count, Constant.StatusOKJson)
}
