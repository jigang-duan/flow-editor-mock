package datamodels

type Style struct {
	Color string `json:"color"`
}

type Status struct {
	Show  bool `json:"show"`
	Label bool `json:"label"`
}

type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
	H int `json:"h"`
	W int `json:"w"`
}

type Processor struct {
	ID          string `json:"id"`
	TypeID      string `json:"typeId"`
	Label       string `json:"label"`
	HasInput    bool   `json:"hasInput"`
	HasOutput   bool   `json:"hasOutput"`
	IconOnRight bool   `json:"iconOnRight"`
	HasErr      bool   `json:"error"`
	Changed     bool   `json:"changed"`
	Status      Status `json:"status"`
	Rect        Rect   `json:"rect"`
	Icon        string `json:"icon"`
	Style       Style  `json:"style"`
}
