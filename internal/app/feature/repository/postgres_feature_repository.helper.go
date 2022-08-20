package feature_repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"github.com/jmoiron/sqlx"
)

// Build the Filter Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func (r *PostgresFeatureRepository) Filter(filter domain.FeatureFilterArgs) (query string, args []interface{}, err error) {
	queries := []string{}

	if filter.Q != "" {
		filter.Q = fmt.Sprint("%", filter.Q, "%")
		queries = append(queries, "((features.name ILIKE :q) OR (features.label ILIKE :q) OR (features.name ILIKE :q))")
	}

	if filter.Enabled {
		queries = append(queries, "(features.enabled IS TRUE)")
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

// Build the Sort Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func (r *PostgresFeatureRepository) Sort(sortStr string, allowedCols map[string]string) (query string, err error) {
	sorts := strings.Split(strings.ReplaceAll(string(sortStr), " ", ""), ",")
	stmts := []string{}
	for _, sortExpr := range sorts {
		col := strings.TrimPrefix(sortExpr, "-")
		isDesc := strings.HasPrefix(sortExpr, "-")

		if stmt, ok := allowedCols[col]; ok {
			if isDesc {
				stmt = fmt.Sprint(stmt, " DESC")
			}

			stmts = append(stmts, stmt)
		}
	}

	if len(stmts) == 0 {
		return "", nil
	}

	query = fmt.Sprintf("ORDER BY %s", strings.Join(stmts, ", "))

	return query, nil
}

// Build the Pagination Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func (r *PostgresFeatureRepository) Paginate(limit uint32, offset uint32) (query string, args []interface{}) {
	query = "LIMIT ? OFFSET ?"
	args = []interface{}{limit, offset}
	return query, args
}

// Scan the query result, map it into Feature entity
func (r *PostgresFeatureRepository) Scan(rows *sql.Rows) (*domain.Feature, error) {
	feature := &domain.Feature{}

	if err := rows.Scan(
		&feature.Name,
		&feature.Label,
		&feature.Enabled,
		&feature.HasAudience,
		&feature.HasAudienceGroup,
		&feature.CreatedAt,
		&feature.UpdatedAt,
		&feature.EnabledAt,
	); err != nil {
		r.Logger.Errorf("[PostgresFeatureRepository] failed to scan: %s", err.Error())
		return nil, err
	}

	return feature, nil
}
