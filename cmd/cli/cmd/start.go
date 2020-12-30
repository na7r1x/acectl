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

	"github.com/na7r1x/acectl/internal/core/service/brokersrv"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start broker",
	Long:  `Connects to the host of the specified broker and attempts to start said broker, reporting result.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		brokerId := args[0]
		timeout, _ := cmd.Flags().GetInt("timeout")
		repo := brokerrepo.NewSqliteRepo("acectl.db")
		srv := brokersrv.New(repo)
		response, exitState, err := srv.Start(brokerId, timeout)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if exitState != "" {
			fmt.Println(exitState)
		}
		fmt.Println(response)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntP("timeout", "t", 60, "Timeout")

}
