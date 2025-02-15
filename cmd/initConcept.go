/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/redradrat/kable/pkg/concepts"

	"github.com/spf13/cobra"
)

var conceptTypeInput string

// initConceptCmd represents the initConcept command
var initConceptCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a concept in the current folder",
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		wd, err := os.Getwd()
		if err != nil {
			PrintError("unable to initialize concept dir: %s", err)
		}
		name := filepath.Base(wd)
		PrintMsg("Initializing concept '%s' of type '%s'...", name, conceptTypeInput)

		ct := concepts.ConceptType(conceptTypeInput)
		if !ct.IsSupported() {
			PrintError("given concept type is unsupported")
		}

		if err := concepts.InitConcept(".", name, concepts.ConceptType(conceptTypeInput)); err != nil {
			PrintError("unable to initialize concept dir: %s", err)
		}
		PrintSuccess("Successfully initialized Concept!")
	},
}

func init() {
	rootCmd.AddCommand(initConceptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conceptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	initConceptCmd.Flags().StringVarP(
		&conceptTypeInput,
		"type",
		"t",
		"jsonnet",
		"the type this concept should have",
	)
}
