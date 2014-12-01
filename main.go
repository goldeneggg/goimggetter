package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/goldeneggg/goimggetter/imggetter"
	"github.com/jessevdk/go-flags"
)

const (
	Version = "0.1.0"
)

// element names need to Uppercase
type options struct {
	Site        string `short:"s" long:"site" description:"Target site" default:""`
	Offset      int    `short:"o" long:"offset" description:"Search result offset index" default:"1"`
	LimitPage   int    `short:"l" long:"limitpage" description:"Search result limit page" default:"1"`
	Concurrency int    `short:"c" long:"concurrency" description:"Concurrency count for save img" default:"1"`
	Debug       bool   `short:"d" long:"debug" description:"Debug detail information"`
	Version     bool   `short:"v" long:"version" description:"Print version"`
	Help        bool   `short:"h" long:"help" description:"Show help message"` // not "help" but "Help", because cause error using "-h" option
}

func main() {
	// handler for return
	var status int
	defer func() { os.Exit(status) }()

	// parse option args
	opts := &options{}
	parser := flags.NewParser(opts, flags.PrintErrors)
	args, err := parser.Parse()
	if err != nil {
		//fmt.Fprintln(os.Stderr, err)
		printHelp()
		status = 1
		return
	}

	// print help
	if opts.Help {
		printHelp()
		return
	}

	// print version
	if opts.Version {
		fmt.Fprintf(os.Stderr, "goimggetter: version %s (%s)\n", Version, runtime.GOARCH)
		return
	}

	// validate site
	if len(opts.Site) == 0 {
		fmt.Fprintf(os.Stderr, "-s or --site parameter is required\n")
		printHelp()
		status = 1
		return
	}

	// validate query
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "query arg is not assigned: %+v\n", args)
		printHelp()
		status = 1
		return
	}

	// download
	for p := opts.Offset; p <= opts.LimitPage; p++ {
		if err := imggetter.Download(opts.Site, args[0], p, opts.Concurrency, opts.Debug); err != nil {
			fmt.Println("Download error", err)
		}
	}
}

func printHelp() {
	h := `
Usage:
  goimggetter -s <SITE> [OTHER OPTIONS] <QUERY>

Application Options:
  -s, --site=        Target site (ex. flickr, and more...)
  -o, --offset=      Offset index of search result *default = 1
  -l, --limitpage=   Limit page of search result *default = 1
  -c, --concurrency= Concurrency count for save img *default = 1
  -d, --debug        Debug detail information
  -v, --version      Print version

Help Options:
  -h, --help         Show this help message
`
	os.Stderr.Write([]byte(h))
}
