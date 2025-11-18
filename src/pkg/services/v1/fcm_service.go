package services

import (
	"context"
	"fmt"
	"jk-api/api/http/controllers/v1/dto"
	"jk-api/internal/config"
	"time"

	"firebase.google.com/go/messaging"
)

func SendNotification(topic string, title string, message string) error {
	if topic == "" {
		return fmt.Errorf("topic cannot be empty")
	}
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	fcmClient := config.GetFirebaseMessaging()
	if fcmClient == nil {
		return fmt.Errorf("firebase messaging client is not initialized")
	}

	fcmReq := dto.SendToTopicDto{
		Notification: dto.FcmNotificationDto{
			Title: title,
			Body:  message,
		},
		Topic: topic,
	}

	err := sendToTopic(context.Background(), fcmClient, fcmReq)
	if err != nil {
		return err
	}

	return nil
}

func sendToTopic(ctx context.Context, client *messaging.Client, req dto.SendToTopicDto) error {
	ttl := 30 * time.Second
	msg := &messaging.Message{
		Notification: &messaging.Notification{
			Title: req.Notification.Title,
			Body:  req.Notification.Body,
		},
		Topic: req.Topic,
		Android: &messaging.AndroidConfig{
			TTL: &ttl,
		},
		Webpush: &messaging.WebpushConfig{
			Headers: map[string]string{
				"TTL": "30",
			},
		},
	}

	response, err := client.Send(ctx, msg)
	if err != nil {
		return err
	}

	fmt.Println("Successfully sent message:", response)
	return nil
}
