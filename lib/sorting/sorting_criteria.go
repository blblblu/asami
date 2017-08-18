package sorting

//type byRed rgbaPixels
//func (p byRed) Less(i, j int) bool { return p[i*4] < p[j*4] }
//type byGreen rgbaPixels
//func (p byGreen) Less(i, j int) bool { return p[i*4+1] < p[j*4+1] }
//type byBlue rgbaPixels
//func (p byBlue) Less(i, j int) bool { return p[i*4+2] < p[j*4+2] }

type bySum rgbaPixels

func (p bySum) Len() int { return len(p) / 4 }
func (p bySum) Swap(i, j int) {
	for k := 0; k < 4; k++ {
		p[i*4+k], p[j*4+k] = p[j*4+k], p[i*4+k]
	}
}

func (p bySum) Less(i, j int) bool {
	return p[i*4]+p[i*4+1]+p[i*4+2]+p[i*4+3] < p[j*4]+p[j*4+1]+p[j*4+2]+p[j*4+3]
}

type byGrayscale rgbaPixels

func (p byGrayscale) Len() int { return len(p) / 4 }
func (p byGrayscale) Swap(i, j int) {
	for k := 0; k < 4; k++ {
		p[i*4+k], p[j*4+k] = p[j*4+k], p[i*4+k]
	}
}
func (p byGrayscale) Less(i, j int) bool {
	return p[i*4]>>2+p[i*4+1]>>1+p[i*4]>>3+p[i*4+2]>>3 < p[j*4]>>2+p[j*4+1]>>1+p[j*4]>>3+p[j*4+2]>>3
}
