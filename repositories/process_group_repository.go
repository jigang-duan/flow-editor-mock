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
	UpdateProcessors(gid string, processors []datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error)

	InsertConnection(gid string, connection datamodels.Connection) (updatedProcessGroup datamodels.ProcessGroup, err error)

	CloneProcessorsAndConnections(gid string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error)

	InsertProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error)
	DeleteContent(parentID string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error)
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

func (r *processGroupRepository) DeleteContent(parentID string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error) {
	group, found := r.Select(parentID)
	if !found {
		return group, errors.New("不存在的组")
	}

	if len(processors) > 0 {
		group.Processors = deleteProcessors(group, processors)
	}
	if len(connections) > 0 {
		group.Connections = deleteConnections(group, connections)
	}
	if len(processGroups) > 0 {
		group.ProcessGroups = deleteProcessGroups(group, processGroups)
	}

	r.mu.Lock()
	r.source[parentID] = group
	r.mu.Unlock()

	return group, nil
}

func (r *processGroupRepository) InsertProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	group, found := r.Select(parentID)
	if !found {
		return group, errors.New("不存在的组")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return group, err
	}

	var (
		minX = 4950
		minY = 4950
	)
	for _, processor := range processors {
		if processor.Rect.X < minX {
			minX = processor.Rect.X
		}
		if processor.Rect.Y < minY {
			minY = processor.Rect.Y
		}
	}
	var processGroup = datamodels.ProcessGroup{
		ID:          id.String(),
		Processors:  processors,
		Connections: connections,
		Label:       "新建分组",
		Rect: datamodels.Rect{
			X: minX,
			Y: minY,
			W: 140,
			H: 70,
		},
		Icon: "icons/node-red/folder.png",
		Style: datamodels.Style{
			Color: "RGBA(242, 244, 245, 1.00)",
		},
	}
	group.ProcessGroups = append(group.ProcessGroups, processGroup)

	if len(processors) > 0 {
		var processorsIDs []string
		for _, processor := range processors {
			processorsIDs = append(processorsIDs, processor.ID)
		}
		group.Processors = deleteProcessors(group, processorsIDs)
	}
	if len(connections) > 0 {
		var connIDs []string
		for _, c := range connections {
			connIDs = append(connIDs, c.ID)
		}
		group.Connections = deleteConnections(group, connIDs)
	}

	r.mu.Lock()
	r.source[parentID] = group
	r.mu.Unlock()

	return group, nil
}

func (r *processGroupRepository) CloneProcessorsAndConnections(gid string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("不存在的组")
	}

	for _, p := range processors {
		id, err := uuid.NewV4()
		if err != nil {
			return group, err
		}
		p.OldID = p.ID
		p.ID = id.String()
		p.Rect.Y = p.Rect.Y + 10
		group.Processors = append(group.Processors, p)
	}

	for _, c := range connections {
		id, err := uuid.NewV4()
		if err != nil {
			return group, err
		}
		connection := datamodels.Connection{
			ID:         id.String(),
			SourcePort: c.SourcePort,
		}
		for _, p := range group.Processors {
			if p.OldID == c.SourceID {
				connection.SourceID = p.ID
			} else if p.OldID == c.TargetID {
				connection.TargetID = p.ID
			}
		}
		group.Connections = append(group.Connections, connection)
	}

	r.mu.Lock()
	r.source[gid] = group
	r.mu.Unlock()

	return group, nil
}

func (r *processGroupRepository) InsertConnection(gid string, connection datamodels.Connection) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("不存在的组")
	}

	id, err := uuid.NewV4()
	if err != nil {
		return group, err
	}
	connection.ID = id.String()
	group.Connections = append(group.Connections, connection)

	r.mu.Lock()
	r.source[gid] = group
	r.mu.Unlock()

	return group, nil
}

func (r *processGroupRepository) UpdateProcessors(gid string, processors []datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("不存在的组")
	}

	var ps []datamodels.Processor
	var has = false
	for _, gp := range group.Processors {
		has = false
		for _, p := range processors {
			if p.ID == gp.ID {
				ps = append(ps, p)
				has = true
				break
			}
		}
		if !has {
			ps = append(ps, gp)
		}
	}
	group.Processors = ps

	r.mu.Lock()
	r.source[gid] = group
	r.mu.Unlock()

	return group, nil

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

func deleteProcessors(group datamodels.ProcessGroup, ids []string) []datamodels.Processor {
	var results []datamodels.Processor
	var has = false
	for _, item := range group.Processors {
		has = false
		for _, id := range ids {
			if item.ID == id {
				has = true
				break
			}
		}
		if !has {
			results = append(results, item)
		}
	}

	return results
}

func deleteConnections(group datamodels.ProcessGroup, ids []string) []datamodels.Connection {
	var results []datamodels.Connection
	var has = false
	for _, item := range group.Connections {
		has = false
		for _, id := range ids {
			if item.ID == id {
				has = true
				break
			}
		}
		if !has {
			results = append(results, item)
		}
	}

	return results
}

func deleteProcessGroups(group datamodels.ProcessGroup, ids []string) []datamodels.ProcessGroup {
	var results []datamodels.ProcessGroup
	var has = false
	for _, item := range group.ProcessGroups {
		has = false
		for _, id := range ids {
			if item.ID == id {
				has = true
				break
			}
		}
		if !has {
			results = append(results, item)
		}
	}

	return results
}