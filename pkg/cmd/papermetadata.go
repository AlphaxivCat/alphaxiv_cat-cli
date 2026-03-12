// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/apiquery"
	"github.com/stainless-sdks/alphaxiv_cat-cli/internal/requestflag"
	"github.com/stainless-sdks/alphaxiv_cat-go"
	"github.com/stainless-sdks/alphaxiv_cat-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var papersMetadataRetrieveLatestMetadata = cli.Command{
	Name:    "retrieve-latest-metadata",
	Usage:   "Get metadata for latest paper version (deprecated, used by browser extension)",
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
	Action:          handlePapersMetadataRetrieveLatestMetadata,
	HideHelpCommand: true,
}

var papersMetadataRetrieveVersionMetadata = cli.Command{
	Name:    "retrieve-version-metadata",
	Usage:   "Get metadata for a specific paper version (deprecated, used by browser\nextension)",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "upid",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "version-order",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "prevent-tracking",
			QueryPath: "preventTracking",
		},
	},
	Action:          handlePapersMetadataRetrieveVersionMetadata,
	HideHelpCommand: true,
}

func handlePapersMetadataRetrieveLatestMetadata(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("upid") && len(unusedArgs) > 0 {
		cmd.Set("upid", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperMetadataGetLatestMetadataParams{}

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
	_, err = client.Papers.Metadata.GetLatestMetadata(
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
	return ShowJSON(os.Stdout, "papers:metadata retrieve-latest-metadata", obj, format, transform)
}

func handlePapersMetadataRetrieveVersionMetadata(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-order") && len(unusedArgs) > 0 {
		cmd.Set("version-order", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperMetadataGetVersionMetadataParams{
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
	_, err = client.Papers.Metadata.GetVersionMetadata(
		ctx,
		cmd.Value("version-order").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:metadata retrieve-version-metadata", obj, format, transform)
}
