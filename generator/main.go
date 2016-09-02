package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const ianaURL = "http://www.iana.org/assignments/tls-parameters/tls-parameters-4.csv"

func main() {
	var (
		loc     = flag.String("url", ianaURL, "location of IANA TLS CSV")
		outfile = flag.String("out", "ciphermap.go", "name of output file")
	)
	flag.Parse()
	resp, err := http.Get(*loc)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r := csv.NewReader(resp.Body)

	buf := bytes.Buffer{}
	buf.WriteString("package tlstext\n")
	buf.WriteString("\n")
	buf.WriteString("var cipherMap = map[uint16]string{\n")

	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("unable to read: %s", err)
		}
		if strings.HasPrefix(records[1], "TLS") {
			val := records[0][0:4] + records[0][7:]
			buf.WriteString(fmt.Sprintf("%s: %q,\n", val, records[1]))
		}
	}
	buf.WriteString("}\n")
	out, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatalf("invalid source code: %s", err)
	}

	// write output to file
	fo, err := os.Create(*outfile)
	if err != nil {
		log.Fatalf("unable to create: %s", err)
	}
	_, err = fo.Write(out)
	if err != nil {
		log.Fatalf("unable to write: %s", err)
	}
	fo.Close()
}
