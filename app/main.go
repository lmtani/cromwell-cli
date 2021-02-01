package main

import (
	_cromwellDelivery "github.com/lmtani/cromwell-cli/cromwell/delivery/cli"
	_cromwellRepo "github.com/lmtani/cromwell-cli/cromwell/repository/http"
	_cromwellUsecase "github.com/lmtani/cromwell-cli/cromwell/usecase"
)

func main() {
	cromwellRepo := _cromwellRepo.NewHTTPCromwellRepository("http://127.0.0.1:8000", "")
	cromwellUsecase := _cromwellUsecase.NewCromwellUsecase(cromwellRepo)
	_cromwellDelivery.NewCromwellHandler(cromwellUsecase)
}
