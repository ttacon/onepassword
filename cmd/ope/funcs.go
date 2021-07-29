package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/ttacon/onepassword"
	"github.com/urfave/cli/v2"
)

func getClient() (onepassword.Client, error) {
	token := os.Getenv("ONEPASS_TOKEN")
	if len(token) == 0 {
		return nil, errors.New("no authentication token provided")
	}

	client, err := onepassword.NewHTTPClient(
		nil,
		token,
		"events.1password.com",
	)
	return client, err
}

func introspectFunc(c *cli.Context) error {
	client, err := getClient()
	if err != nil {
		fmt.Println("failed to create client: ", err)
		os.Exit(1)
	}

	val, err := client.Service().Introspect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf(`
Token UUID: %s
Issued At:  %s
Features:
 - %s

`,
		val.UUID,
		val.IssuedAt,
		strings.Join(val.Features, "\n - "),
	)

	return nil
}

func itemUsageFunc(c *cli.Context) error {
	client, err := getClient()
	if err != nil {
		fmt.Println("failed to create client: ", err)
		os.Exit(1)
	}

	start := time.Now().Add(-1 * 24 * time.Hour)
	resetCursor := &onepassword.ResetCursor{
		Limit:     100,
		StartTime: &start,
	}
	currCursor := ""

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{
		"UUID",
		"Timestamp",
		"Version",
		"VaultUUID",
		"User",
		"ClientInfo",
	})
	tw.SetTitle("Item usage report")

	for {
		resp, err := client.Service().GetItemUsages(resetCursor, currCursor)
		if err != nil {
			fmt.Println("failed to retrieve item usage report, err: ", err)
			os.Exit(1)
		}

		for _, item := range resp.Items {
			tw.AppendRows([]table.Row{
				{
					item.UUID,
					item.Timestamp,
					item.UsedVersion,
					item.VaultUUID,
					item.ItemUUID,
					item.User.Email,
					item.ClientInfo.IPAddress,
				},
			})
		}

		if resetCursor != nil {
			// We only need this on the first call to reset our token's place
			// in the event stream. So once we've used it, set it to nil so
			// that we don't accidentally reset our position other than in the
			// first call.
			resetCursor = nil
		}

		if !resp.HasMore {
			break
		}
		currCursor = resp.Cursor
	}

	fmt.Println(tw.Render())

	return nil
}

func loginFunc(c *cli.Context) error {
	client, err := getClient()
	if err != nil {
		fmt.Println("failed to create client: ", err)
		os.Exit(1)
	}

	start := time.Now().Add(-1 * 24 * time.Hour)
	resetCursor := &onepassword.ResetCursor{
		Limit:     100,
		StartTime: &start,
	}
	currCursor := ""

	tw := table.NewWriter()
	tw.AppendHeader(table.Row{
		"UUID",
		"SessionUUID",
		"Timestamp",
		"Category",
		"Type",
		"TargetUser",
		"IP Address",
	})
	tw.SetTitle("Sign in attempts report")

	for {
		resp, err := client.Service().GetSignInAttempts(resetCursor, currCursor)
		if err != nil {
			fmt.Println("failed to retrieve item usage report, err: ", err)
			os.Exit(1)
		}

		for _, item := range resp.Items {
			tw.AppendRows([]table.Row{
				{
					item.UUID,
					item.SessionUUID,
					item.Timestamp,
					item.Category,
					item.Type,
					item.TargetUser.Name,
					item.ClientInfo.IPAddress,
				},
			})
		}

		if resetCursor != nil {
			// We only need this on the first call to reset our token's place
			// in the event stream. So once we've used it, set it to nil so
			// that we don't accidentally reset our position other than in the
			// first call.
			resetCursor = nil
		}

		if !resp.HasMore {
			break
		}
		currCursor = resp.Cursor
	}

	fmt.Println(tw.Render())

	return nil
}
