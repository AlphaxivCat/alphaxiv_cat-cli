// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var papersV3Retrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve paper version metadata. Fetches from ArXiv if needed.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "unresolved",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
	},
	Action:          handlePapersV3Retrieve,
	HideHelpCommand: true,
}

var papersV3Comment = cli.Command{
	Name:    "comment",
	Usage:   "Create a public comment or private note on a paper.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "version",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tag",
			Usage:    `Allowed values: "anonymous", "general", "personal", "research", "resources".`,
			Required: true,
			BodyPath: "tag",
		},
		&requestflag.Flag[any]{
			Name:     "body",
			BodyPath: "body",
		},
		&requestflag.Flag[any]{
			Name:     "parent",
			BodyPath: "parent",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			BodyPath: "title",
		},
	},
	Action:          handlePapersV3Comment,
	HideHelpCommand: true,
}

var papersV3DeleteVotes = cli.Command{
	Name:    "delete-votes",
	Usage:   "Remove votes from many papers at once",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "body",
			BodyRoot: true,
		},
	},
	Action:          handlePapersV3DeleteVotes,
	HideHelpCommand: true,
}

var papersV3Implementation = cli.Command{
	Name:    "implementation",
	Usage:   "Create or update an implementation for a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "url",
			Required: true,
			BodyPath: "url",
		},
	},
	Action:          handlePapersV3Implementation,
	HideHelpCommand: true,
}

var papersV3KickoffPaperCountries = cli.Command{
	Name:    "kickoff-paper-countries",
	Usage:   "Kickoff paper countries processing for hot papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "batch",
			Usage:    "Number of papers to process in each batch",
			Default:  1,
			BodyPath: "batch",
		},
		&requestflag.Flag[float64]{
			Name:     "max-papers",
			Usage:    "Maximum number of papers to process",
			Default:  1000000,
			BodyPath: "maxPapers",
		},
		&requestflag.Flag[float64]{
			Name:     "months",
			Usage:    "Only process papers at least this many months old",
			BodyPath: "months",
		},
	},
	Action:          handlePapersV3KickoffPaperCountries,
	HideHelpCommand: true,
}

var papersV3KickoffPaperFullText = cli.Command{
	Name:    "kickoff-paper-full-text",
	Usage:   "Kickoff paper full text processing for recent papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[float64]{
			Name:     "max-papers",
			Usage:    "Maximum number of paper versions to process",
			Default:  100,
			BodyPath: "maxPapers",
		},
	},
	Action:          handlePapersV3KickoffPaperFullText,
	HideHelpCommand: true,
}

var papersV3KickoffPaperPodcasts = cli.Command{
	Name:            "kickoff-paper-podcasts",
	Usage:           "Kickoff paper podcasts on Uptash for a subset of paper groups",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersV3KickoffPaperPodcasts,
	HideHelpCommand: true,
}

var papersV3KickoffThumbnailsTrendingPapers = cli.Command{
	Name:            "kickoff-thumbnails-trending-papers",
	Usage:           "Kickoff background job to generate thumbnails for trending papers",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersV3KickoffThumbnailsTrendingPapers,
	HideHelpCommand: true,
}

var papersV3KickoffXMentionsSync = cli.Command{
	Name:    "kickoff-x-mentions-sync",
	Usage:   "Kickoff X mentions sync for hot papers. Uses x-mentions-sync-queue with\nparallelism=1 and built-in delays.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "dry-run",
			Usage:    "If true, only logs papers without queuing",
			Default:  false,
			BodyPath: "dryRun",
		},
		&requestflag.Flag[int64]{
			Name:     "limit",
			Usage:    "Number of hot papers to sync (default: 500)",
			Default:  500,
			BodyPath: "limit",
		},
	},
	Action:          handlePapersV3KickoffXMentionsSync,
	HideHelpCommand: true,
}

var papersV3Like = cli.Command{
	Name:    "like",
	Usage:   "Toggle your like status on a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
	},
	Action:          handlePapersV3Like,
	HideHelpCommand: true,
}

var papersV3Podcast = cli.Command{
	Name:    "podcast",
	Usage:   "Generates a podcast for a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Usage:    "Paper Group ID to generate a podcast for",
			Required: true,
		},
	},
	Action:          handlePapersV3Podcast,
	HideHelpCommand: true,
}

var papersV3ProcessAI = cli.Command{
	Name:    "process-ai",
	Usage:   "Generates AI overviews for a paper version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-version-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "preferred-language",
			Usage:     `Allowed values: "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es", "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja", "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa", "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".`,
			QueryPath: "preferredLanguage",
		},
	},
	Action:          handlePapersV3ProcessAI,
	HideHelpCommand: true,
}

var papersV3ProcessCountries = cli.Command{
	Name:    "process-countries",
	Usage:   "Processes and generates country metadata for papers based on institutional\naffiliations",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "universal-paper-id",
			Usage:    "Array of universal paper IDs (versionless)",
			Required: true,
			BodyPath: "universalPaperIds",
		},
	},
	Action:          handlePapersV3ProcessCountries,
	HideHelpCommand: true,
}

var papersV3ProcessFullText = cli.Command{
	Name:    "process-full-text",
	Usage:   "Processes and extracts full text from paper PDFs for indexing and search",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-version-id",
			Usage:    "Paper version ID to process for full text extraction",
			Required: true,
			BodyPath: "paperVersionId",
		},
	},
	Action:          handlePapersV3ProcessFullText,
	HideHelpCommand: true,
}

var papersV3PruneEmbeddingsByDate = cli.Command{
	Name:            "prune-embeddings-by-date",
	Usage:           "Clear 'is_last_X_days' flags from paper embeddings that have become too old",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersV3PruneEmbeddingsByDate,
	HideHelpCommand: true,
}

var papersV3RequestImplementation = cli.Command{
	Name:    "request-implementation",
	Usage:   "Toggle your implementation request status on a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "paper-title",
			Required: true,
			BodyPath: "paperTitle",
		},
		&requestflag.Flag[string]{
			Name:     "universal-paper-id",
			Required: true,
			BodyPath: "universalPaperId",
		},
		&requestflag.Flag[string]{
			Name:     "additional-info",
			BodyPath: "additionalInfo",
		},
	},
	Action:          handlePapersV3RequestImplementation,
	HideHelpCommand: true,
}

var papersV3RequestPodcast = cli.Command{
	Name:    "request-podcast",
	Usage:   "Request podcast generation for a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Usage:    "Paper Group ID to generate a podcast for",
			Required: true,
		},
	},
	Action:          handlePapersV3RequestPodcast,
	HideHelpCommand: true,
}

var papersV3RetrieveAll = cli.Command{
	Name:    "retrieve-all",
	Usage:   "Get all paper universal IDs sorted by most recent publication date",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "limit",
			Default:   "1000",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "skip",
			Default:   "0",
			QueryPath: "skip",
		},
	},
	Action:          handlePapersV3RetrieveAll,
	HideHelpCommand: true,
}

var papersV3RetrieveDiversePapers = cli.Command{
	Name:    "retrieve-diverse-papers",
	Usage:   "Get an initial batch of diverse papers on the given topics for recommendations",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "topics",
			Required:  true,
			QueryPath: "topics",
		},
	},
	Action:          handlePapersV3RetrieveDiversePapers,
	HideHelpCommand: true,
}

var papersV3RetrieveFeed = cli.Command{
	Name:    "retrieve-feed",
	Usage:   "Get an optionally filtered list of papers for the main feed",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "interval",
			Usage:     `Allowed values: "3 Days", "7 Days", "30 Days", "90 Days", "All time".`,
			Required:  true,
			QueryPath: "interval",
		},
		&requestflag.Flag[string]{
			Name:      "page-num",
			Required:  true,
			QueryPath: "pageNum",
		},
		&requestflag.Flag[string]{
			Name:      "page-size",
			Required:  true,
			QueryPath: "pageSize",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     `Allowed values: "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)", "Recommended".`,
			Required:  true,
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "exclude-seen-briefs",
			Usage:     `Allowed values: "true", "false".`,
			QueryPath: "excludeSeenBriefs",
		},
		&requestflag.Flag[string]{
			Name:      "organizations",
			QueryPath: "organizations",
		},
		&requestflag.Flag[string]{
			Name:      "require-summary",
			Usage:     `Allowed values: "true", "false".`,
			QueryPath: "requireSummary",
		},
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     `Allowed values: "GitHub", "Twitter (X)".`,
			QueryPath: "source",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			QueryPath: "topics",
		},
		&requestflag.Flag[string]{
			Name:      "universal-id",
			Usage:     "A versionless universal paper ID (e.g. 1706.03762)",
			QueryPath: "universalId",
		},
	},
	Action:          handlePapersV3RetrieveFeed,
	HideHelpCommand: true,
}

var papersV3RetrieveFigures = cli.Command{
	Name:    "retrieve-figures",
	Usage:   "Get list of figure URLs for a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
		},
	},
	Action:          handlePapersV3RetrieveFigures,
	HideHelpCommand: true,
}

var papersV3RetrieveFullText = cli.Command{
	Name:    "retrieve-full-text",
	Usage:   "Get the full extracted text of a paper, page by page",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-version",
			Required: true,
		},
	},
	Action:          handlePapersV3RetrieveFullText,
	HideHelpCommand: true,
}

var papersV3RetrieveGeoTrends = cli.Command{
	Name:    "retrieve-geo-trends",
	Usage:   "Retrieve geographical trends and analytics data for papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "collaboration-limit",
			QueryPath: "collaborationLimit",
		},
		&requestflag.Flag[string]{
			Name:      "paper-limit",
			QueryPath: "paperLimit",
		},
		&requestflag.Flag[string]{
			Name:      "past-months",
			QueryPath: "pastMonths",
		},
		&requestflag.Flag[string]{
			Name:      "repo-limit",
			QueryPath: "repoLimit",
		},
		&requestflag.Flag[string]{
			Name:      "top-countries",
			QueryPath: "topCountries",
		},
	},
	Action:          handlePapersV3RetrieveGeoTrends,
	HideHelpCommand: true,
}

var papersV3RetrieveMetrics = cli.Command{
	Name:    "retrieve-metrics",
	Usage:   "Retrieve metrics for a paper (comments count, upvotes, views)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "unresolved",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
	},
	Action:          handlePapersV3RetrieveMetrics,
	HideHelpCommand: true,
}

var papersV3RetrievePapersByCountry = cli.Command{
	Name:    "retrieve-papers-by-country",
	Usage:   "Retrieve top papers by country with optional country filter",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "country",
			QueryPath: "country",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			QueryPath: "limit",
		},
	},
	Action:          handlePapersV3RetrievePapersByCountry,
	HideHelpCommand: true,
}

var papersV3RetrievePreview = cli.Command{
	Name:    "retrieve-preview",
	Usage:   "Retrieve paper data for paper preview cards",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
	},
	Action:          handlePapersV3RetrievePreview,
	HideHelpCommand: true,
}

var papersV3RetrieveSimilarPapers = cli.Command{
	Name:    "retrieve-similar-papers",
	Usage:   "Get papers semantically similar to the selected one",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "exclude",
			QueryPath: "exclude",
		},
		&requestflag.Flag[string]{
			Name:      "exclude-likes",
			Usage:     `Allowed values: "false", "true".`,
			QueryPath: "excludeLikes",
		},
		&requestflag.Flag[string]{
			Name:      "interval",
			Usage:     `Allowed values: "3 Days", "7 Days", "30 Days", "90 Days", "All time".`,
			Default:   "All time",
			QueryPath: "interval",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			QueryPath: "limit",
		},
	},
	Action:          handlePapersV3RetrieveSimilarPapers,
	HideHelpCommand: true,
}

var papersV3RetrieveUnrelated = cli.Command{
	Name:    "retrieve-unrelated",
	Usage:   "Get some papers on the provided topics that are unrelated to the provided papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "limit",
			Required:  true,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "papers",
			Required:  true,
			QueryPath: "papers",
		},
		&requestflag.Flag[string]{
			Name:      "topics",
			Required:  true,
			QueryPath: "topics",
		},
	},
	Action:          handlePapersV3RetrieveUnrelated,
	HideHelpCommand: true,
}

var papersV3View = cli.Command{
	Name:    "view",
	Usage:   "Track paper view event for analytics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
	},
	Action:          handlePapersV3View,
	HideHelpCommand: true,
}

func handlePapersV3Retrieve(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("unresolved") && len(unusedArgs) > 0 {
		cmd.Set("unresolved", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Get(ctx, cmd.Value("unresolved").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve", obj, format, transform)
}

func handlePapersV3Comment(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3CommentParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Comment(
		ctx,
		cmd.Value("version").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 comment", obj, format, transform)
}

func handlePapersV3DeleteVotes(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3DeleteVotesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.DeleteVotes(ctx, params, options...)
}

func handlePapersV3Implementation(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ImplementationParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Implementation(
		ctx,
		cmd.Value("paper-group-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 implementation", obj, format, transform)
}

func handlePapersV3KickoffPaperCountries(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3KickoffPaperCountriesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.KickoffPaperCountries(ctx, params, options...)
}

func handlePapersV3KickoffPaperFullText(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3KickoffPaperFullTextParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.KickoffPaperFullText(ctx, params, options...)
}

func handlePapersV3KickoffPaperPodcasts(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.KickoffPaperPodcasts(ctx, options...)
}

func handlePapersV3KickoffThumbnailsTrendingPapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.KickoffThumbnailsTrendingPapers(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 kickoff-thumbnails-trending-papers", obj, format, transform)
}

func handlePapersV3KickoffXMentionsSync(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3KickoffXMentionsSyncParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.KickoffXMentionsSync(ctx, params, options...)
}

func handlePapersV3Like(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Like(ctx, cmd.Value("group").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 like", obj, format, transform)
}

func handlePapersV3Podcast(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.Podcast(ctx, cmd.Value("paper-group-id").(string), options...)
}

func handlePapersV3ProcessAI(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-version-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-version-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ProcessAIParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.ProcessAI(
		ctx,
		cmd.Value("paper-version-id").(string),
		params,
		options...,
	)
}

func handlePapersV3ProcessCountries(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ProcessCountriesParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.ProcessCountries(ctx, params, options...)
}

func handlePapersV3ProcessFullText(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3ProcessFullTextParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.ProcessFullText(ctx, params, options...)
}

func handlePapersV3PruneEmbeddingsByDate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.PruneEmbeddingsByDate(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 prune-embeddings-by-date", obj, format, transform)
}

func handlePapersV3RequestImplementation(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3RequestImplementationParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.RequestImplementation(
		ctx,
		cmd.Value("group").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 request-implementation", obj, format, transform)
}

func handlePapersV3RequestPodcast(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.RequestPodcast(ctx, cmd.Value("paper-group-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 request-podcast", obj, format, transform)
}

func handlePapersV3RetrieveAll(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetAllParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetAll(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-all", obj, format, transform)
}

func handlePapersV3RetrieveDiversePapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetDiversePapersParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetDiversePapers(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-diverse-papers", obj, format, transform)
}

func handlePapersV3RetrieveFeed(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetFeedParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetFeed(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-feed", obj, format, transform)
}

func handlePapersV3RetrieveFigures(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-group-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-group-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetFigures(ctx, cmd.Value("paper-group-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-figures", obj, format, transform)
}

func handlePapersV3RetrieveFullText(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-version") && len(unusedArgs) > 0 {
		cmd.Set("paper-version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetFullText(ctx, cmd.Value("paper-version").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-full-text", obj, format, transform)
}

func handlePapersV3RetrieveGeoTrends(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetGeoTrendsParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetGeoTrends(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-geo-trends", obj, format, transform)
}

func handlePapersV3RetrieveMetrics(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("unresolved") && len(unusedArgs) > 0 {
		cmd.Set("unresolved", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetMetrics(ctx, cmd.Value("unresolved").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-metrics", obj, format, transform)
}

func handlePapersV3RetrievePapersByCountry(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetPapersByCountryParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetPapersByCountry(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-papers-by-country", obj, format, transform)
}

func handlePapersV3RetrievePreview(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetPreview(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-preview", obj, format, transform)
}

func handlePapersV3RetrieveSimilarPapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetSimilarPapersParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetSimilarPapers(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-similar-papers", obj, format, transform)
}

func handlePapersV3RetrieveUnrelated(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV3GetUnrelatedParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetUnrelated(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3 retrieve-unrelated", obj, format, transform)
}

func handlePapersV3View(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("group") && len(unusedArgs) > 0 {
		cmd.Set("group", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	return client.Papers.V3.View(ctx, cmd.Value("group").(string), options...)
}
