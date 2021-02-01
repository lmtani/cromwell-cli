package usecase

import "github.com/lmtani/cromwell-cli/refactor/domain"

type cromwellUsecase struct {
	cromwellServer domain.CromwellRepository
}

func NewCromwellUsecase(c domain.CromwellRepository) domain.CromwellUsecase {
	return &cromwellUsecase{cromwellServer: c}
}

func (c cromwellUsecase) Query() (domain.WorkflowQueryResponse, error) {
	return c.cromwellServer.Query()
}
