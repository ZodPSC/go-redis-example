package storage

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Config struct {
	Addr        string        `yaml:"addr"`
	DB          int           `yaml:"db"`
	User        string        `yaml:"user"`
	Password    string        `yaml:"password"`
	MaxRetries  int           `yaml:"max_retries"`
	DialTimeout time.Duration `yaml:"dial_timeout"`
	Timeout     time.Duration `yaml:"timeout"`
}

func NewClient(ctx context.Context, cfg Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		DB:           cfg.DB,
		Password:     cfg.Password,
		Username:     cfg.User,
		MaxRetries:   cfg.MaxRetries,
		DialTimeout:  cfg.DialTimeout,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
	})

	if err := db.Ping(ctx).Err(); err != nil {
		fmt.Printf("failed to connect to redis server: %s\n", err.Error())
		return nil, err
	}

	return db, nil
}
