package service

import "github.com/haiyiyun/config"

type Config struct {
	MongoDNS               string          `json:"mongo_dns"`
	MongoDatabaseName      string          `json:"mongo_database_name"`
	CacheDefaultExpiration config.Duration `json:"cache_default_expiration"`
	CacheCleanupInterval   config.Duration `json:"cache_cleanup_interval"`
}
