// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ride/rider/v1alpha1/user_service.proto

package v1alpha1connect

import (
	context "context"
	errors "errors"
	http "net/http"
	strings "strings"

	connect "connectrpc.com/connect"

	v1alpha1 "github.com/ride-app/user-service/api/ride/rider/v1alpha1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "ride.rider.v1alpha1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/ride.rider.v1alpha1.UserService/GetUser"
	// UserServiceUpdateUserProcedure is the fully-qualified name of the UserService's UpdateUser RPC.
	UserServiceUpdateUserProcedure = "/ride.rider.v1alpha1.UserService/UpdateUser"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/ride.rider.v1alpha1.UserService/DeleteUser"
	// UserServiceCreateSavedLocationProcedure is the fully-qualified name of the UserService's
	// CreateSavedLocation RPC.
	UserServiceCreateSavedLocationProcedure = "/ride.rider.v1alpha1.UserService/CreateSavedLocation"
	// UserServiceListSavedLocationsProcedure is the fully-qualified name of the UserService's
	// ListSavedLocations RPC.
	UserServiceListSavedLocationsProcedure = "/ride.rider.v1alpha1.UserService/ListSavedLocations"
	// UserServiceGetSavedLocationProcedure is the fully-qualified name of the UserService's
	// GetSavedLocation RPC.
	UserServiceGetSavedLocationProcedure = "/ride.rider.v1alpha1.UserService/GetSavedLocation"
	// UserServiceUpdateSavedLocationProcedure is the fully-qualified name of the UserService's
	// UpdateSavedLocation RPC.
	UserServiceUpdateSavedLocationProcedure = "/ride.rider.v1alpha1.UserService/UpdateSavedLocation"
	// UserServiceDeleteSavedLocationProcedure is the fully-qualified name of the UserService's
	// DeleteSavedLocation RPC.
	UserServiceDeleteSavedLocationProcedure = "/ride.rider.v1alpha1.UserService/DeleteSavedLocation"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor                   = v1alpha1.File_ride_rider_v1alpha1_user_service_proto.Services().ByName("UserService")
	userServiceGetUserMethodDescriptor             = userServiceServiceDescriptor.Methods().ByName("GetUser")
	userServiceUpdateUserMethodDescriptor          = userServiceServiceDescriptor.Methods().ByName("UpdateUser")
	userServiceDeleteUserMethodDescriptor          = userServiceServiceDescriptor.Methods().ByName("DeleteUser")
	userServiceCreateSavedLocationMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateSavedLocation")
	userServiceListSavedLocationsMethodDescriptor  = userServiceServiceDescriptor.Methods().ByName("ListSavedLocations")
	userServiceGetSavedLocationMethodDescriptor    = userServiceServiceDescriptor.Methods().ByName("GetSavedLocation")
	userServiceUpdateSavedLocationMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("UpdateSavedLocation")
	userServiceDeleteSavedLocationMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("DeleteSavedLocation")
)

// UserServiceClient is a client for the ride.rider.v1alpha1.UserService service.
type UserServiceClient interface {
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	UpdateUser(context.Context, *connect.Request[v1alpha1.UpdateUserRequest]) (*connect.Response[v1alpha1.UpdateUserResponse], error)
	DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error)
	CreateSavedLocation(context.Context, *connect.Request[v1alpha1.CreateSavedLocationRequest]) (*connect.Response[v1alpha1.CreateSavedLocationResponse], error)
	ListSavedLocations(context.Context, *connect.Request[v1alpha1.ListSavedLocationsRequest]) (*connect.Response[v1alpha1.ListSavedLocationsResponse], error)
	GetSavedLocation(context.Context, *connect.Request[v1alpha1.GetSavedLocationRequest]) (*connect.Response[v1alpha1.GetSavedLocationResponse], error)
	UpdateSavedLocation(context.Context, *connect.Request[v1alpha1.UpdateSavedLocationRequest]) (*connect.Response[v1alpha1.UpdateSavedLocationResponse], error)
	DeleteSavedLocation(context.Context, *connect.Request[v1alpha1.DeleteSavedLocationRequest]) (*connect.Response[v1alpha1.DeleteSavedLocationResponse], error)
}

// NewUserServiceClient constructs a client for the ride.rider.v1alpha1.UserService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		getUser: connect.NewClient[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect.WithSchema(userServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateUser: connect.NewClient[v1alpha1.UpdateUserRequest, v1alpha1.UpdateUserResponse](
			httpClient,
			baseURL+UserServiceUpdateUserProcedure,
			connect.WithSchema(userServiceUpdateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			connect.WithSchema(userServiceDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createSavedLocation: connect.NewClient[v1alpha1.CreateSavedLocationRequest, v1alpha1.CreateSavedLocationResponse](
			httpClient,
			baseURL+UserServiceCreateSavedLocationProcedure,
			connect.WithSchema(userServiceCreateSavedLocationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listSavedLocations: connect.NewClient[v1alpha1.ListSavedLocationsRequest, v1alpha1.ListSavedLocationsResponse](
			httpClient,
			baseURL+UserServiceListSavedLocationsProcedure,
			connect.WithSchema(userServiceListSavedLocationsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSavedLocation: connect.NewClient[v1alpha1.GetSavedLocationRequest, v1alpha1.GetSavedLocationResponse](
			httpClient,
			baseURL+UserServiceGetSavedLocationProcedure,
			connect.WithSchema(userServiceGetSavedLocationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateSavedLocation: connect.NewClient[v1alpha1.UpdateSavedLocationRequest, v1alpha1.UpdateSavedLocationResponse](
			httpClient,
			baseURL+UserServiceUpdateSavedLocationProcedure,
			connect.WithSchema(userServiceUpdateSavedLocationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteSavedLocation: connect.NewClient[v1alpha1.DeleteSavedLocationRequest, v1alpha1.DeleteSavedLocationResponse](
			httpClient,
			baseURL+UserServiceDeleteSavedLocationProcedure,
			connect.WithSchema(userServiceDeleteSavedLocationMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	getUser             *connect.Client[v1alpha1.GetUserRequest, v1alpha1.GetUserResponse]
	updateUser          *connect.Client[v1alpha1.UpdateUserRequest, v1alpha1.UpdateUserResponse]
	deleteUser          *connect.Client[v1alpha1.DeleteUserRequest, v1alpha1.DeleteUserResponse]
	createSavedLocation *connect.Client[v1alpha1.CreateSavedLocationRequest, v1alpha1.CreateSavedLocationResponse]
	listSavedLocations  *connect.Client[v1alpha1.ListSavedLocationsRequest, v1alpha1.ListSavedLocationsResponse]
	getSavedLocation    *connect.Client[v1alpha1.GetSavedLocationRequest, v1alpha1.GetSavedLocationResponse]
	updateSavedLocation *connect.Client[v1alpha1.UpdateSavedLocationRequest, v1alpha1.UpdateSavedLocationResponse]
	deleteSavedLocation *connect.Client[v1alpha1.DeleteSavedLocationRequest, v1alpha1.DeleteSavedLocationResponse]
}

// GetUser calls ride.rider.v1alpha1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return c.getUser.CallUnary(ctx, req)
}

// UpdateUser calls ride.rider.v1alpha1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect.Request[v1alpha1.UpdateUserRequest]) (*connect.Response[v1alpha1.UpdateUserResponse], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// DeleteUser calls ride.rider.v1alpha1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// CreateSavedLocation calls ride.rider.v1alpha1.UserService.CreateSavedLocation.
func (c *userServiceClient) CreateSavedLocation(ctx context.Context, req *connect.Request[v1alpha1.CreateSavedLocationRequest]) (*connect.Response[v1alpha1.CreateSavedLocationResponse], error) {
	return c.createSavedLocation.CallUnary(ctx, req)
}

// ListSavedLocations calls ride.rider.v1alpha1.UserService.ListSavedLocations.
func (c *userServiceClient) ListSavedLocations(ctx context.Context, req *connect.Request[v1alpha1.ListSavedLocationsRequest]) (*connect.Response[v1alpha1.ListSavedLocationsResponse], error) {
	return c.listSavedLocations.CallUnary(ctx, req)
}

// GetSavedLocation calls ride.rider.v1alpha1.UserService.GetSavedLocation.
func (c *userServiceClient) GetSavedLocation(ctx context.Context, req *connect.Request[v1alpha1.GetSavedLocationRequest]) (*connect.Response[v1alpha1.GetSavedLocationResponse], error) {
	return c.getSavedLocation.CallUnary(ctx, req)
}

// UpdateSavedLocation calls ride.rider.v1alpha1.UserService.UpdateSavedLocation.
func (c *userServiceClient) UpdateSavedLocation(ctx context.Context, req *connect.Request[v1alpha1.UpdateSavedLocationRequest]) (*connect.Response[v1alpha1.UpdateSavedLocationResponse], error) {
	return c.updateSavedLocation.CallUnary(ctx, req)
}

// DeleteSavedLocation calls ride.rider.v1alpha1.UserService.DeleteSavedLocation.
func (c *userServiceClient) DeleteSavedLocation(ctx context.Context, req *connect.Request[v1alpha1.DeleteSavedLocationRequest]) (*connect.Response[v1alpha1.DeleteSavedLocationResponse], error) {
	return c.deleteSavedLocation.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the ride.rider.v1alpha1.UserService service.
type UserServiceHandler interface {
	GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error)
	UpdateUser(context.Context, *connect.Request[v1alpha1.UpdateUserRequest]) (*connect.Response[v1alpha1.UpdateUserResponse], error)
	DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error)
	CreateSavedLocation(context.Context, *connect.Request[v1alpha1.CreateSavedLocationRequest]) (*connect.Response[v1alpha1.CreateSavedLocationResponse], error)
	ListSavedLocations(context.Context, *connect.Request[v1alpha1.ListSavedLocationsRequest]) (*connect.Response[v1alpha1.ListSavedLocationsResponse], error)
	GetSavedLocation(context.Context, *connect.Request[v1alpha1.GetSavedLocationRequest]) (*connect.Response[v1alpha1.GetSavedLocationResponse], error)
	UpdateSavedLocation(context.Context, *connect.Request[v1alpha1.UpdateSavedLocationRequest]) (*connect.Response[v1alpha1.UpdateSavedLocationResponse], error)
	DeleteSavedLocation(context.Context, *connect.Request[v1alpha1.DeleteSavedLocationRequest]) (*connect.Response[v1alpha1.DeleteSavedLocationResponse], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceGetUserHandler := connect.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceUpdateUserHandler := connect.NewUnaryHandler(
		UserServiceUpdateUserProcedure,
		svc.UpdateUser,
		connect.WithSchema(userServiceUpdateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserHandler := connect.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(userServiceDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateSavedLocationHandler := connect.NewUnaryHandler(
		UserServiceCreateSavedLocationProcedure,
		svc.CreateSavedLocation,
		connect.WithSchema(userServiceCreateSavedLocationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceListSavedLocationsHandler := connect.NewUnaryHandler(
		UserServiceListSavedLocationsProcedure,
		svc.ListSavedLocations,
		connect.WithSchema(userServiceListSavedLocationsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetSavedLocationHandler := connect.NewUnaryHandler(
		UserServiceGetSavedLocationProcedure,
		svc.GetSavedLocation,
		connect.WithSchema(userServiceGetSavedLocationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceUpdateSavedLocationHandler := connect.NewUnaryHandler(
		UserServiceUpdateSavedLocationProcedure,
		svc.UpdateSavedLocation,
		connect.WithSchema(userServiceUpdateSavedLocationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteSavedLocationHandler := connect.NewUnaryHandler(
		UserServiceDeleteSavedLocationProcedure,
		svc.DeleteSavedLocation,
		connect.WithSchema(userServiceDeleteSavedLocationMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/ride.rider.v1alpha1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserProcedure:
			userServiceUpdateUserHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		case UserServiceCreateSavedLocationProcedure:
			userServiceCreateSavedLocationHandler.ServeHTTP(w, r)
		case UserServiceListSavedLocationsProcedure:
			userServiceListSavedLocationsHandler.ServeHTTP(w, r)
		case UserServiceGetSavedLocationProcedure:
			userServiceGetSavedLocationHandler.ServeHTTP(w, r)
		case UserServiceUpdateSavedLocationProcedure:
			userServiceUpdateSavedLocationHandler.ServeHTTP(w, r)
		case UserServiceDeleteSavedLocationProcedure:
			userServiceDeleteSavedLocationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect.Request[v1alpha1.GetUserRequest]) (*connect.Response[v1alpha1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect.Request[v1alpha1.UpdateUserRequest]) (*connect.Response[v1alpha1.UpdateUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect.Request[v1alpha1.DeleteUserRequest]) (*connect.Response[v1alpha1.DeleteUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.DeleteUser is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateSavedLocation(context.Context, *connect.Request[v1alpha1.CreateSavedLocationRequest]) (*connect.Response[v1alpha1.CreateSavedLocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.CreateSavedLocation is not implemented"))
}

func (UnimplementedUserServiceHandler) ListSavedLocations(context.Context, *connect.Request[v1alpha1.ListSavedLocationsRequest]) (*connect.Response[v1alpha1.ListSavedLocationsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.ListSavedLocations is not implemented"))
}

func (UnimplementedUserServiceHandler) GetSavedLocation(context.Context, *connect.Request[v1alpha1.GetSavedLocationRequest]) (*connect.Response[v1alpha1.GetSavedLocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.GetSavedLocation is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateSavedLocation(context.Context, *connect.Request[v1alpha1.UpdateSavedLocationRequest]) (*connect.Response[v1alpha1.UpdateSavedLocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.UpdateSavedLocation is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteSavedLocation(context.Context, *connect.Request[v1alpha1.DeleteSavedLocationRequest]) (*connect.Response[v1alpha1.DeleteSavedLocationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("ride.rider.v1alpha1.UserService.DeleteSavedLocation is not implemented"))
}
