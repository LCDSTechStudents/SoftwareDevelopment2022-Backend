package _default

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

func ToBinary(path string) interface{} {
	raw := gocv.IMRead(path, gocv.IMReadUnchanged)
	if raw.Empty() {
		fmt.Printf("Invalid read of Source Mat in test\n")
		return nil
	}
	gocv.CvtColor(raw, &raw, gocv.ColorBGRToRGB)
	gocv.Resize(raw, &raw, image.Point{
		X: 416,
		Y: 416,
	}, 0, 0, gocv.InterpolationDefault)
	raw.DivideFloat(255)
	return raw.ToBytes()
}
