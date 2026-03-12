// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/apiquery"
	"github.com/AlphaxivCat/alphaxiv_cat-cli/internal/requestflag"
	"github.com/AlphaxivCat/alphaxiv_cat-go"
	"github.com/urfave/cli/v3"
)

var commentsEdit = requestflag.WithInnerFlags(cli.Command{
	Name:    "edit",
	Usage:   "Edit a comment by UUID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "comment",
			Required: true,
		},
		&requestflag.Flag[any]{
			Name:     "anchor-position",
			Required: true,
			BodyPath: "anchorPosition",
		},
		&requestflag.Flag[string]{
			Name:     "body",
			Required: true,
			BodyPath: "body",
		},
		&requestflag.Flag[any]{
			Name:     "focus-position",
			Required: true,
			BodyPath: "focusPosition",
		},
		&requestflag.Flag[any]{
			Name:     "highlight-color",
			Required: true,
			BodyPath: "highlightColor",
		},
		&requestflag.Flag[any]{
			Name:     "highlight-rect",
			Required: true,
			BodyPath: "highlightRects",
		},
		&requestflag.Flag[any]{
			Name:     "selected-text",
			Required: true,
			BodyPath: "selectedText",
		},
		&requestflag.Flag[any]{
			Name:     "title",
			Required: true,
			BodyPath: "title",
		},
	},
	Action:          handleCommentsEdit,
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

func handleCommentsEdit(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("comment") && len(unusedArgs) > 0 {
		cmd.Set("comment", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.CommentEditParams{}

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

	return client.Comments.Edit(
		ctx,
		cmd.Value("comment").(string),
		params,
		options...,
	)
}
