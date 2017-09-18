package pdf

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"golang.org/x/image/bmp"
)

const suzanneBytes = 512 * 512 * 3

func loadSuzanneRGBA() (*image.RGBA, error) {
	f, err := os.Open("testdata/suzanne.bmp")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := bmp.Decode(f)
	if err != nil {
		return nil, err
	}
	return img.(*image.RGBA), nil
}

func loadSuzanneNRGBA() (*image.NRGBA, error) {
	f, err := os.Open("testdata/suzanne.png")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return img.(*image.NRGBA), nil
}

func loadSuzanneYCbCr() (*image.YCbCr, error) {
	f, err := os.Open("testdata/suzanne.jpg")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		return nil, err
	}
	return img.(*image.YCbCr), nil
}

func BenchmarkEncodeRGBAGeneric(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneRGBA()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeImageStream(ioutil.Discard, img)
	}
}

func BenchmarkEncodeRGBA(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneRGBA()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeRGBAStream(ioutil.Discard, img)
	}
}

func BenchmarkEncodeNRGBAGeneric(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneNRGBA()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeImageStream(ioutil.Discard, img)
	}
}

func BenchmarkEncodeNRGBA(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneNRGBA()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeNRGBAStream(ioutil.Discard, img)
	}
}

func BenchmarkEncodeYCbCrGeneric(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneYCbCr()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeImageStream(ioutil.Discard, img)
	}
}

func BenchmarkEncodeYCbCr(b *testing.B) {
	b.StopTimer()
	img, _ := loadSuzanneYCbCr()
	b.SetBytes(suzanneBytes)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		encodeYCbCrStream(ioutil.Discard, img)
	}
}

func expectImageBuffer(buf []byte, r image.Rectangle, c color.Color, t *testing.T) {
	nc := color.NRGBAModel.Convert(c).(color.NRGBA)
	if n := 3 * r.Dx() * r.Dy(); len(buf) != n {
		t.Errorf("stream length = %d; want %d", len(buf), n)
	}
	for i := 0; i+2 < len(buf); i += 3 {
		r, g, b := uint8(buf[i]), uint8(buf[i+1]), uint8(buf[i+2])
		if r != nc.R || g != nc.G || b != nc.B {
			t.Errorf("buf[%d:%d] = [%#02x %#02x %#02x]; want [%#02x %#02x %#02x]", i, i+3, r, g, b, nc.R, nc.G, nc.B)
		}
	}
}

func TestEncodeRGBAStream(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewRGBA(r)
	c := color.RGBA{R: 40, G: 9, B: 33, A: 85}
	draw.Draw(img, r, image.NewUniform(c), image.ZP, draw.Src)

	var buf bytes.Buffer
	encodeImageStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}

func TestEncodeRGBAStreamGeneric(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewRGBA(r)
	c := color.RGBA{R: 40, G: 9, B: 33, A: 85}
	draw.Draw(img, r, image.NewUniform(c), image.ZP, draw.Src)

	var buf bytes.Buffer
	encodeRGBAStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}

func TestEncodeNRGBAStream(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewNRGBA(r)
	c := color.NRGBA{R: 120, G: 27, B: 99, A: 85}
	draw.Draw(img, r, image.NewUniform(c), image.ZP, draw.Src)

	var buf bytes.Buffer
	encodeNRGBAStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}

func TestEncodeNRGBAStreamGeneric(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewNRGBA(r)
	c := color.NRGBA{R: 120, G: 27, B: 99, A: 85}
	draw.Draw(img, r, image.NewUniform(c), image.ZP, draw.Src)

	var buf bytes.Buffer
	encodeImageStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}

func TestEncodeYCbCrStream(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
	c := color.YCbCr{Y: 70, Cb: 146, Cr: 164}
	for i := range img.Y {
		img.Y[i] = c.Y
	}
	for i := range img.Cb {
		img.Cb[i] = c.Cb
	}
	for i := range img.Cr {
		img.Cr[i] = c.Cr
	}

	var buf bytes.Buffer
	encodeYCbCrStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}

func TestEncodeYCbCrStreamGeneric(t *testing.T) {
	r := image.Rect(0, 0, 16, 16)
	img := image.NewYCbCr(r, image.YCbCrSubsampleRatio444)
	c := color.YCbCr{Y: 70, Cb: 146, Cr: 164}
	for i := range img.Y {
		img.Y[i] = c.Y
	}
	for i := range img.Cb {
		img.Cb[i] = c.Cb
	}
	for i := range img.Cr {
		img.Cr[i] = c.Cr
	}

	var buf bytes.Buffer
	encodeImageStream(&buf, img)
	expectImageBuffer(buf.Bytes(), r, c, t)
}
