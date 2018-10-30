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

// quitMeetingCmd represents the quitMeeting command
var quitMeetingCmd = &cobra.Command{
	Use:   "quitMeeting",
	Short: "Quit a meeting you participated in",
	Long:  "Usage：agenda quitMeeting -t [title]",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("quitMeeting called")
		title, _ := cmd.Flags().GetString("title")
		entity.Init()
		curUser := entity.GetCurrentUser()
		if curUser == "" {
			log.Println("Please log in first!")
		} else {
			p := entity.DeleteMeetingParticipators(title, curUser)
			if p == 0 {
				log.Printf("Quit meeting " + title + " successfully")
				entity.UpdateLib()
			} else if p == 1 {
				log.Println("No meeting or " + curUser + " is not a participator")
			} else {
				log.Println("No user call " + curUser)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(quitMeetingCmd)
	quitMeetingCmd.Flags().StringP("title", "t", "", "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// quitMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// quitMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
