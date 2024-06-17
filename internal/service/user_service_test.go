package service

import (
	context "context"
	"testing"

	"github.com/alex-v08/inventory-fw/internal/entity"
	"github.com/alex-v08/inventory-fw/internal/repository"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name    string
		Email   string
		UserName string
		Password string
		ExpectedError error
	}{
		{
			Name: "RegisterUser_Success",
			Email: "test@test.com",
			UserName: "test",
			Password: "validpassword",
			ExpectedError: nil,
		},
		{ 
			Name: "RegisterUser_UserAlreadyExists",
			Email: "test@exist.com",
			UserName: "test",
			Password: "validpassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	ctx := context.Background()

	repo := &repository.MockRepository{}
	
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil,nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exist.com").Return(&entity.User{Email: "test@exist.com"},nil)
	repo.On("SaveUser", mock.Anything,mock.Anything,mock.Anything,mock.Anything).Return(nil) 




	for i:= range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			
				t.Parallel()
				repo.Mock.Test(t)

				s:= New(repo)

				err:=s.RegisterUser(ctx, tc.Email, tc.UserName, tc.Password)

				if err != tc.ExpectedError {
					t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
				}


		})

}
	}





