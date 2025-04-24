package exchanges

type ExchangeRates struct {
	Id             uint
	CryptoCurrency string
	DecimalPlaces  int
	USDRate        float64
}
