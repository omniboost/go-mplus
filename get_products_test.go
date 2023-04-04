package mplus_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestGetProducts(t *testing.T) {
	req := client.NewGetProductsRequest()
	// client.SetDebug(false)
	req.RequestBody().Request.OnlyActive = false
	req.RequestBody().Request.SyncMarker = -1
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))

	syncMarker := 0
	for _, v := range resp.ProductList.Product {
		if v.SyncMarker > syncMarker {
			syncMarker = v.SyncMarker
		}
	}

	log.Fatal(syncMarker)
}
