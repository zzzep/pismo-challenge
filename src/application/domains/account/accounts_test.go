package account

import (
	"github.com/zzzep/pismo-challenge/src/helpers"
	"net/http"
	"testing"
)

func TestAccount_CreateAccount(t *testing.T) {
	testCases := helpers.TestCases{
		{
			Name:       "New Account with no error on database",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{AccRepo: helpers.NewMockAccount(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
		{
			Name:       "New Account with database error",
			StatusCode: http.StatusInternalServerError,
			Repos:      helpers.Repos{AccRepo: helpers.NewMockAccount(true)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			a := &Account{
				repo: tt.Repos.AccRepo,
			}
			a.CreateAccount(tt.Args.C)

		})
	}
}

func TestAccount_GetAccount(t *testing.T) {
	testCases := helpers.TestCases{
		{
			Name:       "Get Account with no error on database",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{AccRepo: helpers.NewMockAccount(false)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
		{
			Name:       "Get Account with database error",
			StatusCode: http.StatusInternalServerError,
			Repos:      helpers.Repos{AccRepo: helpers.NewMockAccount(true)},
			Args:       helpers.GinArgs{C: helpers.CreatePostContext()},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			account := &Account{
				repo: tc.Repos.AccRepo,
			}
			account.GetAccount(tc.Args.C)
		})
	}
}

func TestNewAccount(t *testing.T) {
	mockRepo := helpers.NewMockAccount(false)
	tests := []struct {
		Name       string
		StatusCode int
		Repos      helpers.Repos
	}{
		{
			Name:       "New Account Instance correctly",
			StatusCode: http.StatusOK,
			Repos:      helpers.Repos{AccRepo: mockRepo},
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			got := NewAccount(tc.Repos.AccRepo)
			if got.repo != mockRepo {
				t.Errorf("NewAccount() = %v, want %v", got, mockRepo)
			}
		})
	}
}
