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

package create

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/rudeigerc/dexctl/pkg/cmd/client"

	pb "github.com/dexidp/dex/api/v2"
)

func NewCreatePasswordCommand() *cli.Command {
	return &cli.Command{
		Name:   "password",
		Usage:  "Create password",
		Action: createPasswordAction,
	}
}

func createPasswordAction(c *cli.Context) error {
	var questions = []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "Please type the username:",
			},
			Validate: survey.Required,
		},
		{
			Name: "email",
			Prompt: &survey.Input{
				Message: "Please type the E-mail address:",
			},
			Validate: survey.Required,
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Please type the password:",
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		Username string
		Email    string
		Password string
	}{}

	err := survey.Ask(questions, &answers)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(answers.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	client, conn, err := client.NewDexClient(true)
	if err != nil {
		return err
	}
	defer conn.Close()

	resp, err := client.CreatePassword(context.Background(),
		&pb.CreatePasswordReq{
			Password: &pb.Password{
				Email:    answers.Email,
				Hash:     hash,
				Username: answers.Username,
				UserId:   uuid.New().String(),
			},
		},
	)
	if err != nil {
		return err
	}
	if resp.GetAlreadyExists() {
		return fmt.Errorf("Password already exists")
	}

	return nil
}
