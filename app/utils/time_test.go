package utils

import (
	"testing"
	"time"
)

func TestCalDateGap(t *testing.T) {
	dates := CalDateGap(1644217701, 1644390504)
	if len(dates) != 3 {
		t.Error(`len(CalDateGap(1644217701, 1644390504)) != 3`)
	}
}

func TestIsSameDay(t *testing.T) {
	b := IsSameDay(1644199701, 1644217701)
	if !b {
		t.Error("Is Same Day!")
	}

	b = IsSameDay(1644199701, 1644145701)
	if b {
		t.Error("Is not same day")
	}
}

func TestCalTimeDiff(t *testing.T) {
	CalTimeDiff(time.Date(2022, 4, 25, 9, 0, 0, 0, time.Local).Unix(),
		time.Date(2022, 4, 27, 9, 0, 0, 0, time.Local).Unix())
}
