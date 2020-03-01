package decimal

import "github.com/shopspring/decimal"

type springDecimal struct {
	decimal.Decimal
}

func (d springDecimal) Int64() int64 {
	return d.Decimal.IntPart()
}

func (d springDecimal) Float64() float64 {
	f, _ := d.Decimal.Float64()
	return f
}

func (d springDecimal) Add(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Add(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Sub(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Sub(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Neg() Decimal {
	return springDecimal{d.Decimal.Neg()}
}

func (d springDecimal) Mul(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Mul(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Div(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Div(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Mod(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Mod(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Pow(d2 Decimal) Decimal {
	return springDecimal{d.Decimal.Pow(d2.(springDecimal).Decimal)}
}

func (d springDecimal) Round(places int32) Decimal {
	return springDecimal{d.Decimal.Round(places)}
}

func (d springDecimal) Cmp(d2 Decimal) int {
	return d.Decimal.Cmp(d2.(springDecimal).Decimal)
}

type springDecimalCreator struct {
}

func (*springDecimalCreator) NewFromFloat32(val float32) Decimal {
	return springDecimal{decimal.NewFromFloat32(val)}
}

func (*springDecimalCreator) NewFromFloat64(val float64) Decimal {
	return springDecimal{decimal.NewFromFloat(val)}
}


func (*springDecimalCreator) NewFromString(val string) (Decimal, error) {
	d , err :=decimal.NewFromString(val)
	return springDecimal{d}, err
}
