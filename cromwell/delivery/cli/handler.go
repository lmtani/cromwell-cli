package cli

import (
	"github.com/lmtani/cromwell-cli/domain"
	"github.com/urfave/cli"
)

type CromwellHandler struct {
	CromwellUsecase domain.CromwellUsecase
}

func NewCromwellHandler(a *cli.App, us domain.CromwellUsecase) {
	handler := CromwellHandler{
		CromwellUsecase: us,
	}
	a.Commands = append(a.Commands, handler.Query())

}

func (c *CromwellHandler) Query() cli.Command {
	cmd := cli.Command{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "Query workflows",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Required: false},
		},
		Action: c.CromwellUsecase.Query,
	}
	return cmd

}
