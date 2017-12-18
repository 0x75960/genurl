package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var urlFileTemplate = `[InternetShortcut]
URL=%s`

var outputDir = flag.String("o", ".", "output dir path")
var name = flag.String("n", "link", "output name")
var fmode = flag.Bool("f", false, "make urlfile link to filepath")

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

func FilePathToURI(fp string) (uri string, err error) {

	target, err := filepath.Abs(fp)
	if err != nil {
		return uri, err
	}

	if strings.HasPrefix(target, "/") == false {
		// proc as Windows Path
		target = "/" + strings.Replace(target, `\`, `/`, -1)
	}

	return fmt.Sprintf("file://%s", target), nil
}

func main() {

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		log.Fatalln("")
	}

	link := flag.Arg(0)

	if *fmode {
		t, err := FilePathToURI(link)
		if err != nil {
			log.Fatalln(err)
		}
		link = t
	}

	f, err := os.Create(filepath.Join(*outputDir, *name+".url"))
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf(urlFileTemplate, link))
}
