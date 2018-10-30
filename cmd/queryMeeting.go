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

// queryMeetingCmd represents the queryMeeting command
var queryMeetingCmd = &cobra.Command{
	Use:   "queryMeeting",
	Short: "A brief description of your command",
	Long:  "Usage：agenda queryMeeting -s [start time]  -e [end time]",
	Run: func(cmd *cobra.Command, args []string) {
		start, _ := cmd.Flags().GetString("start")
		end, _ := cmd.Flags().GetString("end")
		meetings := entity.QueryMeetings(entity.GetCurrentUser(), start, end)
		fmt.Printf("title\t startTime\tendTime\tsponsor\tparticipators\n")
		for i := 0; i < len(meetings); i++ {
			fmt.Printf("%s\t %s\t%s\t%s\t", meetings[i].Title, meetings[i].StartTime, meetings[i].EndTime, meetings[i].Sponsor)
			for j := 0; j < len(meetings[i].Participators); j++ {
				fmt.Printf("%s", meetings[i].Participators[j])
				if j == len(meetings[i].Participators[j])-1 {
					fmt.Printf("\t\n")
				} else {
					fmt.Printf("&")
				}
			}
		}
		log.Printf("Query meetings of %s.", entity.GetCurrentUser())
	},
}

func init() {
	rootCmd.AddCommand(queryMeetingCmd)
	entity.Init()
	queryMeetingCmd.Flags().StringP("start", "s", "", "")
	queryMeetingCmd.Flags().StringP("end", "e", "", "")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryMeetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryMeetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
