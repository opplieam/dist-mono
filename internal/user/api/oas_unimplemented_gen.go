// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateUser implements createUser operation.
//
// Create a new user.
//
// POST /user
func (UnimplementedHandler) CreateUser(ctx context.Context, req *User) (r CreateUserRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAllUsers implements getAllUsers operation.
//
// Get all users.
//
// GET /user
func (UnimplementedHandler) GetAllUsers(ctx context.Context) (r GetAllUsersRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserById implements getUserById operation.
//
// Get a user by ID.
//
// GET /user/{id}
func (UnimplementedHandler) GetUserById(ctx context.Context, params GetUserByIdParams) (r GetUserByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
