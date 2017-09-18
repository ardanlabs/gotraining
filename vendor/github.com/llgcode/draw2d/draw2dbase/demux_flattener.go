package draw2dbase

type DemuxFlattener struct {
	Flatteners []Flattener
}

func (dc DemuxFlattener) MoveTo(x, y float64) {
	for _, flattener := range dc.Flatteners {
		flattener.MoveTo(x, y)
	}
}

func (dc DemuxFlattener) LineTo(x, y float64) {
	for _, flattener := range dc.Flatteners {
		flattener.LineTo(x, y)
	}
}

func (dc DemuxFlattener) LineJoin() {
	for _, flattener := range dc.Flatteners {
		flattener.LineJoin()
	}
}

func (dc DemuxFlattener) Close() {
	for _, flattener := range dc.Flatteners {
		flattener.Close()
	}
}

func (dc DemuxFlattener) End() {
	for _, flattener := range dc.Flatteners {
		flattener.End()
	}
}
