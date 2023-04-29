package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const api_path = "https://www.bungie.net/Platform"

func BungieGet(uri string) []byte {
	apikey := os.Getenv("APIKEY")
	composite_uri := fmt.Sprintf("%s/%s", api_path, uri)
	req, err := http.NewRequest("GET", composite_uri, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("X-API-Key", apikey)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	return body
}

func BungiePost(uri string) []byte {
	apikey := os.Getenv("APIKEY")
	composite_uri := fmt.Sprintf("%s/%s", api_path, uri)
	req, err := http.NewRequest("POST", composite_uri, nil)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("X-API-Key", apikey)
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}
	return body
}

type BungieResponse struct {
	ErrorCode       int32  `json:"ErrorCode"`
	ThrottleSeconds int32  `json:"ThrottleSeconds"`
	ErrorStatus     string `json:"ErrorStatus"`
	Message         string `json:"Message"`
}

type GroupResponse struct {
	BungieResponse
	Response struct {
		Detail struct {
			GroupId string `json:"groupId"`
		} `json:"detail"`
	} `json:"Response"`
}

type GroupMembers struct {
	BungieResponse
	Response struct {
		Results []GroupMember `json:"results"`
	} `json:"Response"`
}

type GroupMember struct {
	MemberType      int32 `json:"memberType"`
	DestinyUserInfo struct {
		LastSeenDisplayName string `json:"lastSeenDisplayName"`
		MembershipId        string `json:"membershipId"`
		MembershipType      int32  `json:"membershipType"`
	} `json:"destinyUserInfo"`
}

type Profile struct {
	BungieResponse
	Response struct {
		Profile struct {
			Data struct {
				CharacterIds []string
			} `json:"data"`
		} `json:"profile"`
	} `json:"Response"`
}

type DestinyActivityHistoryResults struct {
	BungieResponse
	Response struct {
		Activities []struct {
			Period          string `json:"period"`
			ActivityDetails struct {
				ReferenceId          int    `json:"referenceId"`
				InstanceId           string `json:"instanceId"`
				DirectorActivityHash int    `json:"directorActivityHash"`
			} `json:"activityDetails"`
			Values struct {
				ActivityDurationSeconds struct {
					StatId string `json:"statId"`
					Basic  struct {
						Value float64 `json:"value"`
					} `json:"basic"`
				} `json:"activityDurationSeconds"`
			} `json:"values"`
		} `json:"activities"`
	} `json:"Response"`
}

type EntityDefinition struct {
	BungieResponse
	Response struct {
		DisplayProperties struct {
			Name string `json:"name"`
		} `json:"displayProperties"`
	} `json:"Response"`
}

func GetGroupByName(name string, grouptype int32) string {
	groupResponse := &GroupResponse{}
	group_uri := fmt.Sprintf("/GroupV2/Name/%s/%d/", name, grouptype)
	bget := BungieGet(group_uri)
	json.Unmarshal(bget, groupResponse)
	return groupResponse.Response.Detail.GroupId
}

func GetGroupMembers(groupId string) GroupMembers {
	groupMembers := &GroupMembers{}
	group_members_uri := fmt.Sprintf("/GroupV2/%s/Members/", groupId)
	bget := BungieGet(group_members_uri)
	json.Unmarshal(bget, &groupMembers)
	return *groupMembers
}

func GetProfile(membership_type int32, membership_id string) Profile {
	profile := &Profile{}
	get_profile_uri := fmt.Sprintf("/Destiny2/%d/Profile/%s/?components=Profiles,Characters", membership_type, membership_id)
	bget := BungieGet(get_profile_uri)
	json.Unmarshal(bget, &profile)
	return *profile
}

func GetPlayerActivityHistory(membershipType int32, membershipId string, characterId string) []DestinyActivityHistoryResults {
	actHistories := []DestinyActivityHistoryResults{}

	page := 0
	for {
		actHistory := &DestinyActivityHistoryResults{}
		act_history_uri := fmt.Sprintf("/Destiny2/%d/Account/%s/Character/%s/Stats/Activities/?count=100&mode=4&page=%d", membershipType, membershipId, characterId, page)
		log.Println(act_history_uri)
		bget := BungieGet(act_history_uri)
		json.Unmarshal(bget, &actHistory)
		actHistories = append(actHistories, *actHistory)
		if len(actHistory.Response.Activities) < 100 {
			break
		} else {
			page++
		}
	}
	return actHistories
}

func GetDestinyEntityDefinition(refernceId string) EntityDefinition {
	edef := EntityDefinition{}
	uri := fmt.Sprintf("/Destiny2/Manifest/DestinyActivityDefinition/%s", refernceId)
	bget := BungieGet(uri)
	fmt.Println(string(bget))
	json.Unmarshal(bget, &edef)
	return edef
}
