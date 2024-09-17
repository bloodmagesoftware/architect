package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/bloodmagesoftware/architect/internal/version"
	"github.com/charmbracelet/log"
	"os"
	"runtime/debug"
)

var (
	Upgrade bool
	Version version.Version
)

func init() {
	flag.BoolVar(&Upgrade, "upgrade", false, "upgrade old wold files")
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()

	if info, ok := debug.ReadBuildInfo(); ok {
		var err error
		Version, err = version.Parse(info.Main.Version)
		if err != nil {
			panic(errors.Join(fmt.Errorf("failed to parse version %s", info.Main.Version), err))
		}
		if Version.Major < 0 {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.WarnLevel)
		}
	} else {
		panic("failed to get build info")
	}

	if *showVersion {
		fmt.Println(Version.String())
		os.Exit(0)
		return
	}
}
