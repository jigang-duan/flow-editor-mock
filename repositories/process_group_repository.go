package repositories

import (
	"errors"
	"flow-editor-mock/datamodels"
	"github.com/iris-contrib/go.uuid"
	"sync"
)

type ProcessGroupRepository interface {
	Select(id string) (processGroup datamodels.ProcessGroup, found bool)

	InsertProcessor(gid string, processor datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error)
}

func NewProcessGroupRepository(source map[string]datamodels.ProcessGroup) ProcessGroupRepository {
	return &processGroupRepository{
		source: source,
	}
}

type processGroupRepository struct {
	source map[string]datamodels.ProcessGroup
	mu     sync.RWMutex
}

func (r *processGroupRepository) InsertProcessor(gid string, processor datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {


	group, found := r.Select(gid)
	if !found {
		return group, errors.New("不存在的组")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return group, err
	}
	processor.ID = id.String()
	group.Processors = append(group.Processors, processor)

	r.mu.Lock()
	r.source[gid] = group
	r.mu.Unlock()

	return group, nil
}

func (r *processGroupRepository) Select(id string) (processGroup datamodels.ProcessGroup, found bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for key, g := range r.source {
		if key == id {
			processGroup = g
			return processGroup, true
		}
	}

	return
}


