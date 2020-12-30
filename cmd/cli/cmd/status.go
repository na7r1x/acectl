/*
Copyright © 2020 Ivanin Ivanov <ivanin.val.ivanov@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/na7r1x/acectl/internal/core/service/brokersrv"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report broker status",
	Long:  `Connects to the host of the specified broker and attempts to determine the status of said broker, reporting result.`,
	Run: func(cmd *cobra.Command, args []string) {
		brokerId := args[0]
		timeout, _ := cmd.Flags().GetInt("timeout")
		verbose, _ := cmd.Flags().GetBool("verbose")
		repo := brokerrepo.NewSqliteRepo("acectl.db")
		srv := brokersrv.New(repo)
		response, exitState, err := srv.Status(brokerId, timeout)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if exitState != "" {
			fmt.Println(exitState)
		}
		if strings.Contains(response, "BIP1284") { // running
			fmt.Println("RUNNING")
		} else if strings.Contains(response, "BIP1285") { // stopped
			fmt.Println("STOPPED")
		} else {
			fmt.Println("INDETERMINATE ")
		}
		if verbose {
			fmt.Println(response)
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	statusCmd.Flags().IntP("timeout", "t", 60, "Timeout")
	statusCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

}
