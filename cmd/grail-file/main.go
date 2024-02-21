package main

import (
	"context"
	"flag"
	"os"

	"github.com/Schaudge/grailbase/cmd/grail-file/cmd"
	"github.com/Schaudge/grailbase/file"
	"github.com/Schaudge/grailbase/file/s3file"
	"github.com/Schaudge/grailbase/log"
)

func main() {
	help := flag.Bool("help", false, "Display help about this command")
	flag.Parse()
	if *help {
		cmd.PrintHelp()
		os.Exit(0)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	file.RegisterImplementation("s3", func() file.Implementation {
		return s3file.NewImplementation(s3file.NewDefaultProvider(), s3file.Options{})
	})
	err := cmd.Run(context.Background(), os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
