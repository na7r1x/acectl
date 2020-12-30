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
	"time"

	"github.com/na7r1x/acectl/internal/core/domain"
	"github.com/na7r1x/acectl/internal/core/service/brokersrv"
	"github.com/na7r1x/acectl/internal/repositories/brokerrepo"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a broker",
	Long:  `Registers a broker in persistent storage.`,
	Run: func(cmd *cobra.Command, args []string) {
		var broker domain.Broker
		broker.Id, _ = cmd.Flags().GetString("id")
		broker.Created = time.Now()
		broker.Host, _ = cmd.Flags().GetString("host")
		broker.Port, _ = cmd.Flags().GetString("port")
		broker.Username, _ = cmd.Flags().GetString("username")
		broker.Password, _ = cmd.Flags().GetString("password")
		repo := brokerrepo.NewSqliteRepo("acectl.db")
		srv := brokersrv.New(repo)
		err := srv.Register(broker)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("successfully registered broker [" + broker.Id + "]")
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("id", "i", "", "ID")
	registerCmd.Flags().StringP("host", "", "", "Host")
	registerCmd.Flags().StringP("port", "", "", "Port")
	registerCmd.Flags().StringP("username", "u", "", "Username")
	registerCmd.Flags().StringP("password", "p", "", "Password")
}
