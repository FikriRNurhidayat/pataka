package feature

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

// Build the Filter Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func (r *PostgresFeatureRepository) Filter(filter FeatureFilterArgs) (query string, args []interface{}, err error) {
	queries := []string{}

	if filter.Q != "" {
		filter.Q = fmt.Sprint("%", filter.Q, "%")
		queries = append(queries, "((features.name ILIKE :q) OR (features.label ILIKE :q) OR (features.name ILIKE :q))")
	}

	if filter.Enabled != nil {
		queries = append(queries, "(features.enabled = :enabled)")
	}

	if filter.Name != "" {
		queries = append(queries, "(features.name = :name)")
	}

	if filter.Label != "" {
		queries = append(queries, "(features.label = :label)")
	}

	query = strings.Join(queries, " AND ")
	query, args, err = sqlx.Named(query, filter)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

// Scan the query result, map it into Feature entity
func (r *PostgresFeatureRepository) Scan(rows *sql.Rows) (*Feature, error) {
	feature := &Feature{}
	ea := sql.NullTime{}

	if err := rows.Scan(
		&feature.Name,
		&feature.Label,
		&feature.Enabled,
		&feature.HasAudience,
		&feature.HasAudienceGroup,
		&feature.CreatedAt,
		&feature.UpdatedAt,
		&ea,
	); err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to scan: %s", err.Error())
		return nil, err
	}

	if ea.Valid {
		feature.EnabledAt = ea.Time
	}

	return feature, nil
}
