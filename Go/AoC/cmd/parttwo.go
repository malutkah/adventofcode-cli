/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"AoC/Helper"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var partTwoCmd = &cobra.Command{
	Use:   "partTwo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runPartTwoCmd,
}

func runPartTwoCmd(cmd *cobra.Command, args []string) {
	if len(args) == 2 {
		Helper.AddPart2(args[0], args[1])
	} else {
		fmt.Println("Invalid arguments! Enter only 2")
	}
}

func init() {
	rootCmd.AddCommand(partTwoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
