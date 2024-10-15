package tool

import "github.com/shopspring/decimal"

type Decimal struct {
	d decimal.Decimal
}

func NewDecimal(f float64) Decimal {
	return Decimal{
		d: decimal.NewFromFloat(f),
	}
}

func (o Decimal) Add(f float64) Decimal {
	return Decimal{
		d: o.d.Add(decimal.NewFromFloat(f)),
	}
}

func (o Decimal) Sub(f float64) Decimal {
	return Decimal{
		d: o.d.Sub(decimal.NewFromFloat(f)),
	}
}

func (o Decimal) Mul(f float64) Decimal {
	return Decimal{
		d: o.d.Mul(decimal.NewFromFloat(f)),
	}
}

func (o Decimal) Truncate(precision int32) Decimal {
	return Decimal{
		d: o.d.Truncate(precision),
	}
}

func (o Decimal) Float64() (f float64) {
	f, _ = o.d.Float64()
	return
}
