package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ttacon/onepassword"
)

var (
	host  = flag.String("host", "", "1password host")
	token = flag.String("token", "", "1password token to use")
)

func main() {
	flag.Parse()

	client, err := onepassword.NewHTTPClient(
		nil,
		*token,
		*host,
	)
	if err != nil {
		fmt.Println("failed to create client: ", err)
		os.Exit(1)
	}

	val, err := client.Service().Introspect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(val.IssuedAt)

	startTime := time.Now().Add(-1 * time.Hour * 24)
	intro, err := client.Service().GetSignInAttempts(&onepassword.ResetCursor{
		Limit:     100,
		StartTime: &startTime,
	}, "")
	if err != nil {
		fmt.Println("failed to make introspection request", err)
		os.Exit(1)
	}

	fmt.Println(intro)

}
