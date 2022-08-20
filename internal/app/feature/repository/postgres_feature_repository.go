package feature_repository

import (
	"context"
	"fmt"

	"github.com/fikrirnurhidayat/ffgo/internal/domain"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/grpclog"
)

type DB interface {
	sqlx.ExtContext
	sqlx.PreparerContext
}

type PostgresFeatureRepository struct {
	Logger grpclog.LoggerV2
	DB
}

var SORT_MAP = map[string]string{
	"name":               "features.name",
	"label":              "features.label",
	"enabled":            "features.enabled",
	"has_audience":       "features.has_audience",
	"has_audience_group": "features.has_audience_group",
	"created_at":         "features.created_at",
	"updated_at":         "features.updated_at",
	"enabled_at":         "features.enabled_at",
}

func (r *PostgresFeatureRepository) Get(ctx context.Context, name string) (feature *domain.Feature, err error) {
	stmt, err := r.PrepareContext(ctx, GET_SQL)
	if err != nil {
		r.Logger.Errorf("[PostgresFeatureRepository] failed to prepare get statement: %s", err.Error())
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, name)
	if err != nil {
		r.Logger.Errorf("[PostgresFeatureRepository] failed to run get statement: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		feature, err = r.Scan(rows)
		if err != nil {
			r.Logger.Errorf("Failed to scan query result: %s", err.Error())
			return nil, err
		}
	}

	return feature, nil
}

func (r *PostgresFeatureRepository) List(ctx context.Context, args *domain.FeatureListArgs) (features []domain.Feature, err error) {
	var (
		query string = LIST_SQL
		qargs []interface{}
	)

	if args.Filter != nil {
		filterQuery, filterArgs, err := r.Filter(*args.Filter)
		if err != nil {
			r.Logger.Errorf("[postgres-feature-repository] failed to build filter list query: %s", err.Error())
			return features, err
		}

		if filterQuery != "" {
			query = fmt.Sprint(query, "AND ", filterQuery)
			qargs = append(qargs, filterArgs...)
		}
	}

	if args.Sort != "" {
		sortQuery, err := r.Sort(args.Sort, SORT_MAP)
		if err != nil {
			r.Logger.Errorf("[postgres-feature-repository] failed to build sort list query: %s", err.Error())
			return features, err
		}

		query = fmt.Sprint(query, " ", sortQuery)
	}

	paginationQuery, paginationArgs := r.Paginate(args.Limit, args.Offset)
	query = fmt.Sprint(query, paginationQuery)
	qargs = append(qargs, paginationArgs...)

	query = r.Rebind(query)

	r.Logger.Info(query)

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to prepare list statement: %s", err.Error())
		return features, err
	}

	rows, err := stmt.QueryContext(ctx, qargs...)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to query list: %s", err.Error())
		return features, err
	}

	features = []domain.Feature{}

	for rows.Next() {
		feature, err := r.Scan(rows)
		if err != nil {
			r.Logger.Errorf("[postgres-feature-repository] failed to scan list: %s", err.Error())
			return nil, err
		}
		features = append(features, *feature)
	}

	return features, nil
}

func (r *PostgresFeatureRepository) Save(ctx context.Context, feature *domain.Feature) error {
	query := SAVE_SQL
	query, qargs, err := r.BindNamed(query, feature)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to bind query for save operation: %s", err.Error())
		return err
	}

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to prepare save query: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, qargs...)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to save: %s", err.Error())
		return err
	}

	return nil
}

func (r *PostgresFeatureRepository) Delete(ctx context.Context, name string) error {
	stmt, err := r.PrepareContext(ctx, DELETE_SQL)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to prepare delete statement: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, name)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to run delete statement: %s", err.Error())
		return err
	}

	return nil
}

func (r *PostgresFeatureRepository) Size(ctx context.Context, args *domain.FeatureFilterArgs) (uint32, error) {
	var (
		query string = SIZE_SQL
		qargs []interface{}
	)

	if args != nil {
		filterQuery, filterArgs, err := r.Filter(*args)
		if err != nil {
			r.Logger.Errorf("[postgres-feature-repository] failed to build filter count query: %s", err.Error())
			return 0, err
		}

		if filterQuery != "" {
			query = fmt.Sprint(query, "AND ", filterQuery)
			qargs = append(qargs, filterArgs...)
		}
	}

	query = r.Rebind(query)

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to prepare count statement: %s", err.Error())
		return 0, err
	}

	rows, err := stmt.QueryContext(ctx, qargs...)
	if err != nil {
		r.Logger.Errorf("[postgres-feature-repository] failed to count: %s", err.Error())
		return 0, err
	}

	count := 0

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			r.Logger.Errorf("[postgres-feature-repository] failed to scan count result: %s", err.Error())
			return 0, err
		}
	}

	return uint32(count), nil
}

func NewPostgresFeatureRepository(db DB, Logger grpclog.LoggerV2) domain.FeatureRepository {
	r := new(PostgresFeatureRepository)

	r.Logger = Logger
	r.DB = db

	return r
}
