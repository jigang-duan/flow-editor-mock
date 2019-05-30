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
