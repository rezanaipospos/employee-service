package Controller

import (
	"EmployeeService/Library/Helper/Response"
	"EmployeeService/Repository/dashboard"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewEmployeeData_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	var resultResp []dashboard.NewEmployeeData

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  resultResp,
	}

	target := "/dahsboards/employee/new"
	req := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	controller.NewEmployeeData(w, req)

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
}

func TestTotalReligionSummary_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	var resultResp []dashboard.TotalReligionSummary

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  resultResp,
	}

	target := "/dahsboards/employee/religion/summary"
	req := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	controller.TotalReligionSummary(w, req)

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
}

func TestTotalWorkStatusSummary_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	var resultResp []dashboard.TotalWorkStatusSummary

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  resultResp,
	}

	target := "/dahsboards/employee/work-status/summary"
	req := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	controller.TotalWorkStatusSummary(w, req)

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
}

func TestTotalWillExpireEmployeeContract_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	var resultResp dashboard.TotalWillExpireEmployeeContract

	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  resultResp,
	}

	target := "/dahsboards/employee/work-status/will-expire/summary"
	req := httptest.NewRequest("GET", target, nil)
	w := httptest.NewRecorder()
	controller.TotalWillExpireEmployeeContract(w, req)

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
	assert.NotNil(t, Resp.RestResult)
}

func TestTotalEmployeeByLengthOfWork_InvalidParameter(t *testing.T) {
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

	req := httptest.NewRequest("GET", "/dashboards/employee/length-of-work/abc", nil)
	claims := map[string]interface{}{
		"Email":        "muhammad.nasrul@pt-ssss.com",
		"EmployeeID":   "01.18001855",
		"EmployeeName": "Muhammad Nasrul",
		"PhoneNo":      "081221922901",
		"UserID":       1,
	}
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("numberOfYear", "abc")
	claimsctx := context.WithValue(req.Context(), "claims_value", claims)
	req = req.WithContext(claimsctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.TotalEmployeeByLengthOfWork(w, req)

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

func TestTotalEmployeeByLengthOfWork_Success(t *testing.T) {
	var controller Controller
	var Resp Response.RespResultStruct
	expectResp := Response.RespResultStruct{
		RestTitle:   http.StatusText(http.StatusOK),
		RestStatus:  2000,
		RestMessage: "Success",
		RestResult:  true,
	}

	req := httptest.NewRequest("GET", fmt.Sprintf("/dashboards/employee/length-of-work/%d", 10), nil)
	// req := httptest.NewRequest("GET", "/attendanceMachines", nil)
	claims := map[string]interface{}{
		"Email":        "muhammad.nasrul@pt-ssss.com",
		"EmployeeID":   "01.18001855",
		"EmployeeName": "Muhammad Nasrul",
		"PhoneNo":      "081221922901",
		"UserID":       1,
	}
	ctx := context.WithValue(req.Context(), "claims_value", claims)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("numberOfYear", "10")
	claimsctx := context.WithValue(req.Context(), "claims_value", claims)
	req = req.WithContext(claimsctx)
	ctx = req.Context()
	req = req.WithContext(context.WithValue(ctx, chi.RouteCtxKey, rctx))
	w := httptest.NewRecorder()
	controller.TotalEmployeeByLengthOfWork(w, req)

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
