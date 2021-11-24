package ulink

import "fmt"

func NewSimpleUlink(teamID string, bundleID string) *AppleUniversalLink {
	return &AppleUniversalLink{
		Applinks: Applinks{
			Apps: []interface{}{},
			Details: []AppleDetail{
				{
					AppID: fmt.Sprintf("%v.%v", teamID, bundleID),
					Paths: []string{"*"},
				},
			},
		},
	}
}

type AppleUniversalLink struct {
	Applinks Applinks `json:"applinks"`
}

type Applinks struct {
	Apps    []interface{} `json:"apps"`
	Details []AppleDetail `json:"details"`
}

type AppleDetail struct {
	AppID string   `json:"appID"`
	Paths []string `json:"paths"`
}
