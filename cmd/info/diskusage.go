/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// diskusageCmd represents the diskusage command
var diskusageCmd = &cobra.Command{
	Use:   "diskusage",
	Short: "A command that will output Disk Usage information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("diskusage called")
	},
}

func init() {
	InfoCmd.AddCommand(diskusageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskusageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskusageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
