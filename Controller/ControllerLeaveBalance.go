package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Response"
	validator "EmployeeService/Library/Helper/Validator"
	"EmployeeService/Library/Logging"
	leaveBalancePolicy "EmployeeService/Repository/leaveBalancePolicy"
	Services "EmployeeService/Services/leaveBalance"
	ServiceLBP "EmployeeService/Services/leaveBalancePolicy"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

type LeaveBalance interface {
	LeaveBalanceCreate(w http.ResponseWriter, r *http.Request)
	LeaveBalanceAdjusmentCreate(w http.ResponseWriter, r *http.Request)
	LeaveBalanceData(w http.ResponseWriter, r *http.Request)
	LeaveBalanceDetail(w http.ResponseWriter, r *http.Request)
	LeaveBalanceAdjustmentDetail(w http.ResponseWriter, r *http.Request)
	LeaveBalanceReset(w http.ResponseWriter, r *http.Request)
}

// @Tags         Leave Balance
// @Summary      Create Leave Balance Adjusment
// @Description  Sample Payload: `{"employeeId": 36,"employeeCode":"10.0001","startDate":"2022-02-10T07:48:48.243Z","endDate":"2022-02-10T07:48:48.243Z","type":"increase","quantity": 5,"reason":"Potongan Cuti 15 Hari"}`
// @Accept       json
// @Produce      json
// @Param        "payload" body Dto.LeaveBalanceAdjustmentDTO true "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance [post]
func (c Controller) LeaveBalanceAdjusmentCreate(w http.ResponseWriter, r *http.Request) {

	var params Dto.LeaveBalanceAdjustmentDTO
	var paramsLB Dto.LeaveBalanceDataDTO

	err := json.NewDecoder(r.Body).Decode(&params)
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	params.CreatedBy = EmployeeName
	params.CreatedTime = time.Now()

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	typeLeave := strings.ToLower(params.Type)
	if typeLeave == "increase" {
		params.Type = typeLeave

	} else if typeLeave == "decrease" {
		params.Type = typeLeave

	} else {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidType)
		return
	}

	result, err := Services.LeaveBalance.SaveLeaveBalanceAdjusment(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	paramsLB.EmployeeId = result.EmployeeId
	resultData, err := Services.LeaveBalance.SelectEmployeeLeaveBalance(paramsLB)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	StartBalances := resultData.StartBalances
	IncreaseBalance := resultData.IncreaseBalance
	DecreaseBalance := resultData.DecreaseBalance
	LastPeriodBalance := resultData.LastPeriodBalance

	if result.Type == "increase" {

		IncreaseBalance = result.Quantity
		paramsLB.IncreaseBalance = IncreaseBalance
		paramsLB.DecreaseBalance = DecreaseBalance

	} else if result.Type == "decrease" {

		DecreaseBalance = result.Quantity
		paramsLB.DecreaseBalance = DecreaseBalance
		paramsLB.IncreaseBalance = IncreaseBalance
	}

	CurrentBalance := (StartBalances + IncreaseBalance + LastPeriodBalance) - DecreaseBalance
	paramsLB.CurrentBalance = CurrentBalance
	paramsLB.EmployeeId = result.EmployeeId

	err = Services.LeaveBalance.UpdateLeaveBalance(paramsLB)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Leave Balance
// @Summary      Show Leave Balance Data
// @Description  Sample Parameter: `?sortColumn=id&sortOrder=desc&pageSize=20?companyName=com&locationName=loc&departmentName=dept`
// @Accept       json
// @Produce      json
// @Param        sortColumn      query     string                       false  "sortColumn"
// @Param        sortOrder       query     string                       false  "sortOrder"
// @Param        pageSize        query     int64                        false  "pageSize"
// @Param        pageNumber      query     int64                        false  "pageNumber"
// @Param        companyName     query     string                       false  "companyName"
// @Param        locationName    query     string                       false  "locationName"
// @Param        departmentName  query     string                       false  "departmentName"
// @Success      200             {object}  Response.RespResultStruct{}  "OK"
// @Failure      500             {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance [get]
// Keterangan : function ini digunakan untuk mmenampilkan data cuti karyawan -> bisa lebih dari 1
func (c Controller) LeaveBalanceData(w http.ResponseWriter, r *http.Request) {
	var params Dto.LeaveBalanceParams
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.LeaveBalance.DataLeaveBalance(params)

	if err != nil {
		if strings.Contains(err.Error(), Constant.ErrorNoRows) {
			Response.ResponseJson(w, nil, Constant.StatusOKJson)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Leave Balance
// @Summary      Show Leave Balance Detail Data
// @Description  Sample Parameter: `2`
// @Accept       json
// @Produce      json
// @Param        employeeId  path      int64                        true  "employeeId"
// @Success      200         {object}  Response.RespResultStruct{}  "OK"
// @Failure      500         {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance/{employeeId} [get]
// Keterangan : function ini digunakan untuk menampilkan detail data cuti karyawan -> hanya 1 data
func (c Controller) LeaveBalanceDetail(w http.ResponseWriter, r *http.Request) {
	var EmployeeId int64
	var err error

	LeaveBalanceStr := chi.URLParam(r, "employeeId")
	EmployeeId, err = strconv.ParseInt(LeaveBalanceStr, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.LeaveBalance.DetailLeaveBalance(EmployeeId)
	if err != nil {
		Logging.LogError(err, r)
		if strings.Contains(err.Error(), Constant.ErrorNoRows) {
			Response.ResponseJson(w, nil, Constant.StatusOKJson)
			return
		}
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Leave Balance
// @Summary      Show Leave Balance Adjustment Detail
// @Description  Sample Parameter: `123/2021`
// @Accept       json
// @Produce      json
// @Param        employeeId  path      int64                        true  "employeeId"
// @Param        year        path      int64                        true  "year"
// @Success      200         {object}  Response.RespResultStruct{}  "OK"
// @Failure      500         {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance/{employeeId}/{year}/history [get]
// Keterangan : function ini digunakan untuk menampilkan Detail data cuti karyawan berdasarkan tahun
func (c Controller) LeaveBalanceAdjustmentDetail(w http.ResponseWriter, r *http.Request) {
	var Tahun, EmployeeId int64
	var err error

	employeeID := chi.URLParam(r, "employeeId")
	EmployeeId, err = strconv.ParseInt(employeeID, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	year := chi.URLParam(r, "year")
	Tahun, err = strconv.ParseInt(year, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.LeaveBalance.DetailLeaveBalanceAdjustment(Tahun, EmployeeId)
	if err != nil {
		Logging.LogError(err, r)
		if strings.Contains(err.Error(), Constant.ErrorNoRows) {
			Response.ResponseJson(w, nil, Constant.StatusOKJson)
			return
		}
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags      Leave Balance
// @Tags      Leave Balance
// @Summary   Reset Leave Balance
// @Accept    json
// @Produce   json
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /leave-balance/reset-leave-balance [get]
// Penjadwalan Untuk mereser Saldo Cuti Pegawai apabila tanggal hari ini dan tanggal expired sama
// Maka data berkaitan akan diubah status aktif nya menjadi false
// dan akan di tambahkan data baru dengan saldo cuti yang sudah di reset.
// Scheduler terletak di library/Scheduler/Scheduler.go
func (c Controller) LeaveBalanceReset(w http.ResponseWriter, r *http.Request) {

	var params Dto.LeaveBalanceDataDTO
	var bonusLists []leaveBalancePolicy.BonusList
	var bonus int64
	var lastPeriodBalance int64

	// tampilkan seluruh data yang status cuti nya aktif dan expired date nya hari ini
	resultDataActive, err := Services.LeaveBalance.SelectLeaveBalanceActive()
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	for _, res := range resultDataActive {
		// Update data cuti is_active => false
		err = Services.LeaveBalance.UpdateLeaveBalanceActive(res.ID)
		if err != nil {
			Logging.LogError(err, r)
			Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
			return
		}

		yearJoinDate := res.JoinDate.Year()
		yearNow := time.Now().Year()
		workingDay := yearNow - yearJoinDate

		checkDataExist, err := ServiceLBP.LeaveBalancePolicy.CheckValidationCompanyName(res.CompanyName)
		if err != nil {
			Logging.LogError(err, r)
			Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
			return
		}

		if checkDataExist.LeaveBalanceBonusByLenghtOfWork {
			jsonData := checkDataExist.LeaveBalanceBonusList
			json.Unmarshal([]byte(jsonData), &bonusLists)
			for i := range bonusLists {
				if int64(workingDay) >= bonusLists[i].WorkingPeriodStart && int64(workingDay) <= bonusLists[i].WorkingPeriodEnd {
					bonus = bonusLists[i].Reward
				} else {
					continue
				}
			}
		}

		if checkDataExist.LeaveBalanceAccumulation {
			lastPeriodBalance = res.CurrentBalance
		}

		params.EmployeeId = res.EmployeeId
		params.EmployeeCode = res.EmployeeCode
		params.CompanyName = res.CompanyName
		params.LocationName = res.LocationName
		params.DepartmentName = res.DepartmentName
		params.StartBalances = 12
		params.IncreaseBalance = bonus
		params.DecreaseBalance = 0
		params.LastPeriodBalance = lastPeriodBalance
		params.CurrentBalance = 0
		params.IsActive = true
		resPeriod, _ := strconv.Atoi(res.Period)
		params.Period = strconv.Itoa(resPeriod + 1)
		params.JoinDate = res.JoinDate
		params.ExpiredDate = time.Now().AddDate(1, 0, 0)
		params.CreatedBy = "System"
		params.CreatedTime = time.Now()
		params.Deleted = false
		err = Services.LeaveBalance.SaveLeaveBalance(params)
		if err != nil {
			Logging.LogError(err, r)
			Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
			return
		}

	}
	Response.ResponseJson(w, resultDataActive, Constant.StatusOKJson)
}

// Dieksekusi setelah EmployeeAdd dijalankan pada service employee -> Untuk menambahkan data saldo cuti pertama kali
func (c Controller) LeaveBalanceCreate(w http.ResponseWriter, r *http.Request) {
	var params Dto.LeaveBalanceDataDTO
	err := json.NewDecoder(r.Body).Decode(&params)
	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	params.CreatedBy = EmployeeName
	params.CreatedTime = time.Now()

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidData)
		return
	}

	exists, err := Services.LeaveBalance.ValidationDuplicateDataLeaveBalance(params.EmployeeId)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if exists {
		Response.ResponseError(w, err, Constant.StatusBadRequestAlreadyExists)
		return
	}

	result := Services.LeaveBalance.SaveLeaveBalance(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

func (c Controller) DepartmentHardDelete(w http.ResponseWriter, r *http.Request) {
	var EmployeeId int64
	var err error

	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	res, err := Services.LeaveBalance.HardDeleteLeaveBalance(EmployeeId)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	Response.ResponseJson(w, res, Constant.StatusOKJson)
}
