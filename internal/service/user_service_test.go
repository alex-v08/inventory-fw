package service

import (
	context "context"
	"os"
	"testing"

	"github.com/alex-v08/inventory-fw/internal/encryption"
	"github.com/alex-v08/inventory-fw/internal/entity"
	"github.com/alex-v08/inventory-fw/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var s Service

func TestMain(m *testing.M) {

	validpassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validpassword)
	u := &entity.User{
		Email:    "test@exists.com",
		Password: encryptedPassword,
	}

	repo = &repository.MockRepository{}

	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(&u, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	s = New(repo)
	code := m.Run()
	os.Exit(code)

}

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validpassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validpassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {

			t.Parallel()
			repo.Mock.Test(t)

			err := s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}

		})

	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validpassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidCredentials",
			Email:         "test@exists.com",
			Password:      "invalidpassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {

			t.Parallel()
			repo.Mock.Test(t)

			_, err := s.LoginUser(ctx, tc.Email, tc.Password)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})

	}
}
