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

	"github.com/MrFive5555/GO_agenda/entity"
	"github.com/spf13/cobra"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "log out from Agenda",
	Long: `log out from Agenda. You can only log in or 
register when you aren't log in
`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLog("[command] logout ")

		var state entity.LogState
		entity.GetLogState(&state)
		if state.HasLogin {
			fmt.Printf("[success] account (%s) has been logged out\n", state.UserName)
			state.UserName = ""
			state.HasLogin = false
			entity.SetLogState(&state)
		} else {
			fmt.Printf("[fail] you haven't loged in\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
