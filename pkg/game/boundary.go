package game

type Boundary struct {
	Start     int  `json:"start"`
	End       int  `json:"end"`
	Confirmed bool `json:"confirmed"`
}

func (b *Boundary) Set(start, end int) {
	b.Start = start
	b.End = end
}

func (b *Boundary) Mean() int {
	return (b.Start + b.End) / 2
}
