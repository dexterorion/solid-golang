package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type OnboardingService struct {
	db                *sql.DB
	googleProjectID   string
	googlePubsubTopic string
}

func NewOnboardingService(db *sql.DB, googleProjectID, googlePubsubTopic string) *OnboardingService {
	return &OnboardingService{
		db:                db,
		googleProjectID:   googleProjectID,
		googlePubsubTopic: googlePubsubTopic,
	}
}

func (ns *OnboardingService) SendWelcome(userID, templateType string) error {
	// sending message
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, ns.googleProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(ns.googlePubsubTopic)

	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("UserID: "+userID, ", TemplateType: ", templateType),
	})

	// The Get method blocks until a server-generated ID or
	// an error is returned for the published message.
	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("Failed to publish: %v", err)
	}

	// store on db
	_, err = db.Exec("INSERT INTO notification  (user_id, template_type, created_at) VALUES ($1, $2, $3)", userID, templateType, time.Now())
	if err != nil {
		return fmt.Errorf("Failed to store on db: %v", err)
	}

	return nil
}
