package nifi

import (
	"errors"
	"flow-editor-mock/datamodels"
	"flow-editor-mock/entitys"
	"flow-editor-mock/repositories"
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
)

func NewProcessGroupRepository(client *resty.Client) repositories.ProcessGroupRepository {
	return &processGroupRepository{
		client: client,
	}
}

type processGroupRepository struct {
	client *resty.Client
}


func (r *processGroupRepository) Select(id string) (processGroup datamodels.ProcessGroup, found bool) {
	resp, err := r.client.R().SetResult(&entitys.ProcessGroupEntity{}).Get(bashURL + "/process-groups/" + id)
	if err != nil {
		return datamodels.ProcessGroup{}, false
	}
	if entity, ok := resp.Result().(*entitys.ProcessGroupEntity); ok {
		processGroup.ID = entity.ProcessGroupFlow.ID
		processGroup.Link.Flow = entity.ProcessGroupFlow
		if entity.ProcessGroupFlow.Flow.Processors != nil {
			processGroup.Processors = []datamodels.Processor{}
			for _, item := range entity.ProcessGroupFlow.Flow.Processors {
				typeID := item.Component.Type
				name := item.Component.Name
				x := int(item.Position.X)
				if x < 75 {
					x = 75
				}
				y := int(item.Position.Y)
				if y < 75 {
					y = 75
				}
				processor := datamodels.Processor{
					ID: item.ID,
					TypeID: typeID,
					Label: name,
					HasInput: item.InputRequirement != "INPUT_FORBIDDEN",
					HasOutput: true,
					IconOnRight: item.InputRequirement != "INPUT_FORBIDDEN",
					Rect:datamodels.Rect{
						X: x,
						Y: y,
						W: 140,
						H: 30,
					},
					Icon:  "icons/node-red/db.png",
					Style:datamodels.Style{
						Color: colorByName(name),
					},
				}
				processGroup.Processors = append(processGroup.Processors, processor)
				if ClientID != item.Revision.ClientID {
					ClientID = item.Revision.ClientID
				}
			}
		}
		if entity.ProcessGroupFlow.Flow.ProcessGroups != nil {
			processGroup.ProcessGroups = []datamodels.ProcessGroup{}
			for _, item := range entity.ProcessGroupFlow.Flow.ProcessGroups {
				x := int(item.Position.X)
				if x < 75 {
					x = 75
				}
				y := int(item.Position.Y)
				if y < 75 {
					y = 75
				}
				group := datamodels.ProcessGroup{
					ID: item.ID,
					Label: item.Component.Name,
					Rect:datamodels.Rect{
						X: x,
						Y: y,
						W: 140,
						H: 70,
					},
					Icon: "icons/node-red/folder.png",
					Style: datamodels.Style{
						Color: "RGBA(242, 244, 245, 1.00)",
					},
					Count: item.RunningCount + item.DisabledCount + item.InvalidCount,
				}
				processGroup.ProcessGroups = append(processGroup.ProcessGroups, group)
				if ClientID != item.Revision.ClientID {
					ClientID = item.Revision.ClientID
				}
			}
		}
		if entity.ProcessGroupFlow.Flow.Connections != nil {
			processGroup.Connections = []datamodels.Connection{}
			for _, item := range entity.ProcessGroupFlow.Flow.Connections {
				conn := datamodels.Connection{
					ID: item.ID,
					SourceID: item.SourceID,
					SourcePort: 0,
					TargetID: item.DestinationID,
				}
				processGroup.Connections = append(processGroup.Connections, conn)
			}
		}
		return processGroup, true
	}

	return datamodels.ProcessGroup{}, false
}

func (r *processGroupRepository) InsertProcessor(gid string, processor datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	for _, item := range TypeCache.ProcessorTypes {
		if item.Id == processor.TypeID {
			ids := strings.Split(processor.TypeID, ".")
			name := ids[len(ids)-1]
			processor := entitys.InsertProcessorEntity{
				Component: entitys.InsertComponent{
					Bundle: item.Bundle,
					Name: name,
					Position: entitys.Position{
						X: float64(processor.Rect.X),
						Y: float64(processor.Rect.Y),
					},
					TypeID: processor.TypeID,
				},
				Revision: entitys.Revision{
					ClientID: ClientID,
					Version: 0,
				},
			}
			_, err = r.client.R().SetBody(processor).Post(apiURL + "/process-groups/" + gid + "/processors")
			break
		}
	}

	group, found := r.Select(gid)
	if !found {
		return group, errors.New("没有找到组")
	}
	return group, err
}

func (*processGroupRepository) UpdateProcessors(gid string, processors []datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	panic("implement UpdateProcessors")
}

func (*processGroupRepository) InsertConnection(gid string, connection datamodels.Connection) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	panic("implement InsertConnection")
}

func (*processGroupRepository) CloneSnippet(gid string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	panic("implement CloneSnippet")
}

func (r *processGroupRepository) InsertProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	group, found := r.Select(parentID)
	if !found {
		return group, errors.New("没有找到组")
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

	insertGroup := entitys.InsertProcessGroupEntity{
		Revision:entitys.Revision{
			ClientID: ClientID,
			Version: 0,
		},
	}
	insertGroup.Component.Name = "新建分组"
	insertGroup.Component.Position = entitys.Position{
		X: float64(minX),
		Y: float64(minY),
	}

	if resp, err := r.client.R().SetResult(&entitys.ProcessGroup{}).SetBody(insertGroup).Post(apiURL + "/process-groups/" + parentID + "/process-groups"); err == nil {
		if newGroup, ok := resp.Result().(*entitys.ProcessGroup); ok {
			var processorsMap = map[string]entitys.Revision{}
			for _, item := range group.Link.Flow.Flow.Processors {
				if includesOnProcessors(processors, item.ID) {
					processorsMap[item.ID] = item.Revision
				}
			}
			snippet := entitys.CreateSnippetEntity{
				Snippet:entitys.CreateSnippet{
					ParentGroupId: parentID,
					Processors: processorsMap,
				},
			}
			if resp, err := r.client.R().SetBody(snippet).SetResult(&entitys.CreatedSnippetEntity{}).Post(apiURL + "/snippets"); err == nil {
				if entity, ok := resp.Result().(*entitys.CreatedSnippetEntity); ok {
					update := entitys.UpdateSnippetEntity{}
					update.Snippet.ID = entity.Snippet.ID
					update.Snippet.ParentGroupId = newGroup.ID
					if _, err := r.client.R().SetBody(update).Put(apiURL + "/snippets/" + entity.Snippet.ID + "?disconnectedNodeAcknowledged=false"); err == nil {
						if group, found := r.Select(parentID); found {
							return group, nil
						}
					}
				}
			}
		}
	}
	group, found = r.Select(parentID)
	if !found {
		return group, errors.New("没有找到组")
	}
	return group, nil
}

func (r *processGroupRepository) UngroupProcessGroup(gid string, parentID string) (datamodels.ProcessGroup, error) {
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("没有找到组")
	}

	var processorsMap = map[string]entitys.Revision{}
	for _, item := range group.Link.Flow.Flow.Processors {
		processorsMap[item.ID] = item.Revision
	}
	snippet := entitys.CreateSnippetEntity{
		Snippet:entitys.CreateSnippet{
			ParentGroupId: gid,
			Processors: processorsMap,
		},
	}
	if resp, err := r.client.R().SetBody(snippet).SetResult(&entitys.CreatedSnippetEntity{}).Post(apiURL + "/snippets"); err == nil {
		if entity, ok := resp.Result().(*entitys.CreatedSnippetEntity); ok {
			update := entitys.UpdateSnippetEntity{}
			update.Snippet.ID = entity.Snippet.ID
			update.Snippet.ParentGroupId = parentID
			if _, err := r.client.R().SetBody(update).Put(apiURL + "/snippets/" + entity.Snippet.ID + "?disconnectedNodeAcknowledged=false"); err == nil {
				group, found = r.Select(parentID)
				if !found {
					return group, errors.New("没有找到组")
				}
				for _, item := range group.Link.Flow.Flow.ProcessGroups {
					if item.ID == gid {
						_, _ = r.client.R().Delete(apiURL + fmt.Sprintf("/process-groups/%s?version=%d&clientId=%s", gid, item.Revision.Version, item.Revision.ClientID))
						break
					}
				}
			}
		}
	}

	group, found = r.Select(parentID)
	if !found {
		return group, errors.New("没有找到组")
	}
	return group, nil
}

func (r *processGroupRepository) DeleteSnippet(gid string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error) {
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("没有找到组")
	}
	var processorsMap = map[string]entitys.Revision{}
	for _, item := range group.Link.Flow.Flow.Processors {
		if includes(processors, item.ID) {
			processorsMap[item.ID] = item.Revision
		}
	}
	var processGroupsMap = map[string]entitys.Revision{}
	for _, item := range group.Link.Flow.Flow.ProcessGroups {
		if includes(processGroups, item.ID) {
			processGroupsMap[item.ID] = item.Revision
		}
	}
	snippet := entitys.CreateSnippetEntity{
		Snippet:entitys.CreateSnippet{
			ParentGroupId: gid,
			Processors: processorsMap,
			ProcessGroups: processGroupsMap,
		},
	}
	if resp, err := r.client.R().SetBody(snippet).SetResult(&entitys.CreatedSnippetEntity{}).Post(apiURL + "/snippets"); err == nil {
		if entity, ok := resp.Result().(*entitys.CreatedSnippetEntity); ok {
			if _, err := r.client.R().Delete(apiURL + "/snippets/" + entity.Snippet.ID + "?disconnectedNodeAcknowledged=false"); err == nil {
				if group, found := r.Select(gid); found {
					return group, nil
				}
			}
		}
	}

	return group, nil
}

func (r *processGroupRepository) UpdateSnippet(gid string, processors []datamodels.Processor, connections []datamodels.Connection, processGroups []datamodels.ProcessGroup) (datamodels.ProcessGroup, error) {
	if processors != nil {
		for _, item := range processors {
			resp, err := r.client.R().SetResult(&entitys.UpdateProcessorEntity{}).Get(apiURL + "/processors/" + item.ID)
			if err != nil {
				continue
			}
			if entity, ok := resp.Result().(*entitys.UpdateProcessorEntity); ok {
				entity.Component.Position.X = float64(item.Rect.X)
				entity.Component.Position.Y = float64(item.Rect.Y)
				entity.Revision.Version += 1
				_, _ = r.client.R().SetBody(entity).Put(apiURL + "/processors/" + item.ID)
			}
		}
	}
	if processGroups != nil {
		for _, item := range processGroups {
			resp, err := r.client.R().SetResult(&entitys.UpdateProcessorGroupEntity{}).Get(apiURL + "/process-groups/" + item.ID)
			if err != nil {
				continue
			}
			if entity, ok := resp.Result().(*entitys.UpdateProcessorGroupEntity); ok {
				entity.Component.Position.X = float64(item.Rect.X)
				entity.Component.Position.Y = float64(item.Rect.Y)
				entity.Revision.Version += 1
				_, _ = r.client.R().SetBody(entity).Put(apiURL + "/process-groups/" + item.ID)
			}
		}
	}
	group, found := r.Select(gid)
	if !found {
		return group, errors.New("没有找到组")
	}
	return group, nil
}


func includesOnProcessors(processors []datamodels.Processor, id string) bool {
	for _, item := range processors {
		if item.ID == id {
			return true
		}
	}
	return false
}