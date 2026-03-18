// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestRetoolV1GetCumulativeUsers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-cumulative-users",
		)
	})
}

func TestRetoolV1GetDailyConversations(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-daily-conversations",
		)
	})
}

func TestRetoolV1GetDailyNewAccounts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-daily-new-accounts",
		)
	})
}

func TestRetoolV1GetDailyUserChatMessages(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-daily-user-chat-messages",
		)
	})
}

func TestRetoolV1GetTotalCommentCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-total-comment-count",
		)
	})
}

func TestRetoolV1GetTotalPaperCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-total-paper-count",
		)
	})
}

func TestRetoolV1GetTotalPrivateNotesCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-total-private-notes-count",
		)
	})
}

func TestRetoolV1GetTotalUserCount(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-total-user-count",
		)
	})
}

func TestRetoolV1GetWeeklyMessageCountsByUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-weekly-message-counts-by-user",
		)
	})
}

func TestRetoolV1GetWeeklyPrivateNotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-weekly-private-notes",
		)
	})
}

func TestRetoolV1GetWeeklyPublicComments(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"retool:v1", "get-weekly-public-comments",
		)
	})
}
