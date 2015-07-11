// All material is licensed under the GNU Free Documentation License
// https://github.com/ArdanStudios/gotraining/blob/master/LICENSE

package objects

import "fmt"

// location provides a point in space for an object.
type location struct {
	x int
	y int
}

// Set sets the location of the object in space.
func (l *location) Set(x, y int) {
	l.x = x
	l.y = y
}

// Paint display the location of the object.
func (l *location) Paint() string {
	return fmt.Sprintf("(%d, %d)\n", l.x, l.y)
}

// movable provides the attribute of a location that can
// be moved in space.
type movable struct {
	location
}

// Move provides a basic behavior of moving through space.
func (m *movable) Move(x, y int) {
	m.x += x
	m.y += y
}

// solid provides the attribute of an object that can't
// be passed through.
type solid struct {
	shape string
	color string
}

// Draw provides the behavior of creating objects.
func (s *solid) Draw(shape, color string) {
	s.shape = shape
	s.color = color
}

// Paint display the attribute of the solid object.
func (s *solid) Paint() string {
	return fmt.Sprintf("{Shape: %s, Color: %s}\n", s.shape, s.color)
}
