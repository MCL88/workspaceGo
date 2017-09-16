type Door struct {
	opened bool
}

type OpenCloser interface {
	Open()
	Close()
}

func (d *Door) Open() {
	d.opened = true
}

func (d *Door) Close() {
	d.opened = false
}