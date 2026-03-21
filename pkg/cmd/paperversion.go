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

var papersVersionsRequestAIOverview = cli.Command{
	Name:    "request-ai-overview",
	Usage:   "Request AI overview generation for a paper version",
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
			Name:      "preferred-language",
			Usage:     `Allowed values: "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es", "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja", "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa", "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".`,
			QueryPath: "preferredLanguage",
		},
	},
	Action:          handlePapersVersionsRequestAIOverview,
	HideHelpCommand: true,
}

var papersVersionsRequestAITranslation = cli.Command{
	Name:    "request-ai-translation",
	Usage:   "Request AI overview translation for a paper version",
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
			Name:     "language",
			Usage:    `Allowed values: "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "es", "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja", "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa", "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".`,
			Required: true,
		},
	},
	Action:          handlePapersVersionsRequestAITranslation,
	HideHelpCommand: true,
}

func handlePapersVersionsRequestAIOverview(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version-order") && len(unusedArgs) > 0 {
		cmd.Set("version-order", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperVersionRequestAIOverviewParams{
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
	_, err = client.Papers.Versions.RequestAIOverview(
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
	return ShowJSON(os.Stdout, "papers:versions request-ai-overview", obj, format, transform)
}

func handlePapersVersionsRequestAITranslation(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("language") && len(unusedArgs) > 0 {
		cmd.Set("language", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperVersionRequestAITranslationParams{
		Upid:         cmd.Value("upid").(string),
		VersionOrder: cmd.Value("version-order").(string),
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
	_, err = client.Papers.Versions.RequestAITranslation(
		ctx,
		alphaxivcat.PaperVersionRequestAITranslationParamsLanguage(cmd.Value("language").(string)),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:versions request-ai-translation", obj, format, transform)
}
