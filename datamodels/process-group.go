package datamodels

type ProcessGroup struct {
	ID            string         `json:"id"`
	Processors    []Processor    `json:"processors"`
	Connections   []Connection   `json:"connections"`
	ProcessGroups []ProcessGroup `json:"processGroups"`

	Label   string `json:"label"`
	HasErr  bool   `json:"error"`
	Changed bool   `json:"changed"`
	Status  Status `json:"status"`
	Rect    Rect   `json:"rect"`
	Icon    string `json:"icon"`
	Style   Style  `json:"style"`
}

type DelContent struct {
	Processors    []string `json:"processors"`
	Connections   []string `json:"connections"`
	ProcessGroups []string `json:"processGroups"`
}
