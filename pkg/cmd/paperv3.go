// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

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
			Name:      "unresolved",
			Usage:     "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required:  true,
			PathParam: "unresolved",
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
			Name:      "version",
			Required:  true,
			PathParam: "version",
		},
		&requestflag.Flag[string]{
			Name:     "tag",
			Usage:    `Allowed values: "anonymous", "general", "personal", "research", "resources".`,
			Required: true,
			BodyPath: "tag",
		},
		&requestflag.Flag[*string]{
			Name:     "body",
			BodyPath: "body",
		},
		&requestflag.Flag[*string]{
			Name:     "parent",
			BodyPath: "parent",
		},
		&requestflag.Flag[*string]{
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
			Name:      "paper-group-id",
			Required:  true,
			PathParam: "paperGroupId",
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

var papersV3KickoffPaperPodcasts = cli.Command{
	Name:            "kickoff-paper-podcasts",
	Usage:           "Kickoff paper podcasts on Uptash for a subset of paper groups",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersV3KickoffPaperPodcasts,
	HideHelpCommand: true,
}

var papersV3Like = cli.Command{
	Name:    "like",
	Usage:   "Set your like status on a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "group",
			Required:  true,
			PathParam: "group",
		},
		&requestflag.Flag[string]{
			Name:      "liked",
			Usage:     `Allowed values: "true", "false".`,
			Required:  true,
			QueryPath: "liked",
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
			Name:      "paper-group-id",
			Usage:     "Paper Group ID to generate a podcast for",
			Required:  true,
			PathParam: "paperGroupId",
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
			Name:      "paper-version-id",
			Required:  true,
			PathParam: "paperVersionId",
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

var papersV3RequestImplementation = cli.Command{
	Name:    "request-implementation",
	Usage:   "Toggle your implementation request status on a paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "group",
			Required:  true,
			PathParam: "group",
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
			Name:      "paper-group-id",
			Usage:     "Paper Group ID to generate a podcast for",
			Required:  true,
			PathParam: "paperGroupId",
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
			Usage:     `Allowed values: "Hot", "Comments", "Views", "Likes", "GitHub", "Recommended", "Recent".`,
			Required:  true,
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "include-external-blogs",
			QueryPath: "includeExternalBlogs",
		},
		&requestflag.Flag[string]{
			Name:      "source",
			Usage:     `Allowed values: "GitHub".`,
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
			Name:      "paper-group-id",
			Required:  true,
			PathParam: "paperGroupId",
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
			Name:      "paper-version",
			Required:  true,
			PathParam: "paperVersion",
		},
	},
	Action:          handlePapersV3RetrieveFullText,
	HideHelpCommand: true,
}

var papersV3RetrieveMetrics = cli.Command{
	Name:    "retrieve-metrics",
	Usage:   "Retrieve metrics for a paper (comments count, upvotes, views)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "unresolved",
			Usage:     "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required:  true,
			PathParam: "unresolved",
		},
	},
	Action:          handlePapersV3RetrieveMetrics,
	HideHelpCommand: true,
}

var papersV3RetrievePreview = cli.Command{
	Name:    "retrieve-preview",
	Usage:   "Retrieve paper data for paper preview cards",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "id",
			Usage:     "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required:  true,
			PathParam: "id",
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
			Name:      "id",
			Usage:     "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required:  true,
			PathParam: "id",
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
			Name:      "group",
			Required:  true,
			PathParam: "group",
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve",
		Transform:      transform,
	})
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

	params := alphaxivcat.PaperV3CommentParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 comment",
		Transform:      transform,
	})
}

func handlePapersV3DeleteVotes(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := alphaxivcat.PaperV3DeleteVotesParams{}

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

	params := alphaxivcat.PaperV3ImplementationParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 implementation",
		Transform:      transform,
	})
}

func handlePapersV3KickoffPaperCountries(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

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

	params := alphaxivcat.PaperV3KickoffPaperCountriesParams{}

	return client.Papers.V3.KickoffPaperCountries(ctx, params, options...)
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

	params := alphaxivcat.PaperV3LikeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.Like(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 like",
		Transform:      transform,
	})
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

	params := alphaxivcat.PaperV3ProcessAIParams{}

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

	params := alphaxivcat.PaperV3ProcessCountriesParams{}

	return client.Papers.V3.ProcessCountries(ctx, params, options...)
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

	params := alphaxivcat.PaperV3RequestImplementationParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 request-implementation",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 request-podcast",
		Transform:      transform,
	})
}

func handlePapersV3RetrieveAll(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.PaperV3GetAllParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetAll(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-all",
		Transform:      transform,
	})
}

func handlePapersV3RetrieveDiversePapers(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.PaperV3GetDiversePapersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetDiversePapers(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-diverse-papers",
		Transform:      transform,
	})
}

func handlePapersV3RetrieveFeed(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.PaperV3GetFeedParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetFeed(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-feed",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-figures",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-full-text",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-metrics",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-preview",
		Transform:      transform,
	})
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

	params := alphaxivcat.PaperV3GetSimilarPapersParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-similar-papers",
		Transform:      transform,
	})
}

func handlePapersV3RetrieveUnrelated(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.PaperV3GetUnrelatedParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Papers.V3.GetUnrelated(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "papers:v3 retrieve-unrelated",
		Transform:      transform,
	})
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
