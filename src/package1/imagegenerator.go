package package1

import (
	"image/color"

	"github.com/fogleman/gg"
)

// TreeMap represents the treemap structure.
type TreeMap struct {
	Root Node // Root node of the treemap
}

// Node represents a rectangle in the treemap.
type Node struct {
	X, Y, Width, Height float64     // Position and dimensions of the rectangle
	Color               color.RGBA // Color of the rectangle
	Children            []Node      // Children of the current rectangle
}


// GenerateTreeMap generates a treemap with the specified width and height.
func GenerateTreeMap(width, height int, rectangles []Node) TreeMap {
	// Generate some example rectangles for demonstration

	return TreeMap{
		Root: Node{
			X:        0,
			Y:        0,
			Width:    float64(width),
			Height:   float64(height),
			Children: rectangles,
		},
	}
}


// RenderTreeMap renders the treemap onto the image context.
func RenderTreeMap(dc *gg.Context, treeMap TreeMap) {
	// Render the root node and its children recursively
	renderNode(dc, treeMap.Root)
}

// renderNode renders the current node and its children recursively.
func renderNode(dc *gg.Context, node Node) {
	dc.SetRGB(float64(node.Color.R)/255, float64(node.Color.G)/255, float64(node.Color.B)/255)
	dc.DrawRectangle(node.X, node.Y, node.Width, node.Height)
	dc.Fill()

	// Render the children recursively
	for _, child := range node.Children {
		renderNode(dc, child)
	}
}
