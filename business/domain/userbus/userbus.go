// Package userbus provides business access to user domain.
package userbus

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/ardanlabs/service/business/api/order"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("user not found")
	ErrUniqueEmail           = errors.New("email is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to perists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, usr User) error
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, usr User) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]User, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, userID uuid.UUID) (User, error)
	QueryByIDs(ctx context.Context, userID []uuid.UUID) ([]User, error)
	QueryByEmail(ctx context.Context, email mail.Address) (User, error)
}

// Domain manages the set of APIs for user access.
type Domain struct {
	log    *logger.Logger
	storer Storer
}

// NewDomain constructs a user core API for use.
func NewDomain(log *logger.Logger, storer Storer) *Domain {
	return &Domain{
		log:    log,
		storer: storer,
	}
}

// Create adds a new user to the system.
func (d *Domain) Create(ctx context.Context, nu NewUser) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generatefrompassword: %w", err)
	}

	now := time.Now()

	usr := User{
		ID:           uuid.New(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		Department:   nu.Department,
		Enabled:      true,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := d.storer.Create(ctx, usr); err != nil {
		return User{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

// Update modifies information about a user.
func (d *Domain) Update(ctx context.Context, usr User, uu UpdateUser) (User, error) {
	if uu.Name != nil {
		usr.Name = *uu.Name
	}

	if uu.Email != nil {
		usr.Email = *uu.Email
	}

	if uu.Roles != nil {
		usr.Roles = uu.Roles
	}

	if uu.Password != nil {
		pw, err := bcrypt.GenerateFromPassword([]byte(*uu.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, fmt.Errorf("generatefrompassword: %w", err)
		}
		usr.PasswordHash = pw
	}

	if uu.Department != nil {
		usr.Department = *uu.Department
	}

	if uu.Enabled != nil {
		usr.Enabled = *uu.Enabled
	}
	usr.DateUpdated = time.Now()

	if err := d.storer.Update(ctx, usr); err != nil {
		return User{}, fmt.Errorf("update: %w", err)
	}

	return usr, nil
}

// Delete removes the specified user.
func (d *Domain) Delete(ctx context.Context, usr User) error {
	if err := d.storer.Delete(ctx, usr); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

// Query retrieves a list of existing users.
func (d *Domain) Query(ctx context.Context, filter QueryFilter, orderBy order.By, pageNumber int, rowsPerPage int) ([]User, error) {
	if err := filter.Validate(); err != nil {
		return nil, err
	}

	users, err := d.storer.Query(ctx, filter, orderBy, pageNumber, rowsPerPage)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return users, nil
}

// Count returns the total number of users.
func (d *Domain) Count(ctx context.Context, filter QueryFilter) (int, error) {
	if err := filter.Validate(); err != nil {
		return 0, err
	}

	return d.storer.Count(ctx, filter)
}

// QueryByID finds the user by the specified ID.
func (d *Domain) QueryByID(ctx context.Context, userID uuid.UUID) (User, error) {
	user, err := d.storer.QueryByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("query: userID[%s]: %w", userID, err)
	}

	return user, nil
}

// QueryByIDs finds the users by a specified User IDs.
func (d *Domain) QueryByIDs(ctx context.Context, userIDs []uuid.UUID) ([]User, error) {
	user, err := d.storer.QueryByIDs(ctx, userIDs)
	if err != nil {
		return nil, fmt.Errorf("query: userIDs[%s]: %w", userIDs, err)
	}

	return user, nil
}

// QueryByEmail finds the user by a specified user email.
func (d *Domain) QueryByEmail(ctx context.Context, email mail.Address) (User, error) {
	user, err := d.storer.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	return user, nil
}

// Authenticate finds a user by their email and verifies their password. On
// success it returns a Claims User representing this user. The claims can be
// used to generate a token for future authentication.
func (d *Domain) Authenticate(ctx context.Context, email mail.Address, password string) (User, error) {
	usr, err := d.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	if err := bcrypt.CompareHashAndPassword(usr.PasswordHash, []byte(password)); err != nil {
		return User{}, fmt.Errorf("comparehashandpassword: %w", ErrAuthenticationFailure)
	}

	return usr, nil
}
