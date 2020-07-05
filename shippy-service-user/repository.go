package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/onrooftop/shippy/shippy-service-user/proto/user"
	uuid "github.com/satori/go.uuid"
)

// User -
type User struct {
	ID       string `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Company  string `sql:"company"`
	Password string `sql:"password"`
}

// Repository -
type repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

// MarshalUserCollection -
func MarshalUserCollection(users []*pb.User) []*User {
	u := make([]*User, len(users))
	for _, user := range users {
		u = append(u, MarshalUser(user))
	}
	return u
}

// MarshalUser -
func MarshalUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Company:  user.Company,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}
}

// UnmarshalUserCollection -
func UnmarshalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, len(users))
	for _, user := range users {
		u = append(u, UnmarshalUser(user))
	}
	return u
}

// UnmarshalUser -
func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Company:  user.Company,
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	}
}

// PostgresRepository -
type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository -
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

// GetAll -
func (repository *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := repository.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}
	return users, nil
}

// Get -
func (repository *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	user := &User{}
	if err := repository.db.GetContext(ctx, user, "select * from users where id = $1", id); err != nil {
		return nil, err
	}
	return user, nil
}

// Create -
func (repository *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	query := "insert into users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err := repository.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

// GetByEmail -
func (repository *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := "select * from users where email = $1"
	user := &User{}
	if err := repository.db.GetContext(ctx, user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}
