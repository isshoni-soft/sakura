package sakura

import "time"

type DeltaTicker struct {
	Function func(game *Game, ticker *DeltaTicker)
	Defer    func()

	previous *time.Time
	current  *time.Time
	delta    uint64
}

func (t *DeltaTicker) Run(game *Game) {
	go func() {
		defer t.Defer()

		for Running() {
			t.Function(game, t)
		}
	}()
}

func (t *DeltaTicker) GetDelta() uint64 {
	return t.delta
}

func (t *DeltaTicker) RecalculateDelta() uint64 {
	if t.previous == nil {
		p := time.Now()

		t.previous = &p
	} else {
		t.previous = t.current
	}

	c := time.Now()

	t.current = &c
	t.delta = uint64(c.UnixMilli() - t.previous.UnixMilli())

	return t.delta
}
