package decimal

var Zero = NewFromFloat64(0)

type Decimal interface {
	// Add returns d + d2.
	Add(d2 Decimal) Decimal
	// Sub returns d - d2.
	Sub(d2 Decimal) Decimal
	// Neg returns -d.
	Neg() Decimal
	// Mul returns d * d2.
	Mul(d2 Decimal) Decimal
	// Div returns d / d2.
	Div(d2 Decimal) Decimal
	// Mod returns d % d2.
	Mod(d2 Decimal) Decimal
	// Pow returns d to the power d2
	Pow(d2 Decimal) Decimal
	// IsPositive return
	//
	//	true if d > 0
	//	false if d == 0
	//	false if d < 0
	IsPositive() bool
	// IsNegative return
	//
	//	true if d < 0
	//	false if d == 0
	//	false if d > 0
	IsNegative() bool
	// IsZero return
	//
	//	true if d == 0
	//	false if d > 0
	//	false if d < 0
	IsZero() bool
	// Round rounds the decimal to places decimal places.
	Round(places int32) Decimal
	Int64() int64
	Float64() float64
	String() string
}

type decimalCreator interface {
	NewFromFloat32(val float32) Decimal
	NewFromFloat64(val float64) Decimal
}

var (
	creator decimalCreator = &springDecimalCreator{}
)

func NewFromFloat32(val float32) Decimal {
	return creator.NewFromFloat32(val)
}

func NewFromFloat64(val float64) Decimal {
	return creator.NewFromFloat64(val)
}
