// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/autocomplete"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "alphaxiv-cat",
		Usage:     "CLI for the alphaxiv_cat API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("ALPHAXIV_CAT_API_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "papers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersAddAuthor,
					&papersAdminVote,
					&papersCrxAbstractClick,
					&papersCrxAbstractHit,
					&papersCrxPdfClick,
					&papersCrxPdfHit,
					&papersEmailAuthor,
					&papersGetCrxPaperInfo,
					&papersGetPaperInfo,
					&papersKickoffAbstractEmbed,
					&papersKickoffAI,
					&papersKickoffBibtex,
					&papersKickoffGitHub,
					&papersKickoffPaperCategorization,
					&papersKickoffRecentPapers,
					&papersMarkViewed,
					&papersProcessAbstractEmbed,
					&papersProcessMetadata,
					&papersRequestAILatest,
					&papersRequestAITranslationLatest,
					&papersSetGitHubRepository,
					&papersToggleFollow,
					&papersTranslateAIOverview,
					&papersUnclaim,
					&papersVote,
				},
			},
			{
				Name:     "papers:versions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersVersionsRequestAIOverview,
					&papersVersionsRequestAITranslation,
				},
			},
			{
				Name:     "papers:metadata",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersMetadataRetrieveLatestMetadata,
					&papersMetadataRetrieveVersionMetadata,
				},
			},
			{
				Name:     "papers:ingest",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersIngestIngestLatest,
					&papersIngestIngestVersion,
				},
			},
			{
				Name:     "papers:private",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersPrivateCreate,
					&papersPrivateUpdateMetadata,
				},
			},
			{
				Name:     "papers:kickoff-daily-github-stars",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersKickoffDailyGitHubStarsUpdate,
					&papersKickoffDailyGitHubStarsKickoffDailyGitHubStars,
				},
			},
			{
				Name:     "papers:v3",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV3Retrieve,
					&papersV3Comment,
					&papersV3DeleteVotes,
					&papersV3Implementation,
					&papersV3KickoffPaperCountries,
					&papersV3KickoffPaperFullText,
					&papersV3KickoffPaperPodcasts,
					&papersV3KickoffThumbnailsTrendingPapers,
					&papersV3KickoffXMentionsSync,
					&papersV3Like,
					&papersV3Podcast,
					&papersV3ProcessAI,
					&papersV3ProcessCountries,
					&papersV3ProcessFullText,
					&papersV3PruneEmbeddingsByDate,
					&papersV3RequestImplementation,
					&papersV3RequestPodcast,
					&papersV3RetrieveAll,
					&papersV3RetrieveDiversePapers,
					&papersV3RetrieveFeed,
					&papersV3RetrieveFigures,
					&papersV3RetrieveFullText,
					&papersV3RetrieveGeoTrends,
					&papersV3RetrieveMetrics,
					&papersV3RetrievePapersByCountry,
					&papersV3RetrievePreview,
					&papersV3RetrieveSimilarPapers,
					&papersV3RetrieveUnrelated,
					&papersV3View,
				},
			},
			{
				Name:     "papers:v3:legacy",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV3LegacyRetrieve,
					&papersV3LegacyRetrieveComments,
				},
			},
			{
				Name:     "papers:v3:overview",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV3OverviewRetrieve,
					&papersV3OverviewRetrieveStatus,
				},
			},
			{
				Name:     "papers:v3:implementations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV3ImplementationsCreate,
					&papersV3ImplementationsList,
					&papersV3ImplementationsDelete,
				},
			},
			{
				Name:     "papers:v3:x-mentions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV3XMentionsUpdate,
					&papersV3XMentionsDelete,
				},
			},
			{
				Name:     "papers:v2",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&papersV2Comment,
				},
			},
			{
				Name:     "emails",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&emailsCaptureBouncedEmails,
					&emailsCaptureResendBouncedEmail,
					&emailsKickoffCommentUpdate,
					&emailsKickoffGeneralUpdate,
					&emailsProcessBouncedEmail,
					&emailsProcessCommentUpdate,
					&emailsProcessGeneralUpdate,
				},
			},
			{
				Name:     "users",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersGetPrivateNotes,
					&usersWeighWeeklyReputation,
				},
			},
			{
				Name:     "users:v3",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3AutocompleteProfile,
					&usersV3DeleteBanner,
					&usersV3DeleteOwnUser,
					&usersV3GetActivity,
					&usersV3GetClaimedPapers,
					&usersV3GetCurrentUser,
					&usersV3GetFeaturedActivity,
					&usersV3GetFollowers,
					&usersV3GetLeaderboard,
					&usersV3GetUserByUuid,
					&usersV3GetViewedHistory,
					&usersV3ProcessNotificationEmail,
					&usersV3Search,
					&usersV3ToggleFollowUser,
					&usersV3UpdatePreferences,
					&usersV3UpdateProfile,
					&usersV3UploadAvatar,
				},
			},
			{
				Name:     "users:v3:following",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3FollowingList,
				},
			},
			{
				Name:     "users:v3:following:topics",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3FollowingTopicsList,
					&usersV3FollowingTopicsToggle,
				},
			},
			{
				Name:     "users:v3:following:organizations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3FollowingOrganizationsList,
					&usersV3FollowingOrganizationsToggle,
				},
			},
			{
				Name:     "users:v3:by-username",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3ByUsernameGetProfilePage,
					&usersV3ByUsernameGetUser,
				},
			},
			{
				Name:     "users:v3:semantic-scholar",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3SemanticScholarLink,
					&usersV3SemanticScholarScrape,
				},
			},
			{
				Name:     "users:v3:citations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersV3CitationsGetGraph,
					&usersV3CitationsGetSummary,
				},
			},
			{
				Name:     "users:preferences",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersPreferencesGetFoldersPreferences,
				},
			},
			{
				Name:     "users:preferences:email",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&usersPreferencesEmailUpdate,
					&usersPreferencesEmailGet,
				},
			},
			{
				Name:     "assistant",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&assistantUploadFile,
				},
			},
			{
				Name:     "assistant:v2",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&assistantV2Chat,
					&assistantV2DeleteChat,
					&assistantV2EditChat,
					&assistantV2GetChats,
					&assistantV2GetURLMetadata,
				},
			},
			{
				Name:     "assistant:v2:messages",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&assistantV2MessagesList,
					&assistantV2MessagesSelect,
					&assistantV2MessagesSetFeedback,
				},
			},
			{
				Name:     "folders:v3",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&foldersV3Create,
					&foldersV3List,
					&foldersV3Delete,
					&foldersV3AddPapers,
					&foldersV3MovePapers,
					&foldersV3RemovePapers,
					&foldersV3ToggleSharing,
					&foldersV3UpdateName,
					&foldersV3UpdateParent,
				},
			},
			{
				Name:     "folders:v3:shared",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&foldersV3SharedRetrieve,
					&foldersV3SharedCopy,
				},
			},
			{
				Name:     "comments",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&commentsEdit,
				},
			},
			{
				Name:     "comments:v2",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&commentsV2Delete,
					&commentsV2Downvote,
					&commentsV2Flag,
					&commentsV2ToggleEndorse,
					&commentsV2Upvote,
				},
			},
			{
				Name:     "comments:v2:moderator",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&commentsV2ModeratorSendFeedback,
					&commentsV2ModeratorToggleFlags,
				},
			},
			{
				Name:     "analytics:paper-view-count",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&analyticsPaperViewCountIngestEvent,
					&analyticsPaperViewCountKickoffJob,
					&analyticsPaperViewCountProcessJob,
				},
			},
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchClosestTopic,
					&searchGoogleSearch,
				},
			},
			{
				Name:     "search:v2:paper",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchV2PaperFastSearch,
				},
			},
			{
				Name:     "organizations:v2",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationsV2ListTop,
					&organizationsV2RetrieveByID,
					&organizationsV2RetrieveByName,
					&organizationsV2Search,
					&organizationsV2ToggleFollow,
				},
			},
			{
				Name:     "google-scholar:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&googleScholarV1Connect,
					&googleScholarV1DeleteConnection,
					&googleScholarV1GetReport,
					&googleScholarV1GetStatus,
					&googleScholarV1Resync,
					&googleScholarV1SetEmail,
					&googleScholarV1Sync,
					&googleScholarV1VerifyEmail,
				},
			},
			{
				Name:     "arxiv:v1:labs",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&arxivV1LabsRetrieve,
				},
			},
			{
				Name:     "api-keys:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&apiKeysV1Create,
					&apiKeysV1List,
					&apiKeysV1CreateImpersonation,
					&apiKeysV1Revoke,
				},
			},
			{
				Name:     "admin:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&adminV1GetModeratorFeed,
					&adminV1LookupUserByEmail,
				},
			},
			{
				Name:     "admin:v1:emails",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&adminV1EmailsSendMonthlyDigest,
					&adminV1EmailsSendWeeklyDigest,
				},
			},
			{
				Name:     "wrapped",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wrappedRetrieveByUser,
				},
			},
			{
				Name:     "notifications",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationsSendKickoffNotificationEmails,
				},
			},
			{
				Name:     "notifications:v4",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationsV4ListNotifications,
					&notificationsV4Subscribe,
					&notificationsV4Unsubscribe,
				},
			},
			{
				Name:     "notifications:v4:archive",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&notificationsV4ArchiveCreate,
					&notificationsV4ArchiveList,
				},
			},
			{
				Name:     "sitemaps",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&sitemapsListOverviews,
					&sitemapsListPapers,
					&sitemapsListUsers,
				},
			},
			{
				Name:     "retrieve:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&retrieveV1GetTopPapers,
				},
			},
			{
				Name:     "retool:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&retoolV1GetCumulativeUsers,
					&retoolV1GetDailyConversations,
					&retoolV1GetDailyNewAccounts,
					&retoolV1GetDailyUserChatMessages,
					&retoolV1GetTotalCommentCount,
					&retoolV1GetTotalPaperCount,
					&retoolV1GetTotalPrivateNotesCount,
					&retoolV1GetTotalUserCount,
					&retoolV1GetWeeklyMessageCountsByUser,
					&retoolV1GetWeeklyPrivateNotes,
					&retoolV1GetWeeklyPublicComments,
				},
			},
			{
				Name:     "briefs:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&briefsV1GenerateSpeech,
				},
			},
			{
				Name:     "briefs:v1:seen",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&briefsV1SeenGetSeen,
					&briefsV1SeenMarkSeen,
				},
			},
			{
				Name:     "research",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&researchCreateProject,
				},
			},
			{
				Name:     "mcp:v1",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&mcpV1EstablishConnection,
					&mcpV1SendMessage,
					&mcpV1TerminateSession,
				},
			},
			{
				Name:     "events",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&eventsList,
				},
			},
			{
				Name:     "well-known",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&wellKnownRetrieveOAuthProtectedResource,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "alphaxiv-cat @manpages [-o alphaxiv-cat.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "alphaxiv-cat.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "alphaxiv-cat.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
