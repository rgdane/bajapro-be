// routes/fcm_routes.go
package routes

import (
	"fmt"
	"jk-api/api/http/presenters"
	"jk-api/internal/config"
	"jk-api/internal/container"
	"jk-api/pkg/services/v1"

	"firebase.google.com/go/messaging"
	"github.com/gofiber/fiber/v2"
)

type SubscribeRequest struct {
	Token  string   `json:"token"`
	Topics []string `json:"topics"`
	UserID int64    `json:"user_id"`
}

func FcmRoutes(router fiber.Router, c *container.AppContainer) {
	app := router.Group("fcm")

	app.Get("/creds", getFCMEnv())
	app.Post("/subscribe", subscribeFcm())
	app.Post("/unsubscribe", unsubscribeFcm())

	app.Post("/send", sendFcm())
	app.Post("/send-user", sendFcmUser())
}

func getFCMEnv() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		creds, err := config.GetFirebaseCreds()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.JSON(creds)
	}
}

func subscribeFcm() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SubscribeRequest
		if err := ctx.BodyParser(&req); err != nil {
			return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, err)
		}

		if req.Token == "" || len(req.Topics) == 0 {
			return presenters.ErrorResponseWithMessage(ctx, fiber.StatusBadRequest, "token & topics required")
		}

		fmt.Println("token:", req.Token)
		fmt.Println("topics:", req.Topics)

		client := config.GetFirebaseMessaging()

		results := make(map[string]fiber.Map)

		for _, topic := range req.Topics {
			// 1. Subscribe ke topic global
			respGlobal, err := client.SubscribeToTopic(ctx.Context(), []string{req.Token}, topic)
			if err != nil {
				results[topic] = fiber.Map{
					"error": err.Error(),
				}
				continue
			}

			// 2. Subscribe ke topic personal (topic_user_<id>)
			var respUser *messaging.TopicManagementResponse
			var errUser error
			if req.UserID > 0 {
				userTopic := fmt.Sprintf("%s_user_%d", topic, req.UserID)
				respUser, errUser = client.SubscribeToTopic(ctx.Context(), []string{req.Token}, userTopic)
				if errUser != nil {
					results[topic] = fiber.Map{
						"global": fiber.Map{
							"successCount": respGlobal.SuccessCount,
							"failureCount": respGlobal.FailureCount,
						},
						"user": fiber.Map{
							"error": errUser.Error(),
						},
					}
					continue
				}
				results[topic] = fiber.Map{
					"global": fiber.Map{
						"successCount": respGlobal.SuccessCount,
						"failureCount": respGlobal.FailureCount,
					},
					"user": fiber.Map{
						"successCount": respUser.SuccessCount,
						"failureCount": respUser.FailureCount,
					},
				}
			} else {
				results[topic] = fiber.Map{
					"global": fiber.Map{
						"successCount": respGlobal.SuccessCount,
						"failureCount": respGlobal.FailureCount,
					},
				}
			}
		}

		return presenters.SuccessResponse(ctx, fiber.Map{
			"topics": results,
		})
	}
}

func unsubscribeFcm() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var req SubscribeRequest
		if err := ctx.BodyParser(&req); err != nil {
			return presenters.ErrorResponse(ctx, fiber.StatusBadRequest, err)
		}

		if req.Token == "" || len(req.Topics) == 0 {
			return presenters.ErrorResponseWithMessage(ctx, fiber.StatusBadRequest, "token & topics required")
		}

		fmt.Println("token:", req.Token)
		fmt.Println("topics (unsubscribe):", req.Topics)

		client := config.GetFirebaseMessaging()

		results := make(map[string]fiber.Map)

		for _, topic := range req.Topics {
			// 1. Unsubscribe dari topic global
			respGlobal, err := client.UnsubscribeFromTopic(ctx.Context(), []string{req.Token}, topic)
			if err != nil {
				results[topic] = fiber.Map{
					"error": err.Error(),
				}
				continue
			}

			// 2. Unsubscribe dari topic personal (topic_user_<id>)
			if req.UserID > 0 {
				userTopic := fmt.Sprintf("%s_user_%d", topic, req.UserID)
				respUser, errUser := client.UnsubscribeFromTopic(ctx.Context(), []string{req.Token}, userTopic)
				if errUser != nil {
					results[topic] = fiber.Map{
						"global": fiber.Map{
							"successCount": respGlobal.SuccessCount,
							"failureCount": respGlobal.FailureCount,
						},
						"user": fiber.Map{
							"error": errUser.Error(),
						},
					}
					continue
				}

				results[topic] = fiber.Map{
					"global": fiber.Map{
						"successCount": respGlobal.SuccessCount,
						"failureCount": respGlobal.FailureCount,
					},
					"user": fiber.Map{
						"successCount": respUser.SuccessCount,
						"failureCount": respUser.FailureCount,
					},
				}
			} else {
				results[topic] = fiber.Map{
					"global": fiber.Map{
						"successCount": respGlobal.SuccessCount,
						"failureCount": respGlobal.FailureCount,
					},
				}
			}
		}

		return presenters.SuccessResponse(ctx, fiber.Map{
			"topics": results,
		})
	}
}

func sendFcm() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := services.SendNotification("trigger", "test message", "test"); err != nil {
			return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(ctx, nil)
	}
}

func sendFcmUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := services.SendNotification("testing_user_3", "test message", "test"); err != nil {
			return presenters.ErrorResponse(ctx, fiber.StatusInternalServerError, err)
		}
		return presenters.SuccessResponse(ctx, nil)
	}
}
