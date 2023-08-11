package thirdparty

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/ride-app/user-service/config"
	"github.com/ride-app/user-service/logger"
)

func NewFirebaseApp(log logger.Logger) (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: config.Env.Firebase_Project_Id}
	log.Info("Initializing Firebase App")
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.WithError(err).Fatal("Cannot initialize firebase app")
		return nil, err
	}

	return app, nil
}
