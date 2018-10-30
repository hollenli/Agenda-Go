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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a account.",
	Long:  "Usage :agenda register -u [username]  -p [password] -e [email]  -t [telephone]    ",
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := cmd.Flags().GetString("username")
		p, _ := cmd.Flags().GetString("password")
		e, _ := cmd.Flags().GetString("email")
		t, _ := cmd.Flags().GetString("telephone")
		if entity.CreateUser(u, p, e, t) == 0 {
			entity.UpdateLib()
			log.Println("register succesfully")
			log.Println("username is " + u + " password is " + p + " email is " + e + " telethone is " + t)
		} else if entity.CreateUser(u, p, e, t) == 1 {
			log.Println("create user failed")
		} else if entity.CreateUser(u, p, e, t) == 2 {
			log.Println("username repeat")
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	entity.Init()
	registerCmd.Flags().StringP("username", "u", "", "")
	registerCmd.Flags().StringP("password", "p", "", "")
	registerCmd.Flags().StringP("email", "e", "", "")
	registerCmd.Flags().StringP("telephone", "t", "", "")
}
