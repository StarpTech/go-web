package controllers

import (
	"github.com/go-redis/redis"
	"github.com/starptech/go-web/config"
	"github.com/starptech/go-web/models"
)

type CtrlContext interface {
	GetDB() *models.Model
	GetCache() *redis.Client
	GetConfig() *config.Configuration
}
