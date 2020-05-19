package utils

import (
	"fmt"

	"github.com/xc407/agollo/v3/agcache"
)

const (
	propertiesFormat = "%s=%s\n"

	defaultContentKey = "content"
)

//ContentParser 内容转换
type ContentParser interface {
	Parse(cache agcache.CacheInterface) (string, error)
}

//DefaultParser 默认内容转换器
type DefaultParser struct {
}

//Parse 内存内容默认转换器
func (d *DefaultParser) Parse(cache agcache.CacheInterface) (string, error) {
	if cache == nil {
		return Empty, nil
	}

	value, err := cache.Get(defaultContentKey)
	if err != nil {
		return Empty, err
	}
	return string(value), nil
}

//PropertiesParser properties转换器
type PropertiesParser struct {
}

//Parse 内存内容=>properties文件转换器
func (d *PropertiesParser) Parse(cache agcache.CacheInterface) (string, error) {
	properties := convertToProperties(cache)
	return properties, nil
}

func convertToProperties(cache agcache.CacheInterface) string {
	properties := Empty
	if cache == nil {
		return properties
	}
	cache.Range(func(key, value interface{}) bool {
		properties += fmt.Sprintf(propertiesFormat, key, string(value.([]byte)))
		return true
	})
	return properties
}
