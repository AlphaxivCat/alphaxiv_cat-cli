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

var notificationsV4ArchiveCreate = cli.Command{
	Name:    "create",
	Usage:   "Given an array of IDs to mark as read, move them into the notification archive.\nInvalid notifications are ignored. By checking the returned count, invalid\nclient-side state can be detected and corrected.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "id",
			Required: true,
			BodyPath: "ids",
		},
	},
	Action:          handleNotificationsV4ArchiveCreate,
	HideHelpCommand: true,
}

var notificationsV4ArchiveList = cli.Command{
	Name:    "list",
	Usage:   "The notification archive is paginated, requiring multiple calls to query long\nhistory.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "before",
			Usage:     "Decimal UNIX time in milliseconds",
			QueryPath: "before",
		},
	},
	Action:          handleNotificationsV4ArchiveList,
	HideHelpCommand: true,
}

func handleNotificationsV4ArchiveCreate(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.NotificationV4ArchiveNewParams{}

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
	_, err = client.Notifications.V4.Archive.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications:v4:archive create", obj, format, transform)
}

func handleNotificationsV4ArchiveList(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.NotificationV4ArchiveListParams{}

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
	_, err = client.Notifications.V4.Archive.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "notifications:v4:archive list", obj, format, transform)
}
