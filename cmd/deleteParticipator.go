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

	entity "github.com/Agenda-Go/entity"
	"github.com/spf13/cobra"
)

// deleteParticipatorCmd represents the deleteParticipator command
var deleteParticipatorCmd = &cobra.Command{
	Use:   "deleteParticipator",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("deleteParticipator called")
		title, _ := cmd.Flags().GetString("title")
		parti, _ := cmd.Flags().GetString("participator")
		entity.Init()
		p := entity.DeleteMeetingParticipators(title, parti)
		if p == 0 {
			log.Println("Delete participator " + parti + " successfully")
			entity.UpdateLib()
		} else if p == 1 {
			log.Println("No meeting or " + parti + " is not a participator")
		} else {
			log.Println("No user call " + parti)
		}

	},
}

func init() {
	rootCmd.AddCommand(deleteParticipatorCmd)
	deleteParticipatorCmd.Flags().StringP("title", "t", "", "")
	deleteParticipatorCmd.Flags().StringP("participator", "p", "", "")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteParticipatorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteParticipatorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
