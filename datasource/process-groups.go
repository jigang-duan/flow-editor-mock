package datasource

import "flow-editor-mock/datamodels"

var ProcessGroups = map[string]datamodels.ProcessGroup{
	"root": {
		ID:          "root",
		Processors:  []datamodels.Processor{},
		Connections: []datamodels.Connection{},
	},
}
