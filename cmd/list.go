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
	"os"

	"github.com/na-bot-o/ohp/data"
	"github.com/na-bot-o/ohp/page"
	"github.com/na-bot-o/ohp/util"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "show recorded page list",
	Long: `this command display recorded page list
					you can find what page inserted`,
	Run: func(cmd *cobra.Command, args []string) {

		util.LoadEnv()

		fmt.Println("list called")

		dataFile := data.New(os.Getenv("PAGEFILE"))

		var lines []page.Page

		lines, err := page.GetRows(filePath)

		if err != nil {
			os.Exit(1)
		}

		fmt.Println("-------------------------------------------------")
		fmt.Println("|   " + "page" + "   |   " + "tag" + "   |   " + "url" + "   |   ")
		fmt.Println("-------------------------------------------------")

		for _, line := range lines {

			data := strings.Split(string(line), ",")

			page := data[0]
			tag := data[1]
			url := data[2]

			fmt.Println("| " + page + " | " + tag + " | " + url + " | ")
			fmt.Println("----------------------------------------------")

		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
