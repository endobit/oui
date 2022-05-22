package main

import (
	"bytes"
	"encoding/csv"
	"flag"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

var codeTemplate = `
package {{ .Package }} // generated code - do not edit

type ouiDB struct {
	ouis    map[string]uint
	vendors map[uint]string
}

var database = ouiDB{
    ouis: map[string]uint{
{{- range .Entries }}
"{{ .OUI }}": {{ .VendorID }}, // {{ .Vendor }}
{{- end }}
    },
    vendors: map[uint]string{
{{- range $key, $value := .Vendors }}
{{ $value }}: "{{ $key }}",
{{- end }}
    },
}
`

type templateData struct {
	Package string
	Entries []entry
	Vendors map[string]uint
}

type entry struct {
	OUI      string
	Vendor   string
	VendorID uint
}

func main() {
	var (
		inFile  string
		outFile string
		pkg     string
	)

	flag.StringVar(&inFile, "in", "oui.csv", "name of the oui mac file")
	flag.StringVar(&outFile, "out", "data.go", "name of the output file")
	flag.StringVar(&pkg, "pkg", "oui", "name of the go package")
	flag.Parse()

	fin, err := os.Open(inFile)
	if err != nil {
		log.Fatalf("cannot open %q: %v", inFile, err)
	}

	defer fin.Close()

	data := newTemplateData(fin)
	data.Package = pkg

	tmpl, err := template.New("oui").Parse(codeTemplate)
	if err != nil {
		log.Printf("template parse error: %v", err)
		return
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, "oui", data); err != nil {
		log.Printf("template execute error: %v", err)
		return
	}

	fout, err := os.Create(outFile)
	if err != nil {
		log.Printf("cannot open %q: %v", outFile, err)
		return
	}

	defer fout.Close()

	code, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("cannot format code: %v", err)
		return
	}

	if _, err := fout.Write(code); err != nil {
		log.Printf("cannot write code: %v", err)
		return
	}
}

func newTemplateData(r io.Reader) *templateData {
	var entries []entry

	id := uint(0)
	ouis := make(map[string]string)
	idmap := make(map[string]uint)

	c := csv.NewReader(r)

	_, err := c.Read() // skip header
	if err != nil {
		panic(err)
	}

	for {
		record, err := c.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		o := strings.ToLower(record[1])

		v := strings.TrimSpace(record[2])
		v = strings.ReplaceAll(v, `"`, "")

		if prev, ok := ouis[o]; ok { // 080030 is a known duplicate
			log.Printf("Warning %q:%q is already registered to %q", o, v, prev)
			continue
		}

		ouis[o] = v

		if _, ok := idmap[v]; !ok {
			idmap[v] = id
			id++
		}

		entries = append(entries, entry{OUI: o, Vendor: v, VendorID: idmap[v]})
	}

	return &templateData{
		Entries: entries,
		Vendors: idmap,
	}
}
