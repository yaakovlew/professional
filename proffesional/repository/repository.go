package repository

import (
	"github.com/jmoiron/sqlx"
	"proffesional/model"
)

type Group interface {
	AddGroup(group model.AddGroup) (int, error)
	GetAllGroups() ([]model.AllGroup, error)
	ChangeSomeGroup(id int, group model.AddGroup) (model.AddGroup, error)
	DeleteGroup(groupId int) error
}

type Repository struct {
	Group
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Group: NewGroupPostgres(db),
	}
}
