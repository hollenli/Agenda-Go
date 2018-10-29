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

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		nowUser := entity.GetCurrentUser() 
		if nowUser== "" {
			log.Println("create meeting failed , you haven't logged in" )
		}else{
			nowStr := ""
			var participators []string	
			t , _:= cmd.Flags().GetString("title") 
			p , _:= cmd.Flags().GetString("participator") 
			s , _:= cmd.Flags().GetString("start") 
			e , _:= cmd.Flags().GetString("end") 
			for i := 0; i < len(p); i++{
				if p[i] == ','{
					participators = append(participators , nowStr)
					nowStr = ""
				}else{
					nowStr += string(p[i])
				}
			}
			participators = append(participators , nowStr)
			if !entity.CheckDateValid( entity.StrToDate(s) ){
				log.Println("start time is not vaild" )
			}else if !entity.CheckDateValid( entity.StrToDate(e) ){
				log.Println("end time is not valid")
			}else if s > e{
				log.Println("start time must be not bigger than end time")
			}else if entity.MeetingCheck(t) != -1 {
				log.Println("the meeting title is repeat")
			}else{
				canCreate := true
				if !entity.CheckUserFreeTime(nowUser, s  , e)  {
					log.Println("sponsor " + nowUser + " have meeting during this period ")
					canCreate = false
				}
				for i := 0; i < len(participators); i++ {
					if !entity.CheckUserFreeTime(participators[i] , s  , e)  {
						log.Println("participator " + participators[i] + " have meeting during this period ")
						canCreate = false
					}
				}
				if canCreate{
					log.Println("create Meeting succesfully ")
					entity.CreateMeeting(t , nowUser , s , e , participators)
					entity.UpdateLib()
				}
			}
		}
		
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)
	entity.Init()
	createMeetingCmd.Flags().StringP("title", "t", "", "")
	createMeetingCmd.Flags().StringP("participator", "p", "", "")
	createMeetingCmd.Flags().StringP("start", "s", "", "")
	createMeetingCmd.Flags().StringP("end", "e", "", "")
}
