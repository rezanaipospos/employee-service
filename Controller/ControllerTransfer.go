package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Response"
	validator "EmployeeService/Library/Helper/Validator"
	"EmployeeService/Library/Logging"
	"EmployeeService/Library/PubSub/Publisher"
	Services "EmployeeService/Services/transfer"
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

type Transfer interface {
	TransferCreate(w http.ResponseWriter, r *http.Request)
	TransferDetail(w http.ResponseWriter, r *http.Request)
	TransferData(w http.ResponseWriter, r *http.Request)
}

// @Tags         Transfer
// @Summary      TransferCreate
// @Description  Sample Payload: `{"transferCode":"001","transferDate":"2022-04-01T07:48:48.243Z","employeeId":2,"employeeCode":"10.0002","machineId":1,"machineName":"nama mesin","fingerPrintId":"finger 1","oldEmployeeCode":"000.0002","oldCompanyName":"SSSS","oldLocationName":"Medan","oldDepartmentName":"Acconounting","oldSectionName":"Bagian 1","oldPositionName":"Staff IT","oldCompanyLocationCode":"101010","newEmployeeCode":"10.0002","newCompanyName":"S4","newLocationName":"Medan","newDepartmentName":"MIS / IT","newSectionName":"Bagian 3","newPositionName":"Dev Ops","newCompanyLocationCode":"10101010","reason":"Reason"}`
// @Accept       json
// @Produce      json
// @Param        "payload" body Dto.TransferDataDTO true "Example Payload"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /transfers [post]
// Keterangan : function ini digunakan untuk membuat data karyawan yang akan di mutasi
func (c Controller) TransferCreate(w http.ResponseWriter, r *http.Request) {

	var params Dto.TransferDataDTO
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

	exists, err := Services.Transfer.ValidationDuplicateDataTransfer(params.TransferCode)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	if exists {
		Response.ResponseError(w, err, Constant.StatusBadRequestAlreadyExists)
		return
	}

	tx, result, err := Services.Transfer.SaveTransfer(params)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	result.OldIdCompany = params.OldIdCompany
	result.OldIdlocation = params.OldIdlocation
	result.OldIdDepartment = params.OldIdDepartment
	result.OldIdSection = params.OldIdSection
	result.OldIdPosition = params.OldIdPosition
	result.NewIdCompany = params.NewIdCompany
	result.NewIdLocation = params.NewIdLocation
	result.NewIdDepartment = params.NewIdDepartment
	result.NewIdSection = params.NewIdSection
	result.NewIdPosition = params.NewIdPosition

	client, err := Publisher.TOPIC_EMPLOYEETRANSFERED.Get()
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

// @Tags         Transfer
// @Summary      TransferData
// @Description  Sample Parameter: `?sortColumn=id&sortOrder=desc&pageSize=0&pageNumber=0&employeeCode=empl&transferCode=12`
// @Accept       json
// @Produce      json
// @Param        sortColumn    query     string                       false  "sortColumn"
// @Param        sortOrder     query     string                       false  "sortOrder"
// @Param        pageSize      query     int64                        false  "pageSize"
// @Param        pageNumber    query     int64                        false  "pageNumber"
// @Param        transferCode  query     string                       false  "transferCode"
// @Param        employeeCode  query     string                       false  "employeeCode"
// @Param        employeeName  query     string                       false  "employeeName"
// @Success      200           {object}  Response.RespResultStruct{}  "OK"
// @Failure      500           {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /transfers [get]
// Keterangan : function ini digunakan untuk menampilakan data-data karyawan yang terdaftar untuk di mutasi
func (c Controller) TransferData(w http.ResponseWriter, r *http.Request) {
	var params Dto.TransferParams
	var decoder = schema.NewDecoder()

	err := decoder.Decode(&params, r.URL.Query())
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Transfer.DataTransfer(params)

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

// @Tags         Transfer
// @Summary      TransferDetail
// @Description  Sample Parameter: `3`
// @Accept       json
// @Produce      json
// @Param        id   path      int64                        true  "Id"
// @Success      200  {object}  Response.RespResultStruct{}  "OK"
// @Failure      500  {object}  Response.RespErrorStruct{}   "desc"
// @Security     Bearer
// @Router       /transfers/{id} [get]
// Keterangan : function ini digunakan untuk menampilakan detail data karyawan yang terdaftar untuk di mutasi
func (c Controller) TransferDetail(w http.ResponseWriter, r *http.Request) {
	var ID int64
	var err error

	transferStr := chi.URLParam(r, "id")
	ID, err = strconv.ParseInt(transferStr, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Transfer.DetailTransfer(ID)
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

func (c Controller) TransfersHardDelete(w http.ResponseWriter, r *http.Request) {
	var ID int64
	var err error

	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	res, err := Services.Transfer.HardDeleteTransfer(ID)
	if err != nil {
		Logging.LogError(err, r)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}
	Response.ResponseJson(w, res, Constant.StatusOKJson)
}
