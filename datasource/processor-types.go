package datasource

import "flow-editor-mock/datamodels"

var TypeGroups = []datamodels.TypeGroup{
	{
		ID:    "fdCD6BD6-C5Bb-ECdb-6647-b8FDfe8eBa64",
		Title: "连接器",
		Name:  "Connector",
		Items: []datamodels.TypeItem{
			{
				ID:          "76DA30b6-61BF-75AF-AfFc-d2FdB9AdC874",
				Name:        "MySQL",
				Icon:        "icons/node-red/db.png",
				IconOnRight: false,
				Color:       "#79f299",
				HasInput:    false,
				HasOutput:   true,
			},
			{
				ID:          "EAb1dda2-2B17-79CF-820B-a7D28587De7d",
				Name:        "CSV",
				Icon:        "icons/node-red/parser-csv.png",
				IconOnRight: false,
				Color:       "#f2797c",
				HasInput:    false,
				HasOutput:   true,
			},
			{
				ID:          "6f7F4b9b-8ef4-Ee72-9F8C-faC65DeBd78b",
				Name:        "Redis",
				Icon:        "icons/node-red/redis.png",
				IconOnRight: false,
				Color:       "#799ff2",
				HasInput:    false,
				HasOutput:   true,
			},
		},
	},
	{
		ID:    "18Ed9dB5-F3B1-6c83-4de6-d7BcBbA2c84c",
		Title: "转换器",
		Name:  "Converter",
		Items: []datamodels.TypeItem{
			{
				ID:          "871155A1-e8EE-D3Fc-CFcC-E2c83BB98B31",
				Name:        "sort",
				Icon:        "icons/node-red/sort.png",
				IconOnRight: false,
				Color:       "#f2b779",
				HasInput:    true,
				HasOutput:   true,
			},
			{
				ID:          "fA1Be6A1-1335-9a0e-5a7D-7b5cdBbbcfE3",
				Name:        "split",
				Icon:        "icons/node-red/split.png",
				IconOnRight: false,
				Color:       "#9479f2",
				HasInput:    true,
				HasOutput:   true,
			},
		},
	},
	{
		ID:    "F163FeC2-8189-DC1c-bFB5-c7b5eabCF604",
		Title: "格式转换器",
		Name:  "Converter",
		Items: []datamodels.TypeItem{
			{
				ID:          "ccBE4edc-06ab-6Ab1-A54a-22Ff1A9a5aD0",
				Name:        "Yaml",
				Icon:        "icons/node-red/parser-yaml.png",
				IconOnRight: false,
				Color:       "#c3f279",
				HasInput:    true,
				HasOutput:   true,
			},
			{
				ID:          "8b19F27B-DbAF-c59d-22F4-f9f141b73Ee6",
				Name:        "Json",
				Icon:        "icons/node-red/parser-json.png",
				IconOnRight: false,
				Color:       "#f279e6",
				HasInput:    true,
				HasOutput:   true,
			},
		},
	},
}
