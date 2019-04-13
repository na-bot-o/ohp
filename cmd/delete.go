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

	"github.com/na-bot-o/ohp/data"
	"github.com/na-bot-o/ohp/page"
	"github.com/na-bot-o/ohp/util"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete help page url",
	Long:  `delete page url records matched name indicated tag or page flag`,
	Run: func(cmd *cobra.Command, args []string) {

		tagFlag, _ := cmd.PersistentFlags().GetString("tag")
		nameFlag, _ := cmd.PersistentFlags().GetString("name")

		if !IsTagOrPageFlagUsed(tagFlag, nameFlag) {
			fmt.Println("either page or tag flag must use")
			os.Exit(1)
		}

		const BUFSIZE = 1024

		util.LoadEnv()

		//copy file to archive for revovering
		dataFile := data.New("PAGEFILE")
		archiveFile := data.New("ARCHIVEFILE")

		dataFile.CopyTo(archiveFile)

		var fp *os.File
		fp, err := os.Create(dataFile.Path)
		defer fp.Close()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		var lines []page.Page

		lines, err = archiveFile.GetPages()

		for _, line := range lines {

			if line.Tag != tagFlag && line.Name != nameFlag {

				err = line.WrittenIn(fp)

				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("tag", "t", "", "deleted tag")
	deleteCmd.Flags().StringP("name", "n", "", "deleted page name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//check whether have tag or page flag when execute delete command
func IsTagOrPageFlagUsed(tagFlag string, pageFlag string) bool {
	if tagFlag == "" && pageFlag == "" {
		return false
	}
	if tagFlag != "" && pageFlag != "" {
		return false
	}
	return true
}
