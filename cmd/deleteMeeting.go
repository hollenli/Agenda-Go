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

// deleteMeetingCmd represents the deleteMeeting command
var deleteMeetingCmd = &cobra.Command{
	Use:   "deleteMeeting",
	Short: "Delete a meeting you sponsored",
	Long:  "Usage：agenda deleteMeeting -t [title]",
	Run: func(cmd *cobra.Command, args []string) {
		nowUser := entity.GetCurrentUser()
		t, _ := cmd.Flags().GetString("title")
		solveResult := entity.DeleteMeeting(t, nowUser)
		if solveResult == 0 {
			log.Println("delete succesfully")
		} else if solveResult == 1 {
			log.Println("you are not the sponsor of the meeting or the meeting title is not exist")
		} else if solveResult == 2 {
			log.Println("delete failed , you haven't logged in")
		}
		entity.UpdateLib()
	},
}

func init() {
	rootCmd.AddCommand(deleteMeetingCmd)
	entity.Init()
	deleteMeetingCmd.Flags().StringP("title", "t", "", "")
}
