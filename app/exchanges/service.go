package exchanges

import (
	"github.com/shopspring/decimal"
	"kryptonim-interview/app/errs"
	"kryptonim-interview/app/openexchangerates"
	"strings"
)

type Rate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

type ExchangeData struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type Service struct {
	repo                    *Repository
	openExchangeRatesClient *openexchangerates.Client
}

func NewService(
	repo *Repository,
	openExchangeRatesClient *openexchangerates.Client,
) *Service {
	return &Service{
		repo:                    repo,
		openExchangeRatesClient: openExchangeRatesClient,
	}
}

func (s *Service) Exchange(from, to string, amount decimal.Decimal) (*ExchangeData, error) {
	fromRate, err := s.repo.FindByCryptoCurrency(from)
	if err != nil {
		return nil, errs.Wrap(err, "[21251251] cannot find from rate")
	}

	toRate, err := s.repo.FindByCryptoCurrency(to)
	if err != nil {
		return nil, errs.Wrap(err, "[213434] cannot find to rate")
	}

	fromRateDecimal := decimal.NewFromFloat(fromRate.USDRate)
	toRateDecimal := decimal.NewFromFloat(toRate.USDRate)

	amountInUSD := amount.Mul(fromRateDecimal)
	amountInToCurrency := amountInUSD.Div(toRateDecimal)
	amountInToCurrency = amountInToCurrency.Round(int32(toRate.DecimalPlaces))
	result, _ := amountInToCurrency.Float64()

	return &ExchangeData{
		From:   from,
		To:     to,
		Amount: result,
	}, nil
}

func (s *Service) Rates(currencies []string) (interface{}, error) {
	rates, err := s.openExchangeRatesClient.GetExchangeRates(currencies)
	if err != nil {
		return nil, errs.Wrap(err, "[21328328] cannot get exchange rates")
	}

	return s.computeRates(currencies, rates), nil
}

func (s *Service) computeRates(currencies []string, clientResponse *openexchangerates.RatesResponse) []*Rate {
	var rates []*Rate

	for _, from := range currencies {
		for _, to := range currencies {
			if from != to {
				fUpper := strings.ToUpper(from)
				tUpper := strings.ToUpper(to)

				fromRate := decimal.NewFromFloat(clientResponse.Rates[fUpper])
				toRate := decimal.NewFromFloat(clientResponse.Rates[tUpper])
				rate := fromRate.Div(toRate)
				resultRate, _ := rate.Float64()
				rates = append(rates, &Rate{
					From: fUpper,
					To:   tUpper,
					Rate: resultRate,
				})
			}
		}
	}

	return rates
}
