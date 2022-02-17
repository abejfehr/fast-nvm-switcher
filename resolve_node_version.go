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
 * 2. Determine if that version of node is installed with nvm
 * 3. Output the path to that version of node so that the shell script can set it to $PATH
 */

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

	files, err := ioutil.ReadDir(os.Getenv("NVM_DIR") + "/versions/node/")
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

	return os.Getenv("NVM_DIR") + "/versions/node/" + resolved_version + "/bin"
}

func main() {
	nvmrc_path := get_nvmrc_path()

	version := ""

	if nvmrc_path != "" {
		value, _ := os.ReadFile(nvmrc_path)
		version = string(value)
	}

	if version == "" {
		// Get the nvm "default" alias
		value, _ := os.ReadFile(os.Getenv("NVM_DIR") + "/alias/default")
		version = string(value)
	}

	node_path := get_node_path(version)

	fmt.Println(node_path)
}
