package Dapr

import (
	"EmployeeService/Config"
	"EmployeeService/Constant"
	"EmployeeService/Library/Helper/Response"
	"EmployeeService/Library/Logging"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (c ConfigDapr) Initialize(env *Config.Environment) {
	log.Info().Msgf("env.App.ServiceGroup is %s", env.App.ServiceGroup)
	var okHost, okPort bool
	Dapr = ConfigDapr{
		ApiKey:     env.ApiToken.Token,
		StateStore: env.App.ServiceGroup,
	}

	if env.Dapr.Host == "" {
		if Dapr.Host, okHost = os.LookupEnv("DAPR_HOST"); !okHost {
			Dapr.Host = "http://localhost"
		}
	} else {
		Dapr.Host = env.Dapr.Host
	}
	if env.Dapr.Port == "" {
		if Dapr.Port, okPort = os.LookupEnv("DAPR_HTTP_PORT"); !okPort {
			Dapr.Port = "3500"
		}
	} else {
		Dapr.Port = env.Dapr.Port
	}
}

func Request(config ConfigRequest) (data Response.RespResultStruct, err error) {

	client := &http.Client{}
	req, err := http.NewRequest(config.Method, Dapr.Host+":"+Dapr.Port+config.Path, bytes.NewBuffer(config.Payload))
	if err != nil {
		return
	}
	req.Header.Add("dapr-app-id", config.AppID)
	req.Header.Add("Authorization", Dapr.ApiKey)
	response, err := client.Do(req)
	if err != nil {
		Logging.LogError(map[string]interface{}{"error": err.Error(), "dapr-app-id": config.AppID}, nil)
		log.Info().Msg("error: " + err.Error())
		return
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		Logging.LogError(map[string]interface{}{"error": err.Error(), "dapr-app-id": config.AppID}, nil)
		log.Info().Msg("error: " + err.Error())
		return
	}
	err = json.Unmarshal(result, &data)
	if response.StatusCode != http.StatusOK {
		if data.RestResult == nil {
			err = fmt.Errorf("%s", data.RestMessage)
		} else {
			err = fmt.Errorf("%v", data.RestResult)
		}
	}
	log.Info().Msg("result: " + string(result))
	return
}

//DEFAULT STATE STORE IS SERVICE GROUP IN ENVIRONMENT VARIABLE
func GetConfigState(subs []string, stateStore, key string, state interface{}, services ...StateService) (result ConfigState) {
	if stateStore == "" {
		stateStore = Dapr.StateStore + "-state"
		// stateStore = "statestore"
	}
	if key == "" {
		key = uuid.New().String()
	}
	result.StateStore = stateStore
	result.StateFmt.Key = key
	result.StateFmt.Value.State = state
	result.StateFmt.Value.Services = make([]StateService, 0)
	if services != nil || len(services) > 0 {
		result.StateFmt.Value.Services = services
		return
	} else {
		for _, value := range subs {
			element := StateService{Status: "progress", Name: value}
			result.StateFmt.Value.Services = append(result.StateFmt.Value.Services, element)
		}
	}
	return
}

func StatePublish(config ConfigState) error {
	client := &http.Client{}
	statePayload := make([]StateFmt, 0)
	statePayload = append(statePayload, config.StateFmt)
	state, _ := json.Marshal(statePayload)
	req, err := http.NewRequest(http.MethodPost, Dapr.Host+":"+Dapr.Port+"/v1.0/state/"+config.StateStore, bytes.NewBuffer(state))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	log.Info().Msgf("StatePublish Url is:%s", req.URL)
	response, err := client.Do(req)
	if err != nil {
		Logging.LogError(map[string]interface{}{"error": err.Error(), "dapr-state-management": config}, nil)
		log.Info().Msg("error: " + err.Error())
		return err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Info().Msgf("err StateGet is:%s", err.Error())
		return err
	}
	log.Info().Msgf("response.StatusCode  is:%d", response.StatusCode)
	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		// log.Info().Msgf("err StateGet is:%s", err.Error())
		return fmt.Errorf("%s", string(result))
	}
	// fmt.Println("Saving Order: ", string(result))
	log.Info().Msgf("result StatePublish is:%s", string(result))
	return nil
}

func StateGet(stateStore, key string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, Dapr.Host+":"+Dapr.Port+"/v1.0/state/"+stateStore+"/"+key, nil)
	if err != nil {
		log.Info().Msg("StateGet http.NewRequest, error: " + err.Error())
		return nil, err
	}
	response, err := client.Do(req)
	if err != nil {
		Logging.LogError(map[string]interface{}{"error": err.Error(), "dapr-state-management-stateStore": stateStore, "dapr-state-management-key": key}, nil)
		log.Info().Msg("StateGet client.Do, error: " + err.Error())
		return nil, err
	}
	result, err := ioutil.ReadAll(response.Body)
	log.Info().Msgf("response.StatusCode  is:%d", response.StatusCode)
	if response.StatusCode < http.StatusOK || response.StatusCode > 299 {
		// log.Info().Msgf("err StateGet is:%s", err.Error())
		err = fmt.Errorf("%s", string(result))
	}
	log.Info().Msgf("result StateGet is:%s", string(result))
	return result, err
}

func StateDelete(stateStore, key string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, Dapr.Host+":"+Dapr.Port+"/v1.0/state/"+stateStore+"/"+key, nil)
	if err != nil {
		log.Info().Msg("StateDelete http.NewRequest, error: " + err.Error())
		return err
	}
	response, err := client.Do(req)
	if err != nil {
		Logging.LogError(map[string]interface{}{"error": err.Error(), "dapr-state-management-stateStore": stateStore, "dapr-state-management-key": key}, nil)
		log.Info().Msg("StateDelete client.Do, error: " + err.Error())
		return err
	}
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Info().Msgf("err StateGet is:%s", err.Error())
	}
	log.Info().Msgf("result StateDelete is:%s", string(result))
	return nil
}

func SteteUpdateService(subsName, status string, stateSvc []StateService) []StateService {
	location, _ := time.LoadLocation(Constant.TimeLocation)
	timeStamp := time.Now().In(location)
	log.Info().Msgf("SteteUpdateService is:%s", subsName)
	for i := 0; i < len(stateSvc); i++ {
		if strings.Contains(stateSvc[i].Name, "/"+subsName) {
			stateSvc[i].Status = status
			stateSvc[i].UpdatedAt = &timeStamp
			break
		}
	}
	return stateSvc
}

// NOTED ORCHESTRATOR SUBSCRIPTION NOT INCLUDE FOR CHECKED
func CheckServices(orchestrator string, services []StateService) bool {
	var status = ""
	// var lenOfServices = len(services)
	for _, val := range services {
		if strings.Contains(val.Name, "/"+orchestrator) {
			continue
		}
		if val.Status == "progress" {
			status = "progress"
			break
		} else if val.Status == "failed" {
			status = "failed"
			break
		} else if val.Status == "ok" {
			status = "ok"
		}
	}
	if status == "ok" {
		return true

	} else {
		return false
	}
}

func GetService(name string, services []StateService) bool {
	var status = ""
	for _, val := range services {
		if strings.Contains(val.Name, "/"+name) {
			status = val.Status
			break
		}
	}
	return status == "ok"
}
