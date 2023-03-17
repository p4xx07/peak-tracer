package types

import (
	"flag"
	"fmt"
	"path"
)

type Flags struct {
	Input   *string
	Output  *string
	Concat  *string
	Target  int
	Samples int
	Before  float64
	After   float64
	Help    bool
}

func (f *Flags) Set() error {
	f.Input = flag.String("i", "", "path of the input file")
	f.Output = flag.String("o", "", "path of the output peaks file")
	f.Concat = flag.String("concat", "", "path of the output highlight media file. If not specified no encoding will be done")
	f.Target = *flag.Int("target", 300, "duration target in seconds used for the time range analysis")
	f.Samples = *flag.Int("samples", 44100, "set number of samples per each audio frames")
	f.Before = *flag.Float64("before", 10, "seconds to consider for the highlight before the peak")
	f.After = *flag.Float64("after", 10, "seconds to consider for the highlight after the peak")
	f.Help = *flag.Bool("help", false, "shows help")

	flag.Parse()

	if f.Input == nil {
		return fmt.Errorf("need to specify input")
	}
	if f.Output == nil {
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
	if f.Concat != nil && (path.Ext(*f.Input) != path.Ext(*f.Concat)) {
		return fmt.Errorf("extension of input must match extension of concat")
	}

	return nil
}
