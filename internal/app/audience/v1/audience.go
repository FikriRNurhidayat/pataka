package audience

import (
	"time"
)

type Audience struct {
	FeatureName string    `json:"feature_name" db:"feature_name"`
	AudienceId  string    `json:"audience_id" db:"audience_id"`
	Enabled     bool      `json:"enabled" db:"enabled"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	EnabledAt   time.Time `json:"enabled_at" db:"enabled_at"`
}
