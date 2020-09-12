# Cromwell CLI

Command line interface for Cromwell server.

## Development environment

1. Download Cromwell from its [GitHub repository](https://github.com/broadinstitute/cromwell/releases/tag/53.1) and start it with `java -jar cromwell-<version>.jar server`

1. Build this CLI (`go build`) and interact with the server.
  1. Submit a job with `./cromwell-cli s -w sample/wf.wdl -i sample/wf.inputs.json`
  1. Query jobs with `./cromwell-cli q`
  1. Kill running jobs with `./cromwell-cli k -o <operation>`

## Go ecosystem

- [x] Command line [urfave/cli/v2](https://github.com/urfave/cli)
- [x] Logging  [logrus](https://github.com/uber-go/zap)
- [x] Http request  [net/http](https://golang.org/pkg/net/http/)
- [x] Pretty format terminal tables [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter)

## Cromwell server interactions

- [x] Submit a job (with dependencies)
- [x] Kill a job
- [x] Query job status
- [x] Get jobs by name
- [ ] Query job outputs

