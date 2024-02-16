package thirdparty

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/pkg/logger"
	"github.com/ride-app/user-service/config"
)

func NewFirebaseApp(log logger.Logger, config *config.Config) (*firebase.App, error) {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: config.Project_Id}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return app, nil
}
