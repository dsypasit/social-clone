package user

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	input := UserCreated{
		"ong", "e45680fb-29e3-4679-ab45-a95c7d9a18f4", "a@email.com", "1234",
	}
	var want int64 = 1

	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, "Expected nil from created sqlmock")
	defer db.Close()

	mock.ExpectQuery("SELECT username FROM app_user ").
		WithArgs(input.Username).
		WillReturnRows(sqlmock.NewRows([]string{"username"}))

	mock.ExpectQuery("INSERT INTO app_user").
		WithArgs(input.UUID, input.Username, input.Email, input.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	userRepo := NewUserRepository(db)

	actual, err := userRepo.CreateUser(input)
	assert.Nilf(t, err, "Unexpected error from create user %v", err)
	assert.Equalf(t, want, actual, "Expected %v but got %v", want, actual)
}

func TestCreateUser_DuplicateUser(t *testing.T) {
	input := UserCreated{
		"ong", "e45680fb-29e3-4679-ab45-a95c7d9a18f4", "a@email.com", "1234",
	}

	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, "Expected nil from created sqlmock")
	defer db.Close()

	mock.ExpectQuery("SELECT username FROM app_user ").
		WithArgs(input.Username).
		WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow("ong"))

	userRepo := NewUserRepository(db)

	_, err = userRepo.CreateUser(input)
	assert.Equalf(t, ErrDupUsername, err, "Unexpected error: %v", err)
}

func TestGetUserByUsername(t *testing.T) {
	input := "ong"
	CreatedAt := time.Now()
	want := User{
		ID:        1,
		UUID:      "0870a9ce-78d2-463d-bd88-ad0a0eee0e81",
		Username:  "ong",
		Email:     "a@gmail.com",
		CreatedAt: CreatedAt,
	}

	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, "Expected nil from created sqlmock")
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "uuid", "username", "email", "created_at"}).
		AddRow(1, "0870a9ce-78d2-463d-bd88-ad0a0eee0e81", "ong", "a@gmail.com", CreatedAt)

	mock.ExpectQuery("SELECT").
		WithArgs(input).
		WillReturnRows(row)

	userRepo := NewUserRepository(db)

	actual, err := userRepo.GetUserByUsername(input)
	assert.Nilf(t, err, "Unexpected error from get user %v", err)
	assert.Equalf(t, want, actual, "Expected %v but got %v", want, actual)
}

func TestGetUserByUUID(t *testing.T) {
	CreatedAt := time.Now()
	testTable := []struct {
		title    string
		input    string
		rows     sqlmock.Rows
		wantUser User
		wantErr  error
	}{
		{
			"should return user", "0870a9ce-78d2-463d-bd88-ad0a0eee0e81",
			*sqlmock.NewRows([]string{"id", "uuid", "username", "email", "created_at"}).AddRow("1", "0870a9ce-78d2-463d-bd88-ad0a0eee0e81", "ong", "a@gmail.com", CreatedAt),
			User{
				ID:        1,
				UUID:      "0870a9ce-78d2-463d-bd88-ad0a0eee0e81",
				Username:  "ong",
				Email:     "a@gmail.com",
				CreatedAt: CreatedAt,
			},
			nil,
		},
		{
			"should not found", "0870a9ce-78d2-463d-bd88-ad0a0eee0e81", *sqlmock.NewRows([]string{"id", "uuid", "username", "email", "created_at"}),
			User{},
			ErrUserNotFound,
		},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Expected nil from created sqlmock")
			defer db.Close()

			mock.ExpectQuery("SELECT").
				WithArgs(v.input).
				WillReturnRows(&v.rows)

			userRepo := NewUserRepository(db)

			actual, err := userRepo.GetUserByUUID(v.input)
			assert.Equalf(t, v.wantErr, err, "Unexpected error from get user %v", err)
			assert.Equalf(t, v.wantUser, actual, "Expected %v but got %v", v.wantUser, actual)
		})
	}
}

func TestGetUserByUsername_NotFound(t *testing.T) {
	input := "ong1"
	want := sql.ErrNoRows

	db, mock, err := sqlmock.New()
	assert.Nilf(t, err, "Expected nil from created sqlmock")
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "uuid", "username", "email", "created_at"}).CloseError(sql.ErrNoRows)

	mock.ExpectQuery("SELECT").
		WithArgs(input).
		WillReturnRows(row)

	userRepo := NewUserRepository(db)

	_, err = userRepo.GetUserByUsername(input)
	assert.Equal(t, want, err, "Unexpected error from get user by username %v", err)
}

func TestIsDuplicateUsername(t *testing.T) {
	testTable := []struct {
		title string
		input string
		row   sqlmock.Rows
		want  error
	}{
		{"should duplicate", "ong1", *sqlmock.NewRows([]string{"username"}).AddRow("ong1"), ErrDupUsername},
		{"should not duplicate", "ong1", *sqlmock.NewRows([]string{"username"}), nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Expected nil from created sqlmock")
			defer db.Close()

			mock.ExpectQuery("SELECT username FROM app_user").
				WithArgs(v.input).
				WillReturnRows(&v.row)

			userRepo := NewUserRepository(db)

			err = userRepo.IsDuplicateUsername(v.input)
			assert.Equal(t, v.want, err, "Expected true from unique user %v", err)
		})
	}
}

func TestGetPasswordByUsername(t *testing.T) {
	testTable := []struct {
		title    string
		input    string
		rows     sqlmock.Rows
		wantPass string
		wantErr  error
	}{
		{"should got password", "ong1", *sqlmock.NewRows([]string{"password"}).AddRow("1234"), "1234", nil},
		{"user not found", "ong1", *sqlmock.NewRows([]string{"password"}), "", ErrUserNotFound},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Unexpected error from sqlmock: %v", err)
			defer db.Close()

			mock.ExpectQuery("SELECT password FROM app_user").WithArgs(v.input).WillReturnRows(&v.rows)

			userRepo := NewUserRepository(db)

			pass, err := userRepo.GetPasswordByUsername(v.input)

			assert.Equalf(t, v.wantErr, err, "Unexpected error : %v", err)
			assert.Equalf(t, v.wantPass, pass, "Expected %v but got %v", v.wantPass, pass)
		})
	}
}

func TestGetUserUUIDByUsername(t *testing.T) {
	testTable := []struct {
		title    string
		input    string
		rows     sqlmock.Rows
		wantUUID string
		wantErr  error
	}{
		{"should got uuid", "ong1", *sqlmock.NewRows([]string{"uuid"}).AddRow("fdddfba8-0ac2-45d5-ac2b-8660b69de352"), "fdddfba8-0ac2-45d5-ac2b-8660b69de352", nil},
		{"should not found user", "ong1", *sqlmock.NewRows([]string{"uuid"}), "", ErrUserNotFound},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Unexpected error from sqlmock: %v", err)
			defer db.Close()

			mock.ExpectQuery("SELECT uuid FROM app_user").WithArgs(v.input).WillReturnRows(&v.rows)

			userRepo := NewUserRepository(db)

			userUUID, err := userRepo.GetUserUUIDByUsername(v.input)

			assert.Equalf(t, v.wantErr, err, "Unexpected error from get user uuid: %v", err)
			assert.Equalf(t, v.wantUUID, userUUID, "Want %v but got %v", v.wantUUID, userUUID)
		})
	}
}

func TestUpdateEmailByUUID(t *testing.T) {
	testTable := []struct {
		title   string
		uuid    string
		email   string
		rows    driver.Result
		wantErr error
	}{
		{"should update email", "ad8340fb-656f-492b-aaac-aa773bab7520", "a@gmail.com", sqlmock.NewResult(1, 1), nil},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Unexpected error from sqlmock: %v", err)
			defer db.Close()

			mock.ExpectExec("UPDATE").WithArgs(v.email, v.uuid).WillReturnResult(v.rows)

			userRepo := NewUserRepository(db)

			err = userRepo.UpdateEmailByUUID(v.uuid, v.email)

			assert.Equalf(t, v.wantErr, err, "Unexpected error from get user uuid: %v", err)
		})
	}
}

func TestUpdateEmailByUUID_UserNotFound(t *testing.T) {
	testTable := []struct {
		title   string
		uuid    string
		email   string
		rows    driver.Result
		wantErr error
	}{
		{"should update email", "ad8340fb-656f-492b-aaac-aa773bab7520", "a@gmail.com", sqlmock.NewResult(1, 0), ErrUserNotFound},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			assert.Nilf(t, err, "Unexpected error from sqlmock: %v", err)
			defer db.Close()

			mock.ExpectExec("UPDATE").WithArgs(v.email, v.uuid).WillReturnResult(v.rows)

			userRepo := NewUserRepository(db)

			err = userRepo.UpdateEmailByUUID(v.uuid, v.email)

			assert.Equalf(t, v.wantErr, err, "Unexpected error from get user uuid: %v", err)
		})
	}
}
