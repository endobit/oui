package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/magefile/mage/target"
)

const (
	ouiURL  = "https://standards-oui.ieee.org/oui/oui.csv"
	ouiCSV  = "oui.csv"
	ouiData = "data.go"
)

// Build builds the package
func Build() {
	mg.Deps(mg.F(generate, ouiCSV, ouiData))
}

// Lint runs golangci-lint
func Lint() error {
	return sh.RunV("golangci-lint", "run", "--skip-dirs", "magefiles")
}

// Test runs go test with coverage
func Test() error {
	mg.Deps(Build)

	return sh.RunV("go", "test", "-v", "-coverprofile=coverage.txt", "-covermode=count", "./...")
}

// Clean removes build and test artifacts
func Clean() error {
	for _, file := range []string{ouiCSV, "coverage.txt"} {
		if err := sh.Rm(file); err != nil {
			return err
		}
	}

	return nil
}

// Nuke remove all artifact includes the generated code
func Nuke() error {
	mg.Deps(Clean)

	return sh.Rm(ouiData)
}

func download() error {
	stale, err := target.Path(ouiCSV, "magefiles/mage.go", "magefiles/gen.go")
	if err != nil {
		return err
	}

	if !stale {
		return nil
	}

	log.Printf("downloading %q", ouiURL)

	req, err := http.NewRequest(http.MethodGet, ouiURL, nil)
	if err != nil {
		return err
	}

	// Because default Go-http-client/1.1 will result in 418 error
	req.Header.Set("User-Agent", "Mage-oui-parser")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: bad status code=%s", resp.Status)
	}

	fout, err := os.Create("oui.csv")
	if err != nil {
		return err
	}
	defer fout.Close()

	_, err = io.Copy(fout, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	_ = os.Setenv("MAGEFILE_ENABLE_COLOR", "1")
	_ = os.Setenv("MAGEFILE_VERBOSE", "1")
}
