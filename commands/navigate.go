package commands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func selectDesiredTask(c map[string][]CallItem) (string, error) {
	taskOptions := []string{}
	for key := range c {
		taskName := strings.Split(key, ".")[1]
		if !contains(taskOptions, taskName) {
			taskOptions = append(taskOptions, taskName)
		}
	}
	prompt := promptui.Select{
		Label: "Select a task",
		Items: taskOptions,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}
	return result, nil
}

func selectDesiredShard(shards []CallItem) (CallItem, error) {
	if len(shards) == 1 {
		return shards[0], nil
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "✔ {{ .ShardIndex  | green }} ({{ .ExecutionStatus | red }}) CallCaching: {{ .CallCaching.Hit}}",
		Inactive: "  {{ .ShardIndex | faint }} ({{ .ExecutionStatus | red }})",
		Selected: "✔ {{ .ShardIndex | green }}",
	}

	searcher := func(input string, index int) bool {
		shard := shards[index]
		name := strconv.Itoa(shard.ShardIndex)
		return name == input
	}

	prompt := promptui.Select{
		Label:     "Witch shard?",
		Items:     shards,
		Templates: templates,
		Size:      6,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return CallItem{}, err
	}

	return shards[i], err
}

func Navigate(c *cli.Context) error {
	cromwellClient := FromInterface(c.Context.Value("cromwell"))
	resp, err := cromwellClient.Metadata(c.String("operation"))
	if err != nil {
		return err
	}
	task, err := selectDesiredTask(resp.Calls)
	if err != nil {
		return err
	}
	selectedTask := resp.Calls[fmt.Sprintf("%s.%s", resp.WorkflowName, task)]
	item, err := selectDesiredShard(selectedTask)
	if err != nil {
		return err
	}

	fmt.Printf("Command status: %s\n", item.ExecutionStatus)
	if item.ExecutionStatus == "QueuedInCromwell" {
		return nil
	}
	if item.CallCaching.Hit {
		color.Cyan(item.CallCaching.Result)
	} else {
		color.Cyan(item.CommandLine)
	}

	fmt.Printf("Logs:\n")
	color.Cyan("%s\n%s\n", item.Stderr, item.Stdout)
	if item.MonitoringLog != "" {
		color.Cyan("%s\n", item.MonitoringLog)
	}

	fmt.Printf("🐋 Docker image:\n")
	color.Cyan("%s\n", item.RuntimeAttributes.Docker)
	return nil
}