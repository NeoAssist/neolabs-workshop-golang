package interfaces

import "github.com/xesina/golang-echo-realworld-example-app/model"

type Store interface {
	GetByID(uint) (*model.Wanted, error)
	GetByEmail(string) (*model.Wanted, error)
	GetByUsername(string) (*model.Wanted, error)
	Create(*model.Wanted) error
	Update(*model.Wanted) error
	AddFollower(user *model.Wanted, followerID uint) error
	RemoveFollower(user *model.Wanted, followerID uint) error
	IsFollower(userID, followerID uint) (bool, error)
}
