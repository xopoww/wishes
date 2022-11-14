package fixtures

import "github.com/xopoww/wishes/internal/models"

func User() *models.User {
	return &models.User{
		ID: 1,
		Name: "user",
	}
}

func TwoUsers() (alice, bob *models.User) {
	alice = &models.User{
		ID: 1,
		Name: "alice",
	}
	bob = &models.User{
		ID: 2,
		Name: "bob",
	}
	return
}