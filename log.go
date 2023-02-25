package main

import (
	"io"
	"log"
	"os"
	"runtime"
)

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

var (
	LogDebug = log.New(io.Discard, Yellow+"DEB "+Reset, log.Default().Flags()|log.Lmicroseconds)
	LogInfo  = log.New(os.Stdout, Green+"INF "+Reset, log.Default().Flags()|log.Lmicroseconds)
	LogError = log.New(os.Stdout, Red+"ERR "+Reset, log.Default().Flags()|log.Lmicroseconds)
)

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
	}
}
