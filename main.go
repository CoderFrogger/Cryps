package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"cryps/cypher"
)

var (
	library string
	file    string
	output  string
	decrypt bool
)

var rootCmd = &cobra.Command{
	Use:   "cryps [flags] <text>",
	Short: "Use a defined library to encrypt your text",
	Long: `A CLI program that takes a user defined encryption library json file and encrypts text using said library
	Encrypts the given text using the key file passed via -l/--library
	and prints the textInput to stdout.

	Examples:
	  cryps -l key.json "hello there"
	  cryps -l key.json -f message.txt
	  cryps -l key.json -f message.txt -o out.txt
	  cryps -l key.json -d -f response.txt`,
	Args: cobra.ArbitraryArgs,
	Run:  runCryps,
}

func init() {
	rootCmd.Flags().
		StringVarP(&library, "library", "l", "", "(required) Path to the encryption key/library JSON file")
	rootCmd.Flags().
		StringVarP(&file, "file", "f", "", "Read input from a file instead of the command line")
	rootCmd.Flags().
		StringVarP(&output, "output", "o", "", "Write output to this file instead of printing to the terminal (only used with -f)")
	rootCmd.Flags().
		BoolVarP(&decrypt, "decrypt", "d", false, "Decrypt instead of encrypt")
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
			os.Exit(1)
		}
		textInput = input
	}

	result := cypher.Encrypt(textInput)

	fmt.Println(result)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
