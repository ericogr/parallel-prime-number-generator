package main

import (
	"io"
	"log"
	"os"
)

var LogDebug = log.New(io.Discard, "DEBUG ", log.Default().Flags()|log.Lmicroseconds)
var LogInfo = log.New(os.Stdout, "INFO ", log.Default().Flags()|log.Lmicroseconds)
var LogError = log.New(os.Stdout, "ERR ", log.Default().Flags()|log.Lmicroseconds)
