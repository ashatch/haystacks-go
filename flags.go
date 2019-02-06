package main

import (
	"errors"
	"flag"
)

// NeedleHayFlags represents flags given to epoch program
type NeedleHayFlags struct {
	needles string
}

// GetNeedleHayFlags gives an EpochFlags based on command line args
func GetNeedleHayFlags() (NeedleHayFlags, error) {
	flags := NeedleHayFlags{}
	flag.StringVar(&flags.needles, "needles", "", "Things to find")
	flag.Parse()

	if flags.needles == "" {
		flag.Usage()
		return flags, errors.New("Missing --needles parameter")
	}

	return flags, nil
}
