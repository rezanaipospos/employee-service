package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Response"
	validator "EmployeeService/Library/Helper/Validator"
	"EmployeeService/Library/Logging"
	Services "EmployeeService/Services/leaveBalancePolicy"
	"bytes"
	"database/sql"
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

type LeaveBalancePolicy interface {
	SelectLeaveBalancePolicy(w http.ResponseWriter, r *http.Request)
	SaveLeaveBalancePolicy(w http.ResponseWriter, r *http.Request)
	UpdateLeaveBalancePolicy(w http.ResponseWriter, r *http.Request)
	CheckValidationCompanyName(w http.ResponseWriter, r *http.Request)
}

// @Tags      Leave Balance Policy
// @Summary   Select All Data
// @Accept    json
// @Produce   json
// @Param     companyName  query     string                       false  "companyName"
// @Success   200          {object}  Response.RespResultStruct{}  "OK"
// @Failure   500          {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /leave-balance/policy [get]
func (c Controller) SelectLeaveBalancePolicy(w http.ResponseWriter, r *http.Request) {

	var params Dto.LeaveBalancePolicyDTO
	var decoder = schema.NewDecoder()
	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.LeaveBalancePolicy.SelectLeaveBalancePolicy(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags         Leave Balance Policy
// @Summary      Save Data
// @Description  Sample Payload: `{"companyId":2,"companyName":"PT. SSSS","autoCutLeaveWeekly":true,"leaveBalanceAccumulation":true,"leaveBalanceBonusByLenghtOfWork":true,"leaveBalanceBonusList":[{"workingPeriodStart":5,"workingPeriodEnd":15,"reward":5},{"workingPeriodStart":10,"workingPeriodEnd":15,"reward":5}]}`
// @Accept       json
// @Produce      json
// @Param        "payload" body Dto.LeaveBalancePolicyDTO true "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance/policy [post]
func (c Controller) SaveLeaveBalancePolicy(w http.ResponseWriter, r *http.Request) {

	var params Dto.LeaveBalancePolicyDTO
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	// Check Company jika duplikat
	companyIDExist, err := Services.LeaveBalancePolicy.ValidationDuplicateCompany(params.CompanyId)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if companyIDExist {
		Response.ResponseError(w, err, Constant.StatusBadRequestCompanyAlreadyExists)
		return
	}

	params.CreatedBy = EmployeeName
	params.CreatedTime = time.Now()

	pagesJson, err := json.Marshal(params.LeaveBalanceBonusList)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}
	LeaveBalanceBonusList := bytes.NewBuffer(pagesJson).String()

	err = Services.LeaveBalancePolicy.SaveLeaveBalancePolicy(LeaveBalanceBonusList, params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, nil, Constant.StatusOKJson)
}

// @Tags         Leave Balance Policy
// @Summary      Update Data
// @Description  Sample Payload: `{"autoCutLeaveWeekly":false,"leaveBalanceAccumulation":false,"leaveBalanceBonusByLenghtOfWork":false,"leaveBalanceBonusList":[{"workingPeriodStart":10,"workingPeriodEnd":10,"reward":10},{"workingPeriodStart":11,"workingPeriodEnd":11,"reward":11}]}`
// @Accept       json
// @Produce      json
// @Param        id  path  int64  true  "Id"
// @Param        "payload" body Dto.LeaveBalancePolicyDTO true "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /leave-balance/policy/{id} [put]
func (c Controller) UpdateLeaveBalancePolicy(w http.ResponseWriter, r *http.Request) {

	var params Dto.LeaveBalancePolicyDTO
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}

	contextMap := r.Context().Value("claims_value").(map[string]interface{})
	EmployeeName := fmt.Sprintf("%v", contextMap["EmployeeName"])

	idlbp := chi.URLParam(r, "id")
	params.ID, err = strconv.ParseInt(idlbp, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	if errValidator := validator.Validate(params); errValidator != "" {
		Response.ResponseError(w, errors.New(errValidator), Constant.StatusBadRequestInvalidData)
		return
	}

	params.ModifiedBy = EmployeeName
	params.ModifiedTime = time.Now()

	pagesJson, err := json.Marshal(params.LeaveBalanceBonusList)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestJson)
		return
	}
	LeaveBalanceBonusList := bytes.NewBuffer(pagesJson).String()

	err = Services.LeaveBalancePolicy.UpdateLeaveBalancePolicy(LeaveBalanceBonusList, params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	Response.ResponseJson(w, nil, Constant.StatusOKJson)
}

func (c Controller) CheckValidationCompanyName(w http.ResponseWriter, r *http.Request) {
	companyName := chi.URLParam(r, "companyName")
	result, err := Services.LeaveBalancePolicy.CheckValidationCompanyName(companyName)
	if err != nil {
		if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
			Response.ResponseError(w, err, Constant.StatusBadRequestNotExist)
			return
		}
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}
