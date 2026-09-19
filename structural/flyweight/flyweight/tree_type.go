package flyweight

import "fmt"

type TreeType struct {
	name    string
	color   string
	texture string
}

func (t *TreeType) Draw(x, y int) string {
	return fmt.Sprintf("%s at (%d,%d) in %s", t.name, x, y, t.color)
}
