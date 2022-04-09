package cache

import (
	"errors"
	"log"
	"os"
	"strzybug/utils"
	"strzybug/weather"
	"sync"
	"time"
)

type Cache struct {
	// place in filesystem where cached value is beeing stored
	Filename string

	// request information used to fetch new value
	Request weather.Request

	value weather.Response
	mutex sync.Mutex
}

func (c *Cache) Init() error {
	if err := c.loadFromFile(); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := c.ForceUpdate(); err != nil {
				return utils.Wrap(err)
			}
		} else {
			return utils.Wrap(err)
		}
	}

	return utils.Wrap(c.KeepUpdated())
}

func (c *Cache) KeepUpdated() error {
	if c.needsUpdate() {
		return c.ForceUpdate()
	}
	return nil
}

func (c *Cache) ForceUpdate() (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.value, err = c.Request.Run(); err != nil {
		return utils.Wrap(err)
	}
	return utils.Wrap(c.value.ToFile(c.Filename))
}

func (c *Cache) Access() *weather.Response {
	if err := c.KeepUpdated(); err != nil {
		log.Fatalln(utils.Wrap(err))
	}
	return &c.value
}

func (c *Cache) needsUpdate() bool {
	year, month, day := time.Now().Local().Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 1, time.Local)
	return startOfDay.After(c.value.FindFirstDate())
}

func (c *Cache) loadFromFile() (err error) {
	v, err := weather.FromFile(c.Filename)
	if err == nil {
		c.value = v
	}
	return utils.Wrap(err)
}
