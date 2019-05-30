package datamodels

type ProcessGroup struct {
	ID         string      `json:"id"`
	Processors []Processor `json:"processors"`
}
