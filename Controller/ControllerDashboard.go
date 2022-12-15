package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Library/Helper/Response"
	"EmployeeService/Library/Logging"
	Services "EmployeeService/Services/dashboard"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type Dashboard interface {
	NewEmployeeData(w http.ResponseWriter, r *http.Request)
	TotalReligionSummary(w http.ResponseWriter, r *http.Request)
	TotalWorkStatusSummary(w http.ResponseWriter, r *http.Request)
	TotalWillExpireEmployeeContract(w http.ResponseWriter, r *http.Request)
	TotalEmployeeByLengthOfWork(w http.ResponseWriter, r *http.Request)
}

// @Tags      Dashboards
// @Summary   NewEmployeeData
// @Accept    json
// @Produce   json
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /dashboards/employee/new [get]
// Keterangan : function ini digunakan untuk Menampikan 10 Data karyawn yang terbaru
func (c Controller) NewEmployeeData(w http.ResponseWriter, r *http.Request) {

	var err error
	result, err := Services.Dashboard.NewEmployeeData()
	if err != nil {
		Logging.LogError(err, nil)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags      Dashboards
// @Summary   TotalReligionSummary
// @Accept    json
// @Produce   json
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /dashboards/employee/religion/summary [get]
// Keterangan : function ini digunakan untuk menampilkan jumlah rekap data agama semua karyawan yang ada
func (c Controller) TotalReligionSummary(w http.ResponseWriter, r *http.Request) {
	var err error

	result, err := Services.Dashboard.TotalReligionSummary()
	if err != nil {
		if strings.Contains(err.Error(), Constant.ErrorNoRows) {
			Response.ResponseJson(w, nil, Constant.StatusOKJson)
			return
		}
		Logging.LogError(err, nil)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags      Dashboards
// @Summary   TotalWorkStatusSummary
// @Accept    json
// @Produce   json
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /dashboards/employee/work-status/summary [get]
// Keterangan : function ini digunakan untuk menampilkan jumlah data status kerja semua karyawan
func (c Controller) TotalWorkStatusSummary(w http.ResponseWriter, r *http.Request) {
	var err error

	result, err := Services.Dashboard.TotalWorkStatusSummary()
	if err != nil {
		if strings.Contains(err.Error(), Constant.ErrorNoRows) {
			Response.ResponseJson(w, nil, Constant.StatusOKJson)
			return
		}
		Logging.LogError(err, nil)
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDB)
		return
	}

	Response.ResponseJson(w, result, Constant.StatusOKJson)
}

// @Tags      Dashboards
// @Summary   TotalWillExpireEmployeeContract
// @Accept    json
// @Produce   json
// @Success   200  {object}  Response.RespResultStruct{}  "OK"
// @Failure   500  {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /dashboards/employee/work-status/will-expire/summary [get]
// Keterangan : function ini digunakan untuk menampilkan jumlah data kontrak kerja yang akan expired semua karyawan
func (c Controller) TotalWillExpireEmployeeContract(w http.ResponseWriter, r *http.Request) {
	var err error

	result, err := Services.Dashboard.TotalWillExpireEmployeeContract()
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

// @Tags      Dashboards
// @Summary   TotalEmployeeByLengthOfWork
// @Accept    json
// @Produce   json
// @Param     numberOfYear  path      int64                        true  "numberOfYear"
// @Success   200           {object}  Response.RespResultStruct{}  "OK"
// @Failure   500           {object}  Response.RespErrorStruct{}   "desc"
// @Security  Bearer
// @Router    /dashboards/employee/length-of-work/{numberOfYear} [get]
// Keterangan : function ini digunakan untuk menampilkan jumlah data kontrak kerja yang akan expired semua karyawan
func (c Controller) TotalEmployeeByLengthOfWork(w http.ResponseWriter, r *http.Request) {
	var err error
	var numberOfYear int64

	numberOfYearStr := chi.URLParam(r, "numberOfYear")
	numberOfYear, err = strconv.ParseInt(numberOfYearStr, 10, 64)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusBadRequestInvalidParameter)
		return
	}

	result, err := Services.Dashboard.TotalEmployeeByLengthOfWork(numberOfYear)
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
