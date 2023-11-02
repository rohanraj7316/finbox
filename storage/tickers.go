package storage

import "github.com/google/uuid"

type Ticker struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Symbol            string    `json:"symbol,omitempty" gorm:"index"`
	Timestamp         int64     `json:"timestamp,omitempty" gorm:"index"`
	ReadableTimestamp string    `json:"readableTimestamp,omitempty" gorm:"index"`
	OpenPrice         string    `json:"openPrice,omitempty"`
	HighPrice         string    `json:"highPrice,omitempty"`
	LowPrice          string    `json:"lowPrice,omitempty"`
	ClosePrice        string    `json:"closePrice,omitempty"`
	Volume            string    `json:"volume,omitempty"`

	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
