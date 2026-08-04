package main

import (
	"fmt"
	"reflect"
	"runtime"

	"github.com/kevinronu/dessign-patterns-go/structural/flyweight/extrinsic"
	"github.com/kevinronu/dessign-patterns-go/structural/flyweight/flyweight"
)

// fatTree is a tree without the pattern: every one carries its own copy of the state that never
// changes.
type fatTree struct {
	x       int
	y       int
	name    string
	color   string
	texture string
}

// describe shows what a fatTree holds: the same three values that every oak repeats.
func (t fatTree) describe() string {
	return fmt.Sprintf("%s/%s/%s at (%d,%d)", t.name, t.color, t.texture, t.x, t.y)
}

func main() {
	const trees = 1_000_000

	factory := flyweight.NewFactory()

	oak := factory.Get("oak", "green", "rough")
	maple := factory.Get("maple", "orange", "smooth")

	// The same key always gives back the same pointer, so every oak in the forest shares one
	// TreeType.
	fmt.Printf("same pointer for the same key: %t\n\n", oak == factory.Get("oak", "green", "rough"))

	forest := make([]extrinsic.Tree, 0, trees)

	for i := range trees {
		kind := oak
		if i%2 == 1 {
			kind = maple
		}

		forest = append(forest, extrinsic.Tree{X: i % 500, Y: i / 500, Type: kind})
	}

	// Only the first few, because a million lines is not a demo.
	for _, tree := range forest[:3] {
		fmt.Println(tree.Draw())
	}

	fmt.Println("...")

	// A string header is two words, a pointer to the bytes and a length. The three of them are what
	// make a fatTree 40 B heavier than a Tree.
	fat := fatTree{x: 0, y: 0, name: "oak", color: "green", texture: "rough"}

	fmt.Printf("\nthe first tree without the pattern: %s\n", fat.describe())

	const megabyte = 1024 * 1024

	small := int(reflect.TypeFor[extrinsic.Tree]().Size())
	big := int(reflect.TypeFor[fatTree]().Size())

	fmt.Printf("\n%d trees share %d types\n\n", len(forest), factory.Size())
	fmt.Printf("  Tree     %2d B  =  2 ints + 1 pointer\n", small)
	fmt.Printf("  fatTree  %2d B  =  2 ints + 3 string headers\n\n", big)
	fmt.Printf("  %d MB of trees, instead of %d MB with one copy per tree\n",
		len(forest)*small/megabyte, len(forest)*big/megabyte)

	var stats runtime.MemStats

	runtime.ReadMemStats(&stats)

	fmt.Printf("  %d MB on the heap, which is those trees plus what the runtime holds\n",
		stats.HeapAlloc/megabyte)
}
