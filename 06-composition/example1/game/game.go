// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package game

type Object struct {
	Length float64
	Volume float64
	Mass   float64

	Texture string
	Color   string

	Visible bool
}

type Drawer interface {
	Draw()
}

type Changer interface {
	Change(length float64, volume float64, mass float64)
}

type Hider interface {
	Hide(b bool)
}

// *****************************************************************************

type Location struct {
	X int
	Y int
}

type Mover interface {
	Move(x int, y int)
}

// *****************************************************************************

func Draw(d Drawer) {
	d.Draw()
}

func Change(c Changer, length float64, volume float64, mass float64) {
	c.Change(length, volume, mass)
}

func Move(m Mover, x int, y int) {
	m.Move(x, y)
}

func Hide(h Hider, b bool) {
	h.Hide(b)
}
