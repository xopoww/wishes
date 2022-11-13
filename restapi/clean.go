//go:build ignore
// +build ignore

// This program removes all .go files generated by go-swagger from ./, ./operations/ and ../models/

package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"os"
	"path"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func checkIfGenerated(r *bufio.Reader) bool {
	for {
		line, err := r.ReadString('\n')
		if errors.Is(err, io.EOF) {
			return false
		}
		if err != nil {
			log.Fatal().Err(err).Msg("r.ReadString")
		}
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}
		if line == `// Code generated by go-swagger; DO NOT EDIT.` {
			return true
		}
		if strings.HasPrefix(line, "//") {
			continue
		}
		return false
	}
}

func handleDirectory(dir string) {
	log.Debug().Str("dir", dir).Msg("reading directory")
	des, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal().Str("dir", dir).Err(err).Msg("os.ReadDir")
	}
	for _, de := range des {
		entry := path.Join(dir, de.Name())
		if de.IsDir() {
			handleDirectory(entry)
			continue
		}
		f, err := os.Open(entry)
		if err != nil {
			log.Fatal().Str("file", entry).Err(err).Msg("os.Open")
		}
		r := bufio.NewReader(f)
		generated := checkIfGenerated(r)
		log.Debug().Str("file", entry).Bool("generated", generated).Msg("checked file")
		if err := f.Close(); err != nil {
			log.Warn().Str("file", entry).Err(err).Msg("os.Close")
		}
		if generated {
			if err := os.Remove(entry); err != nil {
				log.Fatal().Str("file", entry).Err(err).Msg("os.Remove")
			}
		}
	}
}

var (
	args struct {
		quiet bool
	}
)

func main() {
	flag.BoolVar(&args.quiet, "quiet", false, "reduce logs to error or higher")
	flag.Parse()

	var lvl zerolog.Level
	if args.quiet {
		lvl = zerolog.ErrorLevel
	} else {
		lvl = zerolog.DebugLevel
	}
	log.Logger = zerolog.New(zerolog.ConsoleWriter{
		Out: os.Stderr,
	}).Level(lvl)

	root := path.Dir(os.Getenv("GOFILE"))

	dirs := make([]string, 0)
	dirs = append(dirs, ".")

	for _, dir := range dirs {
		handleDirectory(path.Clean(path.Join(root, dir)))
	}
}
