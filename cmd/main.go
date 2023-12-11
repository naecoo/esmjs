package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func Build(input, output string, minify, isBrowser bool) {
	platform := api.PlatformNode
	if isBrowser {
		platform = api.PlatformBrowser
	}

	result := api.Build(api.BuildOptions{
		Bundle:            true,
		Write:             true,
		LogLevel:          api.LogLevelInfo,
		Format:            api.FormatESModule,
		MinifyWhitespace:  minify,
		MinifyIdentifiers: minify,
		MinifySyntax:      minify,
		Platform:          platform,
		EntryPoints:       []string{input},
		Outfile:           output,
	})

	if len(result.Errors) > 0 {
		os.Exit(1)
	}

	if len(result.Warnings) > 0 {
		for _, message := range result.Warnings {
			fmt.Printf("warning: %+v\n", message)
		}
	}
}

func Execute() {
	for _, arg := range os.Args[1:] {
		if strings.Contains(arg, "-h") || strings.Contains(arg, "--help") {
			fmt.Println("One Command transform Javascript to ESM module")
			fmt.Println("Usage: esmjs <input> <output> [flags...]")
			fmt.Println("-m    requires minify, default is false")
			fmt.Println("-b    build script for browser platform, default is node platform")
			fmt.Println("Eg: esmjs input.ts dist.js -m -b")
			os.Exit(0)
		}
		if strings.Contains(arg, "-v") || strings.Contains(arg, "--version") {
			fmt.Println("v0.0.1")
			os.Exit(0)
		}
	}

	var input, output string
	argLen := len(os.Args)
	if argLen <= 1 {
		fmt.Println("esmjs: try 'esmjs --help' for more information")
		os.Exit(0)
	}
	input = os.Args[1]
	if _, error := os.Stat(input); error != nil {
		if errors.Is(error, os.ErrNotExist) {
			fmt.Println("input file not exists")
			os.Exit(-1)
		} else {
			panic(error)
		}
	}

	if len(os.Args) <= 2 {
		fmt.Println("no output file")
		os.Exit(-1)
	}
	output = os.Args[2]

	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	minify := flagSet.Bool("m", false, "requires minify, default is false")
	isBrowser := flagSet.Bool("b", false, "build script for browser platform, default is node platform")
	flagSet.Parse(os.Args[3:])

	Build(input, output, *minify, *isBrowser)
}
