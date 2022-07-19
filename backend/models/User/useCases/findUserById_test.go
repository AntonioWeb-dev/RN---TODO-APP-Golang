package useCases

import (
	mock "api/infra/repository/user"
	"testing"
)

func TestFindUserById(t *testing.T) {
	userRepoMock := mock.InitUserRepositoryMock()
	useCase := InitFindUserByIdCase(userRepoMock)
	_, _, err := useCase.Handler("userId")
	if err != nil {
		t.Errorf("User not found")
	}
}
