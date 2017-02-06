package sorting

// +gen criteria
type byRed rgbaPixels

func (p byRed) Less(i, j int) bool { return p[i*4] < p[j*4] }

// +gen criteria
type byGreen rgbaPixels

func (p byGreen) Less(i, j int) bool { return p[i*4+1] < p[j*4+1] }

// +gen criteria
type byBlue rgbaPixels

func (p byBlue) Less(i, j int) bool { return p[i*4+2] < p[j*4+2] }

// +gen criteria
type bySum rgbaPixels

func (p bySum) Less(i, j int) bool {
	return p[i*4]+p[i*4+1]+p[i*4+2]+p[i*4+3] < p[j*4]+p[j*4+1]+p[j*4+2]+p[j*4+3]
}

// +gen criteria
type byGrayscale rgbaPixels

func (p byGrayscale) Less(i, j int) bool {
	return p[i*4]>>2+p[i*4+1]>>1+p[i*4]>>3+p[i*4+2]>>3 < p[j*4]>>2+p[j*4+1]>>1+p[j*4]>>3+p[j*4+2]>>3
}
