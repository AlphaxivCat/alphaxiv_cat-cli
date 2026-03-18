// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
)

func TestAnalyticsPaperViewCountIngestEvent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"analytics:paper-view-count", "ingest-event",
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
			t, pipeData,
			"--api-key", "string",
			"analytics:paper-view-count", "ingest-event",
		)
	})
}

func TestAnalyticsPaperViewCountKickoffJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"analytics:paper-view-count", "kickoff-job",
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
			t, pipeData,
			"--api-key", "string",
			"analytics:paper-view-count", "kickoff-job",
		)
	})
}

func TestAnalyticsPaperViewCountProcessJob(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"analytics:paper-view-count", "process-job",
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
			t, pipeData,
			"--api-key", "string",
			"analytics:paper-view-count", "process-job",
		)
	})
}
