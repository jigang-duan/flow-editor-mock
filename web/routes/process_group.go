package routes

import (
	"errors"
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

func CloneSnippet(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processGroup datamodels.ProcessGroup
	if err := ctx.ReadJSON(&processGroup); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.CloneSnippet(gid, processGroup.Processors, processGroup.Connections)
}

func CreateProcessGroup(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processGroup datamodels.ProcessGroup
	if err := ctx.ReadJSON(&processGroup); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.CreateProcessGroup(gid, processGroup.Processors, processGroup.Connections)
}

func DeleteSnippet(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var delContent datamodels.DelContent
	if err := ctx.ReadJSON(&delContent); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.DeleteSnippet(gid, delContent.Processors, delContent.Connections, delContent.ProcessGroups)
}

func UpdateSnippet(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	var processGroup datamodels.ProcessGroup
	if err := ctx.ReadJSON(&processGroup); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, err
	}
	return service.UpdateSnippet(gid, processGroup.Processors, processGroup.Connections, processGroup.ProcessGroups)
}

func UngroupProcessGroup(ctx iris.Context, service services.ProcessGroupService, gid string) (datamodels.ProcessGroup, error) {
	parentID := ctx.URLParam("parentID")
	if parentID == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		return datamodels.ProcessGroup{}, errors.New("需要parentID参数")
	}
	return service.UngroupProcessGroup(gid, parentID)
}