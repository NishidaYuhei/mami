// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"strings"

	"github.com/NishidaYuuhei/mami/lib"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "make path specificationv",
	Long:  `make path specificationv`,
	Run: func(cmd *cobra.Command, args []string) {
		var filePath string
		fmt.Println("Please enter the file path to save")
		fmt.Scanf("%v", &filePath)
		_, statDirErr := os.Stat(lib.GetHomePath() + "/.mami")
		filePath = strings.Replace(filePath, "~", lib.GetHomePath(), 1)
		filePath = strings.TrimSuffix(filePath, "/")
		if statDirErr != nil {
			lib.Mkdir(lib.GetHomePath() + "/.mami")
		}
		_, statFileErr := os.Stat(lib.GetHomePath() + "/.mami/config")
		if statFileErr == nil {
			lib.RemoveFile(lib.GetHomePath() + "/.mami/config")
		}
		lib.CreateNewFile(lib.GetHomePath()+"/.mami/config", filePath)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
