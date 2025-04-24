package app

import (
	"gorm.io/gorm"
	"kryptonim-interview/app/exchanges"
	"kryptonim-interview/app/openexchangerates"
	"kryptonim-interview/app/store"
	"kryptonim-interview/config"
	"sync"
)

type Container struct {
	config                      *config.App
	db                          *gorm.DB
	dbOnce                      sync.Once
	repositories                *Repositories
	repositoriesOnce            sync.Once
	exchangeRatesService        *exchanges.Service
	exchangeRatesServiceOnce    sync.Once
	openExchangeRatesClient     *openexchangerates.Client
	openExchangeRatesClientOnce sync.Once
}

type Repositories struct {
	ExchangeRates *exchanges.Repository
}

func NewContainer() *Container {
	return &Container{
		config: config.LoadConfig(),
	}
}

func (c *Container) NewRepositories() *Repositories {
	c.repositoriesOnce.Do(func() {
		c.repositories = NewRepositories(c.DB())
	})
	return c.repositories
}

func (c *Container) DB() *gorm.DB {
	c.dbOnce.Do(func() {
		c.db = store.InitDriver(c.config)
	})
	return c.db
}

func (c *Container) NewExchangeRatesService() *exchanges.Service {
	c.exchangeRatesServiceOnce.Do(func() {
		c.exchangeRatesService = exchanges.NewService(
			c.NewRepositories().ExchangeRates,
			c.NewOpenExchangeRatesClient(),
		)
	})
	return c.exchangeRatesService
}

func (c *Container) NewOpenExchangeRatesClient() *openexchangerates.Client {
	c.openExchangeRatesClientOnce.Do(func() {
		c.openExchangeRatesClient = openexchangerates.NewClient(
			c.config.OpenExchangeRates.Url,
			c.config.OpenExchangeRates.AppId,
		)
	})
	return c.openExchangeRatesClient
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		ExchangeRates: exchanges.NewRepository(db),
	}
}
