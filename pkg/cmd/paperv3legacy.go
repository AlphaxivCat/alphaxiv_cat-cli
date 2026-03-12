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

var papersV3LegacyRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieve paper version metadata and comments. Fetches from ArXiv if needed. This\nruns the old id resolution implementation, old fetching service, and old JSON\nformat. Do not write new code with this please.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "unresolved",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
	},
	Action:          handlePapersV3LegacyRetrieve,
	HideHelpCommand: true,
}

var papersV3LegacyRetrieveComments = cli.Command{
	Name:    "retrieve-comments",
	Usage:   "Retrieve paper group comments. Does not distinguish \"paper group does not exist\"\nfrom \"there are no comments\"",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "group",
			Required: true,
		},
	},
	Action:          handlePapersV3LegacyRetrieveComments,
	HideHelpCommand: true,
}

func handlePapersV3LegacyRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.V3.Legacy.Get(ctx, cmd.Value("unresolved").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3:legacy retrieve", obj, format, transform)
}

func handlePapersV3LegacyRetrieveComments(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Papers.V3.Legacy.GetComments(ctx, cmd.Value("group").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:v3:legacy retrieve-comments", obj, format, transform)
}
