// Copyright (C) 2011, Ross Light

package pdf

import (
	"image"
	"image/color"
	"io"
)

const (
	deviceRGBColorSpace name = "DeviceRGB"
)

type imageStream struct {
	*stream
	Width            int
	Height           int
	BitsPerComponent int
	ColorSpace       name
}

type imageStreamInfo struct {
	Type             name
	Subtype          name
	Length           int
	Filter           name `pdf:",omitempty"`
	Width            int
	Height           int
	BitsPerComponent int
	ColorSpace       name
}

func newImageStream(filter name, w, h int) *imageStream {
	return &imageStream{
		stream:           newStream(filter),
		Width:            w,
		Height:           h,
		BitsPerComponent: 8,
		ColorSpace:       deviceRGBColorSpace,
	}
}

func (st *imageStream) marshalPDF(dst []byte) ([]byte, error) {
	return marshalStream(dst, imageStreamInfo{
		Type:             xobjectType,
		Subtype:          imageSubtype,
		Length:           st.Len(),
		Filter:           st.filter,
		Width:            st.Width,
		Height:           st.Height,
		BitsPerComponent: st.BitsPerComponent,
		ColorSpace:       st.ColorSpace,
	}, st.Bytes())
}

// encodeImageStream writes RGB data from an image in PDF format.
func encodeImageStream(w io.Writer, img image.Image) error {
	bd := img.Bounds()
	row := make([]byte, bd.Dx()*3)
	for y := bd.Min.Y; y < bd.Max.Y; y++ {
		i := 0
		for x := bd.Min.X; x < bd.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			if a != 0 {
				row[i+0] = uint8((r * 65535 / a) >> 8)
				row[i+1] = uint8((g * 65535 / a) >> 8)
				row[i+2] = uint8((b * 65535 / a) >> 8)
			} else {
				row[i+0] = 0
				row[i+1] = 0
				row[i+2] = 0
			}
			i += 3
		}
		if _, err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func encodeRGBAStream(w io.Writer, img *image.RGBA) error {
	buf := make([]byte, 3*img.Rect.Dx()*img.Rect.Dy())
	var a uint16
	for i, j := 0, 0; i < len(img.Pix); i, j = i+4, j+3 {
		a = uint16(img.Pix[i+3])
		if a != 0 {
			buf[j+0] = byte(uint16(img.Pix[i+0]) * 0xff / a)
			buf[j+1] = byte(uint16(img.Pix[i+1]) * 0xff / a)
			buf[j+2] = byte(uint16(img.Pix[i+2]) * 0xff / a)
		}
	}
	_, err := w.Write(buf)
	return err
}

func encodeNRGBAStream(w io.Writer, img *image.NRGBA) error {
	buf := make([]byte, 3*img.Rect.Dx()*img.Rect.Dy())
	for i, j := 0, 0; i < len(img.Pix); i, j = i+4, j+3 {
		buf[j+0] = img.Pix[i+0]
		buf[j+1] = img.Pix[i+1]
		buf[j+2] = img.Pix[i+2]
	}
	_, err := w.Write(buf)
	return err
}

func encodeYCbCrStream(w io.Writer, img *image.YCbCr) error {
	var yy, cb, cr uint8
	var i, j int
	dx, dy := img.Rect.Dx(), img.Rect.Dy()
	buf := make([]byte, 3*dx*dy)
	bi := 0
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			i, j = x, y
			switch img.SubsampleRatio {
			case image.YCbCrSubsampleRatio420:
				j /= 2
				fallthrough
			case image.YCbCrSubsampleRatio422:
				i /= 2
			}
			yy = img.Y[y*img.YStride+x]
			cb = img.Cb[j*img.CStride+i]
			cr = img.Cr[j*img.CStride+i]

			buf[bi+0], buf[bi+1], buf[bi+2] = color.YCbCrToRGB(yy, cb, cr)
			bi += 3
		}
	}
	_, err := w.Write(buf)
	return err
}
