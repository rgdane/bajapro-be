package config

import (
	"context"
	"fmt"
	"jk-api/internal/helper"
	"os"
	"path/filepath"
	"sync"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"gorm.io/datatypes"
)

var (
	fcmApp     *firebase.App
	fcmOnce    sync.Once
	fcmMsgOnce sync.Once

	authClient *auth.Client
	msgClient  *messaging.Client
)

func InitFirebaseApp() {
	opt := option.WithCredentialsFile(filepath.Join("internal", "creds", "gcp_firebase.json"))
	var err error
	fcmApp, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to Firebase: %v", err))
	}
	helper.LogMessage("CONFIG", "✅ Successfully connected to Firebase")
}

func GetFirebaseCreds() (datatypes.JSON, error) {
	path := filepath.Join("internal", "creds", "gcp_firebase.json")
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return datatypes.JSON(data), nil
}

func GetFirebaseAuth() *auth.Client {
	fcmOnce.Do(func() {
		var err error
		authClient, err = fcmApp.Auth(context.Background())
		if err != nil {
			helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to Firebase: %v", err))
		}
	})
	return authClient
}

func GetFirebaseMessaging() *messaging.Client {
	fcmMsgOnce.Do(func() {
		var err error
		msgClient, err = fcmApp.Messaging(context.Background())
		if err != nil {
			helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to Firebase: %v", err))
		}
	})
	return msgClient
}
