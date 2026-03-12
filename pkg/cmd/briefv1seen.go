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

var briefsV1SeenGetSeen = cli.Command{
	Name:            "get-seen",
	Usage:           "Returns the list of paper group IDs the current user has already viewed as\nbriefs",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleBriefsV1SeenGetSeen,
	HideHelpCommand: true,
}

var briefsV1SeenMarkSeen = cli.Command{
	Name:    "mark-seen",
	Usage:   "Records that the current user has viewed a brief for the given paper group",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Required: true,
			BodyPath: "paperGroupId",
		},
	},
	Action:          handleBriefsV1SeenMarkSeen,
	HideHelpCommand: true,
}

func handleBriefsV1SeenGetSeen(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Briefs.V1.Seen.GetSeen(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "briefs:v1:seen get-seen", obj, format, transform)
}

func handleBriefsV1SeenMarkSeen(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.BriefV1SeenMarkSeenParams{}

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

	return client.Briefs.V1.Seen.MarkSeen(ctx, params, options...)
}
