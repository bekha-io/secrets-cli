package commands

import (
	"fmt"

	"github.com/bekha-io/secrets/services"
	"github.com/bekha-io/secrets/utils"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)


func CreateGroup(ctx *cli.Context) {
	groupName := ctx.Args().First()
	if groupName == "" {
		groupName = ctx.String("name")
	}

	if groupName == "" {
		color.Red("Group's name cannot be empty")
		return
	}

	group, err := services.Services.CreateGroup(groupName)
	if err != nil {
		color.Red(err.Error())
		return
	}

	color.Green("Created group: %v", group.Name)
}

func DeleteGroup(ctx *cli.Context) {
	groupName := ctx.Args().First()
	if groupName == "" {
		groupName = ctx.String("name")
	}

	if groupName == "" {
		color.Red("Group's name cannot be empty")
		return
	}

	if err := services.Services.DeleteGroup(groupName); err != nil {
		color.Red(err.Error())
		return
	}

	// Also show deleted secrets and its amount
	color.Green("Successfully deleted group %v", groupName)
}

func ListGroups(ctx *cli.Context) {
	groups, err := services.Services.GetAllGroups()
	if err != nil {
		color.Red(err.Error())
		return
	}

	headers := []string{"Name", "Secrets count", "Created at"}
	rows := [][]string{}
	for _, row := range groups {
		rows = append(rows, []string{row.Name, fmt.Sprintf("%v", len(row.Secrets)), row.CreatedAt.String()})
	}

	utils.RenderTable("", headers, rows)
}


func AddCommands_Groups(app *cli.App) {

	groupNameFlag := cli.StringFlag{
			Name: "name",
			Usage: "group name",
	}

	commands := []cli.Command{
		{
			Name: "groups",
			Aliases: []string{"g", "gr"},
			Usage: "Group management commands (group contains one or more secrets and pack them together)",
			Action: ListGroups,
			Subcommands: []cli.Command{
				{
					Name: "create",
					Aliases: []string{"add"},
					Flags: []cli.Flag{groupNameFlag},
					Action: CreateGroup,
				},
				{
					Name: "delete",
					Aliases: []string{"rm", "remove"},
					Flags: []cli.Flag{groupNameFlag},
					Action: DeleteGroup,
				},
				{
					Name: "list",
					Aliases: []string{"show", "l", "ls"},
					Flags: []cli.Flag{groupNameFlag},
					Action: ListGroups,
				},
			},
		},
	}
	app.Commands = append(app.Commands, commands...)
}