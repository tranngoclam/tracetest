package cmd

import (
	"context"

	"github.com/kubeshop/tracetest/cli/actions"
	"github.com/kubeshop/tracetest/cli/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var testListCmd = &cobra.Command{
	Use:    "list",
	Short:  "List all tests",
	Long:   "List all tests",
	PreRun: setupCommand(),
	Run: WithResultHandler(func(_ *cobra.Command, _ []string) (string, error) {
		ctx := context.Background()
		cliLogger.Debug("Retrieving list of tests", zap.String("endpoint", cliConfig.Endpoint))
		client := utils.GetAPIClient(cliConfig)
		listTestsAction := actions.NewListTestsAction(cliConfig, cliLogger, client)

		actionArgs := actions.ListTestConfig{}
		err := listTestsAction.Run(ctx, actionArgs)
		return "", err
	}),
	PostRun: teardownCommand,
}

func init() {
	testCmd.AddCommand(testListCmd)
}
