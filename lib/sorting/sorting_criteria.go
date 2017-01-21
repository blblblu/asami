package sorting

type byRed rgbaPixels

func (p byRed) Len() int           { return len(p) }
func (p byRed) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byRed) Less(i, j int) bool { return p[i][0] < p[j][0] }

type byGreen rgbaPixels

func (p byGreen) Len() int           { return len(p) }
func (p byGreen) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byGreen) Less(i, j int) bool { return p[i][1] < p[j][1] }

type byBlue rgbaPixels

func (p byBlue) Len() int           { return len(p) }
func (p byBlue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p byBlue) Less(i, j int) bool { return p[i][2] < p[j][2] }

type bySum rgbaPixels

func (p bySum) Len() int      { return len(p) }
func (p bySum) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p bySum) Less(i, j int) bool {
	return p[i].sum() < p[j].sum()
}

func (p *rgbaPixel) sum() uint32 {
	return uint32(p[0]) + uint32(p[1]) + uint32(p[2])
}

type byGrayscale rgbaPixels

func (p byGrayscale) Len() int      { return len(p) }
func (p byGrayscale) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p byGrayscale) Less(i, j int) bool {
	return p[i].gray() < p[j].gray()
}

func (p *rgbaPixel) gray() uint8 {
	return p[0]>>2 + p[1]>>1 + p[1]>>3 + p[2]>>3
}
