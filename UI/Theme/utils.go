package theme

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"image/color"
)

//TODO: convet xIndex and yIndex to 1 index
func corners(imd *imdraw.IMDraw, index int, cornerSize float64, d dimensions){
	x1, y1 := d.getTilePos(index, 0.0, 0.0)
	x2, y2 := d.getTilePos(index, 1.0, 1.0)

	//Bottom left corner
	imd.Push(pixel.V(x1+cornerSize,y1), pixel.V(x1,y1), pixel.V(x1,y1+cornerSize))
	imd.Polygon(0)

	//Bottom right corner
	imd.Push(pixel.V(x2-cornerSize,y1), pixel.V(x2,y1), pixel.V(x2,y1+cornerSize))
	imd.Polygon(0)

	//Top left corner
	imd.Push(pixel.V(x1+cornerSize,y2), pixel.V(x1,y2), pixel.V(x1,y2-cornerSize))
	imd.Polygon(0)

	//Top right corner
	imd.Push(pixel.V(x2-cornerSize,y2), pixel.V(x2,y2), pixel.V(x2,y2-cornerSize))
	imd.Polygon(0)
}

type dimensions struct {
	Height float64
	Border float64
	Gaps float64
}

func (d *dimensions) getTileSpace() float64 {
	return (float64(d.Height) - (2*d.Border) + d.Gaps)/ 8.0
}

func (d *dimensions) getTileSize() float64{
	return (d.getTileSpace() - d.Gaps)
}

func (d *dimensions) getTilePosBL(index int) (float64, float64) {
	return d.getTilePos(index,0.0,0.0)
	
}

func (d *dimensions) getTilePosCenter(index int) (float64, float64) {
	return d.getTilePos(index,0.5,0.5)
}

func (d *dimensions) getTilePos(index int, spanX, spanY float64) (float64, float64) {
	x, y := index%8, index/8
	return (float64(x)*d.getTileSpace()) + (spanX * d.getTileSize()) + d.Border, (float64(7-y)*d.getTileSpace()) + (spanY * d.getTileSize()) + d.Border

}

func (d *dimensions) getMouseData(win *pixelgl.Window) (bool, bool, int) {
	mPos := win.MousePosition()
	if mPos.X > d.Height - d.Border - d.Gaps { return false, false, -1 }
	if mPos.Y > d.Height - d.Border - d.Gaps { return false, false, -1 }

	tileX := int((mPos.X - d.Border) / d.getTileSpace())
	tileY := 7 - int((mPos.Y - d.Border) / d.getTileSpace())
	
	return win.JustPressed(pixelgl.MouseButtonLeft), win.JustReleased(pixelgl.MouseButtonLeft), (tileY * 8) + tileX
}


func hue2rgb(p, q, t float64) float64 {
	if t < 0 { t += 1 }
	if t > 1 { t -= 1 }
	if t < 0.16666 { return p + (q - p) * 6 * t }
	if t < 0.5 { return q }
	if t < 0.66666 { return p + (q - p) * (2/3 - t) * 6 }
	return p;
}

func HSLToRGB(h, s, l float64) color.Color {
	var r, g, b float64
	//if(s == 0){
	//	r = l
	//	g = l
	//	b = l
		 // achromatic
	//}else{
	var q float64
	if l < 0.5 {
		q = l * (1+s)
	} else {
		q = l + s - (l * s)
	}
	p := (2 * l) - q
	r = hue2rgb(p, q, h + 0.333)
	g = hue2rgb(p, q, h)
	b = hue2rgb(p, q, h - 0.333)

	return color.RGBA { uint8(r * 255), uint8(g * 255), uint8(b * 255), 0xFF }
}