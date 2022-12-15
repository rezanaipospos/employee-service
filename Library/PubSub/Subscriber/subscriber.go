package Subscriber

import (
	"context"
	"strings"

	"cloud.google.com/go/pubsub"
	"github.com/rs/zerolog/log"
)

const (
	//Employee
	SUBS_EMPLOYEEADDED               PubSubSubscription = "Employee-EmployeeAdded"
	SUBS_EMPLOYEEPERSONALINFOUPDATED PubSubSubscription = "Employee-EmployeePersonalInfoUpdated"
	SUBS_EMPLOYEEDELETED             PubSubSubscription = "Employee-EmployeeDeleted"

	SUBS_EMPLOYEEWORKSTATUSUPDATED         PubSubSubscription = "Employee-EmployeeWorkStatusUpdated"
	SUBS_ROLLBACKEMPLOYEEWORKSTATUSUPDATED PubSubSubscription = "Employee-RollbackEmployeeWorkStatusUpdated"

	SUBS_EMPLOYEERESIGNED         PubSubSubscription = "Employee-EmployeeResigned"
	SUBS_ROLLBACKEMPLOYEERESIGNED PubSubSubscription = "Employee-RollbackEmployeeResigned"

	SUBS_EMPLOYEEFINGERUPDATED PubSubSubscription = "Employee-EmployeeFingerUpdated"

	// Employee - Company Stucuture
	SUBS_COMPANYLOCATIONUPDATED PubSubSubscription = "Employee-CompanyLocationUpdated"
	SUBS_COMPANYUPDATED         PubSubSubscription = "Employee-CompanyUpdated"
	SUBS_DEPARTMENTUPDATED      PubSubSubscription = "Employee-DepartmentUpdated"
	SUBS_LOCATIONUPDATED        PubSubSubscription = "Employee-LocationUpdated"
	SUBS_SECTIONUPDATED         PubSubSubscription = "Employee-SectionUpdated"
	SUBS_POSITIONUPDATED        PubSubSubscription = "Employee-PositionUpdated"

	//Employee & Leave Balance
	SUBS_EMPLOYEETRANSFERED PubSubSubscription = "Employee-EmployeeTransfered"

	//Leave Balance
	SUBS_EMPLOYEELEAVEAPPROVED PubSubSubscription = "Employee-LeaveApproved"

	//ORCHESTRATOR
	SUBS_ORCHESTRATORPERSONALINFOUPDATED = "Orchestrator-EmployeePersonalnfoUpdated"
	SUBS_ROLLBACKPERSONALINFOUPDATE      = "Employee-RollbackPersonalInfoUpdated"
)

func (d PubSubSubscriberConfig) SetClient(client *pubsub.Client) {
	pubsubClient = client
}
func (c PubSubSubscriberConfig) CollectSubscriber() bool {
	//Employee
	c.CreateTopicClientSubscription(SUBS_EMPLOYEEADDED, c.SubscriberController.SubsEmployeeAdded)
	c.CreateTopicClientSubscription(SUBS_EMPLOYEEPERSONALINFOUPDATED, c.SubscriberController.SubsEmployeePersonalInfoUpdated)
	c.CreateTopicClientSubscription(SUBS_EMPLOYEEDELETED, c.SubscriberController.SubsEmployeeDeleted)

	c.CreateTopicClientSubscription(SUBS_EMPLOYEEWORKSTATUSUPDATED, c.SubscriberController.SubsEmployeeWorkStatusUpdated)
	c.CreateTopicClientSubscription(SUBS_ROLLBACKEMPLOYEEWORKSTATUSUPDATED, c.SubscriberController.SubsRollbackEmployeeWorkStatusUpdated)

	c.CreateTopicClientSubscription(SUBS_EMPLOYEERESIGNED, c.SubscriberController.SubsEmployeeResigned)
	c.CreateTopicClientSubscription(SUBS_ROLLBACKEMPLOYEERESIGNED, c.SubscriberController.SubsEmployeeFingerUpdated)

	c.CreateTopicClientSubscription(SUBS_EMPLOYEEFINGERUPDATED, c.SubscriberController.SubsEmployeeFingerUpdated)

	c.CreateTopicClientSubscription(SUBS_COMPANYLOCATIONUPDATED, c.SubscriberController.SubsCompanyLocationUpdated)
	c.CreateTopicClientSubscription(SUBS_COMPANYUPDATED, c.SubscriberController.SubsCompanyUpdated)
	c.CreateTopicClientSubscription(SUBS_DEPARTMENTUPDATED, c.SubscriberController.SubsDepartmentUpdate)
	c.CreateTopicClientSubscription(SUBS_LOCATIONUPDATED, c.SubscriberController.SubsLocationUpdate)
	c.CreateTopicClientSubscription(SUBS_SECTIONUPDATED, c.SubscriberController.SubsSectionUpdate)
	c.CreateTopicClientSubscription(SUBS_POSITIONUPDATED, c.SubscriberController.SubsPositionUpdate)
	//Employee & Leave Balance
	c.CreateTopicClientSubscription(SUBS_EMPLOYEETRANSFERED, c.SubscriberController.SubsLeaveBalanceUpdate)
	//Leave Balance
	c.CreateTopicClientSubscription(SUBS_EMPLOYEELEAVEAPPROVED, c.SubscriberController.SubsLeaveBalanceApproved)

	//ORCHESTRATOR
	c.CreateTopicClientSubscription(SUBS_ORCHESTRATORPERSONALINFOUPDATED, c.SubscriberController.SubsOrchestratorEmployeePersonalInfoUpdated)
	c.CreateTopicClientSubscription(SUBS_ROLLBACKPERSONALINFOUPDATE, c.SubscriberController.SubsRollbackPersonalInfoUpdated)

	return true
}

func (c PubSubSubscriberConfig) Subscribe() {
	if ok := c.CollectSubscriber(); ok {
		subscriptionList := make([]string, 0)
		for key, val := range TopicSubscriptionController {
			subscriptionList = append(subscriptionList, string(key))
			go func(key PubSubSubscription, subscriber PubSubTopicSubscription) {
				ctx := context.WithValue(context.Background(), "subscriptionID", string(key))
				err := subscriber.Subscription.Receive(ctx, subscriber.SubcriberContoller)
				if err != nil {
					log.Fatal().Msgf("Failed to Sub  Receive: %v", err)
				}
			}(key, val)
		}
		subscriptions := strings.Join(subscriptionList, ", ")
		log.Info().Msgf("Success to PubSub Subscription Receive: %s", subscriptions)
	}
}

func (c PubSubSubscriberConfig) CreateTopicClientSubscription(subscriptionId PubSubSubscription, controller func(ctx context.Context, m *pubsub.Message)) {
	if _, ok := TopicSubscriptionController[subscriptionId]; !ok {
		subs := pubsubClient.Subscription(string(subscriptionId))
		exists, err := subs.Exists(context.Background())
		if err != nil {
			log.Fatal().Msgf("Failed to connect client Subscription: %s, error: %v", subscriptionId, err)
		}
		if !exists {
			log.Fatal().Msgf("PubSub Subscription: %s, Not Found", subscriptionId)
		}
		TopicSubscriptionController[subscriptionId] = PubSubTopicSubscription{SubscriptionID: subscriptionId, Subscription: subs, SubcriberContoller: controller}
	}
}
