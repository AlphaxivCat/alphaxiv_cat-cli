// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestUsersPreferencesEmailUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences:email", "update",
			"--api-key", "string",
			"--direct-notifications=true",
			"--relevant-activity=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"directNotifications: true\n" +
			"relevantActivity: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:preferences:email", "update",
			"--api-key", "string",
		)
	})
}

func TestUsersPreferencesEmailGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:preferences:email", "get",
			"--api-key", "string",
		)
	})
}
