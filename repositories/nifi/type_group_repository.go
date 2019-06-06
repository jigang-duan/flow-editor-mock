package nifi

import (
	"flow-editor-mock/datamodels"
	"flow-editor-mock/entitys"
	"flow-editor-mock/repositories"
	"gopkg.in/resty.v1"
	"strings"
)

func NewTypeGroupRepository(client *resty.Client) repositories.TypeGroupRepository {
	return &typeGroupMemoryRepository{
		client: client,
	}
}

type typeGroupMemoryRepository struct {
	client *resty.Client
}

const (
	apiURL = "http://192.168.20.104:8080/nifi-api"
	bashURL = "http://192.168.20.104:8080/nifi-api/flow"
)

var (
	colors     = []string{"#79f299", "#f2797c", "#799ff2", "#f2b779", "#9479f2", "#c3f279", "#f279e6"}
	processors = []string{"ExecuteScript", "QueryDatabaseTableRecord", "PutDatabaseRecord", "UpdateRecord", "SplitJson", "PublishKafka_2_0", "GenerateFlowFile", "PutDistributedMapCache", "InvokeHTTP"}
	TypeCache *entitys.ProcessorTypes
	ClientID string
)

func (r *typeGroupMemoryRepository) SelectAll() (results []datamodels.TypeGroup) {
	resp, err := r.client.R().SetResult(&entitys.ProcessorTypes{}).Get(bashURL + "/processor-types")
	if err != nil {
		return []datamodels.TypeGroup{}
	}
	var typesResp *entitys.ProcessorTypes
	var ok = false
	if TypeCache != nil {
		typesResp, ok = TypeCache, true
	} else {
		typesResp, ok = resp.Result().(*entitys.ProcessorTypes)
	}
	if ok {
		TypeCache = typesResp
		typeGroup := datamodels.TypeGroup{
			ID:    "fdCD6BD6-C5Bb-ECdb-6647-b8FDfe8eBa64",
			Title: "处理器",
			Name:  "processor",
		}
		for _, it := range typesResp.ProcessorTypes {
			ids := strings.Split(it.Id, ".")
			name := ids[len(ids)-1]
			if includes(processors, name) {
				item := datamodels.TypeItem{
					ID:    it.Id,
					Name:  name,
					Icon:  "icons/node-red/db.png",
					Color: colorByName(name),
				}
				typeGroup.Items = append(typeGroup.Items, item)
			}
		}
		results = append(results, typeGroup)
		return results
	}
	return []datamodels.TypeGroup{}
}

func (r *typeGroupMemoryRepository) SelectTypeByID(id string) (datamodels.TypeItem, bool) {
	for _, it := range TypeCache.ProcessorTypes {
		ids := strings.Split(it.Id, ".")
		name := ids[len(ids)-1]
		if includes(processors, name) {
			item := datamodels.TypeItem{
				ID:    it.Id,
				Name:  name,
				Icon:  "icons/node-red/db.png",
				Color: colorByName(name),
			}
			return item, true
		}
	}
	return datamodels.TypeItem{}, false
}

func (r *typeGroupMemoryRepository) SelectProcessorByID(id string) (entitys.ProcessorTypeItem, bool) {
	if TypeCache == nil {
		return entitys.ProcessorTypeItem{}, false
	}

	for _, item := range TypeCache.ProcessorTypes {
		if item.Id == id {
			return item, true
		}
	}
	return entitys.ProcessorTypeItem{}, false
}

func colorByName(name string) string {
	first := name[len(name)-1]
	return colors[int(first)%len(colors)]
}

func includes(aars []string, name string) bool {
	for _, it := range aars {
		if it == name {
			return true
		}
	}
	return false
}
