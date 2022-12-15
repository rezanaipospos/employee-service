package main

import "EmployeeService/Application"

// @title Employee Service
// @version 1.0
// @description Employee Serviceclea
// @BasePath /employee
// @securityDefinitions.apikey Bearer
// @in Header
// @name Authorization
// @description Example: Bearer abcdefghijklmnopqrstuvwxyz12345678901
func main() {
	Application.AppInitialization()
	//CreateService.CreateFile()
	//CreateService.DeleteFile()
}
