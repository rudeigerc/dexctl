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

package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/rudeigerc/dexctl/pkg/cmd/version"
)

func main() {
	app := &cli.App{
		Name:  "dexctl",
		Usage: "A command line tool for Dex IdP gRPC interface",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "insercure",
				Value: true,
				Usage: "Disable server certificate verification",
			},
		},
		Commands: []*cli.Command{
			version.NewVersionCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
