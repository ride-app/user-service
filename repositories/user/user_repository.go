//go:generate go run github.com/golang/mock/mockgen -destination ../../mocks/$GOFILE -package mocks . UserRepository

package userrepository

import (
	"context"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	pb "github.com/ride-app/user-service/api/gen/ride/user/v1alpha1"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*pb.User, error)

	UpdateUser(ctx context.Context, user *pb.User) (createTime *time.Time, err error)

	DeleteUser(ctx context.Context, id string) (createTime *time.Time, err error)
}

type FirebaseImpl struct {
	auth *auth.Client
}

func NewFirebaseUserRepository(firebaseApp *firebase.App) (*FirebaseImpl, error) {
	auth, err := firebaseApp.Auth(context.Background())

	if err != nil {
		return nil, err
	}

	return &FirebaseImpl{
		auth: auth,
	}, nil
}

func (r *FirebaseImpl) GetUser(ctx context.Context, id string) (*pb.User, error) {
	userRecord, err := r.auth.GetUser(ctx, id)

	if err != nil {
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

func (r *FirebaseImpl) UpdateUser(ctx context.Context, user *pb.User) (updateTime *time.Time, err error) {
	params := (&auth.UserToUpdate{}).
		DisplayName(user.DisplayName).
		Email(*user.Email).
		PhotoURL(user.PhotoUrl)

	if _, err := r.auth.UpdateUser(ctx, strings.Split(user.Name, "/")[1], params); err != nil {
		return nil, err
	}

	res := time.Now()

	return &res, nil
}

func (r *FirebaseImpl) DeleteUser(ctx context.Context, id string) (deleteTime *time.Time, err error) {
	if err := r.auth.DeleteUser(ctx, id); err != nil {
		return nil, err
	}

	res := time.Now()

	return &res, nil
}
