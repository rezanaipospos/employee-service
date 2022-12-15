package Controller

import (
	"EmployeeService/Constant"
	"EmployeeService/Library/Dapr"
	"EmployeeService/Library/Helper/Jwt"
	"EmployeeService/Library/Helper/Response"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type ControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
	DaprState(w http.ResponseWriter, r *http.Request)
	DaprStateGet(w http.ResponseWriter, r *http.Request)
	DaprStateDelete(w http.ResponseWriter, r *http.Request)
	SwaggerController(w http.ResponseWriter, r *http.Request)
	Employee
	LeaveBalance
	Transfer
	Dashboard
	LeaveBalancePolicy
}

type Controller struct {
	ControllerInterface
	Jwt Jwt.JwtServices
}

// @Tags     Sample
// @Accept   json
// @Produce  json
// @Success  200  {object}  Response.RespResultStruct{}  "OK"
// @Failure  500  {object}  Response.RespErrorStruct{}   "desc"
// @Router   /ping [get]
func (c Controller) Ping(w http.ResponseWriter, r *http.Request) {
	Response.ResponseJson(w, true, Constant.StatusOKJson)
}

// @Tags     Sample
// @Accept   json
// @Produce  json
// @Param        stateStore           path     string                       true  "stateStore"
// @Param        payload           body     map[string]interface{}                       true  "payload"
// @Success  200  {object}  Response.RespResultStruct{}  "OK"
// @Failure  500  {object}  Response.RespErrorStruct{}   "desc"
// @Router   /dapr-state/{stateStore} [post]
func (c Controller) DaprState(w http.ResponseWriter, r *http.Request) {
	var state map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&state)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
		return
	}
	services := []Dapr.StateService{
		{Name: "projects/s4-super-apps-355506/subscriptions/CompanyStructure-EmployeePersonalInfoUpdated", Status: "progress"},
		{Name: "projects/s4-super-apps-355506/subscriptions/Employee-EmployeePersonalInfoUpdated", Status: "progress"},
	}
	stateStore := chi.URLParam(r, "stateStore")
	log.Info().Msgf("state store is:%s", stateStore)
	stateCfg := Dapr.GetConfigState(nil, stateStore, "", state, services...)
	log.Info().Msgf("stateCfg is:%#v", stateCfg)
	err = Dapr.StatePublish(stateCfg)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
		return
	}
	// bytesState, err := Dapr.StateGet(stateCfg.StateStore, stateCfg.StateFmt.Key)
	// if err != nil {
	// 	Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
	// 	return
	// }
	// log.Info().Msgf("bytesState is:%s", string(bytesState))
	// json.Unmarshal(bytesState, &state)
	// log.Info().Msgf("state is:%#v", state)
	// err = Dapr.StateDelete(stateCfg.StateStore, stateCfg.StateFmt.Key)
	// if err != nil {
	// 	Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
	// 	return
	// }
	Response.ResponseJson(w, state, Constant.StatusOKJson)
}

// @Tags     Sample
// @Accept   json
// @Produce  json
// @Param        stateStore           path     string                       true  "stateStore"
// @Param        key           path     string                       true  "key"
// @Success  200  {object}  Response.RespResultStruct{}  "OK"
// @Failure  500  {object}  Response.RespErrorStruct{}   "desc"
// @Router   /dapr-state/{stateStore}/{key} [get]
func (c Controller) DaprStateGet(w http.ResponseWriter, r *http.Request) {
	var state interface{}
	stateStore := chi.URLParam(r, "stateStore")
	key := chi.URLParam(r, "key")
	bytesState, err := Dapr.StateGet(stateStore, key)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
		return
	}
	log.Info().Msgf("bytesState is:%s", string(bytesState))
	json.Unmarshal(bytesState, &state)

	Response.ResponseJson(w, state, Constant.StatusOKJson)
}

// @Tags     Sample
// @Accept   json
// @Produce  json
// @Param        stateStore           path     string                       true  "stateStore"
// @Param        key           path     string                       true  "key"
// @Success  200  {object}  Response.RespResultStruct{}  "OK"
// @Failure  500  {object}  Response.RespErrorStruct{}   "desc"
// @Router   /dapr-state/{stateStore}/{key} [delete]
func (c Controller) DaprStateDelete(w http.ResponseWriter, r *http.Request) {
	// var state map[string]interface{}
	stateStore := chi.URLParam(r, "stateStore")
	key := chi.URLParam(r, "key")
	err := Dapr.StateDelete(stateStore, key)
	if err != nil {
		Response.ResponseError(w, err, Constant.StatusInternalServerErrorDapr)
		return
	}

	Response.ResponseJson(w, true, Constant.StatusOKJson)
}
