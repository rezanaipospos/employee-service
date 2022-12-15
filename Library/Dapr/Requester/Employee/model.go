package Employee

const (
	EmployeeService     string = "employee-service"
	NotificationService string = "notification-service"
)

type CheckEmployeeExistData struct {
	ID            int64  `json:"id"`
	Code          string `json:"code" `
	Name          string `json:"name" `
	MobilePhoneNo string `json:"mobilePhoneNo"`
	Email         string `json:"email"`
	ActiveStatus  bool   `json:"activeStatus" `
}

type CheckExistEmployeeID struct {
	ID            int64  `json:"id" `
	Code          string `json:"code"`
	Name          string `json:"name"`
	MobilePhoneNo string `json:"mobilePhoneNo" `
	Email         string `json:"email" `
	ActiveStatus  bool   `json:"activeStatus" `
	MachineId     int64  `json:"machineId" `
	MachineName   string `json:"machineName" `
}
