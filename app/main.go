package main

import (
	"fmt"
	"os"

	_cromwellRepo "github.com/lmtani/cromwell-cli/cromwell/repository/http"
	_cromwellUsecase "github.com/lmtani/cromwell-cli/cromwell/usecase"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func main() {
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
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("cromwell.command.error",
			zap.NamedError("err", err))
	}

	cromwellRepo := _cromwellRepo.NewHTTPCromwellRepository("http://127.0.0.1:8000", "")
	cromwellUsecase := _cromwellUsecase.NewCromwellUsecase(cromwellRepo)
}
