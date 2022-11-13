package conv

import (
	"github.com/xopoww/wishes/internal/models"
	"github.com/xopoww/wishes/restapi/apimodels"
)

func User(u *apimodels.User) *models.User {
	return &models.User{
		Name:  string(u.Username),
		Fname: *u.Fname,
		Lname: *u.Lname,
	}
}

func SwagUser(u *models.User) *apimodels.User {
	return &apimodels.User{
		Username: apimodels.UserName(u.Name),
		UserInfo: apimodels.UserInfo{
			Fname: &u.Fname,
			Lname: &u.Lname,
		},
	}
}

func Client(p *apimodels.Principal) *models.User {
	return &models.User{
		ID:   p.ID,
		Name: p.Username,
	}
}

func Principal(c *models.User) *apimodels.Principal {
	return &apimodels.Principal{
		ID:       c.ID,
		Username: c.Name,
	}
}
