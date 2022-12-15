package Controller

import (
	"EmployeeService/Controller/Dto"
	"EmployeeService/Library/Helper/Response"
	"EmployeeService/Repository/transfer"
	Services "EmployeeService/Services/transfer"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

var resultTf transfer.Transfers
var claims = map[string]interface{}{
	"Email":        "muhammad.nasrul@pt-ssss.com",
	"EmployeeID":   "01.18001855",
	"EmployeeName": "Muhammad Nasrul",
	"PhoneNo":      "081221922901",
	"UserID":       1,
}

var payloadCreate = Dto.TransferDataDTO{
	TransferCode:           "12345",
	TransferDate:           time.Now(),
	EmployeeId:             71,
	EmployeeCode:           "800.800",
	OldEmployeeCode:        "oldEmployeeCode",
	OldIdCompany:           10,
	OldCompanyName:         "oldCompanyName",
	OldIdlocation:          10,
	OldLocationName:        "oldLocationName",
	OldIdDepartment:        10,
	OldDepartmentName:      "oldDepartmentName",
	OldIdSection:           10,
	OldSectionName:         "oldSectionName",
	OldIdPosition:          10,
	OldPositionName:        "oldPositionName",
	OldCompanyLocationCode: "oldCompanyLocationCode",
	NewEmployeeCode:        "111222333",
	NewIdCompany:           20,
	NewCompanyName:         "GOOGLE",
	NewIdLocation:          20,
	NewLocationName:        "USA",
	NewIdDepartment:        20,
	NewDepartmentName:      "DEPT IT",
	NewIdSection:           20,
	NewSectionName:         "IT",
	NewIdPosition:          20,
	NewPositionName:        "PROGRAMER",
	NewCompanyLocationCode: "2000",
	Reason:                 "reason",
}

func TestTransferCreate_BadJson(t *testing.T) {
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
	req := httptest.NewRequest("POST", "/transfer", nil)
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	w := httptest.NewRecorder()
	controller.TransferCreate(w, req.WithContext(ctx))

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

func TestTransferCreate_InvalidData(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct
	expectRestResultErr := []string{"error when validation, error:TransferCode required"}
	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4004,
		RestMessage:   "invalid data",
		RestResultErr: make([]interface{}, len(expectRestResultErr)),
	}
	for index, data := range expectRestResultErr {
		expectResp.RestResultErr[index] = data
	}
	var payloadCreate = Dto.TransferDataDTO{
		TransferCode:           "",
		TransferDate:           time.Now(),
		EmployeeId:             71,
		EmployeeCode:           "800.800",
		OldEmployeeCode:        "oldEmployeeCode",
		OldIdCompany:           10,
		OldCompanyName:         "oldCompanyName",
		OldIdlocation:          10,
		OldLocationName:        "oldLocationName",
		OldIdDepartment:        10,
		OldDepartmentName:      "oldDepartmentName",
		OldIdSection:           10,
		OldSectionName:         "oldSectionName",
		OldIdPosition:          10,
		OldPositionName:        "oldPositionName",
		OldCompanyLocationCode: "oldCompanyLocationCode",
		NewEmployeeCode:        "111222333",
		NewIdCompany:           20,
		NewCompanyName:         "GOOGLE",
		NewIdLocation:          20,
		NewLocationName:        "USA",
		NewIdDepartment:        20,
		NewDepartmentName:      "DEPT IT",
		NewIdSection:           20,
		NewSectionName:         "IT",
		NewIdPosition:          20,
		NewPositionName:        "PROGRAMER",
		NewCompanyLocationCode: "2000",
		Reason:                 "reason",
	}
	bytePayload, _ := json.Marshal(payloadCreate)

	req := httptest.NewRequest("POST", "/transfers", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	w := httptest.NewRecorder()
	controller.TransferCreate(w, req.WithContext(ctx))

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

func TestTransferCreate_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  nil,
	}
	bytePayload, _ := json.Marshal(payloadCreate)

	req := httptest.NewRequest("POST", "/transfers", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	w := httptest.NewRecorder()
	controller.TransferCreate(w, req.WithContext(ctx))

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
		json.Unmarshal(bytes, &resultTf)
	}
}

func TestTransferCreate_AlreadyExists(t *testing.T) {
	var controller Controller
	var Resp Response.RespErrorStruct

	expectResp := Response.RespErrorStruct{
		RestTitle:     http.StatusText(http.StatusBadRequest),
		RestStatus:    4003,
		RestMessage:   "data already exists",
		RestResultErr: nil,
	}
	bytePayload, _ := json.Marshal(payloadCreate)

	req := httptest.NewRequest("POST", "/transfers", bytes.NewReader(bytePayload))
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	w := httptest.NewRecorder()
	controller.TransferCreate(w, req.WithContext(ctx))

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

func TestTransferData_InvalidParameter(t *testing.T) {
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
	req := httptest.NewRequest("GET", "/transfers?id=asal", nil)
	w := httptest.NewRecorder()
	controller.TransferData(w, req)

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

func TestTransferData_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	var resultResp transfer.TransfersData

	var resultRespData []transfer.Transfers
	expectResultResp := transfer.TransfersData{
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

	fmt.Println("EmployeeCode : ", resultTf.EmployeeCode)

	target := fmt.Sprintf("/transfers?employeeCode=%s", resultTf.EmployeeCode)
	req := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	controller.TransferData(w, req)

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

func TestTransferDetail_InvalidParameter(t *testing.T) {
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

	req := httptest.NewRequest("GET", "/transfers?id=abc", nil)
	claimsctx := context.WithValue(req.Context(), "claims_value", claims)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", "abc")
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.TransferDetail(w, req)

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

func TestTransferDetail_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	fmt.Println("Employe ID:", resultTf.ID)
	req := httptest.NewRequest("GET", fmt.Sprintf("/transfers?id=%d", resultTf.ID), nil)
	claimsctx := context.WithValue(req.Context(), "claims_value", claims)
	rctx := chi.NewRouteContext()

	rctx.URLParams.Add("id", fmt.Sprint(resultTf.ID))
	req = req.WithContext(claimsctx)
	claimsctx = req.Context()
	req = req.WithContext(context.WithValue(claimsctx, chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	controller.TransferDetail(w, req)

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

func TestTransferHardDelete(t *testing.T) {
	var ID int64
	ID = resultTf.ID
	result, err := Services.Transfer.HardDeleteTransfer(ID)
	assert.Nil(t, err, "err must be nill")
	assert.NotNil(t, result, "result must be nill")
}
