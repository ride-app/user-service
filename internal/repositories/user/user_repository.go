//go:generate go run github.com/golang/mock/mockgen -destination ./mock/$GOFILE . UserRepository

package userrepository

import (
	"context"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/dragonfish-tech/go/pkg/logger"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string, log logger.Logger) (*pb.User, error)

	UpdateUser(ctx context.Context, user *pb.User, log logger.Logger) (createTime *time.Time, err error)

	DeleteUser(ctx context.Context, id string, log logger.Logger) (createTime *time.Time, err error)
}

type FirebaseImpl struct {
	auth *auth.Client
}

func NewFirebaseUserRepository(firebaseApp *firebase.App, log logger.Logger) (*FirebaseImpl, error) {
	auth, err := firebaseApp.Auth(context.Background())

	if err != nil {
		log.WithError(err).Error("Failed to initialize firebase auth")
		return nil, err
	}

	return &FirebaseImpl{
		auth: auth,
	}, nil
}

func (r *FirebaseImpl) GetUser(ctx context.Context, id string, log logger.Logger) (*pb.User, error) {
	log.Info("Getting user record")
	log.Debug("id: ", id)
	userRecord, err := r.auth.GetUser(ctx, id)

	if err != nil {
		log.WithError(err).Error("Failed to get user record")
		return nil, err
	}

	user := &pb.User{
		Name:        "users/" + userRecord.UID,
		DisplayName: userRecord.DisplayName,
		PhoneNumber: userRecord.PhoneNumber,
		Email:       &userRecord.Email,
		PhotoUrl:    userRecord.PhotoURL,
	}

	return user, nil
}

func (r *FirebaseImpl) UpdateUser(ctx context.Context, user *pb.User, log logger.Logger) (updateTime *time.Time, err error) {
	log.Info("Updating user record")
	params := (&auth.UserToUpdate{}).
		DisplayName(user.DisplayName).
		Email(*user.Email).
		PhotoURL(user.PhotoUrl)

	if _, err := r.auth.UpdateUser(ctx, strings.Split(user.Name, "/")[1], params); err != nil {
		log.WithError(err).Error("Failed to update user record")
		return nil, err
	}

	res := time.Now()

	return &res, nil
}

func (r *FirebaseImpl) DeleteUser(ctx context.Context, id string, log logger.Logger) (deleteTime *time.Time, err error) {
	log.Info("Deleting user record")
	if err := r.auth.DeleteUser(ctx, id); err != nil {
		log.WithError(err).Error("Failed to delete user record")
		return nil, err
	}

	res := time.Now()

	return &res, nil
}
