package main

import "testing"

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := rec.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf(" got %.2f want %.2f", got, want)

	}
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rec := &Rectangle{10.0, 10.0}
		checkArea(t, rec, 100.0)
	})

	t.Run("circles", func(t *testing.T) {
		cir := &Circle{10}
		checkArea(t, cir, 314.1592653589793)
	})
}
