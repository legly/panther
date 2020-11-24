// Code generated by apigen; DO NOT EDIT.
// package logtypesapi documents github.com/panther-labs/panther/internal/core/logtypesapi.LogTypesAPI
package logtypesapi

// LogTypesAPI available endpoints
type LogTypesAPI interface {
	ListAvailableLogTypes() (ListAvailableLogTypesResponse, error)
}

// Models for LogTypesAPI

// LogTypesAPIPayload is the payload for calls to LogTypesAPI endpoints.
type LogTypesAPIPayload struct {
	ListAvailableLogTypes *struct{}
}

type ListAvailableLogTypesResponse struct {
	LogTypes []string `json:"logTypes"`
}
