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
	"sort"
	"strings"
)

const ianaURL = "http://www.iana.org/assignments/tls-parameters/tls-parameters-4.csv"

func readIANA(loc string) (map[string]string, error) {

	resp, err := http.Get(loc)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ciphers := map[string]string{}

	r := csv.NewReader(resp.Body)
	for {
		records, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if !strings.HasPrefix(records[1], "TLS") {
			continue
		}
		hex := string(records[0][0:4] + records[0][7:])
		ciphers[hex] = records[1]
	}
	return ciphers, nil
}

func main() {
	var (
		loc     = flag.String("url", ianaURL, "location of IANA TLS CSV")
		outfile = flag.String("out", "ciphermap.go", "name of output file")
	)
	flag.Parse()
	ciphers, err := readIANA(*loc)
	if err != nil {
		log.Fatalf("Unable to read IANA file: %s", err)
	}

	inverse := map[string]string{}
	keys := []string{}
	values := []string{}

	for k, v := range ciphers {
		inverse[v] = k
		keys = append(keys, k)
		values = append(values, v)
	}
	sort.Strings(keys)
	sort.Strings(values)

	buf := bytes.Buffer{}
	buf.WriteString("package tlstext\n")
	buf.WriteString("\n")
	buf.WriteString("var cipherMap = map[uint16]string{\n")
	for _, k := range keys {
		buf.WriteString(fmt.Sprintf("%s: %q,\n", k, ciphers[k]))
	}
	buf.WriteString("}\n")
	buf.WriteString("var cipherStringMap = map[string]uint16{\n")
	for _, v := range values {
		buf.WriteString(fmt.Sprintf("%q: %s,\n", v, inverse[v]))
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
