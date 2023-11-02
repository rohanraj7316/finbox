package storage

import "github.com/google/uuid"

type AdjustedTicker struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Symbol            string    `json:"symbol,omitempty" gorm:"index"`
	Timestamp         int64     `json:"timestamp,omitempty" gorm:"index"`
	ReadableTimestamp string    `json:"readableTimestamp,omitempty" gorm:"index"`
	OpenPrice         string    `json:"openPrice,omitempty"`
	HighPrice         string    `json:"highPrice,omitempty"`
	LowPrice          string    `json:"lowPrice,omitempty"`
	ClosePrice        string    `json:"closePrice,omitempty"`
	Volume            string    `json:"volume,omitempty"`

	// added one to many relation for security
	IntradayDataID uuid.UUID `json:"intradayDataID" gorm:"foreignKey:IntradayDataID;references:ID"`
	IntradayData   Ticker    `json:"intradayData" gorm:"foreignKey:IntradayDataID;references:ID"`

	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
