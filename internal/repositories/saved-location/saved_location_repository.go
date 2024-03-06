//go:generate go run github.com/golang/mock/mockgen -destination ./mock/$GOFILE . SavedLocationRepository

package savedlocationrepository

import (
	"context"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"github.com/dragonfish/go/v2/pkg/logger"
	pb "github.com/ride-app/user-service/api/ride/rider/v1alpha1"

	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SavedLocationRepository interface {
	CreateSavedLocation(ctx context.Context, savedlocation *pb.SavedLocation, log logger.Logger) (createTime *time.Time, err error)

	GetSavedLocation(ctx context.Context, uid string, id string, log logger.Logger) (*pb.SavedLocation, error)

	GetSavedLocations(ctx context.Context, uid string, log logger.Logger) ([]*pb.SavedLocation, error)

	UpdateSavedLocation(ctx context.Context, savedlocation *pb.SavedLocation, log logger.Logger) (createTime *time.Time, err error)

	DeleteSavedLocation(ctx context.Context, uid string, id string, log logger.Logger) (createTime *time.Time, err error)
}

type FirebaseImpl struct {
	firestore *firestore.Client
}

func NewFirebaseSavedLocationRepository(firebaseApp *firebase.App, log logger.Logger) (*FirebaseImpl, error) {
	firestore, err := firebaseApp.Firestore(context.Background())

	if err != nil {
		log.WithError(err).Error("Failed to initialize firebase firestore")
		return nil, err
	}

	return &FirebaseImpl{
		firestore: firestore,
	}, nil
}

func (r *FirebaseImpl) CreateSavedLocation(ctx context.Context, savedlocation *pb.SavedLocation, log logger.Logger) (createTime *time.Time, err error) {
	log.Info("Writing saved location to firestore")
	writeResult, err := r.firestore.Collection("users").Doc(strings.Split(savedlocation.Name, "/")[1]).Collection("savedlocations").Doc(strings.Split(savedlocation.Name, "/")[3]).Set(ctx, map[string]interface{}{
		"displayName": savedlocation.DisplayName,
		"location": map[string]interface{}{
			"latitude":  savedlocation.Location.Latitude,
			"longitude": savedlocation.Location.Longitude,
		},
		"address": savedlocation.Address,
	})

	if err != nil {
		log.WithError(err).Error("Failed to write saved location to firestore")
		return nil, err
	}

	return &writeResult.UpdateTime, nil
}

func (r *FirebaseImpl) GetSavedLocation(ctx context.Context, uid string, id string, log logger.Logger) (*pb.SavedLocation, error) {
	log.Info("Getting saved location from firestore")
	doc, err := r.firestore.Collection("users").Doc(uid).Collection("savedlocations").Doc(id).Get(ctx)

	if err != nil {
		log.WithError(err).Error("Failed to get saved location from firestore")
		return nil, err
	}

	if !doc.Exists() {
		log.Info("Saved location does not exist")
		return nil, nil
	}

	savedlocation := &pb.SavedLocation{
		Name:        "users/" + uid + "/savedlocations/" + id,
		DisplayName: doc.Data()["displayName"].(string),
		Location: &latlng.LatLng{
			Latitude:  doc.Data()["location"].(map[string]interface{})["latitude"].(float64),
			Longitude: doc.Data()["location"].(map[string]interface{})["longitude"].(float64),
		},
		Address:    doc.Data()["address"].(string),
		CreateTime: timestamppb.New(doc.CreateTime),
		UpdateTime: timestamppb.New(doc.UpdateTime),
	}

	return savedlocation, nil
}

func (r *FirebaseImpl) GetSavedLocations(ctx context.Context, uid string, log logger.Logger) ([]*pb.SavedLocation, error) {
	log.Info("Getting saved locations from firestore")
	docs, err := (r.firestore.Collection("users").Doc(uid).Collection("savedlocations").Documents(ctx)).GetAll()

	if err != nil {
		log.WithError(err).Error("Failed to get saved locations from firestore")
		return nil, err
	}

	var savedlocations []*pb.SavedLocation

	if len(docs) == 0 {
		return savedlocations, nil
	}

	for _, doc := range docs {
		savedlocation := &pb.SavedLocation{
			Name:        "users/" + uid + "/savedlocations/" + doc.Ref.ID,
			DisplayName: doc.Data()["displayName"].(string),
			Location: &latlng.LatLng{
				Latitude:  doc.Data()["location"].(map[string]interface{})["latitude"].(float64),
				Longitude: doc.Data()["location"].(map[string]interface{})["longitude"].(float64),
			},
			Address:    doc.Data()["address"].(string),
			CreateTime: timestamppb.New(doc.CreateTime),
			UpdateTime: timestamppb.New(doc.UpdateTime),
		}

		savedlocations = append(savedlocations, savedlocation)
	}

	return savedlocations, nil
}

func (r *FirebaseImpl) UpdateSavedLocation(ctx context.Context, savedlocation *pb.SavedLocation, log logger.Logger) (createTime *time.Time, err error) {
	log.Info("Updating saved location in firestore")
	writeResult, err := r.firestore.Collection("users").Doc(strings.Split(savedlocation.Name, "/")[1]).Collection("savedLocations").Doc(strings.Split(savedlocation.Name, "/")[3]).Set(ctx, map[string]interface{}{
		"displayName": savedlocation.DisplayName,
		"location": map[string]interface{}{
			"latitude":  savedlocation.Location.Latitude,
			"longitude": savedlocation.Location.Longitude,
		},
		"address": savedlocation.Address,
	}, firestore.MergeAll)

	if err != nil {
		log.WithError(err).Error("Failed to update saved location in firestore")
		return nil, err
	}

	return &writeResult.UpdateTime, nil
}

func (r *FirebaseImpl) DeleteSavedLocation(ctx context.Context, uid string, id string, log logger.Logger) (deleteTime *time.Time, err error) {
	log.Info("Deleting saved location from firestore")
	writeResult, err := r.firestore.Collection("users").Doc(uid).Collection("savedLocationa").Doc(id).Delete(ctx)

	if err != nil {
		log.WithError(err).Error("Failed to delete saved location from firestore")
		return nil, err
	}

	return &writeResult.UpdateTime, nil
}
