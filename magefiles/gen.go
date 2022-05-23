package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"go/format"
	"io"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

var codeTemplate = `
package oui // generated code - do not edit

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
	Entries []entry
	Vendors map[string]uint
}

type entry struct {
	OUI      string
	Vendor   string
	VendorID uint
}

func generate(src, dst string) error {
	mg.Deps(download)

	fin, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fin.Close()

	data := newTemplateData(fin)

	tmpl, err := template.New("oui").Parse(codeTemplate)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, "oui", data); err != nil {
		return err
	}

	fout, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer fout.Close()

	code, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if _, err := fout.Write(code); err != nil {
		return err
	}

	return nil
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
		if errors.Is(err, io.EOF) {
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
