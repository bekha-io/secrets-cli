package commands

import (
	"fmt"

	"github.com/bekha-io/secrets/services"
	"github.com/bekha-io/secrets/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func CreateSecret(ctx *cli.Context) {

	if len(ctx.Args()) < 2 {
		color.Red("Arguments length must be 2 (name, value)")
		return
	}

	secretName := ctx.Args()[0]
	secretValue := ctx.Args()[1]
	isHidden := ctx.Bool("hidden")
	groupName := ctx.String("group")

	secret, err := services.Services.CreateSecret(secretName, secretValue, groupName, isHidden)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Successfully created secret %v in group %v", secret.Name, groupName)
}

func DeleteSecret(ctx *cli.Context) {

}

func ListSecrets(ctx *cli.Context) {
	unmask := ctx.Bool("unmask")
	groupName := ctx.String("group")

	secrets, err := services.Services.GetGroupSecrets(groupName)
	if err != nil {
		color.Red(err.Error())
		return
	}

	headers := []string{"Name", "Value", "Hidden"}
	rows := [][]string{}
	for _, secret := range secrets {
		var hiddenString string
		if secret.IsHidden {
			hiddenString = "X"
		}
		
		var secretValue string
		if secret.IsHidden && !unmask {
			secretValue = "*******"
		} else {
			secretValue = secret.Value
		}

		rows = append(rows, []string{secret.Name, secretValue, hiddenString})
	}
	utils.RenderTable(fmt.Sprintf("%v's group secrets", groupName), headers, rows)
}

func AddCommands_Secrets(app *cli.App) {

	isHiddenFlag := cli.BoolFlag{
		Name:  "hidden",
		Usage: "Whether a secret should be hidden while creating",
	}
	groupName := cli.StringFlag{
		Name:     "group",
		Usage:    "Group name",
		Required: true,
	}
	unmask := cli.BoolFlag{
		Name: "unmask",
		Usage: "Unmask hidden secrets",
	}

	commands := []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"add"},
			Usage:   "Create a new secret",
			Action:  CreateSecret,
			Flags:   []cli.Flag{isHiddenFlag, groupName},
		},
		{
			Name:    "delete",
			Aliases: []string{"rm"},
			Usage:   "Delete a secret",
			Action:  DeleteSecret,
		},
		{
			Name:    "list",
			Aliases: []string{"ls", "l"},
			Usage:   "Lists all secrets with the given conditions",
			Flags:   []cli.Flag{groupName, unmask},
			Action:  ListSecrets,
		},
	}
	app.Commands = append(app.Commands, commands...)
}
