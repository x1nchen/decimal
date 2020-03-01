package decimal

import (
	sdecimal "github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestShopringDecimalCmp(t *testing.T) {
	Convey(`shopspring cmp`, t, func() {
		zero := springDecimal{sdecimal.Zero}

		Convey(`greater`, func() {
			s := springDecimal{sdecimal.NewFromFloat(1.0)}
			So(s.Cmp(zero), ShouldBeGreaterThan, 0)
		})

		Convey(`less`, func() {
			s := springDecimal{sdecimal.NewFromFloat(-1.0)}
			So(s.Cmp(zero), ShouldBeLessThan, 0)
		})

		Convey(`equal`, func() {
			s := springDecimal{sdecimal.NewFromFloat(0)}
			So(s.Cmp(zero), ShouldEqual, 0)
		})
	})
}