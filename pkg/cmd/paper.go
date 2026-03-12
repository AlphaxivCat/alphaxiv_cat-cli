// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/AlphaxivCat/alphaxiv_cat-go/option"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var papersAddAuthor = cli.Command{
	Name:    "add-author",
	Usage:   "Add a new author to a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "author-email",
			Required: true,
			BodyPath: "authorEmail",
		},
		&requestflag.Flag[bool]{
			Name:     "should-email",
			BodyPath: "shouldEmail",
		},
	},
	Action:          handlePapersAddAuthor,
	HideHelpCommand: true,
}

var papersAdminVote = cli.Command{
	Name:    "admin-vote",
	Usage:   "Set paper vote count (admin only)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
		&requestflag.Flag[float64]{
			Name:     "entry",
			Required: true,
			BodyPath: "entry",
		},
	},
	Action:          handlePapersAdminVote,
	HideHelpCommand: true,
}

var papersCrxAbstractClick = cli.Command{
	Name:    "crx-abstract-click",
	Usage:   "Legacy route for v1 browser extensions to track abstract page clicks",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			Usage:    "Referrer",
			Required: true,
		},
	},
	Action:          handlePapersCrxAbstractClick,
	HideHelpCommand: true,
}

var papersCrxAbstractHit = cli.Command{
	Name:    "crx-abstract-hit",
	Usage:   "Legacy route for v1 browser extensions to track abstract page hits",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
	},
	Action:          handlePapersCrxAbstractHit,
	HideHelpCommand: true,
}

var papersCrxPdfClick = cli.Command{
	Name:    "crx-pdf-click",
	Usage:   "Legacy route for v1 browser extensions to track PDF page clicks",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "ref",
			Usage:    "Referrer",
			Required: true,
		},
	},
	Action:          handlePapersCrxPdfClick,
	HideHelpCommand: true,
}

var papersCrxPdfHit = cli.Command{
	Name:    "crx-pdf-hit",
	Usage:   "Legacy route for v1 browser extensions to track PDF page hits",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
	},
	Action:          handlePapersCrxPdfHit,
	HideHelpCommand: true,
}

var papersEmailAuthor = cli.Command{
	Name:    "email-author",
	Usage:   "Send email to individual author about paper comments or trending",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "author-individual-email",
			Required: true,
			BodyPath: "authorIndividualEmail",
		},
		&requestflag.Flag[string]{
			Name:     "email-author-name",
			Required: true,
			BodyPath: "emailAuthorName",
		},
		&requestflag.Flag[string]{
			Name:     "paper-name",
			Required: true,
			BodyPath: "paperName",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[bool]{
			Name:     "ignore-duplicate-error",
			BodyPath: "ignoreDuplicateError",
		},
	},
	Action:          handlePapersEmailAuthor,
	HideHelpCommand: true,
}

var papersGetCrxPaperInfo = cli.Command{
	Name:    "get-crx-paper-info",
	Usage:   "Legacy route for v1 browser extensions to get paper information",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
	},
	Action:          handlePapersGetCrxPaperInfo,
	HideHelpCommand: true,
}

var papersGetPaperInfo = cli.Command{
	Name:    "get-paper-info",
	Usage:   "Legacy route for getting paper information from arXiv abstract pages",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "pid",
			Usage:    "Paper ID",
			Required: true,
		},
	},
	Action:          handlePapersGetPaperInfo,
	HideHelpCommand: true,
}

var papersKickoffAbstractEmbed = cli.Command{
	Name:            "kickoff-abstract-embed",
	Usage:           "Kickoff background job to generate abstract embeddings for paper versions",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffAbstractEmbed,
	HideHelpCommand: true,
}

var papersKickoffAI = cli.Command{
	Name:            "kickoff-ai",
	Usage:           "Kickoff background job to generate AI overviews for papers",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffAI,
	HideHelpCommand: true,
}

var papersKickoffBibtex = cli.Command{
	Name:            "kickoff-bibtex",
	Usage:           "Kickoff background job to generate bibtex for papers",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffBibtex,
	HideHelpCommand: true,
}

var papersKickoffGitHub = cli.Command{
	Name:            "kickoff-github",
	Usage:           "Kickoff background job to link papers with GitHub repositories",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffGitHub,
	HideHelpCommand: true,
}

var papersKickoffPaperCategorization = cli.Command{
	Name:    "kickoff-paper-categorization",
	Usage:   "Kickoff background job to categorize papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "all",
			Required: true,
		},
	},
	Action:          handlePapersKickoffPaperCategorization,
	HideHelpCommand: true,
}

var papersKickoffRecentPapers = cli.Command{
	Name:            "kickoff-recent-papers",
	Usage:           "Kickoff background job to ingest recent papers from arXiv",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handlePapersKickoffRecentPapers,
	HideHelpCommand: true,
}

var papersMarkViewed = cli.Command{
	Name:    "mark-viewed",
	Usage:   "Track paper view event for analytics",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
	},
	Action:          handlePapersMarkViewed,
	HideHelpCommand: true,
}

var papersProcessAbstractEmbed = cli.Command{
	Name:    "process-abstract-embed",
	Usage:   "Process abstract embedding for a paper version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-version-id",
			Required: true,
			BodyPath: "paperVersionId",
		},
	},
	Action:          handlePapersProcessAbstractEmbed,
	HideHelpCommand: true,
}

var papersProcessMetadata = requestflag.WithInnerFlags(cli.Command{
	Name:    "process-metadata",
	Usage:   "Process various metadata for a paper (thumbnail, github, bibtex, etc.)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "universal-paper-id",
			Required: true,
			BodyPath: "universalPaperId",
		},
	},
	Action:          handlePapersProcessMetadata,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.bibtex",
			InnerField: "bibtex",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.custom-categories",
			InnerField: "custom_categories",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.embedding",
			InnerField: "embedding",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.github",
			InnerField: "github",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.organizations",
			InnerField: "organizations",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.overview",
			InnerField: "overview",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "metadata.thumbnail",
			InnerField: "thumbnail",
		},
	},
})

var papersRequestAILatest = cli.Command{
	Name:    "request-ai-latest",
	Usage:   "Request AI overview generation for the latest paper version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "preferred-language",
			QueryPath: "preferredLanguage",
		},
	},
	Action:          handlePapersRequestAILatest,
	HideHelpCommand: true,
}

var papersRequestAITranslationLatest = cli.Command{
	Name:    "request-ai-translation-latest",
	Usage:   "Request AI overview translation for the latest paper version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Required: true,
		},
	},
	Action:          handlePapersRequestAITranslationLatest,
	HideHelpCommand: true,
}

var papersSetGitHubRepository = cli.Command{
	Name:    "set-github-repository",
	Usage:   "Set GitHub repository for a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "github",
			Required: true,
			BodyPath: "github",
		},
	},
	Action:          handlePapersSetGitHubRepository,
	HideHelpCommand: true,
}

var papersToggleFollow = cli.Command{
	Name:    "toggle-follow",
	Usage:   "Toggle paper follow status (add to Want to read folder or remove from all\nfolders)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
	},
	Action:          handlePapersToggleFollow,
	HideHelpCommand: true,
}

var papersTranslateAIOverview = cli.Command{
	Name:    "translate-ai-overview",
	Usage:   "Translate AI overview to specified language",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-version-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "language",
			Required: true,
		},
	},
	Action:          handlePapersTranslateAIOverview,
	HideHelpCommand: true,
}

var papersUnclaim = cli.Command{
	Name:    "unclaim",
	Usage:   "Remove authorship claim from a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
	},
	Action:          handlePapersUnclaim,
	HideHelpCommand: true,
}

var papersVote = cli.Command{
	Name:    "vote",
	Usage:   "Toggle vote for a paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Required: true,
		},
	},
	Action:          handlePapersVote,
	HideHelpCommand: true,
}

func handlePapersAddAuthor(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperAddAuthorParams{}

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
	_, err = client.Papers.AddAuthor(
		ctx,
		cmd.Value("paper-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers add-author", obj, format, transform)
}

func handlePapersAdminVote(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperAdminVoteParams{}

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
	_, err = client.Papers.AdminVote(
		ctx,
		cmd.Value("paper-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers admin-vote", obj, format, transform)
}

func handlePapersCrxAbstractClick(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ref") && len(unusedArgs) > 0 {
		cmd.Set("ref", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperCrxAbstractClickParams{
		Pid: cmd.Value("pid").(string),
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
	_, err = client.Papers.CrxAbstractClick(
		ctx,
		cmd.Value("ref").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers crx-abstract-click", obj, format, transform)
}

func handlePapersCrxAbstractHit(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pid") && len(unusedArgs) > 0 {
		cmd.Set("pid", unusedArgs[0])
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
	_, err = client.Papers.CrxAbstractHit(ctx, cmd.Value("pid").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers crx-abstract-hit", obj, format, transform)
}

func handlePapersCrxPdfClick(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("ref") && len(unusedArgs) > 0 {
		cmd.Set("ref", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperCrxPdfClickParams{
		Pid: cmd.Value("pid").(string),
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
	_, err = client.Papers.CrxPdfClick(
		ctx,
		cmd.Value("ref").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers crx-pdf-click", obj, format, transform)
}

func handlePapersCrxPdfHit(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pid") && len(unusedArgs) > 0 {
		cmd.Set("pid", unusedArgs[0])
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
	_, err = client.Papers.CrxPdfHit(ctx, cmd.Value("pid").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers crx-pdf-hit", obj, format, transform)
}

func handlePapersEmailAuthor(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperEmailAuthorParams{}

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
	_, err = client.Papers.EmailAuthor(
		ctx,
		cmd.Value("paper-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers email-author", obj, format, transform)
}

func handlePapersGetCrxPaperInfo(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pid") && len(unusedArgs) > 0 {
		cmd.Set("pid", unusedArgs[0])
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
	_, err = client.Papers.GetCrxPaperInfo(ctx, cmd.Value("pid").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers get-crx-paper-info", obj, format, transform)
}

func handlePapersGetPaperInfo(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pid") && len(unusedArgs) > 0 {
		cmd.Set("pid", unusedArgs[0])
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
	_, err = client.Papers.GetPaperInfo(ctx, cmd.Value("pid").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers get-paper-info", obj, format, transform)
}

func handlePapersKickoffAbstractEmbed(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffAbstractEmbed(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-abstract-embed", obj, format, transform)
}

func handlePapersKickoffAI(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffAI(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-ai", obj, format, transform)
}

func handlePapersKickoffBibtex(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffBibtex(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-bibtex", obj, format, transform)
}

func handlePapersKickoffGitHub(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffGitHub(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-github", obj, format, transform)
}

func handlePapersKickoffPaperCategorization(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("all") && len(unusedArgs) > 0 {
		cmd.Set("all", unusedArgs[0])
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
	_, err = client.Papers.KickoffPaperCategorization(ctx, cmd.Value("all").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-paper-categorization", obj, format, transform)
}

func handlePapersKickoffRecentPapers(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.KickoffRecentPapers(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers kickoff-recent-papers", obj, format, transform)
}

func handlePapersMarkViewed(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("upid") && len(unusedArgs) > 0 {
		cmd.Set("upid", unusedArgs[0])
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
	_, err = client.Papers.MarkViewed(ctx, cmd.Value("upid").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers mark-viewed", obj, format, transform)
}

func handlePapersProcessAbstractEmbed(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperProcessAbstractEmbedParams{}

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
	_, err = client.Papers.ProcessAbstractEmbed(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers process-abstract-embed", obj, format, transform)
}

func handlePapersProcessMetadata(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperProcessMetadataParams{}

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
	_, err = client.Papers.ProcessMetadata(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers process-metadata", obj, format, transform)
}

func handlePapersRequestAILatest(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("upid") && len(unusedArgs) > 0 {
		cmd.Set("upid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperRequestAILatestParams{}

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
	_, err = client.Papers.RequestAILatest(
		ctx,
		cmd.Value("upid").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers request-ai-latest", obj, format, transform)
}

func handlePapersRequestAITranslationLatest(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("language") && len(unusedArgs) > 0 {
		cmd.Set("language", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperRequestAITranslationLatestParams{
		Upid: cmd.Value("upid").(string),
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
	_, err = client.Papers.RequestAITranslationLatest(
		ctx,
		alphaxivcat.PaperRequestAITranslationLatestParamsLanguage(cmd.Value("language").(string)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers request-ai-translation-latest", obj, format, transform)
}

func handlePapersSetGitHubRepository(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperSetGitHubRepositoryParams{}

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
	_, err = client.Papers.SetGitHubRepository(
		ctx,
		cmd.Value("paper-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers set-github-repository", obj, format, transform)
}

func handlePapersToggleFollow(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
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
	_, err = client.Papers.ToggleFollow(ctx, cmd.Value("paper-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers toggle-follow", obj, format, transform)
}

func handlePapersTranslateAIOverview(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("language") && len(unusedArgs) > 0 {
		cmd.Set("language", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperTranslateAIOverviewParams{
		PaperVersionID: cmd.Value("paper-version-id").(string),
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
	_, err = client.Papers.TranslateAIOverview(
		ctx,
		alphaxivcat.PaperTranslateAIOverviewParamsLanguage(cmd.Value("language").(string)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers translate-ai-overview", obj, format, transform)
}

func handlePapersUnclaim(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
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
	_, err = client.Papers.Unclaim(ctx, cmd.Value("paper-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers unclaim", obj, format, transform)
}

func handlePapersVote(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
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
	_, err = client.Papers.Vote(ctx, cmd.Value("paper-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers vote", obj, format, transform)
}
