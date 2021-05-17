package formats

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"

	"github.com/daroay/behavioural/strategy"
)

type ImageSquare struct {
	strategy.PrintOutput
	DestinationFilePath string
}

func (t *ImageSquare) Print() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{
		R: 70, G: 70, B: 70, A: 0,
	}}

	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)
	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{
		R: 255, G: 0, B: 0, A: 1,
	}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	if t.Writer == nil {
		return fmt.Errorf("No Writer Implemented")
	}
	if err := jpeg.Encode(t.Writer, bgImage, quality); err != nil {
		return fmt.Errorf("Error writing to disk")
	}
	if t.LogWriter != nil {
		r := bytes.NewReader([]byte("Image written\n"))
		io.Copy(t.LogWriter, r)
	}
	return nil
}
