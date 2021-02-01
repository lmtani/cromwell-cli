package cli

import (
	"fmt"
	"os"

	"github.com/lmtani/cromwell-cli/domain"
	"github.com/urfave/cli/v2"
)

type CromwellHandler struct {
	CromwellUsecase domain.CromwellUsecase
}

func NewCromwellHandler(us domain.CromwellUsecase) {
	handler := CromwellHandler{
		CromwellUsecase: us,
	}

	app := &cli.App{
		Name:  "cromwell-cli",
		Usage: "Command line interface for Cromwell Server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Aliases:  []string{"t"},
				Required: false,
				Usage:    "Bearer token to be included in HTTP requsts",
			},
			&cli.StringFlag{
				Name:  "host",
				Value: "http://127.0.0.1:8000",
				Usage: "Url for your Cromwell Server",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Query workflows",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: false},
				},
				Action: handler.QueryWorkflow,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("cromwell.command.error: ", err)
	}
}

func (h *CromwellHandler) QueryWorkflow(c *cli.Context) error {
	_, err := h.CromwellUsecase.Query()
	return err
}
