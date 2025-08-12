package interfaces

import db "github.com/FrancoMusolino/go-todoapp/db/schema"

type IUsersRepo interface {
	CreateUser(u *db.User) (*db.User, error)
}
