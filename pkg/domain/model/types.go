package model

import (
	"time"
)

// Audit represents the audit log entity.
type Audit struct {
	Date      time.Time `json:"date"`
	Operation string    `json:"operation"`
}
