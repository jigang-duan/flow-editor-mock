package repositories

import "flow-editor-mock/datamodels"

type TypeGroupRepository interface {
	SelectAll() (results []datamodels.TypeGroup)
	SelectTypeByID(id string) (datamodels.TypeItem, bool)
}

func NewTypeGroupRepository(source []datamodels.TypeGroup) TypeGroupRepository {
	return &typeGroupMemoryRepository{source:source}
}

type typeGroupMemoryRepository struct {
	source []datamodels.TypeGroup
}

func (r *typeGroupMemoryRepository) SelectTypeByID(id string) (datamodels.TypeItem, bool) {
	for _, g := range r.source {
		for _, item := range g.Items {
			if item.ID == id {
				return item, true
			}
		}
	}
	return datamodels.TypeItem{}, false
}

func (r *typeGroupMemoryRepository) SelectAll() (results []datamodels.TypeGroup) {
	return r.source
}
