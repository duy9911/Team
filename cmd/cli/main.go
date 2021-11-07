package main

import (
	"log"
	"os"
	"sort"

	"github.com/duy9911/Team/handler"
	"github.com/duy9911/Team/models"
	"github.com/urfave/cli/v2"
)

type TeamReceive struct {
	ID    string
	Name  string
	Staff string
}

func main() {
	teamrc := TeamReceive{}
	team := models.Team{}
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "create team, eg:  -n <name> -d <1999-09-15> -s <153.0> create  ",
				Action: func(c *cli.Context) error {
					handler.CreateTeam(team)
					return nil
				},
			},
			{
				Name:    "getall",
				Aliases: []string{"g"},
				Usage:   "return all staffs, eg: getall <domain> ",
				Action: func(c *cli.Context) error {
					domain := c.Args().First()
					handler.ReturnTeams(domain)
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update, eg: -flag update <team_id> ",
				Action: func(c *cli.Context) error {
					key := c.Args().First()
					handler.UpdateTeam(key, team)
					return nil
				},
			},
			{
				Name:    "delete",
				Aliases: []string{"d"},
				Usage:   "delete, eg: delete <team_id>",
				Action: func(c *cli.Context) error {
					key := c.Args().First()
					handler.Deletestaff(key)
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       " ",
				Usage:       "name for team",
				Destination: &team.Name,
			},
			&cli.StringFlag{
				Name:        "staffs",
				Aliases:     []string{"s"},
				Value:       " ",
				Usage:       "staffs for your team, eg: '/staff1/staff2/staff3'",
				Destination: &teamrc.Staff,
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
