package audience

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/inspector"
	"github.com/fikrirnurhidayat/ffgo/internal/pkg/queryhelper"
	"google.golang.org/grpc/grpclog"
)

type PostgresAudienceRepository struct {
	logger grpclog.LoggerV2
	driver.DB
}

// DeleteBy implements domain.AudienceRepository
func (*PostgresAudienceRepository) DeleteBy(ctx context.Context, args *domain.AudienceFilterArgs) error {
	panic("unimplemented")
}

// GetBy implements domain.AudienceRepository
func (*PostgresAudienceRepository) GetBy(ctx context.Context, args *domain.AudienceFilterArgs) (*domain.Audience, error) {
	panic("unimplemented")
}

var SORT_MAP = map[string]string{
	"id":           "audiences.id",
	"feature_name": "audiences.feature_name",
	"user_id":      "audiences.user_id",
	"created_at":   "audiences.created_at",
	"updated_at":   "audiences.updated_at",
}

func (r *PostgresAudienceRepository) Get(ctx context.Context, fn string, ui string) (audience *domain.Audience, err error) {
	stmt, err := r.PrepareContext(ctx, GET_SQL)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to prepare get statement: %s", err.Error())
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, fn, ui)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to run get statement: %s", err.Error())
		return nil, err
	}

	for rows.Next() {
		audience, err = r.Scan(rows)
		if err != nil {
			r.logger.Errorf("Failed to scan query result: %s", err.Error())
			return nil, err
		}
	}

	return audience, nil
}

func (r *PostgresAudienceRepository) List(ctx context.Context, args *domain.AudienceListArgs) (audiences []domain.Audience, err error) {
	var (
		query string = LIST_SQL
		qargs []interface{}
	)

	// Filter if filter argument is specified
	if args.Filter != nil {
		filterQuery, filterArgs, err := r.Filter(*args.Filter)
		if err != nil {
			r.logger.Errorf("[postgres-audience-repository] failed to build filter list query: %s", err.Error())
			return audiences, err
		}

		if !inspector.IsEmpty(filterQuery) {
			query = fmt.Sprint(query, "WHERE ", filterQuery)
			qargs = append(qargs, filterArgs...)
		}
	}

	// Sort if sort argument is specified
	if !inspector.IsEmpty(args.Sort) {
		sortQuery, err := queryhelper.Sort(args.Sort, SORT_MAP)
		if err != nil {
			r.logger.Errorf("[postgres-audience-repository] failed to build sort list query: %s", err.Error())
			return audiences, err
		}

		query = fmt.Sprint(query, " ", sortQuery)
	}

	// Specify pagination statement
	paginationQuery, paginationArgs := queryhelper.Paginate(args.Limit, args.Offset)
	query = fmt.Sprint(query, paginationQuery)
	qargs = append(qargs, paginationArgs...)

	query = r.Rebind(query)

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to prepare list statement: %s", err.Error())
		return audiences, err
	}

	rows, err := stmt.QueryContext(ctx, qargs...)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to query list: %s", err.Error())
		return audiences, err
	}

	audiences = []domain.Audience{}

	for rows.Next() {
		audience, err := r.Scan(rows)
		if err != nil {
			r.logger.Errorf("[postgres-audience-repository] failed to scan list: %s", err.Error())
			return nil, err
		}
		audiences = append(audiences, *audience)
	}

	return audiences, nil
}

func (r *PostgresAudienceRepository) Size(ctx context.Context, args *domain.AudienceFilterArgs) (uint32, error) {
	var (
		query string = SIZE_SQL
		qargs []interface{}
	)

	if args != nil {
		filterQuery, filterArgs, err := r.Filter(*args)
		if err != nil {
			r.logger.Errorf("[postgres-audience-repository] failed to build filter count query: %s", err.Error())
			return 0, err
		}

		if !inspector.IsEmpty(filterQuery) {
			query = fmt.Sprint(query, "WHERE ", filterQuery)
			qargs = append(qargs, filterArgs...)
		}
	}

	query = r.Rebind(query)

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to prepare count statement: %s", err.Error())
		return 0, err
	}

	rows, err := stmt.QueryContext(ctx, qargs...)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to count: %s", err.Error())
		return 0, err
	}

	count := 0

	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			r.logger.Errorf("[postgres-audience-repository] failed to scan count result: %s", err.Error())
			return 0, err
		}
	}

	return uint32(count), nil
}

func (r *PostgresAudienceRepository) Delete(ctx context.Context, aud *domain.Audience) error {
	stmt, err := r.PrepareContext(ctx, DELETE_SQL)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to prepare delete statement: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, aud.FeatureName, aud.AudienceId)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to run delete statement: %s", err.Error())
		return err
	}

	return nil
}

func (r *PostgresAudienceRepository) Save(ctx context.Context, audience *domain.Audience) error {
	query := SAVE_SQL

	args := map[string]interface{}{
		"feature_name": audience.FeatureName,
		"audience_id":  audience.AudienceId,
		"enabled":      audience.Enabled,
		"created_at":   audience.CreatedAt,
		"updated_at":   audience.UpdatedAt,
		"enabled_at":   audience.EnabledAt,
	}

	if !audience.EnabledAt.IsZero() {
		args["enabled_at"] = sql.NullTime{
			Time:  audience.EnabledAt,
			Valid: true,
		}
	} else {
		args["enabled_at"] = sql.NullTime{
			Time:  time.Time{},
			Valid: false,
		}
	}

	query, qargs, err := r.BindNamed(query, args)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to bind query for save operation: %s", err.Error())
		return err
	}

	stmt, err := r.PrepareContext(ctx, query)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to prepare save query: %s", err.Error())
		return err
	}

	_, err = stmt.ExecContext(ctx, qargs...)
	if err != nil {
		r.logger.Errorf("[postgres-audience-repository] failed to save: %s", err.Error())
		return err
	}

	return nil
}

func NewPostgresRepository(db driver.DB, Logger grpclog.LoggerV2) domain.AudienceRepository {
	r := new(PostgresAudienceRepository)

	r.logger = Logger
	r.DB = db

	return r
}
