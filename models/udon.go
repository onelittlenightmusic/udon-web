package models

type Udon struct {
	ID          uint    `json:"id"`
	Type        UdonType  `json:"type"`
}

type UdonType string

const (
	Curry UdonType = "curry"
	Nabeyaki UdonType = "nabeyaki"
)