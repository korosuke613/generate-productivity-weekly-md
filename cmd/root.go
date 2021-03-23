package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/korosuke613/tempura/lib"
	"github.com/spf13/cobra"
)

var (
	version          string
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
		Short: "tempura is a Tool that fills in templates.",
		Long:  `A Fast and Flexible Template Fill Tool built with love by korosuke613 in Go.`,
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

			output, err := t.Fill()
			if err != nil {
				return err
			}

			if outputFilePath != "" {
				file, err := os.Create(outputFilePath)
				if err != nil {
					return err
				}
				defer file.Close()

				_, err = file.WriteString(output)
				if err != nil {
					return err
				}
			} else {
				fmt.Println(output)
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&showVersion, "version", "v", false, "show version")
	cmd.PersistentFlags().StringVarP(&inputFilePath, "input-filepath", "i", "input.json", "input file name")
	cmd.PersistentFlags().StringVarP(&templateFilePath, "template-filepath", "t", "template.tmpl", "template file name")
	cmd.PersistentFlags().StringVarP(&outputFilePath, "output", "o", "", "output file name")
	cmd.PersistentFlags().StringVar(&inputString, "input-string", "", "input string")
	cmd.PersistentFlags().StringVar(&templateString, "template-string", "", "template string")

	return cmd
}

func printVersion() {
	trueVersion := version

	if version == "" {
		info, ok := debug.ReadBuildInfo()
		if ok {
			trueVersion = info.Main.Version
		} else {
			trueVersion = "(devel)"
		}
	}

	fmt.Printf("tempura version: %s\n", trueVersion)
}
