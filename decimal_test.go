package decimal

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func DecimalNewFromFloat64(creator decimalCreator) {
	n := 0.0
	for i := 0; i < 1000; i++ {
		n += 0.01
	}

	d := creator.NewFromFloat64(n)
	So(d, ShouldNotBeNil)
	f := d.Float64()
	So(f, ShouldEqual, n)

	j := 0.1
	n = -2.0
	So(math.Pow(j, n), ShouldNotEqual, 100)
}

func TestNewFromFloat64(t *testing.T) {
	Convey(`测试 NewFromFloat64`, t, func() {
		DecimalNewFromFloat64(&tidbDecimalCreator{})
		DecimalNewFromFloat64(&springDecimalCreator{})
	})
}

func TestNewFromString(t *testing.T) {
	Convey(`tidb NewFromString`, t, func() {
		creator := &tidbDecimalCreator{}
		d, err := creator.NewFromString("1.5")
		So(err, ShouldBeNil)
		So(d, ShouldResemble, creator.NewFromFloat64(1.5))
	})

	Convey(`spring NewFromString`, t, func() {
		creator := &springDecimalCreator{}
		d, err := creator.NewFromString("1.5")
		So(err, ShouldBeNil)
		So(d, ShouldResemble, creator.NewFromFloat64(1.5))
	})
}

func Decimal1Div3Float64(creator decimalCreator) {
	a := creator.NewFromFloat64(1)
	So(a, ShouldNotBeNil)

	b := creator.NewFromFloat64(3)
	So(b, ShouldNotBeNil)

	c := a.Div(b)
	So(c, ShouldNotBeNil)

	d := c.Float64()
	So(d, ShouldNotBeNil)
}

func Test1Div3Float64(t *testing.T) {
	Convey(`测试 1/3`, t, func() {
		Decimal1Div3Float64(&tidbDecimalCreator{})
		Decimal1Div3Float64(&springDecimalCreator{})
	})
}

func TestFloat64Add(t *testing.T) {
	Convey("Float64 小数相加", t, func() {
		n := 0.0
		s := 0.01
		for i := 0; i < 1000; i++ {
			n += s
		}
		So(n, ShouldNotEqual, 10)
	})
}

func AddValue(creator decimalCreator) {
	n := creator.NewFromFloat64(0)
	s := creator.NewFromFloat64(0.01)

	for i := 0; i < 1000; i++ {
		n = n.Add(s)
	}

	val := n.Int64()
	So(val, ShouldEqual, 10)

	fval := n.Float64()
	So(fval, ShouldEqual, 10)
}

func SubValue(creator decimalCreator) {
	n := creator.NewFromFloat64(10)
	s := creator.NewFromFloat64(0.01)

	for i := 0; i < 1000; i++ {
		n = n.Sub(s)
	}

	val := n.Int64()
	So(val, ShouldEqual, 0)

	fval := n.Float64()
	So(fval, ShouldEqual, 0)
}

func MultiplyValue(creator decimalCreator) {
	n := creator.NewFromFloat64(1)
	w := creator.NewFromFloat64(0.1)

	for i := 0; i < 30; i++ {
		n = n.Mul(w)
	}

	fval := n.Float64()
	So(fval, ShouldEqual, 1e-30)
}

func MultiplyValue2(creator decimalCreator) {
	n := creator.NewFromFloat64(0.1)
	w := creator.NewFromFloat64(0.1)

	for i := 0; i < 10; i++ {
		n = n.Mul(w)
	}

	fval := n.Float64()
	So(fval, ShouldEqual, 1e-11)
}

func DivValue(creator decimalCreator) {
	n := creator.NewFromFloat64(100)
	w := creator.NewFromFloat64(4)

	for i := 0; i < 9; i++ {
		n = n.Div(w)
	}

	fval := n.Float64()
	So(fval, ShouldEqual, 0.0003814697265625)

	n = creator.NewFromFloat64(float64(124215))
	w = creator.NewFromFloat64(float64(100000))
	n = n.Div(w)

	fval = n.Float64()
	So(fval, ShouldEqual, 1.24215)
}

func DivValue2(creator decimalCreator) {
	n := creator.NewFromFloat64(0)
	w := creator.NewFromFloat64(1)
	n = n.Div(w)

	fval := n.Float64()
	So(fval, ShouldEqual, 0)
}

func DivValue3(creator decimalCreator) {
	n := creator.NewFromFloat64(1)
	w := creator.NewFromFloat64(0)
	n = n.Div(w)

	// special case: 0 / 0 = 0
	fval := n.Float64()
	So(fval, ShouldEqual, 0)
}

func IsZeroValue(creator decimalCreator) {
	n := creator.NewFromFloat64(100.1)
	m := creator.NewFromFloat64(100)
	k := creator.NewFromFloat64(0.1)

	n = n.Sub(m)
	n = n.Sub(k)

	So(n.IsZero(), ShouldBeTrue)
	So(n.IsPositive(), ShouldBeFalse)

	fval := n.Float64()
	So(fval, ShouldEqual, 0)
}

func IsNegative(creator decimalCreator) {
	n := creator.NewFromFloat64(100.1)
	m := creator.NewFromFloat64(100)
	k := creator.NewFromFloat64(0.11)
	n = n.Sub(m)

	n = n.Sub(k)

	So(n.IsZero(), ShouldBeFalse)
	So(n.IsNegative(), ShouldBeTrue)
	So(n.IsPositive(), ShouldBeFalse)

	fval := n.Float64()
	So(fval, ShouldEqual, -0.01)

	n = n.Neg()
	fval = n.Float64()
	So(fval, ShouldEqual, 0.01)
}

func RoundValue(creator decimalCreator) {
	n := creator.NewFromFloat64(100.155)
	n = n.Round(1)

	fval := n.Float64()
	So(fval, ShouldEqual, 100.2)

	n = creator.NewFromFloat64(100.145)
	n = n.Round(1)
	fval = n.Float64()
	So(fval, ShouldEqual, 100.1)
}

func ModValue(creator decimalCreator) {
	n := creator.NewFromFloat64(100.155)
	m := creator.NewFromFloat64(100)
	d := n.Mod(m)

	fval := d.Float64()
	So(fval, ShouldEqual, .155)
}

func ModValueZero(creator decimalCreator) {
	// special case: Mod(1, 0) = 0
	n := creator.NewFromFloat64(1)
	m := creator.NewFromFloat64(0)

	fval := n.Mod(m).Float64()
	So(fval, ShouldBeZeroValue)
}

func PowValue(creator decimalCreator) {
	n := creator.NewFromFloat64(10)
	m := creator.NewFromFloat64(3)
	d := n.Pow(m)

	fval := d.Float64()
	So(fval, ShouldEqual, 1000)

	n = creator.NewFromFloat64(3)
	m = creator.NewFromFloat64(3)
	d = n.Pow(m)

	fval = d.Float64()
	So(fval, ShouldEqual, 27)

	n = creator.NewFromFloat64(3.3)
	m = creator.NewFromFloat64(3)
	d = n.Pow(m)

	fval = d.Float64()
	So(fval, ShouldEqual, 35.937)

	n = creator.NewFromFloat64(-0.1)
	m = creator.NewFromFloat64(-2)
	fval = n.Pow(m).Float64()
	So(fval, ShouldEqual, 100)
}

func TestSpringDecimalAddSub(t *testing.T) {
	Convey("Spring 小数相加", t, func() {
		creator = &springDecimalCreator{}
		AddValue(creator)
	})
	Convey("Spring 小数相减", t, func() {
		creator = &springDecimalCreator{}
		SubValue(creator)
	})
}

func TestTiDBDecimalAddSub(t *testing.T) {
	Convey("TiDB 小数相加", t, func() {
		creator = &tidbDecimalCreator{}
		AddValue(creator)
	})
	Convey("TiDB 小数相减", t, func() {
		creator = &tidbDecimalCreator{}
		SubValue(creator)
	})
}

func TestSpringDecimalMultiply(t *testing.T) {
	Convey("Spring 小数相乘", t, func() {
		creator = &springDecimalCreator{}
		MultiplyValue(creator)
	})
	Convey("Spring 小数相乘2", t, func() {
		creator = &springDecimalCreator{}
		MultiplyValue2(creator)
	})
}

func TestTiDBDecimalMultiply(t *testing.T) {
	Convey("TiDB 小数相乘", t, func() {
		creator = &tidbDecimalCreator{}
		MultiplyValue(creator)
	})
	Convey("TiDB 小数相乘2", t, func() {
		creator = &tidbDecimalCreator{}
		MultiplyValue2(creator)
	})
}

func TestSpringDecimalDiv(t *testing.T) {
	Convey("Spring 小数相除", t, func() {
		creator = &springDecimalCreator{}
		DivValue(creator)
	})
	Convey("Spring 小数相除2", t, func() {
		creator = &springDecimalCreator{}
		DivValue2(creator)
	})
	Convey("Spring 小数相除3", t, func() {
		creator = &springDecimalCreator{}
		So(func() { DivValue3(creator) }, ShouldPanic)
	})
}

func TestTiDBDecimalDiv(t *testing.T) {
	Convey("TiDB 小数相除", t, func() {
		creator = &tidbDecimalCreator{}
		DivValue(creator)
	})
	Convey("TiDB 小数相除2", t, func() {
		creator = &tidbDecimalCreator{}
		DivValue2(creator)
	})
	Convey("TiDB 小数相除3", t, func() {
		creator = &tidbDecimalCreator{}
		DivValue3(creator)
	})
}

func TestSpringDecimalChecks(t *testing.T) {
	Convey("Spring 数为零判断", t, func() {
		creator = &springDecimalCreator{}
		IsZeroValue(creator)
	})
	Convey("Spring 负数判断", t, func() {
		creator = &springDecimalCreator{}
		IsNegative(creator)
	})
}

func TestTiDBDecimalChecks(t *testing.T) {
	Convey("TiDB 数为零判断", t, func() {
		creator = &tidbDecimalCreator{}
		IsZeroValue(creator)
	})
	Convey("TiDB 负数判断", t, func() {
		creator = &tidbDecimalCreator{}
		IsNegative(creator)
	})
}

func TestSpringDecimalRound(t *testing.T) {
	Convey("Spring 四舍五入", t, func() {
		creator = &springDecimalCreator{}
		RoundValue(creator)
	})
}

func TestTiDBDecimalRound(t *testing.T) {
	Convey("TiDB 四舍五入", t, func() {
		creator = &tidbDecimalCreator{}
		RoundValue(creator)
	})
}

func TestSpringDecimalMod(t *testing.T) {
	Convey("Spring 取余数", t, func() {
		creator = &springDecimalCreator{}
		ModValue(creator)
		Convey(`Spring 1 取 0 余数`, func() {
			So(func() { ModValueZero(creator) }, ShouldPanic)
		})
	})
}

func TestTiDBDecimalMod(t *testing.T) {
	Convey("TiDB 取余数", t, func() {
		creator = &tidbDecimalCreator{}
		ModValue(creator)
		Convey(`TiDB 1 取 0 余数`, func() {
			ModValueZero(creator)
		})
	})
}

func TestSpringDecimalPow(t *testing.T) {
	Convey("Spring 幂", t, func() {
		creator = &springDecimalCreator{}
		PowValue(creator)
	})
}

func TestTiDBDecimalPow(t *testing.T) {
	Convey("TiDB 幂", t, func() {
		creator = &tidbDecimalCreator{}
		PowValue(creator)
	})
}

func TestTiDBDecimalIncrementOfFraction(t *testing.T) {
	Convey(`TiDB 实现精度兼容度测试`, t, func() {
		c := &tidbDecimalCreator{}
		n := 1e-30
		d := c.NewFromFloat64(n)
		So(d.Float64(), ShouldEqual, n)

		spring := &springDecimalCreator{}

		values := [][2]float64{
			{1e-81, 1e-80},
			{1e+80 + 1e-81, 1e+80 + 1e-80},
		}

		for _, v := range values {
			n := v[0]
			m := v[1]
			d = c.NewFromFloat64(n)
			j := c.NewFromFloat64(m)
			So(d.Add(j).Float64(), ShouldEqual,
				spring.NewFromFloat64(n).
					Add(spring.NewFromFloat64(m)).
					Float64(),
			)
		}
	})
}

func ValueIntPart(creator decimalCreator) {
	d := creator.NewFromFloat64(1.0 + 1e-15)
	So(d.Int64(), ShouldEqual, 1)

	d = creator.NewFromFloat64(2.0 - 1e-15)
	So(d.Int64(), ShouldEqual, 1)
}

func TestTiDBIntPart(t *testing.T) {
	Convey(`TiDB 取整数部分测试`, t, func() {
		ValueIntPart(&tidbDecimalCreator{})
	})
}

func TestSpringIntPart(t *testing.T) {
	Convey(`Spring 取整数部分测试`, t, func() {
		ValueIntPart(&tidbDecimalCreator{})
	})
}
