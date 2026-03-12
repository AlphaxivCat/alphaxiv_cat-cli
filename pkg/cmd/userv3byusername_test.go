// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestUsersV3ByUsernameGetProfilePage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:v3:by-username", "get-profile-page",
			"--api-key", "string",
			"--username", "x",
		)
	})
}

func TestUsersV3ByUsernameGetUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users:v3:by-username", "get-user",
			"--api-key", "string",
			"--username", "x",
		)
	})
}
