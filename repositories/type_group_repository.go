package repositories

import "flow-editor-mock/datamodels"

type TypeGroupRepository interface {
	SelectAll() (results []datamodels.TypeGroup)
}

func NewTypeGroupRepository(source []datamodels.TypeGroup) TypeGroupRepository {
	return &typeGroupMemoryRepository{source:source}
}

type typeGroupMemoryRepository struct {
	source []datamodels.TypeGroup
}

func (r *typeGroupMemoryRepository) SelectAll() (results []datamodels.TypeGroup) {
	return r.source
}
