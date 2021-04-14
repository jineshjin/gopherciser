package scenario

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qlik-oss/gopherciser/action"
	"github.com/qlik-oss/gopherciser/connection"
	"github.com/qlik-oss/gopherciser/session"
)

type (
	//SetSenseVariableSettings
	SetSenseVariableSettings struct {
		VariableName  string `json:"name" displayname:"name of the variable" doc-key:"setsensevariable.name"`
		VariableValue string `json:"value" displayname:"value of the variable" doc-key:"setsensevariable.value"`
	}
)

// Validate SetSenseVariable action (Implements ActionSettings interface)
func (settings SetSenseVariableSettings) Validate() ([]string, error) {
	if settings.VariableName == "" {
		return nil, errors.New("No Name specified")
	}
	if settings.VariableValue == "" {
		return nil, errors.New("No Value specified")
	}
	return nil, nil
}

// Execute SetSenseVariable action (Implements ActionSettings interface)
func (settings SetSenseVariableSettings) Execute(sessionState *session.State, actionState *action.State, connection *connection.ConnectionSettings, label string, reset func()) {
	if sessionState.Connection == nil || sessionState.Connection.Sense() == nil {
		actionState.AddErrors(errors.New("Not connected to a Sense environment"))
		return
	}
	uplink := sessionState.Connection.Sense()
	app := uplink.CurrentApp
	if app == nil {
		actionState.AddErrors(errors.New("Not connected to a Sense app"))
		return
	}
	doc := app.Doc

	sessionState.QueueRequest(func(ctx context.Context) error {
		variable, err := varReq(doc.GetVariableByName).WithCache(&uplink.VarCache)(ctx, settings.VariableName)
		if err != nil {
			return errors.WithStack(err)
		}

		if err := variable.SetStringValue(ctx, settings.VariableValue); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}, actionState, true, "Failed to set variable")

	// Send GetApplayout request
	sessionState.QueueRequest(func(ctx context.Context) error {
		_, err := app.Doc.GetAppLayout(ctx)
		return errors.WithStack(err)
	}, actionState, false, "GetAppLayout request failed")

	sessionState.Wait(actionState)
}
