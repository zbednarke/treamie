package main

import (
	"image/color"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/fogleman/gg"

	"treamie/src/package1"
)

const ROOT_WIDTH = 1200
const ROOT_HEIGHT = 1200

func main() {

	rectangles := gen_all_rectangles()

	// Generate the treemap using the GenerateTreeMap function from the package1 package
	treeMap := package1.GenerateTreeMap(ROOT_WIDTH, ROOT_HEIGHT, rectangles)

	dc := gg.NewContext(ROOT_WIDTH, ROOT_HEIGHT) // Create a new image surface
	dc.SetRGB(1, 1, 1)                           // Set the background color to white
	dc.Clear()                                   // Clear the canvas with the background color
	package1.RenderTreeMap(dc, treeMap)          // Render the treemap to the image

	// Save the image as PNG
	fname := "image.png"
	if err := os.Remove(fname); err != nil {
		log.Println(err)
	}
	err := dc.SavePNG("image.png")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Image generated and saved as image.png")

	// Open the generated image using the default image viewer
	err = openImage("image.png")
	if err != nil {
		log.Fatal(err)
	}
}

func gen_all_rectangles() []package1.Node {
	sub_width := float64(ROOT_WIDTH) * .5
	sub_height := float64(ROOT_HEIGHT) * .5
	rectangles := []package1.Node{
		{X: 0, Y: 0, Width: float64(ROOT_WIDTH) * 0.5, Height: float64(ROOT_HEIGHT) * 0.5, Color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, Children: gen_rectangles_to_depth(0, 0, sub_width, sub_height, 6)},
		{X: float64(ROOT_WIDTH) * 0.5, Y: 0, Width: float64(ROOT_WIDTH) * 0.5, Height: float64(ROOT_HEIGHT) * 0.5, Color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, Children: gen_rectangles_to_depth(float64(ROOT_WIDTH)*0.5, 0, sub_width, sub_height, 3)},   // Green
		{X: 0, Y: float64(ROOT_HEIGHT) * 0.5, Width: float64(ROOT_WIDTH) * 0.5, Height: float64(ROOT_HEIGHT) * 0.5, Color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, Children: gen_rectangles_to_depth(0, float64(ROOT_HEIGHT)*0.5, sub_width, sub_height, 1)}, // Blue
		{X: float64(ROOT_WIDTH) * 0.5, Y: float64(ROOT_HEIGHT) * 0.5, Width: float64(ROOT_WIDTH) * 0.5, Height: float64(ROOT_HEIGHT) * 0.5, Color: color.RGBA{R: 255, G: 255, B: 0, A: 255}},                                                                 // Children: gen_4_rectangles(float64(ROOT_WIDTH)*0.5, float64(ROOT_HEIGHT)*0.5, sub_width, sub_height)}, // Yellow
	}

	return rectangles
}

func gen_rectangles_to_depth(x0 float64, y0 float64, width float64, height float64, remaining_depth int64) []package1.Node {
	var quad_1_children []package1.Node
	var quad_2_children []package1.Node
	var quad_3_children []package1.Node
	var quad_4_children []package1.Node
	if remaining_depth > 0 {
		next_width := float64(width) * 0.5
		next_height := float64(height) * 0.5
		quad_1_children = gen_rectangles_to_depth(x0+0, y0+0, next_width, next_height, remaining_depth-1)
		quad_2_children = gen_rectangles_to_depth(x0+float64(next_width), y0+0, next_width, next_height, remaining_depth-1)
		quad_3_children = gen_rectangles_to_depth(x0+0, y0+float64(next_height), next_width, next_height, remaining_depth-1)
		quad_4_children = gen_rectangles_to_depth(x0+float64(next_width), y0+float64(next_height), next_width, next_height, remaining_depth-1)
	}
	rectangles := []package1.Node{
		{X: x0 + 0, Y: y0 + 0, Width: float64(width) * 0.5, Height: float64(height) * 0.5, Color: color.RGBA{R: 255, G: 0, B: 0, A: 255}, Children: quad_1_children},                                      // Red
		{X: x0 + float64(width)*0.5, Y: y0 + 0, Width: float64(width) * 0.5, Height: float64(height) * 0.5, Color: color.RGBA{R: 0, G: 255, B: 0, A: 255}, Children: quad_2_children},                     // Green
		{X: x0 + 0, Y: y0 + float64(height)*0.5, Width: float64(width) * 0.5, Height: float64(height) * 0.5, Color: color.RGBA{R: 0, G: 0, B: 255, A: 255}, Children: quad_3_children},                    // Blue
		{X: x0 + float64(width)*0.5, Y: y0 + float64(height)*0.5, Width: float64(width) * 0.5, Height: float64(height) * 0.5, Color: color.RGBA{R: 255, G: 255, B: 0, A: 255}, Children: quad_4_children}, // Yellow
	}

	return rectangles
}

// openImage opens the specified file using the default image viewer for the operating system.
func openImage(filename string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filename)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", filename)
	case "linux":
		cmd = exec.Command("xdg-open", filename)
	default:
		return nil // Unsupported operating system
	}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}
