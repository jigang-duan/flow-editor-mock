package routes

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/services"
	"github.com/kataras/iris"
)

func ProcessGroups(service services.ProcessGroupService, id string) (processGroup datamodels.ProcessGroup, found bool) {
	return service.GetByID(id)
}

func CreateProcessor(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processor datamodels.Processor
	if err := ctx.ReadJSON(&processor); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}

	return service.CreateProcessorByID(gid, processor)
}

func UpdateProcessors(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processors []datamodels.Processor
	if err := ctx.ReadJSON(&processors); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.UpdateProcessorsByID(gid, processors)
}

func CreateConnection(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var connection datamodels.Connection
	if err := ctx.ReadJSON(&connection); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}

	return service.CreateConnectionByID(gid, connection)
}

func DeleteConnections(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var connections []string
	if err := ctx.ReadJSON(&connections); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}

	return service.DeleteConnectionsByIDs(gid, connections)
}

func DeleteProcessors(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processors []string
	if err := ctx.ReadJSON(&processors); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}

	return service.DeleteProcessorsByIDs(gid, processors)
}

func CloneProcessorsAndConnections(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processGroup datamodels.ProcessGroup
	if err := ctx.ReadJSON(&processGroup); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.CloneProcessorsAndConnections(gid, processGroup.Processors, processGroup.Connections)
}
