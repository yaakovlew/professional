package service

import (
	"proffesional/model"
	"proffesional/repository"
)

type Group interface {
	AddGroup(group model.AddGroup) (int, error)
	GetAllGroups() ([]model.AllGroup, error)
	ChangeSomeGroup(id int, group model.AddGroup) (model.AddGroup, error)
	DeleteGroup(groupId int) error
}

type Service struct {
	Group
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Group: NewGroupService(repo.Group),
	}
}
