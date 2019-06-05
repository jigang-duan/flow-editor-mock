package nifi

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/repositories"
	"gopkg.in/resty.v1"
)

func NewProcessGroupRepository(client *resty.Client) repositories.ProcessGroupRepository {
	return &processGroupRepository{
		client: client,
	}
}

type processGroupRepository struct {
	client *resty.Client
}

//const bashURL = "http://192.168.20.104:8080/nifi-api/process-groups/016b10c5-bef9-188e-46b1-e08ef9485cda/processors"

func (*processGroupRepository) Select(id string) (processGroup datamodels.ProcessGroup, found bool) {
	panic("implement me")
}

func (r *processGroupRepository) InsertProcessor(gid string, processor datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {

}

func (*processGroupRepository) UpdateProcessors(gid string, processors []datamodels.Processor) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	panic("implement me")
}

func (*processGroupRepository) InsertConnection(gid string, connection datamodels.Connection) (updatedProcessGroup datamodels.ProcessGroup, err error) {
	panic("implement me")
}

func (*processGroupRepository) CloneSnippet(gid string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	panic("implement me")
}

func (*processGroupRepository) InsertProcessGroup(parentID string, processors []datamodels.Processor, connections []datamodels.Connection) (datamodels.ProcessGroup, error) {
	panic("implement me")
}

func (*processGroupRepository) UngroupProcessGroup(gid string, parentID string) (datamodels.ProcessGroup, error) {
	panic("implement me")
}

func (*processGroupRepository) DeleteSnippet(parentID string, processors []string, connections []string, processGroups []string) (datamodels.ProcessGroup, error) {
	panic("implement me")
}

func (*processGroupRepository) UpdateSnippet(gid string, processors []datamodels.Processor, connections []datamodels.Connection, processGroups []datamodels.ProcessGroup) (datamodels.ProcessGroup, error) {
	panic("implement me")
}

