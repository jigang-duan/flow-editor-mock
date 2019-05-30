package services

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/repositories"
)

type TypeGroupService interface {
	GetAll() []datamodels.TypeGroup
}

func NewTypeGroupService(repo repositories.TypeGroupRepository) TypeGroupService {
	return &typeGroupService{
		repo: repo,
	}
}

type typeGroupService struct {
	repo repositories.TypeGroupRepository
}

func (s *typeGroupService) GetAll() []datamodels.TypeGroup {
	return s.repo.SelectAll()
}
