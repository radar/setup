package toolversions

import (
	"bufio"

	"os"
	"strings"
)

type Versions struct {
	Versions map[string]string
}

const toolVersions = ".tool-versions"


func ForPackage(pkg string) (string, error) {
	versions, err := Load()
	if err != nil {
		return "", err
	}
	return versions.forPackage(pkg)
}

func (versions Versions) forPackage(pkg string) (string, error) {
	return versions.Versions[pkg], nil
}

func Present() bool {
	if _, err := os.Stat(toolVersions); os.IsNotExist(err) {
		return false
	}

	return true
}

func Load() (Versions, error) {
	f, err := os.Open(toolVersions)
	versions := Versions{Versions: make(map[string]string)}

	if err != nil {
		return versions, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.FieldsFunc(scanner.Text(), splitter)
		versions.Versions[parts[0]] = parts[1]
	}

	return versions, err
}

func splitter(r rune) bool {
	return r == '-' || r == ' '
}
