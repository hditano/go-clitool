/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	// Logic
	client = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command and calls ping functionality in the Run Field
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Ping is a command to ping a url",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("Status Code Response: %v\n", resp)
		}
	},
}

func init() {
	// This adds a Flag for the specific ping command
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "thr url to ping")

	// This adds the Ping Command into the Net Palette ( clitool net <ping> )
	NetCmd.AddCommand(pingCmd)

	// This checks if the Flag above -url has been specified , if not will thrown an error
	var err error
	pingCmd.MarkFlagRequired("url")
	if err != nil {
		log.Fatal(err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
