/*
Copyright 2022 Yuchen Cheng.

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

package list

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/rudeigerc/dexctl/pkg/cmd/client"

	pb "github.com/dexidp/dex/api/v2"
)

func NewListPasswordsCommand() *cli.Command {
	return &cli.Command{
		Name:   "passwords",
		Usage:  "List password",
		Action: listPasswordsAction,
	}
}

func listPasswordsAction(c *cli.Context) error {

	client, conn, err := client.NewDexClient(true)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.ListPasswords(context.Background(),
		&pb.ListPasswordReq{},
	)
	if err != nil {
		return err
	}
	passwords := resp.GetPasswords()
	fmt.Println(passwords)

	return nil
}
