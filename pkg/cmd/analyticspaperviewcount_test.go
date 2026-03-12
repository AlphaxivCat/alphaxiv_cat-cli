// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/mocktest"
)

func TestAnalyticsPaperViewCountIngestEvent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "analytics:paper-view-count", "ingest-event",
			"--api-key", "string",
			"--paper-id", "paperId",
			"--created-at", "createdAt",
			"--user-id", "userId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperId: paperId\n" +
			"createdAt: createdAt\n" +
			"userId: userId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "analytics:paper-view-count", "ingest-event",
			"--api-key", "string",
		)
	})
}

func TestAnalyticsPaperViewCountKickoffJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "analytics:paper-view-count", "kickoff-job",
			"--api-key", "string",
			"--type", "type",
			"--like", "like",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"type: type\n" +
			"like: like\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "analytics:paper-view-count", "kickoff-job",
			"--api-key", "string",
		)
	})
}

func TestAnalyticsPaperViewCountProcessJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "analytics:paper-view-count", "process-job",
			"--api-key", "string",
			"--paper-id", "paperId",
			"--publication-date", "publicationDate",
			"--like=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"paperId: paperId\n" +
			"publicationDate: publicationDate\n" +
			"like: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "analytics:paper-view-count", "process-job",
			"--api-key", "string",
		)
	})
}
