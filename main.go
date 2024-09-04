/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github/hditano/clitool/cmd"
	"github/hditano/clitool/initializers"
)

func init() {
	initializers.DbConnection()
	initializers.DbMigrate()
	initializers.AddData()
}

func main() {
	cmd.Execute()
}
