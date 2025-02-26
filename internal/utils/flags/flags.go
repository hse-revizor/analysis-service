package flags

import (
	"errors"
	"flag"
)

const (
	envModeFlag = "env-mode"
)

type CMDFlags struct {
	EnvMode string
}

func ParseFlags() (*CMDFlags, error) {
	envMode := flag.String(envModeFlag, "", "Environment mode")
	flag.Parse()
	if *envMode == "" {
		return nil, errors.New("Environment mode was not found in application flags")
	}

	return &CMDFlags{EnvMode: *envMode}, nil
}

func MustParseFlags() *CMDFlags {
	flags, err := ParseFlags()
	if err != nil {
		panic(err)
	}

	return flags
}
