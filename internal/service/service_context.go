package service

import (
	"database/sql"
	"scaffold/internal/config"
	"scaffold/internal/model"
	"scaffold/pkg/logx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type ServiceContext struct {
	Config      *config.Config
	RedisClient *redis.Client
	DBClient    *model.Queries
	Log         zerolog.Logger
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := sql.Open(c.Database.DriverName, c.Database.DataSource)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: &c,
		RedisClient: redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password, // no password set
			DB:       c.Redis.DB,       // use default DB
		}),
		DBClient: model.New(db),
		Log:      logx.InitLog(c),
	}
}
