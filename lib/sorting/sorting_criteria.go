package sorting

// +gen criteria
type byRed rgbaPixels

func (p byRed) Less(i, j int) bool { return p[i][0] < p[j][0] }

// +gen criteria
type byGreen rgbaPixels

func (p byGreen) Less(i, j int) bool { return p[i][1] < p[j][1] }

// +gen criteria
type byBlue rgbaPixels

func (p byBlue) Less(i, j int) bool { return p[i][2] < p[j][2] }

// +gen criteria
type bySum rgbaPixels

func (p bySum) Less(i, j int) bool {
	return p[i].sum() < p[j].sum()
}

func (p *rgbaPixel) sum() uint32 {
	return uint32(p[0]) + uint32(p[1]) + uint32(p[2])
}

// +gen criteria
type byGrayscale rgbaPixels

func (p byGrayscale) Less(i, j int) bool {
	return p[i].gray() < p[j].gray()
}

func (p *rgbaPixel) gray() uint8 {
	return p[0]>>2 + p[1]>>1 + p[1]>>3 + p[2]>>3
}
