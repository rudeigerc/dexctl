/*
Copyright 2022 Yuchen Cheng.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	pb "github.com/dexidp/dex/api/v2"
	"github.com/rudeigerc/dexctl/pkg/cmd/client"
)

var (
	gitVersion = "0.0.0"
	buildDate  = "1970-01-01T00:00:00Z"
	gitCommit  = ""
)

func NewVersionCommand() *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print version information",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "client",
				Value: false,
				Usage: "If true, shows client version only",
			},
		},
		Action: versionAction,
	}
}

func versionAction(c *cli.Context) error {
	fmt.Println("Client Version:")
	fmt.Printf("dexctl %s commit: %s\n", gitVersion, gitCommit)

	if c.Bool("client") {
		return nil
	}

	client, conn, err := client.NewDexClient(true)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.GetVersion(context.Background(), &pb.VersionReq{})
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Println("Server Version:")
	fmt.Printf("API: %d, Server: %s\n", resp.Api, resp.Server)

	return nil
}
