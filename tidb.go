package decimal

import (
	types "github.com/x1nchen/tidb-decimal"
)

const (
	MaxDecimalScale   = 30
)

var (
	spring = &springDecimalCreator{}
)

func init() {
	creator = &tidbDecimalCreator{}
}

type tidbDecimal struct {
	*types.MyDecimal
}

func (d tidbDecimal) Int64() int64 {
	x, _ := d.MyDecimal.ToInt()
	return x
}

func (d tidbDecimal) Float64() float64 {
	x, _ := d.MyDecimal.ToFloat64()
	return x
}

func (d tidbDecimal) Add(d2 Decimal) Decimal {
	var sum types.MyDecimal
	_ = types.DecimalAdd(d.MyDecimal, d2.(tidbDecimal).MyDecimal, &sum)
	return tidbDecimal{&sum}
}

func (d tidbDecimal) Sub(d2 Decimal) Decimal {
	var sum types.MyDecimal
	_ = types.DecimalSub(d.MyDecimal, d2.(tidbDecimal).MyDecimal, &sum)
	return tidbDecimal{&sum}
}

func (d tidbDecimal) Neg() Decimal {
	return tidbDecimal{types.DecimalNeg(d.MyDecimal)}
}

func (d tidbDecimal) Mul(d2 Decimal) Decimal {
	var sum types.MyDecimal
	_ = types.DecimalMul(d.MyDecimal, d2.(tidbDecimal).MyDecimal, &sum)
	return tidbDecimal{&sum}
}

func (d tidbDecimal) Div(d2 Decimal) Decimal {
	var sum types.MyDecimal
	_ = types.DecimalDiv(d.MyDecimal, d2.(tidbDecimal).MyDecimal, &sum, MaxDecimalScale)
	return tidbDecimal{&sum}
}

func (d tidbDecimal) Mod(d2 Decimal) Decimal {
	var sum types.MyDecimal
	_ = types.DecimalMod(d.MyDecimal, d2.(tidbDecimal).MyDecimal, &sum)
	return tidbDecimal{&sum}
}

func (d tidbDecimal) Pow(d2 Decimal) Decimal {
	val := d2.Float64()
	mval := d.Float64()

	a := spring.NewFromFloat64(mval)
	b := spring.NewFromFloat64(val)

	var o types.MyDecimal
	_ = o.FromFloat64(a.Pow(b).Float64())
	return tidbDecimal{&o}
}

func (d tidbDecimal) IsPositive() bool {
	return !d.MyDecimal.IsNegative() && !d.MyDecimal.IsZero()
}

func (d tidbDecimal) Round(places int32) Decimal {
	var sum types.MyDecimal
	_ = d.MyDecimal.Round(&sum, int(places), types.ModeHalfEven)
	return tidbDecimal{&sum}
}

type tidbDecimalCreator struct {
}

func (t *tidbDecimalCreator) NewFromFloat32(val float32) Decimal {
	return t.NewFromFloat64(float64(val))
}

func (*tidbDecimalCreator) NewFromFloat64(val float64) Decimal {
	d := new(types.MyDecimal)
	_ = d.FromFloat64(float64(val))
	return tidbDecimal{d}
}
