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

// addParticipatorCmd represents the addParticipator command
var addParticipatorCmd = &cobra.Command{
	Use:   "addParticipator",
	Short: "Add a participator to a meeting",
	Long:  `Usage：agenda addParticipator -t [title]  -p [participator]`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetString("participator")

		switch entity.AddMeetingParticipators(title, participator) {
		case 0:
			entity.UpdateLib()
			log.Printf("User %s participates in meeting %s.", participator, title)
		case 1:
			log.Printf("Meeting %s doesn't exit or you are not the sponsor of it.", title)
		case 2:
			log.Printf("User %s doesn't exit or he/she has planned to attend another meeting at that time.", participator)
		}
	},
}

func init() {
	rootCmd.AddCommand(addParticipatorCmd)
	entity.Init()
	addParticipatorCmd.Flags().StringP("title", "t", "", "")
	addParticipatorCmd.Flags().StringP("participator", "p", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
