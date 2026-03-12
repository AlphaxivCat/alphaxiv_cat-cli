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

var papersIngestIngestLatest = cli.Command{
	Name:    "ingest-latest",
	Usage:   "Ingest latest paper version if it doesn't exist (deprecated, used by browser\nextension)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "prevent-tracking",
			QueryPath: "preventTracking",
		},
	},
	Action:          handlePapersIngestIngestLatest,
	HideHelpCommand: true,
}

var papersIngestIngestVersion = cli.Command{
	Name:    "ingest-version",
	Usage:   "Ingest a paper version if it doesn't exist (deprecated, used by browser\nextension)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-label",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "prevent-tracking",
			QueryPath: "preventTracking",
		},
	},
	Action:          handlePapersIngestIngestVersion,
	HideHelpCommand: true,
}

func handlePapersIngestIngestLatest(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("upid") && len(unusedArgs) > 0 {
		cmd.Set("upid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperIngestIngestLatestParams{}

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
	_, err = client.Papers.Ingest.IngestLatest(
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
	return ShowJSON(os.Stdout, "papers:ingest ingest-latest", obj, format, transform)
}

func handlePapersIngestIngestVersion(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-label") && len(unusedArgs) > 0 {
		cmd.Set("version-label", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperIngestIngestVersionParams{
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
	_, err = client.Papers.Ingest.IngestVersion(
		ctx,
		cmd.Value("version-label").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:ingest ingest-version", obj, format, transform)
}
