package draw2dbase

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"testing"
)

var (
	flatteningThreshold = 0.5
	testsCubicFloat64   = []float64{
		100, 100, 200, 100, 100, 200, 200, 200,
		100, 100, 300, 200, 200, 200, 300, 100,
		100, 100, 0, 300, 200, 0, 300, 300,
		150, 290, 10, 10, 290, 10, 150, 290,
		10, 290, 10, 10, 290, 10, 290, 290,
		100, 290, 290, 10, 10, 10, 200, 290,
	}
	testsQuadFloat64 = []float64{
		100, 100, 200, 100, 200, 200,
		100, 100, 290, 200, 290, 100,
		100, 100, 0, 290, 200, 290,
		150, 290, 10, 10, 290, 290,
		10, 290, 10, 10, 290, 290,
		100, 290, 290, 10, 120, 290,
	}
)

func init() {
	os.Mkdir("test_results", 0666)
	f, err := os.Create("../output/curve/_test.html")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	log.Printf("Create html viewer")
	f.Write([]byte("<html><body>"))
	for i := 0; i < len(testsCubicFloat64)/8; i++ {
		f.Write([]byte(fmt.Sprintf("<div><img src='_test%d.png'/></div>\n", i)))
	}
	for i := 0; i < len(testsQuadFloat64); i++ {
		f.Write([]byte(fmt.Sprintf("<div><img src='_testQuad%d.png'/>\n</div>\n", i)))
	}
	f.Write([]byte("</body></html>"))

}

func drawPoints(img draw.Image, c color.Color, s ...float64) image.Image {
	for i := 0; i < len(s); i += 2 {
		x, y := int(s[i]+0.5), int(s[i+1]+0.5)
		img.Set(x, y, c)
		img.Set(x, y+1, c)
		img.Set(x, y-1, c)
		img.Set(x+1, y, c)
		img.Set(x+1, y+1, c)
		img.Set(x+1, y-1, c)
		img.Set(x-1, y, c)
		img.Set(x-1, y+1, c)
		img.Set(x-1, y-1, c)

	}
	return img
}

func TestCubicCurve(t *testing.T) {
	for i := 0; i < len(testsCubicFloat64); i += 8 {
		var p SegmentedPath
		p.MoveTo(testsCubicFloat64[i], testsCubicFloat64[i+1])
		TraceCubic(&p, testsCubicFloat64[i:], flatteningThreshold)
		img := image.NewNRGBA(image.Rect(0, 0, 300, 300))
		PolylineBresenham(img, color.NRGBA{0xff, 0, 0, 0xff}, testsCubicFloat64[i:i+8]...)
		PolylineBresenham(img, image.Black, p.Points...)
		//drawPoints(img, image.NRGBAColor{0, 0, 0, 0xff}, curve[:]...)
		drawPoints(img, color.NRGBA{0, 0, 0, 0xff}, p.Points...)
		SaveToPngFile(fmt.Sprintf("../output/curve/_test%d.png", i/8), img)
		log.Printf("Num of points: %d\n", len(p.Points))
	}
	fmt.Println()
}

func TestQuadCurve(t *testing.T) {
	for i := 0; i < len(testsQuadFloat64); i += 6 {
		var p SegmentedPath
		p.MoveTo(testsQuadFloat64[i], testsQuadFloat64[i+1])
		TraceQuad(&p, testsQuadFloat64[i:], flatteningThreshold)
		img := image.NewNRGBA(image.Rect(0, 0, 300, 300))
		PolylineBresenham(img, color.NRGBA{0xff, 0, 0, 0xff}, testsQuadFloat64[i:i+6]...)
		PolylineBresenham(img, image.Black, p.Points...)
		//drawPoints(img, image.NRGBAColor{0, 0, 0, 0xff}, curve[:]...)
		drawPoints(img, color.NRGBA{0, 0, 0, 0xff}, p.Points...)
		SaveToPngFile(fmt.Sprintf("../output/curve/_testQuad%d.png", i), img)
		log.Printf("Num of points: %d\n", len(p.Points))
	}
	fmt.Println()
}

func BenchmarkCubicCurve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < len(testsCubicFloat64); i += 8 {
			var p SegmentedPath
			p.MoveTo(testsCubicFloat64[i], testsCubicFloat64[i+1])
			TraceCubic(&p, testsCubicFloat64[i:], flatteningThreshold)
		}
	}
}

// SaveToPngFile create and save an image to a file using PNG format
func SaveToPngFile(filePath string, m image.Image) error {
	// Create the file
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	// Create Writer from file
	b := bufio.NewWriter(f)
	// Write the image into the buffer
	err = png.Encode(b, m)
	if err != nil {
		return err
	}
	err = b.Flush()
	if err != nil {
		return err
	}
	return nil
}
