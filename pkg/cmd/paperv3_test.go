// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestPapersV3Retrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve",
			"--unresolved", "unresolved",
		)
	})
}

func TestPapersV3Comment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "comment",
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
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "comment",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3DeleteVotes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "delete-votes",
			"--body", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("- 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "delete-votes",
		)
	})
}

func TestPapersV3Implementation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "implementation",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--url", "url",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("url: url")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "implementation",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3KickoffPaperCountries(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "kickoff-paper-countries",
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
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "kickoff-paper-countries",
		)
	})
}

func TestPapersV3KickoffPaperPodcasts(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "kickoff-paper-podcasts",
		)
	})
}

func TestPapersV3Like(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "like",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--liked", "true",
		)
	})
}

func TestPapersV3Podcast(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "podcast",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3ProcessAI(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "process-ai",
			"--paper-version-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--preferred-language", "am",
		)
	})
}

func TestPapersV3ProcessCountries(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "process-countries",
			"--universal-paper-id", "string",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"universalPaperIds:\n" +
			"  - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "process-countries",
		)
	})
}

func TestPapersV3RequestImplementation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "request-implementation",
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
			t, pipeData,
			"--api-key", "string",
			"papers:v3", "request-implementation",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RequestPodcast(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "request-podcast",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveAll(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-all",
			"--limit", "limit",
			"--skip", "skip",
		)
	})
}

func TestPapersV3RetrieveDiversePapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-diverse-papers",
			"--topics", "topics",
			"--link-blogs", "linkBlogs",
		)
	})
}

func TestPapersV3RetrieveFeed(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-feed",
			"--interval", "3 Days",
			"--page-num", "pageNum",
			"--page-size", "pageSize",
			"--sort", "Hot",
			"--feed-cursor", "feedCursor",
			"--include-external-blogs", "includeExternalBlogs",
			"--link-blogs", "linkBlogs",
			"--runnable", "runnable",
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
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-figures",
			"--paper-group-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveFullText(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-full-text",
			"--paper-version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}

func TestPapersV3RetrieveMetrics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-metrics",
			"--unresolved", "unresolved",
		)
	})
}

func TestPapersV3RetrievePreview(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-preview",
			"--id", "id",
		)
	})
}

func TestPapersV3RetrieveSimilarPapers(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-similar-papers",
			"--id", "id",
			"--exclude", "exclude",
			"--exclude-likes", "false",
			"--interval", "3 Days",
			"--limit", "limit",
			"--link-blogs", "linkBlogs",
		)
	})
}

func TestPapersV3RetrieveUnrelated(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "retrieve-unrelated",
			"--limit", "limit",
			"--papers", "papers",
			"--topics", "topics",
			"--link-blogs", "linkBlogs",
		)
	})
}

func TestPapersV3View(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:v3", "view",
			"--group", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
