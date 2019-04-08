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

	"github.com/na-bot-o/ohp/file"
	"github.com/na-bot-o/ohp/page"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete help page url",
	Long:  `delete page url records matched name indicated tag or page flag`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		tagFlag, _ := cmd.PersistentFlags().GetString("tag")
		nameFlag, _ := cmd.PersistentFlags().GetString("name")

		if !IsEitherFlagUsed(tagFlag, nameFlag) {
			fmt.Println("either page or tag flag must use")
			os.Exit(1)
		}

		const BUFSIZE = 1024

		dataFile := file.New("./ohp")
		archiveFile := file.New("./ohp_old")

		filePath, err := dataFile.GetPath()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		oldFilePath, err := archiveFile.GetPath()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		_, err = os.Stat(filePath)

		if err != nil {
			log.Fatal(err)
			fmt.Println("no recorded in favorite page")
			os.Exit(1)
		}

		page.ArchiveFile(filePath, oldFilePath)

		var fp *os.File
		fp, err = os.Create(filePath)
		defer fp.Close()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		var lines []string

		lines, err = data.GetRows(oldFilePath)

		for _, line := range lines {

			fields := strings.Split(string(line), ",")

			page := data.New(fields[0], fields[1], fields[2])

			if page.Tag != tagFlag && page.Name != nameFlag {
				err = data.Write(fp, line+"\n")
				//	_, err = file.Write(([]byte)(line + "\n"))

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

func GetFileData(filepath string) (lines []string, err error) {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()
	reader := bufio.NewReaderSize(file, 4096)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		lines = append(lines, string(line))
	}
	return lines, nil
}

//Archive .ohp file for recovering
func ArchiveFile(filepath string, old_filepath string) {

	old_file, err := os.Create(old_filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer old_file.Close()

	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	//file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	_, err = io.Copy(old_file, file)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func IsEitherFlagUsed(tagFlag string, pageFlag string) bool {
	if tagFlag == "" && pageFlag == "" {
		return false
	}
	if tagFlag != "" && pageFlag != "" {
		return false
	}
	return true
}
