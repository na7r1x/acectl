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
	"os"

	"github.com/lensesio/tableprinter"
	"github.com/na7r1x/acectl/internal/core/service/brokersrv"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered brokers",
	Long:  `The command lists all registered brokers.`,
	Run: func(cmd *cobra.Command, args []string) {
		repo := brokerrepo.NewSqliteRepo("acectl.db")
		srv := brokersrv.New(repo)
		brokers, err := srv.List()
		if err != nil {
			fmt.Println(err.Error())
		}
		tableprinter.Print(os.Stdout, brokers)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}