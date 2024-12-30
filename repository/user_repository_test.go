package repository_test

import (
	"context"
	"testing"
	"time"
	"user-crud/model"
	"user-crud/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	dateStr := "2024-12-30"
	layout := "2006-01-02"
	createdAt, err := time.Parse(layout, dateStr)
	if err != nil {
		t.Fatalf("Failed to parse date: %v", err)
	}

	repo := repository.NewUserRepository(db)
	user := model.User{
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "1234567890",
		CreatedAt:   createdAt,
	}

	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO users \\(id, name, surname, email, phone_number, created_at\\) VALUES \\(\\?, \\?, \\?, \\?, \\?, \\?\\)$").
		WithArgs(sqlmock.AnyArg(), user.Name, user.Surname, user.Email, user.PhoneNumber, user.CreatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.Save(context.Background(), user)
	assert.NoError(t, err, "Expected no error while saving user")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}


func TestFindById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userId := uuid.New()

	rows := sqlmock.NewRows([]string{"id", "name", "surname", "email", "phone_number", "created_at"}).
		AddRow(userId, "John", "Doe", "john.doe@example.com", "1234567890", time.Now())

	mock.ExpectBegin()
	mock.ExpectQuery("SELECT id, name, surname, email, phone_number, created_at FROM users WHERE id = ?").
		WithArgs(userId).
		WillReturnRows(rows)
	mock.ExpectCommit()

	user, err := repo.FindById(context.Background(), userId)
	assert.NoError(t, err, "Expected no error while finding user by ID")
	assert.Equal(t, userId, user.Id, "Expected user ID to match")


	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unmet expectations: %v", err)
	}
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userId := uuid.New()
	mock.ExpectBegin()

	mock.ExpectExec("DELETE FROM users WHERE id = ?").
		WithArgs(userId).
		WillReturnResult(sqlmock.NewResult(0, 1))


	mock.ExpectCommit()

	err = repo.Delete(context.Background(), userId)

	assert.NoError(t, err, "Expected no error during delete operation")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	users := []model.User{
		{
			Id:          uuid.New(),
			Name:        "John",
			Surname:     "Doe",
			Email:       "john.doe@example.com",
			PhoneNumber: "1234567890",
			CreatedAt:   time.Now(),
		},
		{
			Id:          uuid.New(),
			Name:        "Jane",
			Surname:     "Doe",
			Email:       "jane.doe@example.com",
			PhoneNumber: "9876543210",
			CreatedAt:   time.Now(),
		},
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT id, name, surname, email, phone_number, created_at FROM users").
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "surname", "email", "phone_number", "created_at"}).
				AddRow(users[0].Id, users[0].Name, users[0].Surname, users[0].Email, users[0].PhoneNumber, users[0].CreatedAt).
				AddRow(users[1].Id, users[1].Name, users[1].Surname, users[1].Email, users[1].PhoneNumber, users[1].CreatedAt),
		)


	mock.ExpectCommit()

	result, err := repo.FindAll(context.Background())

	assert.NoError(t, err, "Expected no error during FindAll operation")
	assert.Equal(t, 2, len(result), "Returned user count should match")
	assert.Equal(t, users[0].Id, result[0].Id, "First user ID should match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}


func TestFindByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	email := "john.doe@example.com"
	user := model.User{
		Id:          uuid.New(),
		Name:        "John",
		Surname:     "Doe",
		Email:       email,
		PhoneNumber: "1234567890",
		CreatedAt:   time.Now(),
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT id, name, surname, email, phone_number, created_at FROM users WHERE email = ?").
		WithArgs(email).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "surname", "email", "phone_number", "created_at"}).
				AddRow(user.Id, user.Name, user.Surname, user.Email, user.PhoneNumber, user.CreatedAt),
		)

	mock.ExpectCommit()

	result, err := repo.FindByEmail(context.Background(), email)

	assert.NoError(t, err, "Expected no error during FindByEmail operation")
	assert.Equal(t, user.Email, result.Email, "Returned user email should match")
	assert.Equal(t, user.Id, result.Id, "Returned user ID should match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}


func TestFindByPhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database connection: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	phoneNumber := "1234567890"
	user := model.User{
		Id:          uuid.New(),
		Name:        "John",
		Surname:     "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
	}

	mock.ExpectBegin()

	mock.ExpectQuery("SELECT id, name, surname, email, phone_number, created_at FROM users WHERE phone_number = ?").
		WithArgs(phoneNumber).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "surname", "email", "phone_number", "created_at"}).
				AddRow(user.Id, user.Name, user.Surname, user.Email, user.PhoneNumber, user.CreatedAt),
		)

	mock.ExpectCommit()

	result, err := repo.FindByPhoneNumber(context.Background(), phoneNumber)

	assert.NoError(t, err, "Expected no error during FindByPhoneNumber operation")
	assert.Equal(t, user.PhoneNumber, result.PhoneNumber, "Returned user phone number should match")
	assert.Equal(t, user.Id, result.Id, "Returned user ID should match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Mock expectations were not met: %v", err)
	}
}
