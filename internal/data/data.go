package data

import (
	"context"

	tintlog "github.com/wxlbd/kit/log"

	"github.com/wxlbd/kratos-pms/internal/data/po"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	redis "github.com/redis/go-redis/v9"
	"github.com/wxlbd/kratos-pms/internal/conf"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGorm, NewRedis, NewProductRepo, NewProductAttributeRepo, NewProductCategoryRepo, NewProductAttributeValueRepo)

// Data .
type Data struct {
	redis.UniversalClient
	*gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, rdb redis.UniversalClient) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		UniversalClient: rdb,
		DB:              db,
	}, cleanup, nil
}

func NewGorm(cfg *conf.Data, l *tintlog.Logger) (*gorm.DB, func()) {
	db, err := gorm.Open(mysql.Open(cfg.Database.Source), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxIdleTime(60)
	sqlDB.SetMaxOpenConns(20)
	f := func() {
		_ = sqlDB.Close()
	}
	// stmt := db.Session(&gorm.Session{DryRun: true}).AutoMigrate(&po.PmsProduct{})
	if err := db.AutoMigrate(&po.PmsProduct{}, &po.PmsProductCategory{}, &po.PmsProductAttribute{}, &po.PmsProductAttributeValue{}, &po.PmsProductSku{}); err != nil {
		// log.NewHelper(logger).Error(err)
		return nil, nil
	}
	return db, f
}

func NewRedis(c *conf.Data) (redis.UniversalClient, error) {
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        c.GetRedis().GetAddrs(),
		DB:           int(c.GetRedis().Db),
		Username:     c.GetRedis().Username,
		Password:     c.GetRedis().Password,
		ReadTimeout:  c.GetRedis().ReadTimeout.AsDuration(),
		WriteTimeout: c.GetRedis().WriteTimeout.AsDuration(),
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
		return nil, err
	}
	return rdb, nil
}
