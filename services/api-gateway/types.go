package main

import "go-ride/shared/types"

type previewTripRequest struct {
	UserId      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
