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

var briefsV1GenerateSpeech = cli.Command{
	Name:    "generate-speech",
	Usage:   "Generates speech audio from brief summary text using ElevenLabs. Returns the\nCloudFront URL. Client should try the CDN URL first and only call this if it\n404s.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-group-id",
			Usage:    "Paper group ID for caching",
			Required: true,
			BodyPath: "paperGroupId",
		},
		&requestflag.Flag[string]{
			Name:     "text",
			Usage:    "Text to convert to speech",
			Required: true,
			BodyPath: "text",
		},
	},
	Action:          handleBriefsV1GenerateSpeech,
	HideHelpCommand: true,
}

func handleBriefsV1GenerateSpeech(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.BriefV1GenerateSpeechParams{}

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
	_, err = client.Briefs.V1.GenerateSpeech(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "briefs:v1 generate-speech", obj, format, transform)
}
