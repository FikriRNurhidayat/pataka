package audience_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/fikrirnurhidayat/ffgo/internal/app/audience/v1"
	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	mdriver "github.com/fikrirnurhidayat/ffgo/internal/mocks/driver"
	mgrpclog "github.com/fikrirnurhidayat/ffgo/mocks/google.golang.org/grpc/grpclog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostgresRepository struct {
	DB     *mdriver.DB
	Logger *mgrpclog.LoggerV2
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
				args := map[string]interface{}{
					"feature_name": i.audience.FeatureName,
					"audience_id":  i.audience.AudienceId,
					"enabled":      i.audience.Enabled,
					"created_at":   i.audience.CreatedAt,
					"updated_at":   i.audience.UpdatedAt,
					"enabled_at": sql.NullTime{
						Time:  time.Time{},
						Valid: false,
					},
				}

				mpr.DB.On("BindNamed", audience.SAVE_SQL, args).Return("", []interface{}{}, o.err)
			},
		},
		{
			name: "EnabledAt is not zero and binding failed",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: fmt.Errorf("sqlx.bindNamedMapper: unsupported map type: %T", ""),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				args := map[string]interface{}{
					"feature_name": i.audience.FeatureName,
					"audience_id":  i.audience.AudienceId,
					"enabled":      i.audience.Enabled,
					"created_at":   i.audience.CreatedAt,
					"updated_at":   i.audience.UpdatedAt,
					"enabled_at": sql.NullTime{
						Time:  i.audience.EnabledAt,
						Valid: true,
					},
				}

				mpr.DB.On("BindNamed", audience.SAVE_SQL, args).Return("", []interface{}{}, o.err)
			},
		},
		{
			name: "Prepare statement failed",
			in: &input{
				ctx: context.Background(),
				audience: &domain.Audience{
					FeatureName: "ios/in-app-payment",
					AudienceId:  "a044ec0f-d04b-49d3-b211-80719a6ccd49",
					Enabled:     true,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					EnabledAt:   time.Now(),
				},
			},
			out: &output{
				err: fmt.Errorf("sql: database is closed"),
			},
			on: func(mpr *MockPostgresRepository, i *input, o *output) {
				args := map[string]interface{}{
					"feature_name": i.audience.FeatureName,
					"audience_id":  i.audience.AudienceId,
					"enabled":      i.audience.Enabled,
					"created_at":   i.audience.CreatedAt,
					"updated_at":   i.audience.UpdatedAt,
					"enabled_at": sql.NullTime{
						Time:  i.audience.EnabledAt,
						Valid: true,
					},
				}

				mpr.DB.On("BindNamed", audience.SAVE_SQL, args).Return("", []interface{}{}, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			m := &MockPostgresRepository{
				DB:     &mdriver.DB{},
				Logger: &mgrpclog.LoggerV2{},
			}

			m.Logger.On("Errorf", mock.AnythingOfType("string"), mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := audience.NewPostgresRepository(m.DB, m.Logger)
			err := subject.Save(tt.in.ctx, tt.in.audience)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}
		})
	}
}

func TestPostgresRepository_List(t *testing.T) {
	t.SkipNow()
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
