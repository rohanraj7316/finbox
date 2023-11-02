package storage

import "github.com/google/uuid"

type NewsSentiment struct {
	ID                uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Symbol            string    `json:"symbol,omitempty" gorm:"index"`
	Timestamp         int64     `json:"timestamp,omitempty" gorm:"index"`
	ReadableTimestamp string    `json:"readableTimestamp,omitempty" gorm:"index"`
	SentimentScore    float64   `json:"sentimentScore,omitempty"`

	// added one to many relation for security
	AssociatedIntradayDataID uuid.UUID `json:"associatedIntradayDataID" gorm:"foreignKey:AssociatedIntradayDataID;references:ID"`
	AssociatedIntradayData   Ticker    `json:"associatedIntradayData" gorm:"foreignKey:AssociatedIntradayDataID;references:ID"`

	CreatedAt int64 `json:"createdAt" gorm:"autoCreateTime:milli"`
	UpdatedAt int64 `json:"updatedAt" gorm:"autoUpdateTime:milli"`
}
