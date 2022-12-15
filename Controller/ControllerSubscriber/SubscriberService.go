package ControllerSubscriber

type SubscriberInterface interface {
	EmployeeInterface
	LeaveBalanceInterface
	OrchestratorInterface
}

const DeliveryAttemptMax int = 1

type Subscriber struct {
	SubscriberInterface
}
