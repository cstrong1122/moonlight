package moonlight

type IRateOptionProviderAdapter interface {
	GetRateOptions(parameters *RateOptionParameters) ([]RateOption, error)
}

type RateOptionParameters struct {
	FICO  int16
	LTV   float32
	CLTV  float32
	HCLTV float32
}

type RateOption struct {
	InterestRate      float32
	OriginationPoints float32
	DiscountPoints    float32
}
