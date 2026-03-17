// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestAdminV1EmailsSendMonthlyDigest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"admin:v1:emails", "send-monthly-digest",
			"--role", "admin",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("role: admin")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"admin:v1:emails", "send-monthly-digest",
		)
	})
}

func TestAdminV1EmailsSendWeeklyDigest(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
			"--event", "{date: date, description: description, link: link, title: title, endTimeRaw: endTimeRaw, startTimeRaw: startTimeRaw}",
			"--intro-text", "introText",
			"--role", "admin",
			"--subject", "subject",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(adminV1EmailsSendWeeklyDigest)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
			"--event.date", "date",
			"--event.description", "description",
			"--event.link", "link",
			"--event.title", "title",
			"--event.end-time-raw", "endTimeRaw",
			"--event.start-time-raw", "startTimeRaw",
			"--intro-text", "introText",
			"--role", "admin",
			"--subject", "subject",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"events:\n" +
			"  - date: date\n" +
			"    description: description\n" +
			"    link: link\n" +
			"    title: title\n" +
			"    endTimeRaw: endTimeRaw\n" +
			"    startTimeRaw: startTimeRaw\n" +
			"introText: introText\n" +
			"role: admin\n" +
			"subject: subject\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"admin:v1:emails", "send-weekly-digest",
		)
	})
}
