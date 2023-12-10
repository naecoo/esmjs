package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

func build(input, output string, minify, isBrowser bool) {
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
	var output string
	flag.StringVar(&output, "o", "", "the full path of the output file")
	minify := flag.Bool("m", false, "requires minify, default is false")
	isBrowser := flag.Bool("b", false, "build script for browser platform, default is node platform")

	flag.Parse()

	if len(output) == 0 {
		fmt.Println("no output file option, please specify -o option")
		os.Exit(-1)
	}

	tail := flag.Args()
	if len(tail) == 0 {
		fmt.Println("no input file")
		os.Exit(-1)
	}
	input := tail[0]
	if _, error := os.Stat(input); error != nil {
		if errors.Is(error, os.ErrNotExist) {
			fmt.Println("input file not exists")
			os.Exit(-1)
		} else {
			panic(error)
		}
	}

	build(input, output, *minify, *isBrowser)
}
