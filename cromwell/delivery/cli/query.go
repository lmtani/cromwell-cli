package cli

import (
	"os"
	"time"

	"github.com/lmtani/cromwell-cli/domain"
)

type QueryTableResponse struct {
	Results           []domain.QueryResponseWorkflow
	TotalResultsCount int
}

func (qtr QueryTableResponse) Header() []string {
	return []string{"Operation", "Name", "Start", "Duration", "Status"}
}

func (qtr QueryTableResponse) Rows() [][]string {
	rows := make([][]string, len(qtr.Results))
	timePattern := "2006-01-02 15h04m"
	for _, r := range qtr.Results {
		if r.End.IsZero() {
			r.End = time.Now()
		}
		elapsedTime := r.End.Sub(r.Start)
		rows = append(rows, []string{
			r.ID,
			r.Name,
			r.Start.Format(timePattern),
			elapsedTime.Round(time.Second).String(),
			r.Status,
		})
	}
	return rows
}

func queryTable(r domain.WorkflowQueryResponse) {
	var qtr = QueryTableResponse(r)
	NewTable(os.Stdout).Render(qtr)
}
