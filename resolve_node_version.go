package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"golang.org/x/mod/semver"
)

/**
 * resolve_node_version
 *
 * This program does the following:
 * 1. Walks up the directory to find the nearest .nvmrc
 * 2. Determine if that node version is installed with nvm
 * 3. Output a $PATH value updated with the resolved node version for the shell integration to set
 */

var nvm_dir string = os.Getenv("NVM_DIR")

func get_nvmrc_path() string {
	pwd, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	nvmrc_path := path.Join(pwd, ".nvmrc")

	if _, err := os.Stat(nvmrc_path); err == nil {
		return nvmrc_path
	}

	for pwd != "/" {
		pwd = path.Dir(pwd) // goes up one

		nvmrc_path = path.Join(pwd, ".nvmrc")

		if _, err := os.Stat(nvmrc_path); err == nil {
			return nvmrc_path
		}
	}

	return "" // TODO: Return an error in this case
}

func get_node_path(_version string) string {
	versions := []string{}
	version := strings.Trim(_version, "\n")
	version = "v" + strings.Trim(version, "v")

	is_fuzzy_version := !strings.Contains(version, ".")

	files, err := ioutil.ReadDir(nvm_dir + "/versions/node/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			versions = append(versions, file.Name())
		}
	}

	semver.Sort(versions)

	resolved_version := ""

	// Reversing the versions slice
	for i, j := 0, len(versions)-1; i < j; i, j = i+1, j-1 {
		versions[i], versions[j] = versions[j], versions[i]
	}

	for _, v := range versions {
		if !is_fuzzy_version {
			if v == version {
				resolved_version = v
				break
			}
		} else {
			if semver.Major(v) == semver.Major(version) {
				resolved_version = v
				break
			}
		}
	}

	if resolved_version == "" {
		fmt.Println("Unable to find node version that matches " + version + ", please run 'nvm install'")
		os.Exit(1)
	}

	return nvm_dir + "/versions/node/" + resolved_version + "/bin"
}

func prepend_to_path(path string) string {
	_paths := strings.Split(os.Getenv("PATH"), ":")
	paths := []string{path}

	for _, p := range _paths {
		if !strings.Contains(p, nvm_dir) {
			paths = append(paths, p)
		}
	}

	return strings.Join(paths, ":")
}

func main() {
	nvmrc_path := get_nvmrc_path()

	version := ""

	if nvmrc_path != "" {
		value, _ := os.ReadFile(nvmrc_path)
		version = string(value)
	}

	if version == "" {
		// Read the nvm "default" alias
		value, _ := os.ReadFile(nvm_dir + "/alias/default")
		version = string(value)
	}

	node_path := get_node_path(version)

	fmt.Println(prepend_to_path(node_path))
}
