package Dto

type NotificationInfoDTO struct {
	SenderUserId   int64  `json:"senderUserId"`
	SenderId       int64  `json:"senderId"`
	SenderCode     string `json:"senderCode"`
	SenderName     string `json:"senderName"`
	SenderPhoto    string `json:"senderPhoto"`
	ReceiverUserId int64  `json:"receiverUserId"`
	ReceiverId     int64  `json:"receiverId" `
	ReceiverCode   string `json:"receiverCode"`
	ReceiverName   string `json:"receiverName"`
	ReceiverPhoto  string `json:"receiverPhoto"`
	Application    string `json:"application"`
	Title          string `json:"title"`
	Message        string `json:"message"`
	Transaction    string `json:"transaction"`
	TransactionId  string `json:"transactionId"`
	UrlPath        string `json:"urlPath"`
}
