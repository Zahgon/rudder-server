package main

import (
	"log"
	"os"
	"slices"

	"gopkg.in/yaml.v3"
)

const workflowFile = "./.github/workflows/tests.yaml"

const jobName = "warehouse-integration"

var IgnorePackages = []string{
	"warehouse/integrations/manager",
	"warehouse/integrations/middleware",
	"warehouse/integrations/testhelper",
	"warehouse/integrations/testdata",
	"warehouse/integrations/config",
	"warehouse/integrations/types",
	"warehouse/integrations/tunnelling",
}

func main() {
	// Read the GitHub Actions YAML file
	data, err := os.ReadFile(workflowFile)
	if err != nil {
		panic(err)
	}

	// Parse the YAML file
	var config struct {
		Jobs map[string]struct {
			Strategy struct {
				Matrix struct {
					Include []struct {
						Package string `yaml:"package"`
					}
				}
			}
		}
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	job := config.Jobs[jobName]

	// Extract the matrix entries

	// Get the subfolders in the current directory
	subfolders, err := getSubfolders("./warehouse/integrations")
	if err != nil {
		panic(err)
	}
	for _, folder := range subfolders {
		match := false

		if slices.Contains(IgnorePackages, folder) {
			continue
		}

		// Compare the matrix entries with the subfolders
		for _, entry := range job.Strategy.Matrix.Include {
			if folder == entry.Package {
				match = true
				break
			}
		}

		if !match {
			log.Printf("Folder %q does not have a matrix entry in %s (%q) \n", folder, jobName, workflowFile)
			os.Exit(1)
		}

	}
}

// getSubfolders returns a list of subfolders in the specified directory.
func getSubfolders(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
