// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestUsersV3AutocompleteProfile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "autocomplete-profile",
			"--user-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("userId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:v3", "autocomplete-profile",
		)
	})
}

func TestUsersV3DeleteBanner(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "delete-banner",
			"--banner-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3DeleteOwnUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "delete-own-user",
		)
	})
}

func TestUsersV3GetActivity(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-activity",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sort", "date",
		)
	})
}

func TestUsersV3GetClaimedPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-claimed-papers",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--sort", "date",
		)
	})
}

func TestUsersV3GetCurrentUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-current-user",
		)
	})
}

func TestUsersV3GetFeaturedActivity(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-featured-activity",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3GetFollowers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-followers",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3GetLeaderboard(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-leaderboard",
		)
	})
}

func TestUsersV3GetUserByUuid(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-user-by-uuid",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3GetViewedHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "get-viewed-history",
			"--offset", "offset",
			"--search", "search",
		)
	})
}

func TestUsersV3ProcessNotificationEmail(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "process-notification-email",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3Search(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "search",
			"--q", "x",
			"--limit", "limit",
		)
	})
}

func TestUsersV3ToggleFollowUser(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "toggle-follow-user",
			"--id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestUsersV3UpdatePreferences(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "update-preferences",
			"--banner", "{link: link, name: name, type: success}",
			"--base", "{assistantCustomStyles: [{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, instructions: x, name: x}], assistantStyleSelection: default, defaultPrivatePaperSidebarTab: assistant, defaultPublicPaperSidebarTab: comments, feedSort: Hot, isDarkModeEnabled: true, isDebugModeEnabled: true, isMembersSidebarVisible: true, preferredLanguage: am, preferredLlmFollowLatestCategory: preferredLlmFollowLatestCategory, preferredLlmModel: preferredLlmModel, preferredLlmThinking: preferredLlmThinking, readingModeEnabled: true, showModelThinking: true, toolingPaneWidth: 0, webSearch: 'off'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(usersV3UpdatePreferences)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "update-preferences",
			"--banner.link", "link",
			"--banner.name", "name",
			"--banner.type", "success",
			"--base.assistant-custom-styles", "[{id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e, instructions: x, name: x}]",
			"--base.assistant-style-selection", "default",
			"--base.default-private-paper-sidebar-tab", "assistant",
			"--base.default-public-paper-sidebar-tab", "comments",
			"--base.feed-sort", "Hot",
			"--base.is-dark-mode-enabled=true",
			"--base.is-debug-mode-enabled=true",
			"--base.is-members-sidebar-visible=true",
			"--base.preferred-language", "am",
			"--base.preferred-llm-follow-latest-category", "preferredLlmFollowLatestCategory",
			"--base.preferred-llm-model", "preferredLlmModel",
			"--base.preferred-llm-thinking", "preferredLlmThinking",
			"--base.reading-mode-enabled=true",
			"--base.show-model-thinking=true",
			"--base.tooling-pane-width", "0",
			"--base.web-search", "off",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"banners:\n" +
			"  - link: link\n" +
			"    name: name\n" +
			"    type: success\n" +
			"base:\n" +
			"  assistantCustomStyles:\n" +
			"    - id: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"      instructions: x\n" +
			"      name: x\n" +
			"  assistantStyleSelection: default\n" +
			"  defaultPrivatePaperSidebarTab: assistant\n" +
			"  defaultPublicPaperSidebarTab: comments\n" +
			"  feedSort: Hot\n" +
			"  isDarkModeEnabled: true\n" +
			"  isDebugModeEnabled: true\n" +
			"  isMembersSidebarVisible: true\n" +
			"  preferredLanguage: am\n" +
			"  preferredLlmFollowLatestCategory: preferredLlmFollowLatestCategory\n" +
			"  preferredLlmModel: preferredLlmModel\n" +
			"  preferredLlmThinking: preferredLlmThinking\n" +
			"  readingModeEnabled: true\n" +
			"  showModelThinking: true\n" +
			"  toolingPaneWidth: 0\n" +
			"  webSearch: 'off'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:v3", "update-preferences",
		)
	})
}

func TestUsersV3UpdateProfile(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "update-profile",
			"--biography", "biography",
			"--bluesky-username", "blueskyUsername",
			"--github-username", "githubUsername",
			"--institution", "institution",
			"--linkedin-username", "linkedinUsername",
			"--location", "location",
			"--public-email", "dev@stainless.com",
			"--real-name", "x",
			"--username", "username",
			"--x-username", "xUsername",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"biography: biography\n" +
			"blueskyUsername: blueskyUsername\n" +
			"githubUsername: githubUsername\n" +
			"institution: institution\n" +
			"linkedinUsername: linkedinUsername\n" +
			"location: location\n" +
			"publicEmail: dev@stainless.com\n" +
			"realName: x\n" +
			"username: username\n" +
			"xUsername: xUsername\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"users:v3", "update-profile",
		)
	})
}

func TestUsersV3UploadAvatar(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"users:v3", "upload-avatar",
		)
	})
}
