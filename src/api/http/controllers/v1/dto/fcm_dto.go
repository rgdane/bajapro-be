package dto

type FcmNotificationDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type SendToTopicDto struct {
	Notification FcmNotificationDto `json:"notification"`
	Topic        string             `json:"topic"`
}
