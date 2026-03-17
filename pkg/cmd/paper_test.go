// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestPapersAddAuthor(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "add-author",
			"--paper-id", "x",
			"--author-email", "dev@stainless.com",
			"--should-email=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorEmail: dev@stainless.com\n" +
			"shouldEmail: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "add-author",
			"--paper-id", "x",
		)
	})
}

func TestPapersAdminVote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "admin-vote",
			"--paper-id", "x",
			"--entry", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("entry: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "admin-vote",
			"--paper-id", "x",
		)
	})
}

func TestPapersCrxAbstractClick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "crx-abstract-click",
			"--pid", "pid",
			"--ref", "ref",
		)
	})
}

func TestPapersCrxAbstractHit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "crx-abstract-hit",
			"--pid", "pid",
		)
	})
}

func TestPapersCrxPdfClick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "crx-pdf-click",
			"--pid", "pid",
			"--ref", "ref",
		)
	})
}

func TestPapersCrxPdfHit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "crx-pdf-hit",
			"--pid", "pid",
		)
	})
}

func TestPapersEmailAuthor(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "email-author",
			"--paper-id", "x",
			"--author-individual-email", "dev@stainless.com",
			"--email-author-name", "emailAuthorName",
			"--paper-name", "paperName",
			"--type", "comment",
			"--ignore-duplicate-error=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"authorIndividualEmail: dev@stainless.com\n" +
			"emailAuthorName: emailAuthorName\n" +
			"paperName: paperName\n" +
			"type: comment\n" +
			"ignoreDuplicateError: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "email-author",
			"--paper-id", "x",
		)
	})
}

func TestPapersGetCrxPaperInfo(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "get-crx-paper-info",
			"--pid", "pid",
		)
	})
}

func TestPapersGetPaperInfo(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "get-paper-info",
			"--pid", "pid",
		)
	})
}

func TestPapersKickoffAbstractEmbed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-abstract-embed",
		)
	})
}

func TestPapersKickoffAI(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-ai",
		)
	})
}

func TestPapersKickoffBibtex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-bibtex",
		)
	})
}

func TestPapersKickoffGitHub(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-github",
		)
	})
}

func TestPapersKickoffPaperCategorization(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-paper-categorization",
			"--all", "all",
		)
	})
}

func TestPapersKickoffRecentPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "kickoff-recent-papers",
		)
	})
}

func TestPapersMarkViewed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "mark-viewed",
			"--upid", "x",
		)
	})
}

func TestPapersProcessAbstractEmbed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "process-abstract-embed",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("paperVersionId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "process-abstract-embed",
		)
	})
}

func TestPapersProcessMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "process-metadata",
			"--metadata", "{bibtex: true, custom_categories: true, embedding: true, github: true, organizations: true, overview: true, thumbnail: true}",
			"--universal-paper-id", "universalPaperId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(papersProcessMetadata)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "process-metadata",
			"--metadata.bibtex=true",
			"--metadata.custom-categories=true",
			"--metadata.embedding=true",
			"--metadata.github=true",
			"--metadata.organizations=true",
			"--metadata.overview=true",
			"--metadata.thumbnail=true",
			"--universal-paper-id", "universalPaperId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  bibtex: true\n" +
			"  custom_categories: true\n" +
			"  embedding: true\n" +
			"  github: true\n" +
			"  organizations: true\n" +
			"  overview: true\n" +
			"  thumbnail: true\n" +
			"universalPaperId: universalPaperId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "process-metadata",
		)
	})
}

func TestPapersRequestAILatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "request-ai-latest",
			"--upid", "x",
			"--preferred-language", "am",
		)
	})
}

func TestPapersRequestAITranslationLatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "request-ai-translation-latest",
			"--upid", "x",
			"--language", "am",
		)
	})
}

func TestPapersSetGitHubRepository(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "set-github-repository",
			"--paper-id", "x",
			"--github", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("github: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers", "set-github-repository",
			"--paper-id", "x",
		)
	})
}

func TestPapersToggleFollow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "toggle-follow",
			"--paper-id", "x",
		)
	})
}

func TestPapersTranslateAIOverview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "translate-ai-overview",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--language", "am",
		)
	})
}

func TestPapersUnclaim(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "unclaim",
			"--paper-id", "x",
		)
	})
}

func TestPapersVote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers", "vote",
			"--paper-id", "x",
		)
	})
}
