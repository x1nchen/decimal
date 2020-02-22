package decimal

import (
	"errors"
	"testing"
)

func BenchValue(creator decimalCreator) error {
	n := creator.NewFromFloat64(.001)
	t := creator.NewFromFloat64(0)

	for i := 0; i < 1000; i++ {
		t = t.Add(n)
	}

	fval := t.Float64()
	if fval != 1 {
		return errors.New("total value is not 1")
	}
	return nil
}

func BenchRoundValue(creator decimalCreator) error {
	n := creator.NewFromFloat64(1.15)
	m := n.Round(1)

	fval := m.Float64()
	if fval != 1.2 {
		return errors.New("total value is not 1.2")
	}
	return nil
}

func BenchmarkSpringDecimalAdd(b *testing.B) {
	creator = &springDecimalCreator{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := BenchValue(creator); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTiDBDecimalAdd(b *testing.B) {
	creator = &tidbDecimalCreator{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := BenchValue(creator); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSpringDecimalRound(b *testing.B) {
	creator = &springDecimalCreator{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := BenchRoundValue(creator); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkTiDBDecimalRound(b *testing.B) {
	creator = &tidbDecimalCreator{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := BenchRoundValue(creator); err != nil {
			b.Fatal(err)
		}
	}
}
