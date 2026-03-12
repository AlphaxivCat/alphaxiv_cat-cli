// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestPapersV2Comment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v2", "comment",
			"--api-key", "string",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--anchor-position", "{offset: 0, pageIndex: 0, spanIndex: 0}",
			"--body", "body",
			"--focus-position", "{offset: 0, pageIndex: 0, spanIndex: 0}",
			"--highlight-rect", "[{pageIndex: 0, rects: [{x1: 0, x2: 0, y1: 0, y2: 0}]}]",
			"--parent-comment-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--selected-text", "selectedText",
			"--tag", "anonymous",
			"--title", "title",
			"--highlight-color", "#26fBC5dF",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(papersV2Comment)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "papers:v2", "comment",
			"--api-key", "string",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--anchor-position.offset", "0",
			"--anchor-position.page-index", "0",
			"--anchor-position.span-index", "0",
			"--body", "body",
			"--focus-position.offset", "0",
			"--focus-position.page-index", "0",
			"--focus-position.span-index", "0",
			"--highlight-rect.page-index", "0",
			"--highlight-rect.rects", "[{x1: 0, x2: 0, y1: 0, y2: 0}]",
			"--parent-comment-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			"--selected-text", "selectedText",
			"--tag", "anonymous",
			"--title", "title",
			"--highlight-color", "#26fBC5dF",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"anchorPosition:\n" +
			"  offset: 0\n" +
			"  pageIndex: 0\n" +
			"  spanIndex: 0\n" +
			"body: body\n" +
			"focusPosition:\n" +
			"  offset: 0\n" +
			"  pageIndex: 0\n" +
			"  spanIndex: 0\n" +
			"highlightRects:\n" +
			"  - pageIndex: 0\n" +
			"    rects:\n" +
			"      - x1: 0\n" +
			"        x2: 0\n" +
			"        y1: 0\n" +
			"        y2: 0\n" +
			"parentCommentId: 182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e\n" +
			"selectedText: selectedText\n" +
			"tag: anonymous\n" +
			"title: title\n" +
			"highlightColor: '#26fBC5dF'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "papers:v2", "comment",
			"--api-key", "string",
			"--version", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		)
	})
}
