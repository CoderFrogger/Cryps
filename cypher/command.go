package cypher

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	library string
	file    string
	output  string
	decrypt bool
)

var rootCmd = &cobra.Command{
	Use:   "cryps -l <path_to_library>",
	Short: "Use a defined library to encrypt your text",
	Long: `A CLI program that takes a user defined encryption library json file and encrypts text using said library

			Usage:
	  		cryps [flags] <text>

			Encrypts the given text using the key file passed via -l/--library
			and prints the textInput to stdout.

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
	Args: cobra.ArbitraryArgs,
	Run:  runCryps,
}

func init() {
	rootCmd.Flags().StringVar(&library, "", "l", "")
	rootCmd.Flags().StringVar(&file, "", "f", "")
	rootCmd.Flags().StringVar(&output, "", "o", "")
	rootCmd.Flags().BoolVarP(&decrypt, "", "d", false, "")
}

func runCryps(cmd *cobra.Command, args []string) {
	var text string

	if len(args) > 0 {
		text = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		text = strings.Join(lines, "\n")
	}

	if text == "" {
		fmt.Println("No input text provided")
		return
	}

	if library == "" {
		fmt.Println(
			"Library json filed not provided. Please define encryption library using the -l or --library flag",
		)
		return
	} // TODO: ADD LIBRARY STUFF

	var textInput []byte

	if file == "" {
		textInput = []byte(text)
	} else {
		input, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		textInput = input
	}

	result := Encrypt(textInput)

	fmt.Println(result)
	return
}
