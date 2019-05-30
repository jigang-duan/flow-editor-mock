package routes

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/services"
)

func TypeGroups(service services.TypeGroupService) []datamodels.TypeGroup {
	return service.GetAll()
}