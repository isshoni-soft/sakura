package sakura

import "time"

type DeltaTicker struct {
	Function func(game Game, ticker DeltaTicker)
	previous *time.Time
	current  *time.Time
}

func (t *DeltaTicker) Delta() uint64 {
	if t.previous == nil {
		p := time.Now()

		t.previous = &p
		
		return 0
	} else {
		t.previous = t.current
	}

	c := time.Now()

	t.current = &c

	return uint64(c.UnixMilli() - t.previous.UnixMilli())
}
