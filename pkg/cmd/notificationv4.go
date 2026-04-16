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

var notificationsV4ListNotifications = cli.Command{
	Name:            "list-notifications",
	Usage:           "Returns all active notifications.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleNotificationsV4ListNotifications,
	HideHelpCommand: true,
}

var notificationsV4Subscribe = requestflag.WithInnerFlags(cli.Command{
	Name:    "subscribe",
	Usage:   "Subscribes the current user to push notifications",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "subscription",
			Required: true,
			BodyPath: "subscription",
		},
	},
	Action:          handleNotificationsV4Subscribe,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"subscription": {
		&requestflag.InnerFlag[string]{
			Name:       "subscription.endpoint",
			InnerField: "endpoint",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "subscription.keys",
			InnerField: "keys",
		},
	},
})

var notificationsV4Unsubscribe = cli.Command{
	Name:    "unsubscribe",
	Usage:   "Unsubscribes the current user to push notifications",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "endpoint",
			Required: true,
			BodyPath: "endpoint",
		},
	},
	Action:          handleNotificationsV4Unsubscribe,
	HideHelpCommand: true,
}

func handleNotificationsV4ListNotifications(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Notifications.V4.ListNotifications(ctx, options...)
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
		Title:          "notifications:v4 list-notifications",
		Transform:      transform,
	})
}

func handleNotificationsV4Subscribe(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.NotificationV4SubscribeParams{}

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

	return client.Notifications.V4.Subscribe(ctx, params, options...)
}

func handleNotificationsV4Unsubscribe(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.NotificationV4UnsubscribeParams{}

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

	return client.Notifications.V4.Unsubscribe(ctx, params, options...)
}
