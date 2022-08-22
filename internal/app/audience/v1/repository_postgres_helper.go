package audience

import (
	"database/sql"
	"strings"

	"github.com/fikrirnurhidayat/ffgo/internal/pkg/inspector"
	"github.com/jmoiron/sqlx"
)

// Build the Filter Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func (r *PostgresAudienceRepository) Filter(filter AudienceFilterArgs) (query string, args []interface{}, err error) {
	queries := []string{}

	if filter.Enabled != nil {
		queries = append(queries, "(feature_audiences.enabled = :enabled)")
	}

	if inspector.IsEmpty(filter.FeatureName) {
		queries = append(queries, "(feature_audiences.feature_name = :feature_name)")
	}

	if !inspector.IsEmptySlice(filter.AudienceIds) {
		queries = append(queries, "(feature_audiences.audience_id IN (:audience_id))")
	}

	query = strings.Join(queries, " AND ")
	query, args, err = sqlx.Named(query, filter)
	if err != nil {
		return "", nil, err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

// Scan the query result, map it into Audience entity
func (r *PostgresAudienceRepository) Scan(rows *sql.Rows) (*Audience, error) {
	audience := &Audience{}
	ea := sql.NullTime{}

	if err := rows.Scan(
		&audience.FeatureName,
		&audience.AudienceId,
		&audience.Enabled,
		&audience.CreatedAt,
		&audience.UpdatedAt,
		&ea,
	); err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to scan: %s", err.Error())
		return nil, err
	}

	if ea.Valid {
		audience.EnabledAt = ea.Time
	}

	return audience, nil
}
