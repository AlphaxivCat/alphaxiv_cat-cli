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

var papersV2Comment = requestflag.WithInnerFlags(cli.Command{
	Name:    "comment",
	Usage:   "Add comment to paper version",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "version",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "anchor-position",
			Required: true,
			BodyPath: "anchorPosition",
		},
		&requestflag.Flag[string]{
			Name:     "body",
			Required: true,
			BodyPath: "body",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "focus-position",
			Required: true,
			BodyPath: "focusPosition",
		},
		&requestflag.Flag[any]{
			Name:     "highlight-rect",
			Required: true,
			BodyPath: "highlightRects",
		},
		&requestflag.Flag[any]{
			Name:     "parent-comment-id",
			Required: true,
			BodyPath: "parentCommentId",
		},
		&requestflag.Flag[any]{
			Name:     "selected-text",
			Required: true,
			BodyPath: "selectedText",
		},
		&requestflag.Flag[any]{
			Name:     "tag",
			Usage:    `Allowed values: "anonymous", "general", "personal", "research", "resources".`,
			Required: true,
			BodyPath: "tag",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
		&requestflag.Flag[any]{
			Name:     "highlight-color",
			BodyPath: "highlightColor",
		},
	},
	Action:          handlePapersV2Comment,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"anchor-position": {
		&requestflag.InnerFlag[float64]{
			Name:       "anchor-position.offset",
			InnerField: "offset",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "anchor-position.page-index",
			InnerField: "pageIndex",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "anchor-position.span-index",
			InnerField: "spanIndex",
		},
	},
	"focus-position": {
		&requestflag.InnerFlag[float64]{
			Name:       "focus-position.offset",
			InnerField: "offset",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "focus-position.page-index",
			InnerField: "pageIndex",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "focus-position.span-index",
			InnerField: "spanIndex",
		},
	},
	"highlight-rect": {
		&requestflag.InnerFlag[float64]{
			Name:       "highlight-rect.page-index",
			InnerField: "pageIndex",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "highlight-rect.rects",
			InnerField: "rects",
		},
	},
})

func handlePapersV2Comment(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("version") && len(unusedArgs) > 0 {
		cmd.Set("version", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.PaperV2CommentParams{}

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
	_, err = client.Papers.V2.Comment(
		ctx,
		cmd.Value("version").(string),
		params,
		options...,
	)
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
		Title:          "papers:v2 comment",
		Transform:      transform,
	})
}
