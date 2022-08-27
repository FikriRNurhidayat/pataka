package audience_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
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

var AudiencePostgresRepositoryCols []string = []string{"feature_name", "audience_id", "enabled", "created_at", "updated_at", "enabled_at"}

func TestAudiencePostgresRepository_Save(t *testing.T) {
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

func TestAudiencePostgresRepository_List(t *testing.T) {
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
			name: "Sort, filter, and but sort statement failed",
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
			name: "Sort, filter, and failed to query statement",
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

				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols).
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
				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols)

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
			} else {
				assert.NoError(t, err)
			}

			if tt.out.audiences != nil {
				assert.NotNil(t, audiences)
				assert.Equal(t, tt.out.audiences, audiences)
			} else {
				assert.Nil(t, audiences)
			}
		})
	}
}

func TestAudiencePostgresRepository_Get(t *testing.T) {
	type input struct {
		ctx context.Context
		fn  string
		ui  string
	}

	type output struct {
		audience *domain.Audience
		err      error
	}

	for _, tt := range []struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresRepository, *input, *output)
	}{
		{
			name: "Failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				fn:  "midtrans-payment",
				ui:  "dca62b2e-db28-438c-ad6e-383bfdecee72",
			},
			out: &output{
				audience: nil,
				err:      sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").WillReturnError(o.err)
			},
		},
		{
			name: "Failed to query statement",
			in: &input{
				ctx: context.Background(),
				fn:  "midtrans-payment",
				ui:  "dca62b2e-db28-438c-ad6e-383bfdecee72",
			},
			out: &output{
				audience: nil,
				err:      sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").ExpectQuery().WillReturnError(o.err)
			},
		},
		{
			name: "Failed to scan rows",
			in: &input{
				ctx: context.Background(),
				fn:  "midtrans-payment",
				ui:  "dca62b2e-db28-438c-ad6e-383bfdecee72",
			},
			out: &output{},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				errStr := "sql: Scan error on column index 2, name \"enabled\": sql/driver: couldn't convert %v (%T) into type bool"
				weirdCol := time.Now()

				o.err = fmt.Errorf(errStr, weirdCol, weirdCol)

				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols).
					AddRow("midtrans-payment", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", weirdCol, time.Now(), time.Now(), sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					})

				mpr.dbmock.ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").ExpectQuery().WillReturnRows(rows)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				fn:  "midtrans-payment",
				ui:  "dca62b2e-db28-438c-ad6e-383bfdecee72",
			},
			out: &output{
				audience: &domain.Audience{
					FeatureName: "midtrans-payment",
					AudienceId:  "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols)

				rows.AddRow(o.audience.FeatureName, o.audience.AudienceId, o.audience.Enabled, o.audience.CreatedAt, o.audience.UpdatedAt, sql.NullTime{
					Time:  o.audience.EnabledAt,
					Valid: true,
				})

				mpr.dbmock.
					ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").
					ExpectQuery().
					WillReturnRows(rows)
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
			audience, err := subject.Get(tt.in.ctx, tt.in.fn, tt.in.ui)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.out.audience != nil {
				assert.NotNil(t, audience)
				assert.Equal(t, tt.out.audience, audience)
			} else {
				assert.Nil(t, audience)
			}
		})
	}
}

func TestPostgresRepository_GetBy(t *testing.T) {
	type input struct {
		ctx  context.Context
		args *domain.AudienceGetByArgs
	}

	type output struct {
		audience *domain.Audience
		err      error
	}

	for _, tt := range []struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresRepository, *input, *output)
	}{
		{
			name: "Failed to bind statement during filtering",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceGetByArgs{
					Sort: "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment-",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audience: nil,
				err:      fmt.Errorf("sqlx.bindNamedMapper: failed to bind argumnets"),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mdb := &mdriver.DB{}
				mdb.On("BindNamed", "(feature_audiences.enabled = :enabled) AND (feature_audiences.feature_name = :feature_name) AND (feature_audiences.audience_id IN (:audience_id))", *i.args.Filter).Return("", []interface{}{}, o.err)

				mpr.db = mdb
			},
		},
		{
			name: "Filter and sort but failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceGetByArgs{
					Sort: "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment-",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audience: nil,
				err:      sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").
					WillReturnError(o.err)
			},
		},
		{
			name: "Statement prepared but query failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceGetByArgs{
					Sort: "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment-",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audience: nil,
				err:      sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").
					ExpectQuery().
					WithArgs(i.args.Filter.Enabled, i.args.Filter.FeatureName, i.args.Filter.AudienceIds[0], i.args.Filter.AudienceIds[1], 1, 0).
					WillReturnError(o.err)
			},
		},
		{
			name: "Query ok but scan failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceGetByArgs{
					Sort: "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment-",
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

				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols).
					AddRow("midtrans-payment", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", weirdCol, time.Now(), time.Now(), sql.NullTime{
						Time:  time.Now(),
						Valid: true,
					})

				mpr.dbmock.
					ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").
					ExpectQuery().
					WithArgs(i.args.Filter.Enabled, i.args.Filter.FeatureName, i.args.Filter.AudienceIds[0], i.args.Filter.AudienceIds[1], 1, 0).
					WillReturnRows(rows)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceGetByArgs{
					Sort: "-created_at",
					Filter: &domain.AudienceFilterArgs{
						FeatureName: "midtrans-payment-",
						AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
						Enabled:     new(bool),
					},
				},
			},
			out: &output{
				audience: &domain.Audience{
					FeatureName: "midtrans-payment",
					AudienceId:  "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				rows := sqlmock.NewRows(AudiencePostgresRepositoryCols)

				rows.AddRow(o.audience.FeatureName, o.audience.AudienceId, o.audience.Enabled, o.audience.CreatedAt, o.audience.UpdatedAt, sql.NullTime{
					Time:  o.audience.EnabledAt,
					Valid: true,
				})

				mpr.dbmock.
					ExpectPrepare("SELECT feature_audiences.feature_name, feature_audiences.audience_id, feature_audiences.enabled, feature_audiences.created_at, feature_audiences.updated_at, feature_audiences.enabled_at FROM feature_audiences").
					ExpectQuery().
					WithArgs(i.args.Filter.Enabled, i.args.Filter.FeatureName, i.args.Filter.AudienceIds[0], i.args.Filter.AudienceIds[1], 1, 0).
					WillReturnRows(rows)
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
			audience, err := subject.GetBy(tt.in.ctx, tt.in.args)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.out.audience != nil {
				assert.NotNil(t, audience)
				assert.Equal(t, tt.out.audience, audience)
			} else {
				assert.Nil(t, audience)
			}
		})
	}
}

func TestAudiencePostgresRepository_Delete(t *testing.T) {
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
			name: "Failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "midtrans-payment",
					AudienceId:  "f554c37a-3ed7-4d31-bc88-c2f4dbbf9ed5",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.ExpectPrepare("DELETE FROM feature_audiences").WillReturnError(o.err)
			},
		},
		{
			name: "Failed to execute statement",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "midtrans-payment",
					AudienceId:  "f554c37a-3ed7-4d31-bc88-c2f4dbbf9ed5",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare("DELETE FROM feature_audiences").
					ExpectExec().
					WithArgs(i.audience.FeatureName, i.audience.AudienceId).
					WillReturnError(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "midtrans-payment",
					AudienceId:  "f554c37a-3ed7-4d31-bc88-c2f4dbbf9ed5",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare("DELETE FROM feature_audiences").
					ExpectExec().
					WithArgs(i.audience.FeatureName, i.audience.AudienceId).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
			err = subject.Delete(tt.in.ctx, tt.in.audience)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestPostgresRepository_Size(t *testing.T) {
	type input struct {
		ctx  context.Context
		args *domain.AudienceFilterArgs
	}

	type output struct {
		size uint32
		err  error
	}

	for _, tt := range []struct {
		name string
		in   *input
		out  *output
		on   func(*MockPostgresRepository, *input, *output)
	}{
		{
			name: "Failed to bind statement during filtering",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				size: 0,
				err:  fmt.Errorf("sqlx.bindNamedMapper: failed to bind argumnets"),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mdb := &mdriver.DB{}
				mdb.On("BindNamed", "(feature_audiences.enabled = :enabled) AND (feature_audiences.feature_name = :feature_name) AND (feature_audiences.audience_id IN (:audience_id))", *i.args).Return("", []interface{}{}, o.err)

				mpr.db = mdb
			},
		},
		{
			name: "Filter but failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				size: 0,
				err:  sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare(regexp.QuoteMeta(audience.SIZE_SQL)).
					WillReturnError(o.err)
			},
		},
		{
			name: "Statement prepared but execution failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				size: 0,
				err:  sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare(regexp.QuoteMeta(audience.SIZE_SQL)).
					ExpectQuery().
					WithArgs(i.args.Enabled, i.args.FeatureName, i.args.AudienceIds[0], i.args.AudienceIds[1]).
					WillReturnError(o.err)
			},
		},
		{
			name: "Execution is ok but scan failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				size: 0,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				o.err = fmt.Errorf("sql: Scan error on column index 0, name \"count\": converting driver.Value type bool (\"true\") to a int: invalid syntax")
				rows := sqlmock.NewRows([]string{"count"}).AddRow(true)

				mpr.dbmock.
					ExpectPrepare(regexp.QuoteMeta(audience.SIZE_SQL)).
					ExpectQuery().
					WithArgs(i.args.Enabled, i.args.FeatureName, i.args.AudienceIds[0], i.args.AudienceIds[1]).
					WillReturnRows(rows)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				size: 2,
				err:  nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				rows := sqlmock.NewRows([]string{"count"}).
					AddRow(2)

				mpr.dbmock.
					ExpectPrepare(regexp.QuoteMeta(audience.SIZE_SQL)).
					ExpectQuery().
					WithArgs(i.args.Enabled, i.args.FeatureName, i.args.AudienceIds[0], i.args.AudienceIds[1]).
					WillReturnRows(rows)
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
			size, err := subject.Size(tt.in.ctx, tt.in.args)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			if tt.out.size != 0 {
				assert.NotZero(t, size)
				assert.Equal(t, tt.out.size, size)
			}
		})
	}
}

func TestPostgresRepository_DeleteBy(t *testing.T) {
	type input struct {
		ctx  context.Context
		args *domain.AudienceFilterArgs
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
			name: "Failed to bind statement during filtering",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				err: fmt.Errorf("sqlx.bindNamedMapper: failed to bind argumnets"),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mdb := &mdriver.DB{}
				mdb.On("BindNamed", "(feature_audiences.enabled = :enabled) AND (feature_audiences.feature_name = :feature_name) AND (feature_audiences.audience_id IN (:audience_id))", *i.args).Return("", []interface{}{}, o.err)

				mpr.db = mdb
			},
		},
		{
			name: "Filter but failed to prepare statement",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				err: sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare("DELETE FROM feature_audiences").
					WillReturnError(o.err)
			},
		},
		{
			name: "Statement prepared but execution failed",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				err: sql.ErrConnDone,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare(audience.DELETE_BY_SQL).
					ExpectExec().
					WithArgs(i.args.Enabled, i.args.FeatureName, i.args.AudienceIds[0], i.args.AudienceIds[1]).
					WillReturnError(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				args: &domain.AudienceFilterArgs{
					FeatureName: "midtrans-payment-",
					AudienceIds: []string{"75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0", "75d0e418-8b7b-4fc4-bd95-3c5cb3da3bb0"},
					Enabled:     new(bool),
				},
			},
			out: &output{
				err: nil,
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				mpr.dbmock.
					ExpectPrepare(audience.DELETE_BY_SQL).
					ExpectExec().
					WithArgs(i.args.Enabled, i.args.FeatureName, i.args.AudienceIds[0], i.args.AudienceIds[1]).
					WillReturnResult(sqlmock.NewResult(1, 1))
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
			err = subject.DeleteBy(tt.in.ctx, tt.in.args)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
