// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3Retrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve",
			"--api-key", "string",
			"--unresolved", "unresolved",
		)
	})
}

func TestPapersV3Comment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "comment",
			"--api-key", "string",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--tag", "anonymous",
			"--body", "body",
			"--parent", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--title", "title",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"tag: anonymous\n" +
			"body: body\n" +
			"parent: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"title: title\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "comment",
			"--api-key", "string",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3DeleteVotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "delete-votes",
			"--api-key", "string",
			"--body", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("- 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "delete-votes",
			"--api-key", "string",
		)
	})
}

func TestPapersV3Implementation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "implementation",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("url: url")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "implementation",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3KickoffPaperCountries(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "kickoff-paper-countries",
			"--api-key", "string",
			"--batch", "1",
			"--max-papers", "1",
			"--months", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"batch: 1\n" +
			"maxPapers: 1\n" +
			"months: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "kickoff-paper-countries",
			"--api-key", "string",
		)
	})
}

func TestPapersV3KickoffPaperFullText(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "kickoff-paper-full-text",
			"--api-key", "string",
			"--max-papers", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("maxPapers: 1")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "kickoff-paper-full-text",
			"--api-key", "string",
		)
	})
}

func TestPapersV3KickoffPaperPodcasts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "kickoff-paper-podcasts",
			"--api-key", "string",
		)
	})
}

func TestPapersV3KickoffThumbnailsTrendingPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "kickoff-thumbnails-trending-papers",
			"--api-key", "string",
		)
	})
}

func TestPapersV3KickoffXMentionsSync(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "kickoff-x-mentions-sync",
			"--api-key", "string",
			"--dry-run=true",
			"--limit", "1",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"dryRun: true\n" +
			"limit: 1\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "kickoff-x-mentions-sync",
			"--api-key", "string",
		)
	})
}

func TestPapersV3Like(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "like",
			"--api-key", "string",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3Podcast(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "podcast",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3ProcessAI(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "process-ai",
			"--api-key", "string",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--preferred-language", "am",
		)
	})
}

func TestPapersV3ProcessCountries(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "process-countries",
			"--api-key", "string",
			"--universal-paper-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"universalPaperIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "process-countries",
			"--api-key", "string",
		)
	})
}

func TestPapersV3ProcessFullText(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "process-full-text",
			"--api-key", "string",
			"--paper-version-id", "paperVersionId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("paperVersionId: paperVersionId")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "process-full-text",
			"--api-key", "string",
		)
	})
}

func TestPapersV3PruneEmbeddingsByDate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "prune-embeddings-by-date",
			"--api-key", "string",
		)
	})
}

func TestPapersV3RequestImplementation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "request-implementation",
			"--api-key", "string",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--paper-title", "paperTitle",
			"--universal-paper-id", "universalPaperId",
			"--additional-info", "additionalInfo",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperTitle: paperTitle\n" +
			"universalPaperId: universalPaperId\n" +
			"additionalInfo: additionalInfo\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v3", "request-implementation",
			"--api-key", "string",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RequestPodcast(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "request-podcast",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-all",
			"--api-key", "string",
			"--limit", "limit",
			"--skip", "skip",
		)
	})
}

func TestPapersV3RetrieveDiversePapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-diverse-papers",
			"--api-key", "string",
			"--topics", "topics",
		)
	})
}

func TestPapersV3RetrieveFeed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-feed",
			"--api-key", "string",
			"--interval", "3 Days",
			"--page-num", "pageNum",
			"--page-size", "pageSize",
			"--sort", "Hot",
			"--exclude-seen-briefs", "true",
			"--organizations", "organizations",
			"--require-summary", "true",
			"--source", "GitHub",
			"--topics", "topics",
			"--universal-id", "universalId",
		)
	})
}

func TestPapersV3RetrieveFigures(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-figures",
			"--api-key", "string",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveFullText(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-full-text",
			"--api-key", "string",
			"--paper-version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveGeoTrends(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-geo-trends",
			"--api-key", "string",
			"--collaboration-limit", "collaborationLimit",
			"--paper-limit", "paperLimit",
			"--past-months", "pastMonths",
			"--repo-limit", "repoLimit",
			"--top-countries", "topCountries",
		)
	})
}

func TestPapersV3RetrieveMetrics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-metrics",
			"--api-key", "string",
			"--unresolved", "unresolved",
		)
	})
}

func TestPapersV3RetrievePapersByCountry(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-papers-by-country",
			"--api-key", "string",
			"--country", "country",
			"--limit", "limit",
		)
	})
}

func TestPapersV3RetrievePreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-preview",
			"--api-key", "string",
			"--id", "id",
		)
	})
}

func TestPapersV3RetrieveSimilarPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-similar-papers",
			"--api-key", "string",
			"--id", "id",
			"--exclude", "exclude",
			"--exclude-likes", "false",
			"--interval", "3 Days",
			"--limit", "limit",
		)
	})
}

func TestPapersV3RetrieveUnrelated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "retrieve-unrelated",
			"--api-key", "string",
			"--limit", "limit",
			"--papers", "papers",
			"--topics", "topics",
		)
	})
}

func TestPapersV3View(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v3", "view",
			"--api-key", "string",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
