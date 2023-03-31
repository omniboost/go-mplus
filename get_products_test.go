package mplus_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetProducts(t *testing.T) {
	req := client.NewGetProductsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
