package types

type Peak struct {
	RMS  float64
	Time float64
}

type Range struct {
	Min float64
	Max float64
}

type PeaksByRMS []Peak

func (p PeaksByRMS) Len() int           { return len(p) }
func (p PeaksByRMS) Less(i, j int) bool { return p[i].RMS > p[j].RMS }
func (p PeaksByRMS) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
