package services

import (
	"errors"
	"flow-editor-mock/datamodels"
	"flow-editor-mock/repositories"
)

type ProcessGroupService interface {
	GetByID(id string) (datamodels.ProcessGroup, bool)

	CreateProcessorByID(id string, processor datamodels.Processor) (datamodels.ProcessGroup, error)
	UpdateProcessorsByID(id string, processors []datamodels.Processor) (datamodels.ProcessGroup, error)

	CreateConnectionByID(id string, connection datamodels.Connection) (datamodels.ProcessGroup, error)

	CloneSnippet(id string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error)

	CreateProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error)
	DeleteSnippet(parentID string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error)

	UpdateSnippet(id string, processors []datamodels.Processor, connections []datamodels.Connection, processGroups []datamodels.ProcessGroup) (datamodels.ProcessGroup, error)
}

func NewProcessGroupService(repo repositories.ProcessGroupRepository, typeRepo repositories.TypeGroupRepository) ProcessGroupService {
	return &processGroupService{
		repo:     repo,
		typeRepo: typeRepo,
	}
}

type processGroupService struct {
	repo     repositories.ProcessGroupRepository
	typeRepo repositories.TypeGroupRepository
}

func (s *processGroupService) UpdateSnippet(id string, processors []datamodels.Processor, connections []datamodels.Connection, processGroups []datamodels.ProcessGroup) (datamodels.ProcessGroup, error) {
	return s.repo.UpdateSnippet(id, processors, connections, processGroups)
}

func (s *processGroupService) DeleteSnippet(parentID string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error) {
	return s.repo.DeleteSnippet(parentID, processors, connections, processGroups)
}

func (s *processGroupService) CreateProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	return s.repo.InsertProcessGroup(parentID, processors, connections)
}

func (s *processGroupService) CloneSnippet(id string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	return s.repo.CloneSnippet(id, processors, connections)
}

func (s *processGroupService) CreateConnectionByID(id string, connection datamodels.Connection) (datamodels.ProcessGroup, error) {
	return s.repo.InsertConnection(id, connection)
}

func (s *processGroupService) UpdateProcessorsByID(id string, processors []datamodels.Processor) (datamodels.ProcessGroup, error) {
	return s.repo.UpdateProcessors(id, processors)
}

func (s *processGroupService) CreateProcessorByID(id string, processor datamodels.Processor) (datamodels.ProcessGroup, error) {
	item, found := s.typeRepo.SelectTypeByID(processor.TypeID)
	if !found {
		return datamodels.ProcessGroup{}, errors.New("没有找到对应的类型")
	}
	return s.repo.InsertProcessor(id, generateNodeByType(item, processor.Rect.X, processor.Rect.Y))
}

func (s *processGroupService) GetByID(id string) (datamodels.ProcessGroup, bool) {
	return s.repo.Select(id)
}

func generateNodeByType(typeItme datamodels.TypeItem, x int, y int) datamodels.Processor {
	//width := 140
	//height := 30
	const (
		maxX   = 5000
		maxY   = 5000
		width  = 140
		height = 30
		minx   = width/2 + 5
		miny   = height/2 + 5
		maxx   = maxX - width/2 - 5
		maxy   = maxY - height/2 - 5
	)
	var (
		cx = x
		cy = y
	)
	if x < minx {
		cx = minx
	} else if x > maxx {
		cx = maxx
	}
	if y < miny {
		cy = miny
	} else if y > maxy {
		cy = maxy
	}

	return datamodels.Processor{
		TypeID:      typeItme.ID,
		Label:       typeItme.Name,
		HasInput:    typeItme.HasInput,
		HasOutput:   typeItme.HasOutput,
		IconOnRight: typeItme.IconOnRight,
		Icon:        typeItme.Icon,
		Style: datamodels.Style{
			Color: typeItme.Color,
		},
		Rect: datamodels.Rect{
			X: cx,
			Y: cy,
			W: width,
			H: height,
		},
	}
}
