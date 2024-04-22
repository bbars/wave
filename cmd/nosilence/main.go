package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"time"

	"github.com/bbars/wave"
)

var (
	multiply    = 1.0
	power       = 1.0
	threshold   = 0.007
	fall        = 0.0005
	rise        = 0.9
	delay       = 1 * time.Second
	logInterval = 1 * time.Second
)

func main() {
	flag.Float64Var(&multiply, "multiply", multiply, "multiply level")
	flag.Float64Var(&power, "power", power, "power level")
	flag.Float64Var(&threshold, "threshold", threshold, "level threshold to pause")
	flag.Float64Var(&fall, "fall", fall, "level EMA fall speed")
	flag.Float64Var(&rise, "rise", rise, "level EMA rise speed")
	flag.DurationVar(&delay, "delay", delay, "delay to pause after level felt down threshold")
	flag.DurationVar(&logInterval, "log-interval", logInterval, "current level EMA value logging interval")
	flag.Parse()

	ctx := context.Background()
	input, err := getInput()
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	header := wave.Header{}
	if _, err = header.ReadFrom(input); err != nil {
		log.Fatal(err)
	}

	data := wave.ChunkData{}
	if _, err = data.ReadFrom(input); err != nil {
		log.Fatal(fmt.Errorf("failed to read data chunk header: %w", err))
	}

	sr, err := wave.NewSampleReader(input, header)
	if err != nil {
		log.Fatal(err)
	}

	output, err := getOutput()
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	sw, err := wave.NewSampleWriter(output, header)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = header.WriteTo(output); err != nil {
		log.Fatal(err)
	}
	if _, err = data.WriteTo(output); err != nil {
		log.Fatal(err)
	}

	if err = processSamples(ctx, header, sr, sw); err != nil {
		log.Fatal(err)
	}
}

func getInput() (io.ReadCloser, error) {
	if flag.NArg() <= 0 || flag.Args()[0] == "-" {
		return os.Stdin, nil
	}

	f, err := os.Open(flag.Args()[0])
	if err != nil {
		return nil, err // TODO: wrap
	}

	return f, nil
}

func getOutput() (io.WriteCloser, error) {
	if flag.NArg() <= 1 || flag.Args()[1] == "-" {
		return os.Stdout, nil
	}

	f, err := os.OpenFile(flag.Args()[1], os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err // TODO: wrap
	}

	return f, nil
}

func processSamples(_ context.Context, header wave.Header, sr wave.SampleReader, sw wave.SampleWriter) error {
	var delayTicks = int64(float64(header.SamplesPerSec) * float64(delay) / float64(time.Second))
	var logIntervalTicks = int64(float64(header.SamplesPerSec) * float64(logInterval) / float64(time.Second))

	var avg float64 = 0.5
	var pos int64 = -1
	var pauseAt int64 = delayTicks

	for {
		pos++

		s, err := sr.ReadSample()
		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err // TODO: wrap
		}

		for i, v := range s.Values {
			if multiply != 1.0 {
				v *= multiply
			}

			switch {
			case power == 1.0:
			case v < 0:
				v = -math.Pow(math.Abs(v), power)
			default:
				v = math.Pow(v, power)
			}

			s.Values[i] = v

			if v > avg {
				avg = avg*(1.0-rise) + v*rise
			} else {
				avg = avg*(1.0-fall) + v*fall
			}
		}

		if logIntervalTicks > 0 && pos%logIntervalTicks == 0 {
			log.Println("avg", avg)
		}

		if avg > threshold {
			if pos > pauseAt {
				log.Println("resume", pos, avg)
			}
			pauseAt = pos + delayTicks
		}

		if pos > pauseAt {
			if pauseAt > 0 {
				log.Println("pause", pos, avg)
				pauseAt = 0
			}
			continue
		}

		if err = sw.WriteSample(s); err != nil {
			return err
		}
	}
}
