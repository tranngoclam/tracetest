package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/kubeshop/tracetest/cli/actions"
	"github.com/kubeshop/tracetest/cli/analytics"
	"github.com/kubeshop/tracetest/cli/config"
	"github.com/kubeshop/tracetest/cli/formatters"
	"github.com/kubeshop/tracetest/cli/parameters"
	"github.com/kubeshop/tracetest/cli/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	cliConfig        config.Config
	cliLogger        *zap.Logger
	versionText      string
	isVersionMatch   bool
	resourceRegistry = actions.NewResourceRegistry()
	resourceParams   = &parameters.ResourceParams{}
)

type setupConfig struct {
	shouldValidateConfig          bool
	shouldValidateVersionMismatch bool
}

type setupOption func(*setupConfig)

func SkipConfigValidation() setupOption {
	return func(sc *setupConfig) {
		sc.shouldValidateConfig = false
	}
}

func SkipVersionMismatchCheck() setupOption {
	return func(sc *setupConfig) {
		sc.shouldValidateVersionMismatch = false
	}
}

func setupCommand(options ...setupOption) func(cmd *cobra.Command, args []string) {
	config := setupConfig{
		shouldValidateConfig:          true,
		shouldValidateVersionMismatch: true,
	}
	for _, option := range options {
		option(&config)
	}

	return func(cmd *cobra.Command, args []string) {
		setupOutputFormat(cmd)
		setupLogger(cmd, args)
		loadConfig(cmd, args)
		overrideConfig()
		setupVersion()

		baseOptions := []actions.ResourceArgsOption{actions.WithLogger(cliLogger), actions.WithConfig(cliConfig)}

		// TODO: remove this client from here when we migrate tests to the resource manager
		openapiClient := utils.GetAPIClient(cliConfig)

		configOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("configs", cliConfig)),
			actions.WithFormatter(formatters.NewConfigFormatter()),
		)
		configActions := actions.NewConfigActions(configOptions...)
		resourceRegistry.Register(configActions)

		analyzerOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("analyzers", cliConfig)),
			actions.WithFormatter(formatters.NewAnalyzerFormatter()),
		)
		analyzerActions := actions.NewAnalyzerActions(analyzerOptions...)
		resourceRegistry.Register(analyzerActions)

		pollingOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("pollingprofiles", cliConfig)),
			actions.WithFormatter(formatters.NewPollingFormatter()),
		)
		pollingActions := actions.NewPollingActions(pollingOptions...)
		resourceRegistry.Register(pollingActions)

		demoOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("demos", cliConfig)),
			actions.WithFormatter(formatters.NewDemoFormatter()),
		)
		demoActions := actions.NewDemoActions(demoOptions...)
		resourceRegistry.Register(demoActions)

		dataStoreOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("datastores", cliConfig)),
			actions.WithFormatter(formatters.NewDatastoreFormatter()),
		)
		dataStoreActions := actions.NewDataStoreActions(dataStoreOptions...)
		resourceRegistry.Register(dataStoreActions)

		environmentOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("environments", cliConfig)),
			actions.WithFormatter(formatters.NewEnvironmentsFormatter()),
		)
		environmentActions := actions.NewEnvironmentsActions(environmentOptions...)
		resourceRegistry.Register(environmentActions)

		transactionOptions := append(
			baseOptions,
			actions.WithClient(utils.GetResourceAPIClient("transactions", cliConfig)),
			actions.WithFormatter(formatters.NewTransactionsFormatter()),
		)
		transactionActions := actions.NewTransactionsActions(openapiClient, transactionOptions...)
		resourceRegistry.Register(transactionActions)

		if config.shouldValidateConfig {
			validateConfig(cmd, args)
		}

		if config.shouldValidateVersionMismatch {
			validateVersionMismatch()
		}

		analytics.Init()
	}
}

func overrideConfig() {
	if overrideEndpoint != "" {
		scheme, endpoint, err := config.ParseServerURL(overrideEndpoint)
		if err != nil {
			msg := fmt.Sprintf("cannot parse endpoint %s", overrideEndpoint)
			cliLogger.Error(msg, zap.Error(err))
			ExitCLI(1)
		}
		cliConfig.Scheme = scheme
		cliConfig.Endpoint = endpoint
	}
}

func setupOutputFormat(cmd *cobra.Command) {
	if cmd.GroupID != "resources" && output == string(formatters.Empty) {
		output = string(formatters.DefaultOutput)
	}

	o := formatters.Output(output)
	if !formatters.ValidOutput(o) {
		fmt.Fprintf(os.Stderr, "Invalid output format %s. Available formats are [%s]\n", output, outputFormatsString)
		ExitCLI(1)
	}
	formatters.SetOutput(o)
}

func loadConfig(cmd *cobra.Command, args []string) {
	config, err := config.LoadConfig(configFile)
	if err != nil {
		cliLogger.Fatal("could not load config", zap.Error(err))
	}

	cliConfig = config
}

func validateConfig(cmd *cobra.Command, args []string) {
	if cliConfig.IsEmpty() {
		cliLogger.Warn("You haven't configured your CLI, some commands might fail!")
		cliLogger.Warn("Run 'tracetest configure' to configure your CLI")
	}
}

func setupLogger(cmd *cobra.Command, args []string) {
	atom := zap.NewAtomicLevel()
	if verbose {
		atom.SetLevel(zap.DebugLevel)
	} else {
		atom.SetLevel(zap.WarnLevel)
	}

	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        zapcore.OmitKey,
		LevelKey:       "level",
		NameKey:        zapcore.OmitKey,
		CallerKey:      zapcore.OmitKey,
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		atom,
	))

	cliLogger = logger
}

func teardownCommand(cmd *cobra.Command, args []string) {
	cliLogger.Sync()
}

func setupVersion() {
	ctx := context.Background()
	options := []actions.ActionArgsOption{
		actions.ActionWithClient(utils.GetAPIClient(cliConfig)),
		actions.ActionWithConfig(cliConfig),
		actions.ActionWithLogger(cliLogger),
	}

	action := actions.NewGetServerVersionAction(options...)
	versionText, isVersionMatch = action.GetVersion(ctx)
}

func validateVersionMismatch() {
	if !isVersionMatch && os.Getenv("TRACETEST_DEV") == "" {
		fmt.Fprintf(os.Stderr, versionText+`
✖️ Error: Version Mismatch
The CLI version and the server version are not compatible. To fix this, you'll need to make sure that both your CLI and server are using compatible versions.
We recommend upgrading both of them to the latest available version. Check out our documentation https://docs.tracetest.io/configuration/upgrade for simple instructions on how to upgrade.
Thank you for using Tracetest! We apologize for any inconvenience caused.
`)
		ExitCLI(1)
	}
}
