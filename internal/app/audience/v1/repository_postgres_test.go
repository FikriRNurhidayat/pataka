package audience_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/driver"
	mdriver "github.com/fikrirnurhidayat/ffgo/internal/mocks/driver"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mgrpclog "github.com/fikrirnurhidayat/ffgo/mocks/google.golang.org/grpc/grpclog"
)

type MockPostgresRepository struct {
	dbmock sqlmock.Sqlmock
	db     driver.DB
	logger *mgrpclog.LoggerV2
}

func TestPostgresRepository_Save(t *testing.T) {
	type input struct {
		ctx      context.Context
		audience *domain.Audience
	}

	type output struct {
		err error
	}

	for _, tt := range []struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresRepository, *input, *output)
	}{
		{
			name: "EnabledAt is zero and binding failed",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Time{},
				},
			},
			out: &output{
				err: fmt.Errorf("sqlx.bindNamedMapper: unsupported map type: %T", ""),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mdb := &mdriver.DB{}
				mdb.On("BindNamed", audience.SAVE_SQL, map[string]interface{}{
					"feature_name": i.audience.FeatureName,
					"audience_id":  i.audience.AudienceId,
					"enabled":      i.audience.Enabled,
					"created_at":   i.audience.CreatedAt,
					"updated_at":   i.audience.UpdatedAt,
					"enabled_at": sql.NullTime{
						Time:  time.Time{},
						Valid: false,
					},
				}).Return("", []interface{}{}, o.err)

				mpr.db = mdb
			},
		},
		{
			name: "EnabledAt is zero and prepare statement failed",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Time{},
				},
			},
			out: &output{
				err: fmt.Errorf("sqlx.bindNamedMapper: unsupported map type: %T", ""),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("INSERT INTO").WillReturnError(o.err)
			},
		},
		{
			name: "EnabledAt is not zero and prepare statement failed",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: fmt.Errorf("sqlx.bindNamedMapper: unsupported map type: %T", ""),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("INSERT INTO").WillReturnError(o.err)
			},
		},
		{
			name: "Failed to exec statement",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: fmt.Errorf("sql.exec: failed to execute statement"),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("INSERT INTO feature_audiences").ExpectExec().WillReturnError(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     false,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("INSERT INTO feature_audiences").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
				mpr.dbmock.ExpectPrepare("INSERT INTO").WillReturnError(o.err)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db, sqlmock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}

			defer db.Close()

			dbx := sqlx.NewDb(db, "sqlmock")

			m := &MockPostgresRepository{
				dbmock: sqlmock,
				db:     dbx,
				logger: &mgrpclog.LoggerV2{},
			}

			m.logger.On("Errorf", mock.AnythingOfType("string"), mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := audience.NewPostgresRepository(m.db, m.logger)
			err = subject.Save(tt.in.ctx, tt.in.audience)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}
		})
	}
}

func TestPostgresRepository_List(t *testing.T) {
	type input struct {
		ctx  context.Context
		args *domain.AudienceListArgs
	}

	type output struct {
		audiences []domain.Audience
		err       error
	}

	for _, tt := range []struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresRepository, *input, *output)
	}{
		{
			name: "Filter specified and failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "",
					Filter: &domain.AudienceFilterArgs{},
				},
			},
			out: &output{
				audiences: nil,
				err:       fmt.Errorf("sqlx.bindNamedMapper: unsupported map type: %T", ""),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mdb := &mdriver.DB{}
				mdb.On("BindNamed", "", *i.args.Filter).Return("", []interface{}{}, o.err)

				mpr.db = mdb
			},
		},
		{
			name: "Filter specified but empty and failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "",
					Filter: &domain.AudienceFilterArgs{},
				},
			},
			out: &output{
				audiences: nil,
				err:       sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").WillReturnError(o.err)
			},
		},
		{
			name: "Filter specified and failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audiences: nil,
				err:       sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").WillReturnError(o.err)
			},
		},
		{
			name: "Sort & filter specified and but sort statement failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audiences: nil,
				err:       sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").WillReturnError(o.err)
			},
		},
		{
			name: "Sort & filter specified and failed to query statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audiences: nil,
				err:       sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").ExpectQuery().WillReturnError(o.err)
			},
		},
		{
			name: "Query executed but scan failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				errStr := "sql: Scan error on column index 2, name \"enabled\": sql/driver: couldn't convert %v (%T) into type bool"
				weirdCol := time.Now()

				o.err = fmt.Errorf(errStr, weirdCol, weirdCol)

				rows := sqlmock.NewRows([]string{"feature_name", "audience_id", "enabled", "created_at", "updated_at", "enabled_at"}).
					AddRow("midtrans-payment", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", weirdCol, time.Now(), time.Now(), sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					})

				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").ExpectQuery().WillReturnRows(rows)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceListArgs{
					Limit:  10,
					Offset: 0,
					Sort:   "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audiences: []domain.Audience{{
					FeatureName: "midtrans-payment",
					AudienceId:  "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				}},
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				rows := sqlmock.NewRows([]string{"feature_name", "audience_id", "enabled", "created_at", "updated_at", "enabled_at"})

				for _, a := range o.audiences {
					rows.AddRow(a.FeatureName, a.AudienceId, a.Enabled, a.CreatedAt, a.UpdatedAt, sql.NullTime{
						Time:  a.EnabledAt,
						Valid: true,
					})
				}

				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.* FROM feature_audiences").ExpectQuery().WillReturnRows(rows)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			db, sqlmock, err := sqlmock.New()
			if err != nil {
				t.Fatal(err)
			}

			defer db.Close()

			dbx := sqlx.NewDb(db, "sqlmock")

			m := &MockPostgresRepository{
				dbmock: sqlmock,
				db:     dbx,
				logger: &mgrpclog.LoggerV2{},
			}

			m.logger.On("Errorf", mock.AnythingOfType("string"), mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := audience.NewPostgresRepository(m.db, m.logger)
			audiences, err := subject.List(tt.in.ctx, tt.in.args)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			}

			if tt.out.audiences != nil {
				assert.NotNil(t, audiences)
				assert.Equal(t, tt.out.audiences, audiences)
			}
		})
	}
}

func TestPostgresRepository_Get(t *testing.T) {
	t.SkipNow()
}

func TestPostgresRepository_GetBy(t *testing.T) {
	t.SkipNow()
}

func TestPostgresRepository_Delete(t *testing.T) {
	t.SkipNow()
}

func TestPostgresRepository_DeleteBy(t *testing.T) {
	t.SkipNow()
}
