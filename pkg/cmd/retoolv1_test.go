// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestRetoolV1GetCumulativeUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-cumulative-users",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetDailyConversations(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-daily-conversations",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetDailyNewAccounts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-daily-new-accounts",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetDailyUserChatMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-daily-user-chat-messages",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetTotalCommentCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-total-comment-count",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetTotalPaperCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-total-paper-count",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetTotalPrivateNotesCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-total-private-notes-count",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetTotalUserCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-total-user-count",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetWeeklyMessageCountsByUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-weekly-message-counts-by-user",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetWeeklyPrivateNotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-weekly-private-notes",
			"--api-key", "string",
		)
	})
}

func TestRetoolV1GetWeeklyPublicComments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "retool:v1", "get-weekly-public-comments",
			"--api-key", "string",
		)
	})
}
