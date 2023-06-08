package models

import (
	"database/sql/driver"
)

type FlowIdType string

const (
	AUTH FlowIdType = "AUTH"
)

func (fit *FlowIdType) Scan(value interface{}) error {
	*fit = FlowIdType([]byte(value.(string)))
	return nil
}

func (fit FlowIdType) Value() (driver.Value, error) {
	return string(fit), nil
}

type FlowId struct {
	FlowId        string     `gorm:"primary_key"`
	FlowIdType    FlowIdType `sql:"flow_id_type"`
	WalletAddress string
}

