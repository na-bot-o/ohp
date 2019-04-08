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
	"log"
	"os"

	"github.com/na-bot-o/ohp/file"
	"github.com/na-bot-o/ohp/util"

	"github.com/spf13/cobra"
)

var (
	tag  string
	name string
	url  string
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert help page",
	Long: `this command registeres help page
				flag is page name, tag name, and page url
				you need page name and url name`,
	Run: func(cmd *cobra.Command, args []string) {

		util.LoadEnv()

		dataFile := file.New(os.Getenv("PAGEFILE"))

		filePath, err := dataFile.GetPath()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		fp, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0755)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer fp.Close()

		output := name + "," + tag + "," + url + "\n"

		_, err = fp.Write(([]byte)(output))

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
	insertCmd.Flags().StringVarP(&name, "name", "n", "", "page names")
	insertCmd.Flags().StringVarP(&tag, "tag", "t", "", "tag name")
	insertCmd.Flags().StringVarP(&url, "url", "u", "", "page url")

	err := insertCmd.MarkFlagRequired("name")
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
