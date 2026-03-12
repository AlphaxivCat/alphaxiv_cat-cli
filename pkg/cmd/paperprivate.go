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

var papersPrivateCreate = cli.Command{
	Name:    "create",
	Usage:   "Upload a private PDF paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "content-type",
			Required: true,
			BodyPath: "contentType",
		},
		&requestflag.Flag[string]{
			Name:     "file",
			Required: true,
			BodyPath: "file",
		},
		&requestflag.Flag[string]{
			Name:     "filename",
			Required: true,
			BodyPath: "filename",
		},
	},
	Action:          handlePapersPrivateCreate,
	HideHelpCommand: true,
}

var papersPrivateUpdateMetadata = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-metadata",
	Usage:   "Update metadata for a private paper",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "paper-id",
			Usage:    "An Unresolved Paper ID (UUID, ArXiv ID, or Versioned ArXiv ID)",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "abstract",
			Required: true,
			BodyPath: "abstract",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "author",
			Required: true,
			BodyPath: "authors",
		},
		&requestflag.Flag[any]{
			Name:     "bibtex",
			Required: true,
			BodyPath: "bibtex",
		},
		&requestflag.Flag[[]string]{
			Name:     "category",
			Required: true,
			BodyPath: "categories",
		},
		&requestflag.Flag[float64]{
			Name:     "publication-date",
			Required: true,
			BodyPath: "publicationDate",
		},
		&requestflag.Flag[string]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
	},
	Action:          handlePapersPrivateUpdateMetadata,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"author": {
		&requestflag.InnerFlag[string]{
			Name:       "author.name",
			InnerField: "name",
		},
	},
})

func handlePapersPrivateCreate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperPrivateNewParams{}

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
	_, err = client.Papers.Private.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "papers:private create", obj, format, transform)
}

func handlePapersPrivateUpdateMetadata(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("paper-id") && len(unusedArgs) > 0 {
		cmd.Set("paper-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperPrivateUpdateMetadataParams{}

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

	return client.Papers.Private.UpdateMetadata(
		ctx,
		cmd.Value("paper-id").(string),
		params,
		options...,
	)
}
