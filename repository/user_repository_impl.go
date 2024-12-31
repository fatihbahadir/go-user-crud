package repository

import (
	"context"
	"database/sql"
	"fmt"
	"user-crud/helper"
	"user-crud/model"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{Db: db}
}

func (repo *UserRepositoryImpl) Save(ctx context.Context, user model.User) error {
	user.Id = uuid.New()

	tx, err := repo.Db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer helper.CommitOrRollback(tx)

	SQL := "INSERT INTO users (id, name, surname, email, phone_number, created_at) VALUES (?, ?, ?, ?, ?, ?)"
	_, err = tx.Exec(SQL, user.Id, user.Name, user.Surname, user.Email, user.PhoneNumber, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %v", err)
	}
	return nil
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, userId uuid.UUID, user model.User) error {
	tx, err := repo.Db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE users SET name = ?, surname = ?, email = ?, phone_number = ? WHERE id = ?"
	_, err = tx.Exec(SQL, user.Name, user.Surname, user.Email, user.PhoneNumber, userId)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %v", err)
	}

	return nil
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, userId uuid.UUID) error {
	tx, err := repo.Db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %v", err)
	}

	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM users WHERE id = ?"
	_, err = tx.Exec(SQL, userId)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %v", err)
	}

	return nil
}

func (repo *UserRepositoryImpl) FindById(ctx context.Context, userId uuid.UUID) (model.User, error) {
	tx, err := repo.Db.Begin()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, surname, email, phone_number, created_at FROM users WHERE id = ?"
	result, err := tx.QueryContext(ctx, SQL, userId)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to execute query to find user by id: %w", err)
	}
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.PhoneNumber, &user.CreatedAt)
		if err != nil {
			return model.User{}, fmt.Errorf("failed to scan user data: %w", err)
		}
		return user, nil
	} else {
		return model.User{}, fmt.Errorf("user with id %s not found", userId)
	}
}

func (repo *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	tx, err := repo.Db.Begin()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, surname, email, phone_number, created_at FROM users WHERE email = ?"
	result, err := tx.QueryContext(ctx, SQL, email)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to execute query to find user by email: %w", err)
	}
	defer result.Close()

	user := model.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.PhoneNumber, &user.CreatedAt)
		if err != nil {
			return model.User{}, fmt.Errorf("failed to scan user data: %w", err)
		}
		return user, nil
	}

	return model.User{}, nil
}

func (repo *UserRepositoryImpl) FindByPhoneNumber(ctx context.Context, phoneNumber string) (model.User, error) {
	tx, err := repo.Db.Begin()
	if err != nil {
		return model.User{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, surname, email, phone_number, created_at FROM users WHERE phone_number = ?"
	result, err := tx.QueryContext(ctx, SQL, phoneNumber)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to execute query to find user by phone number: %w", err)
	}
	defer result.Close()

	user := model.User{}
	if result.Next() {
		err := result.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.PhoneNumber, &user.CreatedAt)
		if err != nil {
			return model.User{}, fmt.Errorf("failed to scan user data: %w", err)
		}
		return user, nil
	}

	return model.User{}, nil
}

func (repo *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	tx, err := repo.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name, surname, email, phone_number, created_at FROM users"
	result, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query to find all users: %w", err)
	}
	defer result.Close()

	var users []model.User
	for result.Next() {
		user := model.User{}
		err := result.Scan(&user.Id, &user.Name, &user.Surname, &user.Email, &user.PhoneNumber, &user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user data: %w", err)
		}
		users = append(users, user)
	}

	return users, nil
}