// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestUsersV3FollowingOrganizationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:v3:following:organizations", "list",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3FollowingOrganizationsToggle(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:v3:following:organizations", "toggle",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--organization", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("organization: x")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "users:v3:following:organizations", "toggle",
			"--api-key", "string",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
