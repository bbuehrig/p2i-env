/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"log"

	"github.com/atotto/clipboard"
	"github.com/bbuehrig/p2i-env/pkg/insomnia"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Printouts the converted Postman-environment-file",
	Long: `This command reads the given Postman-environment-file and converts it to Insomnia-JSON-format. This will print out to the console to paste
	it in Insomnia. It will be copied to the clipboard to paste it directly to Insomnia`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("not enough arguments")
			os.Exit(1)
		}

		jsonContent, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalf("cannot open file: %v\n", err)
			os.Exit(2)
		}

		insomniaEnvironment, err := insomnia.ConvertPostmanEnvironment(jsonContent)
		if err != nil {
			log.Fatalf("cannot convert postman-json: %v\n", err)
			os.Exit(3)
		}

		clipboard.WriteAll(insomniaEnvironment)
		fmt.Println(insomniaEnvironment)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// convertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// convertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
