package store

import (
	"gorm.io/gorm"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/app/exchanges"
)

var AppModels = map[string]interface{}{
	"ExchangeRates": exchanges.ExchangeRates{},
}

func Migrate(db *gorm.DB) error {
	for _, m := range AppModels {
		err := db.AutoMigrate(m)
		if err != nil {
			return errs.Wrap(err, "[21226226] cannot migrate")
		}
	}
	return nil
}
