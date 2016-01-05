package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	args := flag.Args()
	file, err := os.Open(args[0]) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)

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
	fmt.Printf(string(out))
}
