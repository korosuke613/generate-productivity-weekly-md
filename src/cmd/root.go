package cmd

import (
	"fmt"
	"github.com/korosuke613/tempura/src/lib"
	"github.com/spf13/cobra"
	"os"
)

var (
	version          = "dev"
	commit           = "none"
	date             = "unknown"
	inputFilePath    string
	templateFilePath string
	outputFilePath   string
	inputString      string
	templateString   string
)

// Execute runs root command
func Execute() {
	if err := newRootCmd().Execute(); err != nil {
		// Not print an error because cobra.Command prints it.
		os.Exit(1)
	}
}

func newRootCmd() *cobra.Command {
	var showVersion bool

	cmd := &cobra.Command{
		Use:   "tempura",
		Short: "tempura is a generator that fills in templates.",
		Long:  `A Fast and Flexible Template Fill Generator built with love by korosuke613 in Go.`,
		RunE: func(_ *cobra.Command, args []string) error {
			if showVersion {
				printVersion()
				return nil
			}

			t := lib.Tempura{Template: templateString, TemplateFilePath: templateFilePath}

			if inputString != "" {
				if err := t.SetInputFromString(inputString); err != nil {
					return err
				}
			} else {
				if err := t.SetInputFromJSON(inputFilePath); err != nil {
					return err
				}
			}

			var output string
			if err := t.Fill(&output); err != nil {
				return err
			}
			fmt.Println(output)

			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	cmd.PersistentFlags().StringVarP(&inputFilePath, "input-filepath", "i", "input.json", "input file name")
	cmd.PersistentFlags().StringVarP(&templateFilePath, "template-filepath", "t", "template.txt", "template file name")
	cmd.PersistentFlags().StringVarP(&outputFilePath, "output", "o", "output.txt", "output file name")
	cmd.PersistentFlags().StringVar(&inputString, "input-string", "", "input string")
	cmd.PersistentFlags().StringVar(&templateString, "template-string", "", "template string")

	cmd.SetVersionTemplate(fmt.Sprintf("version: %s, commit: %s, date: %s\n", version, commit, date))
	return cmd
}

func printVersion() {
	fmt.Printf("version: %s, commit: %s, date: %s\n", version, commit, date)
}
