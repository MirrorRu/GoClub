package serviceprovider

import (
	"context"
	"errors"

	clubapi "goclub/gears/internal/club_api"
	clubservice "goclub/gears/internal/club_service"
	"goclub/gears/internal/config"
	"goclub/gears/internal/logger"
	"goclub/gears/internal/storage"
)

const pkgLogName = "serviceProvider"

type ServiceProvider struct {
	clubAPI *clubapi.API
	clubSvc clubservice.ClubService
	// services     services
	// repositories repositories
	storage storage.ClubStorage
	cfg     *config.AppConfig
}

// type services struct {
// 	//	gearsService gearsservice.Service
// }

// type repositories struct {

// 	// dbClient     db.Client
// 	// gearsStorage gearsstorage.Repository
// }

func (sp *ServiceProvider) GetClubService(ctx context.Context) (clubservice.ClubService, error) {
	if sp.clubSvc == nil {
		storage, err := sp.GetStorage()
		if err != nil {
			return nil, err
		}
		sp.clubSvc = clubservice.NewClubService(ctx, storage)
	}
	return sp.clubSvc, nil
}

func (sp *ServiceProvider) GetClubAPI(ctx context.Context) (*clubapi.API, error) {
	if sp.clubAPI == nil {
		clubSvc, err := sp.GetClubService(ctx)
		if err != nil {
			return nil, err
		}
		sp.clubAPI = clubapi.NewAPI(
			clubSvc,
		)
	}
	return sp.clubAPI, nil
}

func (sp *ServiceProvider) GetStorage() (_ storage.ClubStorage, err error) {
	if sp.storage == nil {
		return nil, errors.New("Storage not opened")
	}
	return sp.storage, nil
}

func (sp *ServiceProvider) OpenStorage(ctx context.Context) (err error) {
	logger.Info(ctx, "Open storage")
	err = sp.storage.Open(sp.cfg.DSNs)
	if err == nil {
		err = sp.storage.Migrate()
	}
	return err
}

func (sp *ServiceProvider) CloseStorage(ctx context.Context) error {
	logger.Info(ctx, "Closing storage")
	if sp.storage == nil {
		logger.Warn(ctx, "Storage not opened")
		return nil
	}
	return sp.storage.Close()
}

func NewServiceProvider(cfg *config.AppConfig) *ServiceProvider {
	return &ServiceProvider{
		cfg:     cfg,
		storage: storage.NewClubStorage(cfg.DSNs),
	}
}
