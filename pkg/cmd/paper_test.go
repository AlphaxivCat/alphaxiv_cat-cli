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
			t, "papers", "add-author",
			"--api-key", "string",
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
			t, pipeData, "papers", "add-author",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersAdminVote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "admin-vote",
			"--api-key", "string",
			"--paper-id", "x",
			"--entry", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("entry: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers", "admin-vote",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersCrxAbstractClick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "crx-abstract-click",
			"--api-key", "string",
			"--pid", "pid",
			"--ref", "ref",
		)
	})
}

func TestPapersCrxAbstractHit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "crx-abstract-hit",
			"--api-key", "string",
			"--pid", "pid",
		)
	})
}

func TestPapersCrxPdfClick(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "crx-pdf-click",
			"--api-key", "string",
			"--pid", "pid",
			"--ref", "ref",
		)
	})
}

func TestPapersCrxPdfHit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "crx-pdf-hit",
			"--api-key", "string",
			"--pid", "pid",
		)
	})
}

func TestPapersEmailAuthor(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "email-author",
			"--api-key", "string",
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
			t, pipeData, "papers", "email-author",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersGetCrxPaperInfo(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "get-crx-paper-info",
			"--api-key", "string",
			"--pid", "pid",
		)
	})
}

func TestPapersGetPaperInfo(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "get-paper-info",
			"--api-key", "string",
			"--pid", "pid",
		)
	})
}

func TestPapersKickoffAbstractEmbed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-abstract-embed",
			"--api-key", "string",
		)
	})
}

func TestPapersKickoffAI(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-ai",
			"--api-key", "string",
		)
	})
}

func TestPapersKickoffBibtex(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-bibtex",
			"--api-key", "string",
		)
	})
}

func TestPapersKickoffGitHub(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-github",
			"--api-key", "string",
		)
	})
}

func TestPapersKickoffPaperCategorization(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-paper-categorization",
			"--api-key", "string",
			"--all", "all",
		)
	})
}

func TestPapersKickoffRecentPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "kickoff-recent-papers",
			"--api-key", "string",
		)
	})
}

func TestPapersMarkViewed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "mark-viewed",
			"--api-key", "string",
			"--upid", "x",
		)
	})
}

func TestPapersProcessAbstractEmbed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "process-abstract-embed",
			"--api-key", "string",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("paperVersionId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers", "process-abstract-embed",
			"--api-key", "string",
		)
	})
}

func TestPapersProcessMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "process-metadata",
			"--api-key", "string",
			"--metadata", "{bibtex: true, custom_categories: true, embedding: true, github: true, organizations: true, overview: true, thumbnail: true}",
			"--universal-paper-id", "universalPaperId",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(papersProcessMetadata)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "process-metadata",
			"--api-key", "string",
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
			t, pipeData, "papers", "process-metadata",
			"--api-key", "string",
		)
	})
}

func TestPapersRequestAILatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "request-ai-latest",
			"--api-key", "string",
			"--upid", "x",
			"--preferred-language", "am",
		)
	})
}

func TestPapersRequestAITranslationLatest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "request-ai-translation-latest",
			"--api-key", "string",
			"--upid", "x",
			"--language", "am",
		)
	})
}

func TestPapersSetGitHubRepository(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "set-github-repository",
			"--api-key", "string",
			"--paper-id", "x",
			"--github", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("github: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers", "set-github-repository",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersToggleFollow(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "toggle-follow",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersTranslateAIOverview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "translate-ai-overview",
			"--api-key", "string",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--language", "am",
		)
	})
}

func TestPapersUnclaim(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "unclaim",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}

func TestPapersVote(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers", "vote",
			"--api-key", "string",
			"--paper-id", "x",
		)
	})
}
