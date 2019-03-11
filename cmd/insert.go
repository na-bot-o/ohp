// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
)

var (
	tag  string
	page string
	url  string
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("insert called")
		home, err := homedir.Dir()
		filepath := home + "/.ohp"

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		file, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, 0755)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()

		output := page + "," + tag + "," + url + "\n"

		_, err = file.Write(([]byte)(output))

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	insertCmd.Flags().StringVarP(&page, "page", "p", "", "page names")
	insertCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag name")
	insertCmd.Flags().StringVarP(&url, "url", "u", "", "page url")

	err := insertCmd.MarkFlagRequired("page")
	if err != nil {
		log.Println("page name is required")
		os.Exit(1)
	}

	err = insertCmd.MarkFlagRequired("url")
	if err != nil {
		log.Println("url is required")
		os.Exit(1)
	}
}
