// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestUsersGetPrivateNotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "get-private-notes",
			"--api-key", "string",
			"--uid", "uid",
			"--limit", "limit",
			"--page", "page",
		)
	})
}

func TestUsersWeighWeeklyReputation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "users", "weigh-weekly-reputation",
			"--api-key", "string",
		)
	})
}
