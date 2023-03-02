package user

import (
	"github.com/yonsina94/go-generics/commons/service"
	"gorm.io/gorm"
)

type UserService struct {
	service.BaseService[User]
}

func (u *UserService) GetByEmail(email string) (user User, err error) {
	if uu, er := u.GetBy(User{Email: email}); er == nil {
		user = uu[0]
		return
	} else {
		err = er
		return
	}
}

func (u *UserService) GetByUserName(username string) (user User, err error) {
	if uu, er := u.GetBy(User{UserName: username}); er == nil {
		user = uu[0]
		return
	} else {
		err = er
		return
	}
}

func New(db *gorm.DB) *UserService {
	return &UserService{
		BaseService: service.BaseService[User]{DB: db},
	}
}
