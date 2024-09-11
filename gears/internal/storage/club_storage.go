package storage

import (
	"context"
	"errors"
	"goclub/gears/internal/config"
	"goclub/gears/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	clubStorage struct {
		db  *gorm.DB
		dsn config.DSNPair
	}
)

func NewClubStorage(dsn config.DSNPair) *clubStorage {
	return &clubStorage{
		dsn: dsn,
	}
}

func (s *clubStorage) Open(dsns config.DSNPair) (err error) {
	s.db, err = gorm.Open(postgres.Open(dsns.Master), &gorm.Config{PrepareStmt: true})
	return err
}

func (s *clubStorage) Close() (err error) {
	if s.db == nil {
		return errors.New("Database is not opened")
	}
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (s *clubStorage) Migrate() (err error) {
	if err = s.db.AutoMigrate(&model.MemberInfo{}); err != nil {
		return err
	}
	return nil
}

func (s *clubStorage) MemberCreate(ctx context.Context, req *model.MemberCreateReq) (resp *model.MemberCreateResp, err error) {
	tx := s.db.WithContext(ctx).Create(&req.MemberInfo)
	if err = tx.Error; err != nil {
		return nil, err
	}
	resp = &model.MemberCreateResp{
		CRUDResult: model.CRUDResult{RowsAffected: tx.RowsAffected},
		MemberID:   req.MemberInfo.ID,
	}
	return resp, err
}

func (s *clubStorage) MemberUpdate(ctx context.Context, req *model.MemberUpdateReq) (resp *model.MemberUpdateResp, err error) {
	tx := s.db.WithContext(ctx).Select("X").UpdateColumns(&req.MemberInfo) // .Select(...) Избавляет от вызова insert при отсутствии записи в БД
	if err = tx.Error; err != nil {
		return nil, err
	}
	resp = &model.MemberUpdateResp{
		CRUDResult: model.CRUDResult{RowsAffected: tx.RowsAffected},
	}
	return resp, err
}

func (s *clubStorage) MemberDelete(ctx context.Context, req *model.MemberDeleteReq) (resp *model.MemberDeleteResp, err error) {
	tx := s.db.WithContext(ctx).Delete(&model.MemberInfo{ID: req.MemberID})
	if err = tx.Error; err != nil {
		return nil, err
	}
	resp = &model.MemberDeleteResp{
		CRUDResult: model.CRUDResult{RowsAffected: tx.RowsAffected},
	}
	return resp, err
}

func (s *clubStorage) MemberRead(ctx context.Context, req *model.MemberReadReq) (resp *model.MemberReadResp, err error) {
	resp = &model.MemberReadResp{}
	tx := s.db.WithContext(ctx).First(&resp.MemberInfo, req.MemberID)
	if err = tx.Error; err != nil {
		return nil, err
	}
	resp.CRUDResult = model.CRUDResult{RowsAffected: tx.RowsAffected}
	return resp, err
}

func (s *clubStorage) MemberListing(ctx context.Context, req *model.MemberListingReq) (resp *model.MemberListingResp, err error) {
	resp = &model.MemberListingResp{
		Items: make([]model.MemberInfo, 0),
	}
	tx := s.db.WithContext(ctx).Find(&resp.Items)
	if err = tx.Error; err != nil {
		return nil, err
	}
	resp.CRUDResult = model.CRUDResult{RowsAffected: tx.RowsAffected}
	return resp, err
}
