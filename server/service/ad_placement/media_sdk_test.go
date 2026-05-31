package ad_placement

import (
	"fmt"
	media_sdk "github.com/cngamesdk/media-sdk"
	config2 "github.com/cngamesdk/media-sdk/config"
	"github.com/cngamesdk/media-sdk/model"
	"testing"
)

func TestMediaSdkAuth(t *testing.T) {
	config := config2.DefaultConfig(config2.MediaToutiao)
	client, clientErr := media_sdk.NewClient(config)
	if clientErr != nil {
		t.Fatal(clientErr)
	}
	req := &model.AuthReq{}
	req.AppId = 123
	req.RedirectUri = "https://www.test.com"
	resp, err := client.Auth(req)
	if err != nil {
		t.Fatal(err)
	}
	println(fmt.Sprintf("%+v", resp))
}
