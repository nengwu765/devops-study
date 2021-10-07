package main

import (
	"context"
	"fmt"
	"github.com/grafana-tools/sdk"
)

func main() {
	grafanaURL := "http://localhost:3000"
	c, _ := sdk.NewClient(grafanaURL, "eyJrIjoib1A4NUU4RTA0RWhSNmY3ZEEyZUFSUGJpUmROSEtjTGwiLCJuIjoidGVzdC1hcGkiLCJpZCI6MX0=", sdk.DefaultHTTPClient)
	//response, err := c.SetDashboard(context.TODO() ,*board, sdk.SetDashboardParams{
	//	Overwrite: false,
	//})

	board, properties, _ := c.GetDashboardByUID(context.TODO(), "dZJRyjnnz")
	fmt.Printf("board %#v", board)
	fmt.Printf("properties %#v", properties)

	//if err != nil {
	//	fmt.Printf("error on uploading dashboard %s", board.Title)
	//} else {
	//	//fmt.Printf("dashboard URL: %v", grafanaURL+*response.URL)
	//}
}
