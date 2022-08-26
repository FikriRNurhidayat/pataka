package domain

import (
	"time"
)

type Feature struct {
	Name             string    `json:"name" db:"name"`
	Label            string    `json:"label" db:"label"`
	Enabled          bool      `json:"enabled" db:"enabled"`
	HasAudience      bool      `json:"has_audience" db:"has_audience"`
	HasAudienceGroup bool      `json:"has_audience_group" db:"has_audience_group"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	EnabledAt        time.Time `json:"enabled_at" db:"enabled_at"`
}
