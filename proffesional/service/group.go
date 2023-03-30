package service

import (
	"proffesional/model"
	"proffesional/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) AddGroup(group model.AddGroup) (int, error) {
	return s.repo.AddGroup(group)
}

func (s *GroupService) GetAllGroups() ([]model.AllGroup, error) {
	return s.repo.GetAllGroups()
}

func (s *GroupService) ChangeSomeGroup(id int, group model.AddGroup) (model.AddGroup, error) {
	return s.repo.ChangeSomeGroup(id, group)
}

func (s *GroupService) DeleteGroup(groupId int) error {
	return s.repo.DeleteGroup(groupId)
}
