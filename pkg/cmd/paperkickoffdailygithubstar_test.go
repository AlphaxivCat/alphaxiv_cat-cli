// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersKickoffDailyGitHubStarsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:kickoff-daily-github-stars", "update",
			"--max", "max",
		)
	})
}

func TestPapersKickoffDailyGitHubStarsKickoffDailyGitHubStars(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:kickoff-daily-github-stars", "kickoff-daily-github-stars",
		)
	})
}
