// Copyright (C) 2011, Ross Light

/*
	Package pdf implements a Portable Document Format writer, as defined in ISO 32000-1.

	An example of basic usage:

		package main

		import (
			"bitbucket.org/zombiezen/gopdf/pdf"
			"fmt"
			"os"
		)

		func main() {
			doc := pdf.New()
			canvas := doc.NewPage(pdf.USLetterWidth, pdf.USLetterHeight)
			canvas.Translate(100, 100)

			path := new(pdf.Path)
			path.Move(pdf.Point{0, 0})
			path.Line(pdf.Point{100, 0})
			canvas.Stroke(path)

			text := new(pdf.Text)
			text.SetFont(pdf.Helvetica, 14)
			text.Text("Hello, World!")
			canvas.DrawText(text)

			canvas.Close()

			err := doc.Encode(os.Stdout)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}
*/
package pdf
