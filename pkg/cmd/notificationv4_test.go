// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/mocktest"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
)

func TestNotificationsV4ListNotifications(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "notifications:v4", "list-notifications",
			"--api-key", "string",
		)
	})
}

func TestNotificationsV4Subscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "notifications:v4", "subscribe",
			"--api-key", "string",
			"--subscription", "{endpoint: https://example.com, keys: {auth: auth, p256dh: p256dh}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(notificationsV4Subscribe)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t, "notifications:v4", "subscribe",
			"--api-key", "string",
			"--subscription.endpoint", "https://example.com",
			"--subscription.keys", "{auth: auth, p256dh: p256dh}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"subscription:\n" +
			"  endpoint: https://example.com\n" +
			"  keys:\n" +
			"    auth: auth\n" +
			"    p256dh: p256dh\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "notifications:v4", "subscribe",
			"--api-key", "string",
		)
	})
}

func TestNotificationsV4Unsubscribe(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "notifications:v4", "unsubscribe",
			"--api-key", "string",
			"--endpoint", "https://example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("endpoint: https://example.com")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "notifications:v4", "unsubscribe",
			"--api-key", "string",
		)
	})
}
