package main

import sybil "../lib"
import "os"
import "log"

func main() {

	if len(os.Args) < 2 {
		log.Fatal("Command should be one of: ingest, digest, query")
	}

	first_arg := os.Args[1]
	os.Args = os.Args[1:]
	switch first_arg {
	case "ingest":
		sybil.RunIngestCmdLine()
	case "digest":
		sybil.RunDigestCmdLine()
	case "query":
		sybil.RunQueryCmdLine()
	default:
		log.Fatal("Unknown command:", os.Args[1])
	}
}