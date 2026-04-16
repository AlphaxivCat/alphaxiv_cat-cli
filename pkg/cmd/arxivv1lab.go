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

var arxivV1LabsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Gets data necessary to render arXiv labs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "unresolved",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
	},
	Action:          handleArxivV1LabsRetrieve,
	HideHelpCommand: true,
}

func handleArxivV1LabsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Arxiv.V1.Labs.Get(ctx, cmd.Value("unresolved").(string), options...)
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
		Title:          "arxiv:v1:labs retrieve",
		Transform:      transform,
	})
}
