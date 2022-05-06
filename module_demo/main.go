package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Cache struct {
	sync.Map
}

func main() {
	cache := Cache{Map: sync.Map{}}

	cache.Store("i", "1")
	cache.Store("ji", "2")
	cache.Store("yin", "3")

	cache.Range(func(key, value interface{}) bool {
		logrus.Infof("key: %s, value: %s", key, value)
		return true
	})
}
