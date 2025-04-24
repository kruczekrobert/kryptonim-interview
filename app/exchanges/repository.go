package exchanges

import (
	"gorm.io/gorm"
	"kryptonim-interview/app/errs"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) baseQuery() *gorm.DB {
	return r.db.Model(&ExchangeRates{})
}

func (r *Repository) FindByCryptoCurrency(symbol string) (*ExchangeRates, error) {
	var rate ExchangeRates
	//NOTE: gormowe where z = ? jest odporne na SQL Injection jakby co
	err := r.baseQuery().
		Where("crypto_currency = ?", symbol).
		Take(&rate).Error

	if err != nil {
		//TODO: tutaj bym zrobił inny wrap bo wypadało by łapać czy to not found czy inny problem ale no to zadanie na rekrutacje a nie gotowy produkt :D
		return nil, errs.Wrap(err, "[21345345] cannot find crypto currency in db", errs.Context(symbol))
	}

	return &rate, nil
}
