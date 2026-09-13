package cypher

import (
	"github.com/spf13/cobra"
)

var (
	library bool
	file    bool
	output  bool
	decrypt bool
)

var rootCmd = &cobra.Command{
	Use:   "cryps -l <path_to_library>",
	Short: "Use a defined library to encrypt your text",
	Long: `A CLI program that takes a user defined encryption library json file and encrypts text using said library

			Usage:
	  		cryps [flags] <text>

			Encrypts the given text using the key file passed via -l/--library
			and prints the result to stdout.

			Flags:
			  -h, --help              Show this help message
			  -f, --file <path>       Read input from a file instead of the command line
			  -o, --output <path>     Write output to this file instead of printing to
                           the terminal (only used with -f)
			  -l, --library <path>    (required) Path to the encryption key/library JSON file
			  -d, --decrypt           Decrypt instead of encrypt

			Examples:
			  cryps -l key.json "hello there"
			  cryps -l key.json -f message.txt
			  cryps -l key.json -f message.txt -o out.txt
			  cryps -l key.json -d -f response.txt`,
}

func init() {
	rootCmd.Flags().BoolVarP(&library, "", "l", false, "")
	rootCmd.Flags().BoolVarP(&file, "", "f", false, "")
	rootCmd.Flags().BoolVarP(&output, "", "o", false, "")
	rootCmd.Flags().BoolVarP(&decrypt, "", "d", false, "")
}
