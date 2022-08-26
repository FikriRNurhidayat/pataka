package transaction

import (
	"context"

	"github.com/fikrirnurhidayat/ffgo/internal/domain/v1"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

type tx struct {
	db      *sqlx.DB
	logger  grpclog.LoggerV2
	factory *RepositoryFactory
}

func (u *tx) Do(ctx context.Context, block domain.Block) error {
	tx, err := u.db.Beginx()
	if err != nil {
		u.logger.Errorf("[unit-of-work] failed to start transaction: %s", err.Error())
		return err
	}

	repository := makeRepository(tx, u.logger, u.factory)

	if err := block(repository); err != nil {
		if err := tx.Rollback(); err != nil {
			u.logger.Errorf("[unit-of-work] failed to abort transaction: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		u.logger.Errorf("[unit-of-work] failed to commit transaction: %s", err.Error())
		if err := tx.Rollback(); err != nil {
			u.logger.Errorf("[unit-of-work] failed to abort transaction: %s", err.Error())
			return status.Error(codes.Internal, "Internal server error")
		}

		return status.Error(codes.Internal, "Internal server error")
	}

	return nil
}

func New(db *sqlx.DB, logger grpclog.LoggerV2, factory *RepositoryFactory) domain.UnitOfWork {
	return &tx{
		db:      db,
		logger:  logger,
		factory: factory,
	}
}
