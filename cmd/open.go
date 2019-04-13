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
	"github.com/na-bot-o/ohp/util"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// openCmd represents the open command
var openCmd = &cobra.Command{
	Use:   "open",
	Short: "open help page",
	Long: `this command opens registered help page.
				 search tag or page indicated flag, open matched url`,
	Run: func(cmd *cobra.Command, args []string) {

		tagFlag, _ := cmd.PersistentFlags().GetString("tag")
		nameFlag, _ := cmd.PersistentFlags().GetString("name")

		if tagFlag == "" && nameFlag == "" {
			fmt.Println("tag or page flag is required")
			os.Exit(1)
		}

		util.LoadEnv()
		dataFile := data.New(os.Getenv("PAGEFILE"))

		lines, err := dataFile.GetPages()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var count_opened_page int

		for _, line := range lines {
			if line.Name == nameFlag || line.Tag == tagFlag {
				browser.OpenURL(line.Url)
				count_opened_page++
			}
		}
		fmt.Printf("%d pages opened", count_opened_page)
	},
}

func init() {
	rootCmd.AddCommand(openCmd)

	openCmd.Flags().StringP("tag", "t", "", "page tag you want to see")
	openCmd.Flags().StringP("name", "n", "", "page name you want to see")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// openCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// openCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
