package Controller

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Response"
	"EmployeeService/Repository/employee"
	"EmployeeService/Repository/leaveBalance"
	Services "EmployeeService/Services/leaveBalance"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var resultEmp employee.Employees
var resultAdj leaveBalance.LeaveBalanceAdjustment
var resultData leaveBalance.LeaveBalance

var now = time.Now()
var claimsEmployee = map[string]interface{}{
	"Email":        "muhammad.nasrul@pt-ssss.com",
	"EmployeeID":   "01.18001855",
	"EmployeeName": "Muhammad Nasrul",
	"PhoneNo":      "081221922901",
	"UserID":       1,
}

var payloadEmployee = Dto.EmployeeDataDTO{
	Code:                "10203041",
	Type:                "-",
	FingerPrintId:       "10203040",
	FaceId:              "10203040",
	MachineId:           10203040,
	MachineName:         "Mesin",
	DepartmentId:        20,
	DepartmentName:      "MIS / IT",
	SectionId:           1,
	SectionName:         "sectName",
	PositionId:          19,
	PositionName:        "Staff IT",
	CompanyId:           126,
	CompanyName:         "PT.PLN",
	LocationId:          63,
	LocationName:        "Medan",
	CompanyLocationCode: "SSSS",
	Parent:              219,
	ParentName:          "Juliana",
	IdentityNo:          "1271000524163",
	DrivingLicenseNo:    "150404020603",
	NpwpNo:              "0",
	Name:                "PutraAgastiya",
	PlaceOfBirth:        "Medan",
	DateOfBirth:         &now,
	Email:               "test@ssss.com",
	Address:             "Jalan Veteran",
	TemporaryAddress:    "Jalan Veteran",
	NeighbourHoodWardNo: "1020",
	UrbanName:           "Medan",
	SubDistrictName:     "Medan Timur",
	Religion:            "Buddha",
	MaritalStatus:       "Divorced",
	Citizen:             "Indonesia",
	Gender:              "Male",
	Ethnic:              "Etnic",
	MobilePhoneNo:       "0821789456",
	PhoneNo:             "0821789456",
	ShirtSize:           "XL",
	PantSize:            30,
	ShoeSize:            40,
	JoinDate:            &now,
	ResignDate:          nil,
	ResignReason:        "",
	Bank:                "Mandiri",
	BankAccountNo:       "1006001890",
	BankAccountName:     "Test",
	FamilyMobilePhoneNo: "555666999",
	WorkStatus:          "Kontrak",
	ProfilePhoto:        "/Img/test",
	ContractStart:       &now,
	ContractEnd:         &now,
	BpjsNo:              "0",
	JamsostekNo:         "0",
	JamsostekType:       "-",
	JamsostekBalance:    0,
	FamilyCardNo:        "0",
	ActiveStatus:        true,
	Extension:           "123456",
	PostalCode:          "1234",
}

var payloadUpdateEmpl = Dto.EmployeeDataDTO{
	Parent:              219,
	IdentityNo:          "identity_no-edit",
	DrivingLicenseNo:    "driving_license_no-edit",
	NpwpNo:              "npwpno-edit",
	Name:                "PutraAgastiyaEdit",
	PlaceOfBirth:        "place_of_birth-edit",
	DateOfBirth:         &now,
	Email:               "email@email.com-edit",
	Address:             "address-edit",
	TemporaryAddress:    "temporary_address-edit",
	NeighbourHoodWardNo: "neighbour_hood_ward_no-edit",
	UrbanName:           "urban_name-edit",
	SubDistrictName:     "sub_district_name-edit",
	Religion:            "religion-edit",
	MaritalStatus:       "marital_status-edit",
	Citizen:             "citizen-edit",
	Gender:              "Pria",
	Ethnic:              "ethnic-edit",
	MobilePhoneNo:       "mobile_phone_no-edit",
	PhoneNo:             "phone_no-edit",
	ShirtSize:           "L",
	PantSize:            10,
	ShoeSize:            10,
	Bank:                "bank-edit",
	BankAccountNo:       "bank_account_no-edit",
	BankAccountName:     "bank_account_name-edit",
	FamilyMobilePhoneNo: "family_mobile_phone_no-edit",
	BpjsNo:              "bpjs_no-edit",
	JamsostekNo:         "jamsostek_no-edit",
	JamsostekType:       "jamsostek_type-edit",
	JamsostekBalance:    20000,
	FamilyCardNo:        "familiy_card_no-edit",
	Extension:           "ubah",
	LicenseType:         "ubah",
}

var payloadWorkStatusUpdated = Dto.WorkStatusUpdateDTO{
	WorkStatus:             "Pegawai Tetap",
	ContractStart:          &now,
	ContractEnd:            &now,
	WorkStatusChangeReason: "alasan",
	WorkStatusChangeDate:   &now,
}

var payloadFingerUpdated = Dto.FingerprintUpdateDTO{
	FingerPrintId: "Update2 FingerPrintId",
	FaceId:        "Update2 FaceId",
	MachineId:     102,
}

var payloadResign = Dto.ResignStatusUpdateDTO{
	WorkStatus:   "Resign",
	ResignDate:   &now,
	ResignReason: "Pindah Negara",
}

func TestEmployeeCreate_BadJson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/employees", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeAdded(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeCreate_InvalidData(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"error when validation, error:Code max"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4004,
		RestMessage:   "invalid data",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	var now = time.Now()
	var payloadEmployee = Dto.EmployeeDataDTO{
		Code:                "1111111111111111111111111111111111111111111111111111111111111",
		Type:                "-",
		FingerPrintId:       "10203040",
		FaceId:              "10203040",
		MachineId:           10203040,
		MachineName:         "Mesin",
		DepartmentId:        20,
		DepartmentName:      "MIS / IT",
		SectionId:           1,
		SectionName:         "sectName",
		PositionId:          19,
		PositionName:        "Staff IT",
		CompanyId:           126,
		CompanyName:         "PT.MDP",
		LocationId:          63,
		LocationName:        "Medan",
		CompanyLocationCode: "SSSS",
		Parent:              61,
		ParentName:          "Juliana",
		IdentityNo:          "1271000524163",
		DrivingLicenseNo:    "150404020603",
		NpwpNo:              "0",
		Name:                "Staff",
		PlaceOfBirth:        "Medan",
		DateOfBirth:         &now,
		Email:               "test@ssss.com",
		Address:             "Jalan Veteran",
		TemporaryAddress:    "Jalan Veteran",
		NeighbourHoodWardNo: "1020",
		UrbanName:           "Medan",
		SubDistrictName:     "Medan Timur",
		Religion:            "Buddha",
		MaritalStatus:       "Divorced",
		Citizen:             "Indonesia",
		Gender:              "Male",
		Ethnic:              "Etnic",
		MobilePhoneNo:       "0821789456",
		PhoneNo:             "0821789456",
		ShirtSize:           "XL",
		PantSize:            30,
		ShoeSize:            40,
		JoinDate:            &now,
		ResignDate:          nil,
		ResignReason:        "",
		Bank:                "Mandiri",
		BankAccountNo:       "1006001890",
		BankAccountName:     "Test",
		FamilyMobilePhoneNo: "555666999",
		WorkStatus:          "Kontrak",
		ProfilePhoto:        "/Img/test",
		ContractStart:       &now,
		ContractEnd:         &now,
		BpjsNo:              "0",
		JamsostekNo:         "0",
		JamsostekType:       "-",
		JamsostekBalance:    0,
		FamilyCardNo:        "0",
		ActiveStatus:        true,
		Extension:           "123456",
		PostalCode:          "1234",
	}
	bytePayload, _ := json.Marshal(payloadEmployee)

	req := httptest.NewRequest("POST", "/employees", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeAdded(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeCreate_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadEmployee)

	req := httptest.NewRequest("POST", "/employees", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeAdded(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultEmp)
	}
}

func TestEmployeeCreate_AlreadyExists(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct

	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4003,
		RestMessage:   "data already exists",
		RestResultErr: nil,
	}

	bytePayload, _ := json.Marshal(payloadEmployee)
	req := httptest.NewRequest("POST", "/employees", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeAdded(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	// DATA DUPLICATE RestResultErr IS NIL
	if assert.Nil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeUpdate_BadJson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/employees", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeePersonalInfoUpdated(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeUpdate_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	bytePayload, _ := json.Marshal(payloadEmployee)

	req := httptest.NewRequest("POST", "/employees?id=abc", bytes.NewReader(bytePayload))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	claimsctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeePersonalInfoUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeUpdate_InvalidData(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"error when validation, error:Code max"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4004,
		RestMessage:   "invalid data",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	var now = time.Now()
	var payloadEmployee = Dto.EmployeeDataDTO{
		Code:                "1111111111111111111111111111111111111111111111111111111111111111",
		Type:                "-",
		FingerPrintId:       "10203040",
		FaceId:              "10203040",
		MachineId:           10203040,
		MachineName:         "Mesin",
		DepartmentId:        20,
		DepartmentName:      "MIS / IT",
		SectionId:           1,
		SectionName:         "sectName",
		PositionId:          19,
		PositionName:        "Staff IT",
		CompanyId:           126,
		CompanyName:         "PT.MDP",
		LocationId:          63,
		LocationName:        "Medan",
		CompanyLocationCode: "SSSS",
		Parent:              61,
		ParentName:          "Juliana",
		IdentityNo:          "1271000524163",
		DrivingLicenseNo:    "150404020603",
		NpwpNo:              "0",
		Name:                "Staff",
		PlaceOfBirth:        "Medan",
		DateOfBirth:         &now,
		Email:               "test@ssss.com",
		Address:             "Jalan Veteran",
		TemporaryAddress:    "Jalan Veteran",
		NeighbourHoodWardNo: "1020",
		UrbanName:           "Medan",
		SubDistrictName:     "Medan Timur",
		Religion:            "Buddha",
		MaritalStatus:       "Divorced",
		Citizen:             "Indonesia",
		Gender:              "Male",
		Ethnic:              "Etnic",
		MobilePhoneNo:       "0821789456",
		PhoneNo:             "0821789456",
		ShirtSize:           "XL",
		PantSize:            30,
		ShoeSize:            40,
		JoinDate:            &now,
		ResignDate:          nil,
		ResignReason:        "",
		Bank:                "Mandiri",
		BankAccountNo:       "1006001890",
		BankAccountName:     "Test",
		FamilyMobilePhoneNo: "555666999",
		WorkStatus:          "Kontrak",
		ProfilePhoto:        "/Img/test",
		ContractStart:       &now,
		ContractEnd:         &now,
		BpjsNo:              "0",
		JamsostekNo:         "0",
		JamsostekType:       "-",
		JamsostekBalance:    0,
		FamilyCardNo:        "0",
		ActiveStatus:        true,
		Extension:           "123456",
		PostalCode:          "1234",
	}
	bytePayload, _ := json.Marshal(payloadEmployee)

	req := httptest.NewRequest("POST", "/employees?id=3", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(3))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeePersonalInfoUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeUpdate_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadUpdateEmpl)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/employees/%d", resultEmp.ID), bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeePersonalInfoUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultEmp)
		assert.Equal(t, payloadUpdateEmpl.Name, resultEmp.Name, fmt.Sprintf("Name should be equal %s", resultEmp.Name))
	}
}

func TestEmployeeBrowse_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"schema: error converting value for \"id\""}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("GET", "/employees?id=asal", nil)
	w := httptest.NewRecorder()
	controller.EmployeeBrowse(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeBrowse_Success(t *testing.T) {
	TimeDuration := time.Duration(5) * time.Second

	f := func() {
		var controller Controller
		var Resp Response.RespResultStruct
		var resultResp employee.EmployeesData

		var resultRespData []employee.GetEmployees
		expectResultResp := employee.EmployeesData{
			RecordsTotal: 1,
			HasReachMax:  true,
			Data:         resultRespData,
		}
		expectResp := Response.RespResultStruct{
			RestTitle:   http.StatusText(http.StatusOK),
			RestStatus:  2000,
			RestMessage: "Success",
			RestResult:  expectResultResp,
		}

		target := fmt.Sprintf("/employees?code=%s", resultEmp.Code)
		req := httptest.NewRequest("GET", target, nil)
		w := httptest.NewRecorder()
		controller.EmployeeBrowse(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		err = json.Unmarshal(body, &Resp)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		byteData, _ := json.Marshal(Resp.RestResult)
		err = json.Unmarshal(byteData, &resultResp)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
		assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
		assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

		if assert.NotNil(t, Resp.RestResult) {
			assert.Equal(t, expectResultResp.RecordsTotal, resultResp.RecordsTotal)
			assert.Equal(t, expectResultResp.HasReachMax, resultResp.HasReachMax)
		}
	}

	Timer := time.AfterFunc(TimeDuration, f)
	defer Timer.Stop()
}

func TestEmployeeDetail_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("GET", "/employees?id=abc", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeBrowseDetail(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeSubordinatesData_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("GET", "/employees/abc/subordinates", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.SubordinatesData(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeSubordinatesData_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("/employees/%d/subordinates", 219), nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(219))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.SubordinatesData(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))
	assert.NotNil(t, Resp.RestResult)
}

func TestEmployeeDetailSubordinatesData_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("GET", "/employees/abc/subordinates/abc", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_parent", "abc")
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.DetailSubordinatesData(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeDetailSubordinatesData_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	resultEmp.Parent = 219
	// resultEmp.ID = 228

	log.Println(resultEmp.ID)

	req := httptest.NewRequest("GET", fmt.Sprintf("/employees/%d/subordinates/%d", resultEmp.Parent, resultEmp.ID), nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id_parent", fmt.Sprint(resultEmp.Parent))
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.DetailSubordinatesData(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))
	assert.NotNil(t, Resp.RestResult)
}

func TestEmployeeWorkStatusUpdated_Badjson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/employees/id/workstatus", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeWorkStatusUpdated(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeWorkStatusUpdated_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	bytePayload, _ := json.Marshal(payloadWorkStatusUpdated)

	req := httptest.NewRequest("POST", "/employees/abc/workstatus", bytes.NewReader(bytePayload))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	claimsctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeWorkStatusUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeWorkStatusUpdated_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadWorkStatusUpdated)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/employees/%d/workstatus", resultEmp.ID), bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeWorkStatusUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultEmp)
		assert.Equal(t, payloadUpdateEmpl.Name, resultEmp.Name, fmt.Sprintf("Name should be equal %s", resultEmp.Name))
	}
}

func TestEmployeeFingerUpdated_Badjson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/employees/id/fingerprint", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeFingerUpdated(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeFingerUpdated_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	bytePayload, _ := json.Marshal(payloadFingerUpdated)

	req := httptest.NewRequest("POST", "/employees/abc/fingerprint", bytes.NewReader(bytePayload))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	claimsctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeFingerUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeFingerUpdated_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadFingerUpdated)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/employees/%d/fingerprint", resultEmp.ID), bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeFingerUpdated(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultEmp)
		assert.Equal(t, payloadUpdateEmpl.Name, resultEmp.Name, fmt.Sprintf("Name should be equal %s", resultEmp.Name))
	}
}

func TestEmployeeResign_BadJson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/employees/3/resign", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.EmployeeResigned(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeResign_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	bytePayload, _ := json.Marshal(payloadResign)

	req := httptest.NewRequest("POST", "/employees/abc/resign", bytes.NewReader(bytePayload))
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	claimsctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeResigned(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeResign_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadFingerUpdated)

	req := httptest.NewRequest("PUT", fmt.Sprintf("/employees/%d/fingerprint", resultEmp.ID), bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeResigned(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultEmp)
		assert.Equal(t, payloadUpdateEmpl.Name, resultEmp.Name, fmt.Sprintf("Name should be equal %s", resultEmp.Name))
	}
}

/////////////////////////////////////////////////////Leave Balance/////////////////////////////////////////////////////
func TestLeaveBalanceData_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"schema: error converting value for \"id\""}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("GET", "/leaveBalance?id=str", nil)
	w := httptest.NewRecorder()
	controller.LeaveBalanceData(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceData_Success(t *testing.T) {
	TimeDuration := time.Duration(5) * time.Second

	f := func() {
		var controller Controller
		var Resp Response.RespResultStruct
		var resultResp leaveBalance.LeaveBalanceData

		var resultRespData []leaveBalance.LeaveBalance
		expectResultResp := leaveBalance.LeaveBalanceData{
			RecordsTotal: 1,
			HasReachMax:  true,
			Data:         resultRespData,
		}
		expectResp := Response.RespResultStruct{
			RestTitle:   http.StatusText(http.StatusOK),
			RestStatus:  2000,
			RestMessage: "Success",
			RestResult:  expectResultResp,
		}

		fmt.Println("Result Data Company Name : ", resultData.CompanyName)

		target := fmt.Sprintf("/leaveBalance?companyName=%s", resultData.CompanyName)
		req := httptest.NewRequest("GET", target, nil)
		w := httptest.NewRecorder()
		controller.LeaveBalanceData(w, req)

		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		err = json.Unmarshal(body, &Resp)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		byteData, _ := json.Marshal(Resp.RestResult)
		err = json.Unmarshal(byteData, &resultResp)
		if err != nil {
			t.Error(err)
			t.Fail()
			return
		}
		assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
		assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
		assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

		if assert.NotNil(t, Resp.RestResult) {
			assert.Equal(t, expectResultResp.RecordsTotal, resultResp.RecordsTotal)
			assert.Equal(t, expectResultResp.HasReachMax, resultResp.HasReachMax)
		}
	}

	Timer := time.AfterFunc(TimeDuration, f)
	defer Timer.Stop()
}

func TestLeaveBalanceDetail_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("GET", "/leaveBalance?id=abc", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint("abc"))

	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.LeaveBalanceDetail(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceDetail_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("/leaveBalance?employeeId=%d", resultEmp.ID), nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.LeaveBalanceDetail(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))
	assert.NotNil(t, Resp.RestResult)
}

func TestLeaveBalanceAdjusmentCreate_BadJson(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"EOF"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4001,
		RestMessage:   "json error",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	req := httptest.NewRequest("POST", "/leaveBalance", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjusmentCreate(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceAdjusmentCreate_InvalidData(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"error when validation, error:EmployeeCode required"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4004,
		RestMessage:   "invalid data",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	payloadCreate := Dto.LeaveBalanceAdjustmentDTO{
		EmployeeCode: "",
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		Type:         "increase",
		Reason:       "Istirahat",
		EmployeeId:   123,
		Quantity:     6,
		LeaveId:      0,
	}
	bytePayload, _ := json.Marshal(payloadCreate)
	req := httptest.NewRequest("POST", "/leaveBalance", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjusmentCreate(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceAdjusmentCreate_InvalidType(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4005,
		RestMessage:   "Type can only increase or decrease",
		RestResultErr: nil,
	}
	payloadCreateInvalidType := Dto.LeaveBalanceAdjustmentDTO{
		Type:         "increase-asal",
		EmployeeId:   123,
		EmployeeCode: "12345",
		Reason:       "Istirahat",
		Quantity:     6,
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		LeaveId:      0,
	}
	bytePayload, _ := json.Marshal(payloadCreateInvalidType)
	req := httptest.NewRequest("POST", "/leaveBalance", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjusmentCreate(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.Nil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceAdjusmentCreate_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}

	var payloadCreateLB = Dto.LeaveBalanceAdjustmentDTO{
		EmployeeId:   resultEmp.ID,
		EmployeeCode: resultEmp.Code,
		Type:         "increase",
		Reason:       "Istirahat",
		Quantity:     6,
		StartDate:    time.Now(),
		EndDate:      time.Now(),
		LeaveId:      0,
	}
	bytePayload, _ := json.Marshal(payloadCreateLB)

	req := httptest.NewRequest("POST", "/leaveBalance", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjusmentCreate(w, req.WithContext(ctx))

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResult) {
		bytes, _ := json.Marshal(Resp.RestResult)
		json.Unmarshal(bytes, &resultAdj)
	}
}

func TestLeaveBalanceAdjustmentDetail_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("GET", "/leaveBalance/abc/abc", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()

	rctx.URLParams.Add("employeeId", fmt.Sprint("abc"))
	rctx.URLParams.Add("year", fmt.Sprint("abc"))

	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjustmentDetail(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestLeaveBalanceAdjustmentDetail_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}
	req := httptest.NewRequest("GET", fmt.Sprintf("/leaveBalance/%d/2021", resultAdj.EmployeeId), nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()

	rctx.URLParams.Add("employeeId", fmt.Sprint(resultAdj.EmployeeId))
	rctx.URLParams.Add("year", fmt.Sprint("2021"))

	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.LeaveBalanceAdjustmentDetail(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))
	assert.NotNil(t, Resp.RestResult)
}

func TestLeaveBalanceHardDelete(t *testing.T) {
	var params Dto.LeaveBalanceDataDTO
	params.EmployeeId = resultEmp.ID
	result, err := Services.LeaveBalance.HardDeleteLeaveBalance(params.EmployeeId)
	assert.Nil(t, err, "err must be nill")
	assert.NotNil(t, result, "result must be nill")
}

func TestEmployeeDelete_InvalidParameter(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"strconv.ParseInt: parsing \"abc\": invalid syntax"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    40011,
		RestMessage:   "invalid parameter",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}

	req := httptest.NewRequest("DELETE", "/employees/abc", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint("abc"))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeDeleted(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))

	if assert.NotNil(t, Resp.RestResultErr) {
		assert.Equal(t, expectResp.RestResultErr, Resp.RestResultErr)
	}
}

func TestEmployeeDelete_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/employees/%d", resultEmp.ID), nil)
	ctx := context.WithValue(req.Context(), "claims_value", claimsEmployee)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", fmt.Sprint(resultEmp.ID))
	req = req.WithContext(ctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.EmployeeDeleted(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	err = json.Unmarshal(body, &Resp)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	assert.Equal(t, expectResp.RestTitle, Resp.RestTitle, fmt.Sprintf("RestTitle should be equal '%s'", expectResp.RestTitle))
	assert.Equal(t, expectResp.RestStatus, Resp.RestStatus, fmt.Sprintf("RestStatus should be equal %d", expectResp.RestStatus))
	assert.Equal(t, expectResp.RestMessage, Resp.RestMessage, fmt.Sprintf("RestMessage should be equal '%s'", expectResp.RestMessage))
	assert.Equal(t, Resp.RestResult, true)
}
