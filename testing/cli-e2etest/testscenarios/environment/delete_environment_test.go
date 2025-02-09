package environment

import (
	"fmt"
	"testing"

	"github.com/kubeshop/tracetest/cli-e2etest/environment"
	"github.com/kubeshop/tracetest/cli-e2etest/helpers"
	"github.com/kubeshop/tracetest/cli-e2etest/testscenarios/types"
	"github.com/kubeshop/tracetest/cli-e2etest/tracetestcli"
	"github.com/stretchr/testify/require"
)

func TestDeleteEnvironment(t *testing.T) {
	// instantiate require with testing helper
	require := require.New(t)

	// setup isolated e2e environment
	env := environment.CreateAndStart(t)
	defer env.Close(t)

	cliConfig := env.GetCLIConfigPath(t)

	// Given I am a Tracetest CLI user
	// And I have my server recently created

	// When I try to delete an environment that don't exist
	// Then it should return an error and say that this resource does not exist
	result := tracetestcli.Exec(t, "delete environment --id .env", tracetestcli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 1)
	require.Contains(result.StdErr, "Resource environments with ID .env not found") // TODO: update this message to singular

	// When I try to set up a new environment
	// Then it should be applied with success
	newEnvironmentPath := env.GetTestResourcePath(t, "new-environment")

	result = tracetestcli.Exec(t, fmt.Sprintf("apply environment --file %s", newEnvironmentPath), tracetestcli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)

	environmentVars := helpers.UnmarshalYAML[types.EnvironmentResource](t, result.StdOut)
	require.Equal("Environment", environmentVars.Type)
	require.Equal(".env", environmentVars.Spec.ID)

	// When I try to delete the environment
	// Then it should delete with success
	result = tracetestcli.Exec(t, "delete environment --id .env", tracetestcli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)
	require.Contains(result.StdOut, "✔ Environment successfully deleted")

	// When I try to get an environment again
	// Then it should return a message saying that the environment was not found
	result = tracetestcli.Exec(t, "get environment --id .env --output yaml", tracetestcli.WithCLIConfig(cliConfig))
	helpers.RequireExitCodeEqual(t, result, 0)
	require.Contains(result.StdOut, "Resource environment with ID .env not found")
}
