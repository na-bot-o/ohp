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
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		var (
			tag  string
			page string
		)
		fmt.Println(tag + " " + page)

		const BUFSIZE = 1024

		var file *os.File
		home, err := homedir.Dir()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		filepath := home + "/.ohp"
		old_filepath := home + "/.ohp_old"

		_, err = os.Stat(filepath)

		if err != nil {
			log.Fatal(err)
			fmt.Println("no recorded in favorite page")
			os.Exit(1)
		}

		ArchiveFile(filepath, old_filepath)

		var lines []string

		lines, err = GetFileData(old_filepath)

		fmt.Println(lines)

		file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		for _, line := range lines {

			data := strings.Split(string(line), ",")
			fmt.Println(data)

			tag := data[1]
			//	url := data[2]
			fmt.Println(tag)

			if tag != "yahoo" {
				_, err = file.Write(([]byte)(line + "\n"))
			}

			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&tag, "tag", "t", "", "deleted tag")
	deleteCmd.Flags().StringVarP(&page, "page", "p", "", "deleted page")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//Archive .ohp file for recovering
func ArchiveFile(filepath string, old_filepath string) {

	old_file, err := os.Create(old_filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	file, err := os.OpenFile(filepath, os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	_, err = io.Copy(old_file, file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}
