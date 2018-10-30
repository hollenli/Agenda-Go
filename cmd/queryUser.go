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
	"log"

	"github.com/Agenda-Go/entity"
	"github.com/spf13/cobra"
)

// queryUserCmd represents the queryUser command
var queryUserCmd = &cobra.Command{
	Use:   "queryUser",
	Short: "Query the information of all registered users",
	Long:  "Usage: agenda queryUser",
	Run: func(cmd *cobra.Command, args []string) {
		users := entity.GetAllUser()
		fmt.Printf("username\t mail\tphone\n")
		for i := 0; i < len(users); i++ {
			fmt.Printf("%s\t %s\t%s\n", users[i].Username, users[i].Mail, users[i].Phone)
		}
		log.Println("Query all users.")
	},
}

func init() {
	rootCmd.AddCommand(queryUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
