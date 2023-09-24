package utils

import (
	"fmt"
	"os"
	"strings"

	"codeberg.org/Tomkoid/mdhtml/internal/models"
	flag "github.com/spf13/pflag"
)

func ParseArgs() models.Args {
	// create and parse flags using the flag package
	file := flag.StringP("file", "f", "", "The markdown file to convert to HTML")
	out := flag.StringP("out", "o", "", "The destination file to write the HTML to")
	style := flag.StringP("style", "s", "", "Apply extra styling to the HTML using a CSS file")
	watch := flag.BoolP("watch", "w", false, "Watch the source file for changes and reconvert when changes are detected")
	httpServer := flag.BoolP("httpserver", "H", false, "Start a HTTP server to serve the HTML file and reload the page when changes are detected")
	serverPort := flag.IntP("port", "p", 8080, "The port to use for the HTTP server")
	serverHostname := flag.StringP("hostname", "h", "localhost", "The hostname to use for the HTTP server")
	debug := flag.BoolP("debug", "d", false, "Enable debug mode")

	flag.Parse()

	if *file == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *out == "" {
		split := strings.Split(*file, ".md")
		*out = fmt.Sprintf("%s.html", split[0])
	}

	// enable watch mode if httpserver is enabled
	if *httpServer {
		*watch = true
	}

	// return instance of Args
	return models.Args{
		File:           *file,
		Out:            *out,
		Style:          *style,
		Watch:          *watch,
		HttpServer:     *httpServer,
		ServerPort:     *serverPort,
		ServerHostname: *serverHostname,
		Debug:          *debug,
	}
}
