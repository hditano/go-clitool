/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"
	"github/hditano/clitool/cmd/utils"
	"log"

	"github.com/spf13/cobra"
	"golang.org/x/sys/unix"
)

// diskusageCmd represents the diskusage command
var diskusageCmd = &cobra.Command{
	Use:   "diskusage",
	Short: "A command that will output Disk Usage information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Disk Stats")
		var c unix.Statfs_t
		err := unix.Statfs("/", &c)
		if err != nil {
			log.Fatal(err)
			return
		}

		total := c.Blocks * uint64(c.Bsize) / utils.GB
		free := c.Bfree * uint64(c.Bsize) / utils.GB
		used := total - free

		fmt.Printf("Total Disk Space: %d GB\n", total)
		fmt.Printf("Free Disk Space: %d GB\n", free)
		fmt.Printf("Used Disk Space: %d GB\n", used)
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
