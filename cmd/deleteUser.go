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
	"log"

	"github.com/Agenda-Go/entity"
	"github.com/spf13/cobra"
)

// deleteUserCmd represents the deleteUser command
var deleteUserCmd = &cobra.Command{
	Use:   "deleteUser",
	Short: "Delete you account",
	Long:  "Usage：agenda deleteUser",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("deleteUser called")
		entity.Init()
		if entity.GetCurrentUser() == "" {
			log.Println("Please log in first")
		} else {
			flag1 := entity.DeleteAllMeeting(entity.GetCurrentUser())
			flag2 := entity.DeleteUser(entity.GetCurrentUser())
			if flag1 == 0 && flag2 {
				entity.SetCurrentUser("")
				entity.UpdateLib()
				log.Println("Delete user " + entity.GetCurrentUser() + " successfully")
			} else {
				log.Println("Delete user " + entity.GetCurrentUser() + " failed")
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(deleteUserCmd)
}
