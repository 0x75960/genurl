package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var urlFileTemplate = `[InternetShortcut]
URL=%s`

var outputDir = flag.String("o", ".", "output dir path")
var name = flag.String("n", "link", "output name")

func init() {

	flag.Usage = func() {

		fmt.Fprintf(
			os.Stderr,
			"Usage of this:\n\t %s [options] https://url.to/gen/url\n\n",
			os.Args[0],
		)

		flag.PrintDefaults()
	}

}

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		log.Fatalln("")
	}

	link := flag.Arg(0)

	f, err := os.Create(filepath.Join(*outputDir, *name+".url"))
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf(urlFileTemplate, link))
}
