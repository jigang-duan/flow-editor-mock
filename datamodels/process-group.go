package datamodels

type ProcessGroup struct {
	ID         string      `json:"id"`
	Processors []Processor `json:"processors"`
	Connections []Connection `json:"connections"`
}
