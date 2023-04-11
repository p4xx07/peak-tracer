package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Paxx-RnD/go-helper/helpers/slice_helper"
	"github.com/Paxx-RnD/peak-tracer/types"
	"os"
	"os/exec"
	"path"
	"sort"
	"strconv"
	"strings"
	time "time"
)

func main() {
	flags := types.Flags{}
	err := flags.Set()
	if err != nil {
		panic(err)
	}

	if flags.Help {
		flag.Usage()
		return
	}

	fmt.Println("reading flags")
	args := fmt.Sprintf("-y -i %s -vn -af asetnsamples=%d,astats=metadata=1:reset=1,ametadata=print:key=lavfi.astats.Overall.RMS_level:file=log.txt -f null -", flags.Input, flags.Samples)
	split := strings.Split(args, " ")
	command := exec.Command("ffmpeg", split...)
	command.Args = slice_helper.RemoveEmptyEntries(command.Args)

	fmt.Println("execute ffmepg command flags")
	res, err := command.Output()
	if err != nil {
		fmt.Printf("issue -> %s %v", string(res), err.Error())
		return
	}

	fmt.Println("reading log file")
	text, _ := os.ReadFile("log.txt")
	defer os.Remove("log.txt")

	split = strings.Split(string(text), "\n")

	peaks := make([]types.Peak, 0)
	peak := types.Peak{}
	for _, l := range split {
		if l == "" {
			continue
		}
		if !strings.HasPrefix(l, "l") {
			s := strings.Split(l, ":")
			peak.Time, _ = strconv.ParseFloat(s[3], 64)
			continue
		}

		s := strings.Split(l, "=")
		peak.RMS, _ = strconv.ParseFloat(s[1], 64)
		peaks = append(peaks, peak)
		peak = types.Peak{}
	}

	fmt.Println("sorting peaks")
	peaksByRMS := types.PeaksByRMS(peaks)
	sort.Sort(peaksByRMS)

	ranges := make([]types.Range, 0)

	for _, p := range peaksByRMS {
		if InRange(flags.Before, flags.After, p.Time, ranges) {
			continue
		}

		r := types.Range{Min: p.Time - flags.Before, Max: p.Time + flags.After}
		ranges = append(ranges, r)
	}

	var max float64
	i := 0
	for i = range ranges {
		max += flags.After + flags.Before
		if max > float64(flags.Target) {
			break
		}
	}

	if i > len(ranges)-1 {
		i = len(ranges) - 1
	}

	ranges = ranges[:i]

	fmt.Println("slicing out of target ranges")
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Max
	})

	peaksStruct := slice_helper.Map(&ranges, func(p types.Range) float64 {
		return p.Min + (flags.Before)
	})

	fmt.Println("creating the peaks struct")
	json, err := json.Marshal(peaksStruct)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file, err := os.Create(flags.Output)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Printf("writing peaks file to %s", file.Name())
	s := string(json)
	_, err = file.WriteString(s)

	fmt.Printf("completed")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	if flags.Concat == "" {
		return
	}
	fmt.Printf("Joining the videos")
	concat(flags, ranges)
}

func concat(flags types.Flags, ranges []types.Range) {
	t := time.Now()
	defer func() { fmt.Sprintf("Took %s", time.Since(t).String()) }()
	extension := path.Ext(flags.Input)
	file, err := os.Create("concat.txt")
	if err != nil {
		panic(err)
	}
	defer os.Remove(file.Name())
	defer file.Close()

	fmt.Println("creating concat file")
	for i, p := range ranges {
		output := fmt.Sprintf("%d%s", i, extension)
		args := fmt.Sprintf("-y -ss %f -t %f -i %s %s", p.Min, p.Max-p.Min, flags.Input, output)
		split := strings.Split(args, " ")
		split = slice_helper.RemoveEmptyEntries(split)
		command := exec.Command("ffmpeg", split...)
		_, err = command.Output()
		if err != nil {
			panic(err)
		}
		file.WriteString(fmt.Sprintf("file '%s'\n", output))
		defer os.Remove(output)
	}

	fmt.Println("concatenating")
	command := exec.Command("ffmpeg", "-y", "-f", "concat", "-safe", "0", "-i", "concat.txt", flags.Concat)
	_, err = command.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("completed")
	return
}

func InRange(before float64, after float64, time float64, ranges []types.Range) bool {
	for _, r := range ranges {
		if r.Min-before <= time && time <= r.Max+after {
			return true
		}
	}

	return false
}
