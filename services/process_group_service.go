package services

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/repositories"
)

type ProcessGroupService interface {
	GetByID(id string) (datamodels.ProcessGroup, bool)
	CreateProcessorByID(id string, processor datamodels.Processor) (datamodels.ProcessGroup, error)
}

func NewProcessGroupService(repo repositories.ProcessGroupRepository) ProcessGroupService {
	return &processGroupService{
		repo: repo,
	}
}

type processGroupService struct {
	repo repositories.ProcessGroupRepository
}

func (s *processGroupService) CreateProcessorByID(id string, processor datamodels.Processor) (datamodels.ProcessGroup, error) {
	return s.repo.InsertProcessor(id, processor)
}

func (s *processGroupService) GetByID(id string) (datamodels.ProcessGroup, bool) {
	return s.repo.Select(id)
}
