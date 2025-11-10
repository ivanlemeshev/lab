package repos

import (
	"errors"
	"log"

	"github.com/google/uuid"

	users "lem/go/auth/internal/users"
)

var ErrNotFound = errors.New("not found")

type User struct {
	ID       string
	Username string
	Password string
	Role     users.Role
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
}

type FakeUsersRepository struct {
	users []*User
}

func NewFakeUsersRepository() *FakeUsersRepository {
	adminUserID, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate admin user ID")
	}

	userID, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate user ID")
	}

	unknownUserID, err := uuid.NewV7()
	if err != nil {
		panic("failed to generate unknown user ID")
	}

	fakeUsers := []*User{
		{
			ID:       adminUserID.String(),
			Username: "admin",
			Password: "admin",
			Role:     users.RoleAdmin,
		},
		{
			ID:       userID.String(),
			Username: "user",
			Password: "user",
			Role:     users.RoleUser,
		},
		{
			ID:       unknownUserID.String(),
			Username: "unknown",
			Password: "unknown",
		},
	}

	for _, u := range fakeUsers {
		log.Printf("Created fake user: ID=%s, Username=%s\n", u.ID, u.Username)
	}

	return &FakeUsersRepository{
		users: fakeUsers,
	}
}

func (r *FakeUsersRepository) FindByUsername(username string) (*User, error) {
	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, ErrNotFound
}

func (r *FakeUsersRepository) FindByID(userID string) (*User, error) {
	for _, user := range r.users {
		if user.ID == userID {
			return user, nil
		}
	}
	return nil, ErrNotFound
}
