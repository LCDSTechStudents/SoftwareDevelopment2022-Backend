package _default

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
)

func ToBinary(path string) interface{} {
	nraw := &gocv.Mat{}
	raw := gocv.IMRead(path, gocv.IMReadGrayScale)
	if raw.Empty() {
		fmt.Printf("Invalid read of Source Mat in test\n")
		return nil
	}
	gocv.CvtColor(raw, nraw, gocv.ColorBGRToRGB)
	gocv.Resize(raw, nraw, image.Point{
		X: 416,
		Y: 416,
	}, 0, 0, gocv.InterpolationDefault)
	nraw.DivideFloat(255)
	return nraw
}
