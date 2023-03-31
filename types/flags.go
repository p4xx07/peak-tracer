package types

import (
	"flag"
	"fmt"
	"path"
)

type Flags struct {
	Input   string
	Output  string
	Concat  string
	Target  int
	Samples int
	Before  float64
	After   float64
	Help    bool
}

func (f *Flags) Set() error {
	input := flag.String("i", "", "path of the input file")
	output := flag.String("o", "", "path of the output peaks file")
	concat := flag.String("concat", "", "path of the output highlight media file. If not specified no encoding will be done")
	target := flag.Int("target", 300, "duration target in seconds used for the time range analysis")
	samples := flag.Int("samples", 44100, "set number of samples per each audio frames")
	before := flag.Float64("before", 10, "seconds to consider for the highlight before the peak")
	after := flag.Float64("after", 10, "seconds to consider for the highlight after the peak")
	help := flag.Bool("help", false, "shows help")

	flag.Parse()

	f.Input = *input
	f.Output = *output
	f.Concat = *concat
	f.Target = *target
	f.Samples = *samples
	f.Before = *before
	f.After = *after
	f.Help = *help

	if f.Input == "" {
		return fmt.Errorf("need to specify input")
	}
	if f.Output == "" {
		return fmt.Errorf("need to specify output")
	}
	if f.Target <= 0 {
		return fmt.Errorf("need to specify valid target")
	}
	if f.Samples <= 0 {
		return fmt.Errorf("need to specify valid samples")
	}
	if f.Before < 0 {
		return fmt.Errorf("need to specify valid before")
	}
	if f.After < 0 {
		return fmt.Errorf("need to specify valid after")
	}
	if f.Concat != "" && (path.Ext(f.Input) != path.Ext(f.Concat)) {
		return fmt.Errorf("extension of input must match extension of concat")
	}

	return nil
}
