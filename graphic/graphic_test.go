package graphic

import (
	. "launchpad.net/gocheck"
	"testing"
)

const (
	originTestImage         = "testdata/origin_1280x800.jpg"
	originImgWidth          = 1280
	originImgHeight         = 800
	originImgDominantColorH = 205
	originImgDominantColorS = 0.69
	originImgDominantColorV = 0.42
)

func Test(t *testing.T) { TestingT(t) }

type Graphic struct{}

var _ = Suite(&Graphic{})

func delta(x, y float64) float64 {
	if x >= y {
		return x - y
	}
	return y - x
}

func (g *Graphic) TestClipImage(c *C) {
	resultFile := "testdata/test_clipimage_100x200.png"
	err := ClipImage(originTestImage, resultFile, 0, 0, 100, 200, PNG)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 100)
	c.Check(int(h), Equals, 200)
}

func (g *Graphic) TestConvertImage(c *C) {
	resultFile := "testdata/test_convertimage.png"
	err := ConvertImage(originTestImage, resultFile, PNG)
	if err != nil {
		c.Error(err)
	}
}

func (g *Graphic) TestGetDominantColor(c *C) {
	h, s, v, err := GetDominantColorOfImage(originTestImage)
	if err != nil {
		c.Error(err)
	}
	if delta(h, originImgDominantColorH) > 1 ||
		delta(s, originImgDominantColorS) > 0.1 ||
		delta(v, originImgDominantColorV) > 0.1 {
		c.Error("h, s, v = ", h, s, v)
	}
}

// Test that a subset of RGB space can be converted to HSV and back to within
// 1/256 tolerance.
func (g *Graphic) TestHSV(c *C) {
	for r := 0; r < 255; r += 7 {
		for g := 0; g < 255; g += 5 {
			for b := 0; b < 255; b += 3 {
				r0, g0, b0 := uint8(r), uint8(g), uint8(b)
				h, s, v := RGB2HSV(r0, g0, b0)
				r1, g1, b1 := HSV2RGB(h, s, v)
				if delta(float64(r0), float64(r1)) > 1 || delta(float64(g0), float64(g1)) > 1 || delta(float64(b0), float64(b1)) > 1 {
					c.Fatalf("r0, g0, b0 = %d, %d, %d   r1, g1, b1 = %d, %d, %d", r0, g0, b0, r1, g1, b1)
				}
			}
		}
	}
}

func (g *Graphic) TestGetImageSize(c *C) {
	w, h, err := GetImageSize(originTestImage)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, originImgWidth)
	c.Check(int(h), Equals, originImgHeight)
}

func (g *Graphic) TestResizeImage(c *C) {
	resultFile := "testdata/test_resizeimage_500x600.png"
	err := ResizeImage(originTestImage, resultFile, 500, 600, PNG)
	if err != nil {
		c.Error(err)
	}
	w, h, err := GetImageSize(resultFile)
	if err != nil {
		c.Error(err)
	}
	c.Check(int(w), Equals, 500)
	c.Check(int(h), Equals, 600)
}

func (g *Graphic) TestRotateImageLeft(c *C) {
	resultFile := "testdata/test_rotateimageleft.png"
	err := RotateImageLeft(originTestImage, resultFile, PNG)
	if err != nil {
		c.Error(err)
	}
}

func (g *Graphic) TestRotateImageRight(c *C) {
	resultFile := "testdata/test_rotateimageright.png"
	err := RotateImageRight(originTestImage, resultFile, PNG)
	if err != nil {
		c.Error(err)
	}
}
