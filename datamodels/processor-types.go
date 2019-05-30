package datamodels

type TypeItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	IconOnRight bool   `json:"iconOnRight"`
	Color       string `json:"color"`
	HasInput    bool   `json:"hasInput"`
	HasOutput   bool   `json:"hasOutput"`
}

type TypeGroup struct {
	ID    string     `json:"id"`
	Title string     `json:"title"`
	Name  string     `json:"name"`
	Items []TypeItem `json:"items"`
}
