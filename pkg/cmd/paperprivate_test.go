// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestPapersPrivateCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:private", "create",
			"--content-type", "x",
			"--file", "x",
			"--filename", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"contentType: x\n" +
			"file: x\n" +
			"filename: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers:private", "create",
		)
	})
}

func TestPapersPrivateUpdateMetadata(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:private", "update-metadata",
			"--paper-id", "paperId",
			"--abstract", "abstract",
			"--author", "{name: x}",
			"--bibtex", "bibtex",
			"--category", "string",
			"--publication-date", "0",
			"--title", "x",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(papersPrivateUpdateMetadata)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"papers:private", "update-metadata",
			"--paper-id", "paperId",
			"--abstract", "abstract",
			"--author.name", "x",
			"--bibtex", "bibtex",
			"--category", "string",
			"--publication-date", "0",
			"--title", "x",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"abstract: abstract\n" +
			"authors:\n" +
			"  - name: x\n" +
			"bibtex: bibtex\n" +
			"categories:\n" +
			"  - string\n" +
			"publicationDate: 0\n" +
			"title: x\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"papers:private", "update-metadata",
			"--paper-id", "paperId",
		)
	})
}
