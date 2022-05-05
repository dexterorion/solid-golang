package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// publisher
type OnboardingPublisher interface {
	Publish(userID, templateType string) error
}

type OnboardingPubSubPublisher struct {
	googleProjectID   string
	googlePubsubTopic string
}

func (op *OnboardingPublisher) Publish(userID, templateType string) error {
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

	return nil
}

func NewOnboardingPubSubPublisher(googleProjectID, googlePubsubTopic string) OnboardingPublisher {
	return &OnboardingPubSubPublisher{
		googleProjectID:   googleProjectID,
		googlePubsubTopic: googlePubsubTopic,
	}
}

// database manager
type OnboardingRepository interface {
	Store(userID, templateType string)
}

type OnboardingSqlRepository struct {
	db *sql.DB
}

func (osr *OnboardingSqlRepository) Store(userID, templateType string) error {
	_, err = db.Exec("INSERT INTO notification  (user_id, template_type, created_at) VALUES ($1, $2, $3)", userID, templateType, time.Now())
	if err != nil {
		return fmt.Errorf("Failed to store on db: %v", err)
	}

	return nil
}

func NewOnboardingSqlRepository(db *sql.DB) OnboardingRepository {
	return &OnboardingSqlRepository{
		db: db,
	}
}

// service
type OnboardingService struct {
	publisher  OnboardingPublisher
	repository OnboardingRepository
}

func NewOnboardingService(publisher OnboardingPublisher, repository OnboardingRepository) *OnboardingService {
	return &OnboardingService{
		publisher:  publisher,
		repository: repository,
	}
}

func (os *OnboardingService) SendWelcome(userID, templateType string) error {
	err := os.publisher.Publish(userID, templateType)
	if err != nil {
		return err
	}

	return os.repository.Store(userID, templateType)
}
