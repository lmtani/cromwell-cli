package cli

import (
	"github.com/lmtani/cromwell-cli/domain"
	"github.com/urfave/cli/v2"
)

type CromwellHandler struct {
	CromwellUsecase domain.CromwellUsecase
}

func NewCromwellHandler(app *cli.App, us domain.CromwellUsecase) {
	handler := CromwellHandler{
		CromwellUsecase: us,
	}

	app.Commands = []*cli.Command{{
		Name:    "query",
		Aliases: []string{"q"},
		Usage:   "Query workflows",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Aliases: []string{"n"}, Required: false},
		},
		Action: handler.QueryWorkflow,
	}}
}

func (h *CromwellHandler) QueryWorkflow(c *cli.Context) error {
	_, err := h.CromwellUsecase.Query()
	return err
}
