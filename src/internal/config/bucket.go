package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jk-api/internal/helper"
	"os"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	bucketApp      *storage.Client
	credPath       = filepath.Join("internal", "creds", "gcp_bucket.json")
	serviceAccount *ServiceAccount
)

type ServiceAccount struct {
	ClientEmail string `json:"client_email"`
	PrivateKey  string `json:"private_key"`
}

func InitBucketApp() {
	loadServiceAccount()

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credPath))
	if err != nil {
		helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to connect to Google Cloud Storage: %v", err))
	}

	bucketApp = client
	helper.LogMessage("CONFIG", "✅ Successfully connected to Google Cloud Storage")
}

func GetBucketApp() *storage.Client {
	return bucketApp
}

func GetBucketName() string {
	_ = godotenv.Load()
	name := os.Getenv("GCP_BUCKET_NAME")
	return name
}

func GetBucketCreds() string {
	return credPath
}

func GetServiceAccountEmail() string {
	if serviceAccount == nil {
		loadServiceAccount()
	}
	return serviceAccount.ClientEmail
}

func GetServiceAccountPrivateKey() []byte {
	if serviceAccount == nil {
		loadServiceAccount()
	}
	return []byte(serviceAccount.PrivateKey)
}

func loadServiceAccount() {
	if serviceAccount != nil {
		return
	}
	data, err := ioutil.ReadFile(credPath)
	if err != nil {
		helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to read service account file: %v", err))
		os.Exit(1)
	}

	var sa ServiceAccount
	if err := json.Unmarshal(data, &sa); err != nil {
		helper.LogMessage("ERROR", fmt.Sprintf("❌ Failed to parse service account file: %v", err))
		os.Exit(1)
	}
	serviceAccount = &sa
}
