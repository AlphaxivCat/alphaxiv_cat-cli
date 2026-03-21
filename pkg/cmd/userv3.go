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

var usersV3AutocompleteProfile = cli.Command{
	Name:    "autocomplete-profile",
	Usage:   "Generate a biography and institution for a user using their claimed papers",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "user-id",
			Required: true,
			BodyPath: "userId",
		},
	},
	Action:          handleUsersV3AutocompleteProfile,
	HideHelpCommand: true,
}

var usersV3DeleteBanner = cli.Command{
	Name:    "delete-banner",
	Usage:   "Delete the given banner",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "banner-id",
			Required: true,
		},
	},
	Action:          handleUsersV3DeleteBanner,
	HideHelpCommand: true,
}

var usersV3DeleteOwnUser = cli.Command{
	Name:            "delete-own-user",
	Usage:           "Deletes the user's account",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsersV3DeleteOwnUser,
	HideHelpCommand: true,
}

var usersV3GetActivity = cli.Command{
	Name:    "get-activity",
	Usage:   "Retrieve public activity timeline for a user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     `Allowed values: "date", "liked".`,
			Default:   "date",
			QueryPath: "sort",
		},
	},
	Action:          handleUsersV3GetActivity,
	HideHelpCommand: true,
}

var usersV3GetClaimedPapers = cli.Command{
	Name:    "get-claimed-papers",
	Usage:   "Retrieve the claimed papers for a user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     `Allowed values: "date", "liked", "citations".`,
			QueryPath: "sort",
		},
	},
	Action:          handleUsersV3GetClaimedPapers,
	HideHelpCommand: true,
}

var usersV3GetCurrentUser = cli.Command{
	Name:            "get-current-user",
	Usage:           "Retrieve information about yourself",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsersV3GetCurrentUser,
	HideHelpCommand: true,
}

var usersV3GetFeaturedActivity = cli.Command{
	Name:    "get-featured-activity",
	Usage:   "Retrieve highlighted activity for a user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3GetFeaturedActivity,
	HideHelpCommand: true,
}

var usersV3GetFollowers = cli.Command{
	Name:    "get-followers",
	Usage:   "List the users following the specified user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3GetFollowers,
	HideHelpCommand: true,
}

var usersV3GetLeaderboard = cli.Command{
	Name:            "get-leaderboard",
	Usage:           "Retrieve weekly and all-time leaderboards for users ranked by reputation",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsersV3GetLeaderboard,
	HideHelpCommand: true,
}

var usersV3GetUserByUuid = cli.Command{
	Name:    "get-user-by-uuid",
	Usage:   "Retrieve a user's basic information given its UUID",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3GetUserByUuid,
	HideHelpCommand: true,
}

var usersV3GetViewedHistory = cli.Command{
	Name:    "get-viewed-history",
	Usage:   "Retrieve the view history for the current user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "offset",
			QueryPath: "offset",
		},
		&requestflag.Flag[string]{
			Name:      "search",
			QueryPath: "search",
		},
	},
	Action:          handleUsersV3GetViewedHistory,
	HideHelpCommand: true,
}

var usersV3ProcessNotificationEmail = cli.Command{
	Name:    "process-notification-email",
	Usage:   "Send a notification digest email for the given user when necessary.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3ProcessNotificationEmail,
	HideHelpCommand: true,
}

var usersV3Search = cli.Command{
	Name:    "search",
	Usage:   "Search for users by name, username, or institution",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "q",
			Required:  true,
			QueryPath: "q",
		},
		&requestflag.Flag[string]{
			Name:      "limit",
			QueryPath: "limit",
		},
	},
	Action:          handleUsersV3Search,
	HideHelpCommand: true,
}

var usersV3ToggleFollowUser = cli.Command{
	Name:    "toggle-follow-user",
	Usage:   "Follow or unfollow another user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleUsersV3ToggleFollowUser,
	HideHelpCommand: true,
}

var usersV3UpdatePreferences = requestflag.WithInnerFlags(cli.Command{
	Name:    "update-preferences",
	Usage:   "Update base or banner preferences for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "banner",
			BodyPath: "banners",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "base",
			BodyPath: "base",
		},
	},
	Action:          handleUsersV3UpdatePreferences,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"banner": {
		&requestflag.InnerFlag[any]{
			Name:       "banner.link",
			InnerField: "link",
		},
		&requestflag.InnerFlag[string]{
			Name:       "banner.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "banner.type",
			Usage:      `Allowed values: "success", "info", "warning", "error".`,
			InnerField: "type",
		},
	},
	"base": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "base.assistant-custom-styles",
			InnerField: "assistantCustomStyles",
		},
		&requestflag.InnerFlag[string]{
			Name:       "base.assistant-style-selection",
			InnerField: "assistantStyleSelection",
		},
		&requestflag.InnerFlag[any]{
			Name:       "base.default-private-paper-sidebar-tab",
			Usage:      `Allowed values: "assistant", "notes", "similar".`,
			InnerField: "defaultPrivatePaperSidebarTab",
		},
		&requestflag.InnerFlag[any]{
			Name:       "base.default-public-paper-sidebar-tab",
			Usage:      `Allowed values: "comments", "assistant", "similar", "notes", "social".`,
			InnerField: "defaultPublicPaperSidebarTab",
		},
		&requestflag.InnerFlag[string]{
			Name:       "base.feed-sort",
			Usage:      `Allowed values: "Hot", "Comments", "Views", "Likes", "GitHub", "Twitter (X)", "Recommended".`,
			InnerField: "feedSort",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "base.is-dark-mode-enabled",
			InnerField: "isDarkModeEnabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "base.is-debug-mode-enabled",
			InnerField: "isDebugModeEnabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "base.is-members-sidebar-visible",
			InnerField: "isMembersSidebarVisible",
		},
		&requestflag.InnerFlag[any]{
			Name:       "base.preferred-language",
			Usage:      `Allowed values: "am", "ar", "az", "bg", "bn", "ca", "cs", "da", "de", "el", "en", "es", "et", "fa", "fi", "fr", "gu", "ha", "he", "hi", "hr", "hu", "id", "it", "ja", "ka", "kn", "ko", "lt", "lv", "ml", "mr", "ms", "my", "ne", "nl", "no", "pa", "pl", "pt", "ro", "ru", "si", "sk", "sl", "sr", "sv", "sw", "ta", "te", "th", "tl", "tr", "uk", "ur", "uz", "vi", "yo", "zh".`,
			InnerField: "preferredLanguage",
		},
		&requestflag.InnerFlag[any]{
			Name:       "base.preferred-llm-model",
			InnerField: "preferredLlmModel",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "base.reading-mode-enabled",
			InnerField: "readingModeEnabled",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "base.show-model-thinking",
			InnerField: "showModelThinking",
		},
		&requestflag.InnerFlag[any]{
			Name:       "base.tooling-pane-width",
			InnerField: "toolingPaneWidth",
		},
		&requestflag.InnerFlag[string]{
			Name:       "base.web-search",
			Usage:      `Allowed values: "off", "full".`,
			InnerField: "webSearch",
		},
	},
})

var usersV3UpdateProfile = cli.Command{
	Name:    "update-profile",
	Usage:   "Update profile details for the authenticated user",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "biography",
			BodyPath: "biography",
		},
		&requestflag.Flag[any]{
			Name:     "bluesky-username",
			BodyPath: "blueskyUsername",
		},
		&requestflag.Flag[any]{
			Name:     "github-username",
			BodyPath: "githubUsername",
		},
		&requestflag.Flag[any]{
			Name:     "institution",
			BodyPath: "institution",
		},
		&requestflag.Flag[any]{
			Name:     "linkedin-username",
			BodyPath: "linkedinUsername",
		},
		&requestflag.Flag[any]{
			Name:     "location",
			BodyPath: "location",
		},
		&requestflag.Flag[any]{
			Name:     "public-email",
			BodyPath: "publicEmail",
		},
		&requestflag.Flag[string]{
			Name:     "real-name",
			BodyPath: "realName",
		},
		&requestflag.Flag[string]{
			Name:     "username",
			BodyPath: "username",
		},
		&requestflag.Flag[any]{
			Name:     "x-username",
			BodyPath: "xUsername",
		},
	},
	Action:          handleUsersV3UpdateProfile,
	HideHelpCommand: true,
}

var usersV3UploadAvatar = cli.Command{
	Name:            "upload-avatar",
	Usage:           "Upload or remove the authenticated user's avatar image.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleUsersV3UploadAvatar,
	HideHelpCommand: true,
}

func handleUsersV3AutocompleteProfile(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3AutocompleteProfileParams{}

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
	_, err = client.Users.V3.AutocompleteProfile(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 autocomplete-profile", obj, format, transform)
}

func handleUsersV3DeleteBanner(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("banner-id") && len(unusedArgs) > 0 {
		cmd.Set("banner-id", unusedArgs[0])
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

	return client.Users.V3.DeleteBanner(ctx, cmd.Value("banner-id").(string), options...)
}

func handleUsersV3DeleteOwnUser(ctx context.Context, cmd *cli.Command) error {
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

	return client.Users.V3.DeleteOwnUser(ctx, options...)
}

func handleUsersV3GetActivity(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3GetActivityParams{}

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
	_, err = client.Users.V3.GetActivity(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-activity", obj, format, transform)
}

func handleUsersV3GetClaimedPapers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3GetClaimedPapersParams{}

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
	_, err = client.Users.V3.GetClaimedPapers(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-claimed-papers", obj, format, transform)
}

func handleUsersV3GetCurrentUser(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Users.V3.GetCurrentUser(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-current-user", obj, format, transform)
}

func handleUsersV3GetFeaturedActivity(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.GetFeaturedActivity(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-featured-activity", obj, format, transform)
}

func handleUsersV3GetFollowers(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.GetFollowers(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-followers", obj, format, transform)
}

func handleUsersV3GetLeaderboard(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Users.V3.GetLeaderboard(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-leaderboard", obj, format, transform)
}

func handleUsersV3GetUserByUuid(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.GetUserByUuid(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-user-by-uuid", obj, format, transform)
}

func handleUsersV3GetViewedHistory(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3GetViewedHistoryParams{}

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
	_, err = client.Users.V3.GetViewedHistory(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 get-viewed-history", obj, format, transform)
}

func handleUsersV3ProcessNotificationEmail(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.ProcessNotificationEmail(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 process-notification-email", obj, format, transform)
}

func handleUsersV3Search(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3SearchParams{}

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
	_, err = client.Users.V3.Search(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 search", obj, format, transform)
}

func handleUsersV3ToggleFollowUser(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Users.V3.ToggleFollowUser(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 toggle-follow-user", obj, format, transform)
}

func handleUsersV3UpdatePreferences(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3UpdatePreferencesParams{}

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
	_, err = client.Users.V3.UpdatePreferences(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 update-preferences", obj, format, transform)
}

func handleUsersV3UpdateProfile(ctx context.Context, cmd *cli.Command) error {
	client := alphaxivcat.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := alphaxivcat.UserV3UpdateProfileParams{}

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
	_, err = client.Users.V3.UpdateProfile(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 update-profile", obj, format, transform)
}

func handleUsersV3UploadAvatar(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Users.V3.UploadAvatar(ctx, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "users:v3 upload-avatar", obj, format, transform)
}
