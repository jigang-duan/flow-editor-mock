package entitys

type Bundle struct {
	Artifact string `json:"artifact"`
	Group    string `json:"group"`
	Version  string `json:"version"`
}

type ProcessorTypeItem struct {
	Id          string   `json:"type"`
	Tags        []string `json:"tags"`
	Restricted  bool     `json:"restricted"`
	Description string   `json:"description"`
	Bundle      Bundle   `json:"bundle"`
}

type ProcessorTypes struct {
	ProcessorTypes []ProcessorTypeItem `json:"processorTypes"`
}
