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

var sitemapsListOverviews = cli.Command{
	Name:    "list-overviews",
	Usage:   "Get paginated list of paper versions with AI overviews for sitemap generation",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			QueryPath: "page",
		},
	},
	Action:          handleSitemapsListOverviews,
	HideHelpCommand: true,
}

var sitemapsListPapers = cli.Command{
	Name:    "list-papers",
	Usage:   "Get paginated list of original (non-arXiv, non-blog) public papers for sitemap\ngeneration. Uses cursor caching for efficient deep pagination.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "page",
			QueryPath: "page",
		},
	},
	Action:          handleSitemapsListPapers,
	HideHelpCommand: true,
}

func handleSitemapsListOverviews(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.SitemapListOverviewsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sitemaps.ListOverviews(ctx, params, options...)
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
		Title:          "sitemaps list-overviews",
		Transform:      transform,
	})
}

func handleSitemapsListPapers(ctx context.Context, cmd *cli.Command) error {
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

	params := alphaxivcat.SitemapListPapersParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Sitemaps.ListPapers(ctx, params, options...)
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
		Title:          "sitemaps list-papers",
		Transform:      transform,
	})
}
