package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"go/format"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/magefile/mage/mg"
)

var codeTemplate = `
package oui // generated code - do not edit

var ouis = map[string]int{
{{- range .Entries }}
"{{ .OUI }}": {{ .VendorID }}, // {{ .Vendor }}
{{- end }}
}

var vendors = []string{
{{- range .Vendors }}
"{{ . }}",
{{- end }}
}

`

type templateData struct {
	Entries []entry
	Vendors []string
}

type entry struct {
	OUI      string
	VendorID int
	Vendor   string
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
		// if _, err := fout.Write(buf.Bytes()); err != nil {
		return err
	}

	return nil
}

func newTemplateData(r io.Reader) *templateData {
	var (
		entries []entry
		vendors []string
	)

	ouiMap := make(map[string]string)
	vendorMap := make(map[string]int)

	c := csv.NewReader(r)

	_, err := c.Read() // skip header
	if err != nil {
		panic(err)
	}

	for id := 0; ; {
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
		v = simplifyName(v)

		if prev, ok := ouiMap[o]; ok { // 080030 is a known duplicate
			log.Printf("Warning %q:%q is already registered to %q", o, v, prev)
			continue
		}

		ouiMap[o] = v

		if _, ok := vendorMap[v]; !ok {
			vendors = append(vendors, v)
			vendorMap[v] = id
			id++
		}

		entries = append(entries, entry{OUI: o, Vendor: v, VendorID: vendorMap[v]})
	}

	return &templateData{
		Entries: entries,
		Vendors: vendors,
	}
}

var (
	llcRegex  = regexp.MustCompile(`(?i),?\s*(llc|ltd|limited|inc|incorporated)\.?$`)
	coRegex   = regexp.MustCompile(`(?i),?\s*(co|company|corp|corporation)\.?$`)
	gmbhRegex = regexp.MustCompile(`(?i),?\s*gmbh\.?$`)
)

func simplifyName(name string) string {
	b := []byte(name)

	b = llcRegex.ReplaceAll(b, []byte{})
	b = coRegex.ReplaceAll(b, []byte{})
	b = gmbhRegex.ReplaceAll(b, []byte{})

	return string(b)
}
