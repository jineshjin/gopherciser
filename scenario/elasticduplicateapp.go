package scenario

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/qlik-oss/gopherciser/action"
	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/elasticstructs"
	"github.com/qlik-oss/gopherciser/helpers"
	"github.com/qlik-oss/gopherciser/session"
)

type (
	// ElasticDuplicateAppSettingsCore Currently used ElasticDuplicateAppSettingsCore (as opposed to deprecated settings)
	ElasticDuplicateAppSettingsCore struct {
		SpaceID string `json:"spaceid" displayname:"Space ID" doc-key:"elasticduplicateapp.spaceid"`
	}
	// ElasticDuplicateAppSettings specify app to duplicate
	ElasticDuplicateAppSettings struct {
		session.AppSelection
		CanAddToCollection
		ElasticDuplicateAppSettingsCore
	}
)

// UnmarshalJSON unmarshals duplicate app settings from JSON
func (settings *ElasticDuplicateAppSettings) UnmarshalJSON(arg []byte) error {
	// Check for deprecated fields
	if err := helpers.HasDeprecatedFields(arg, []string{
		"/appguid",
		"/appname",
	}); err != nil {
		return errors.Errorf("%s %s, please remove from script", ActionElasticDuplicateApp, err.Error())
	}

	var core ElasticDuplicateAppSettingsCore
	if err := jsonit.Unmarshal(arg, &core); err != nil {
		return errors.Wrapf(err, "failed to unmarshal action<%s>", ActionElasticDuplicateApp)
	}

	var appSelection session.AppSelection
	if err := jsonit.Unmarshal(arg, &appSelection); err != nil {
		return errors.Wrapf(err, "failed to unmarshal action<%s>", ActionElasticDuplicateApp)
	}

	var collection CanAddToCollection
	if err := jsonit.Unmarshal(arg, &collection); err != nil {
		return errors.Wrapf(err, "failed to unmarshal action<%s>", ActionElasticDuplicateApp)
	}

	(*settings).ElasticDuplicateAppSettingsCore = core
	(*settings).AppSelection = appSelection
	(*settings).CanAddToCollection = collection

	return nil
}

// Validate action (Implements ActionSettings interface)
func (settings ElasticDuplicateAppSettings) Validate() error {
	if err := settings.AppSelection.Validate(); err != nil {
		return err
	}
	return nil
}

// Execute action (Implements ActionSettings interface)
func (settings ElasticDuplicateAppSettings) Execute(sessionState *session.State, actionState *action.State, connection *connection.ConnectionSettings, label string, reset func()) {
	host, err := connection.GetRestUrl()
	if err != nil {
		actionState.AddErrors(err)
		return
	}

	entry, err := settings.AppSelection.Select(sessionState)
	if err != nil {
		actionState.AddErrors(errors.Wrap(err, "Failed to perform app selection"))
		return
	}

	newName, err := sessionState.ReplaceSessionVariables(&settings.Title)
	if err != nil {
		actionState.AddErrors(errors.Wrap(err, "failed to expand session variables in newname"))
		return
	}

	copyRequest := elasticstructs.PostCopyApp{}
	copyRequest.Attributes.Name = newName
	copyRequest.Attributes.SpaceID = settings.SpaceID

	copyRequestContent, err := jsonit.Marshal(copyRequest)
	if err != nil {
		actionState.AddErrors(errors.Wrap(err, "failed to prepare payload to copy app"))
	}

	copyRequestRest := session.RestRequest{
		Method:      session.POST,
		ContentType: "application/json",
		Destination: fmt.Sprintf("%s/api/v1/apps/%s/copy", host, entry.ID),
		Content:     copyRequestContent,
	}

	sessionState.Rest.QueueRequest(actionState, true, &copyRequestRest, sessionState.LogEntry)
	if sessionState.Wait(actionState) {
		actionState.AddErrors(errors.Wrap(err, "failed to create request to copy app"))
		return
	}
	if copyRequestRest.ResponseStatusCode != http.StatusOK {
		actionState.AddErrors(errors.Errorf("failed to copy app: %d <%s>", copyRequestRest.ResponseStatusCode, copyRequestRest.ResponseBody))
		return
	}

	appImportResponseRaw := copyRequestRest.ResponseBody
	var appImportResponse elasticstructs.AppImportResponse
	if err := jsonit.Unmarshal(appImportResponseRaw, &appImportResponse); err != nil {
		actionState.AddErrors(errors.Wrapf(err, "failed unmarshaling app copy response data: %s", appImportResponseRaw))
		return
	}

	err = AddAppToCollection(settings.CanAddToCollection, sessionState, actionState, appImportResponse, host)
	if err != nil {
		actionState.AddErrors(err)
	}
}
