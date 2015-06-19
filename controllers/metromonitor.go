package controllers

import (
           	"fmt"
	godClient "github.com/zond/god/client"
	godCommon "github.com/zond/god/common"
	godMurmur "github.com/zond/god/murmur"
	metroapi  "github.com/Skovy/metro-api-go"
)

var _ = fmt.Println

const ProvidersKey = "providers"
var GodConn = godClient.MustConn("localhost:9191")

func Run() {
	listOfProviders := metroapi.GetProviders()
	providersBytes := godCommon.MustJSONEncode(listOfProviders)
	GodConn.Put(godMurmur.HashString(ProvidersKey), providersBytes)

  var providers []metroapi.Providers
	providersData, keyExists := GodConn.Get(godMurmur.HashString(ProvidersKey))
	if keyExists {
		godCommon.MustJSONDecode(providersData, &providers)
		fmt.Println(providers[0])
	}
}
