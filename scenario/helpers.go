package scenario

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qlik-oss/gopherciser/action"
	"github.com/qlik-oss/gopherciser/enigmahandlers"
	"github.com/qlik-oss/gopherciser/logger"
	"github.com/qlik-oss/gopherciser/senseobjects"
	"github.com/qlik-oss/gopherciser/session"
)

func subscribeSheetObjectsAsync(sessionState *session.State, actionState *action.State, app *senseobjects.App, sheetID string) error {
	sheetID = sessionState.IDMap.Get(sheetID)
	sheetEntry, err := GetSheetEntry(sessionState, actionState, app, sheetID)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to objects")
	}

	for _, v := range sheetEntry.Data.Cells {
		sessionState.LogEntry.LogDebugf("subscribe to object<%s> type<%s>", v.Name, v.Type)
		session.GetAndAddObjectAsync(sessionState, actionState, v.Name)
	}

	return nil
}

func GetSheetEntry(sessionState *session.State, actionState *action.State, app *senseobjects.App, sheetid string) (*senseobjects.SheetNxContainerEntry, error) {
	sheetList, err := app.GetSheetList(sessionState, actionState)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return sheetList.GetSheetEntry(sheetid)
}

// GetCurrentSheet from objects
func GetCurrentSheet(uplink *enigmahandlers.SenseUplink) (*senseobjects.Sheet, error) {
	sheets := uplink.Objects.GetObjectsOfType(enigmahandlers.ObjTypeSheet)
	if len(sheets) < 1 {
		return nil, errors.New("no current sheet found")
	}
	if len(sheets) > 1 {
		return nil, errors.Errorf("%d current sheets found", len(sheets))
	}
	sheetObj, ok := sheets[0].EnigmaObject.(*senseobjects.Sheet)
	if !ok {
		return nil, errors.Errorf("failed to cast object id<%s> to sheet object", sheetObj.GenericId)
	}
	return sheetObj, nil
}

// ClearObjectSubscriptions and currently subscribed objects
func ClearObjectSubscriptions(sessionState *session.State) {
	upLink := sessionState.Connection.Sense()
	// Clear subscribed objects
	clearedObjects, errClearObject := upLink.Objects.ClearObjectsOfType(enigmahandlers.ObjTypeGenericObject)
	if errClearObject != nil {
		sessionState.LogEntry.Log(logger.WarningLevel, clearedObjects)
	}
	sessionState.DeRegisterEvents(clearedObjects)

	// Clear any sheets set
	clearedObjects, errClearObject = upLink.Objects.ClearObjectsOfType(enigmahandlers.ObjTypeSheet)
	if errClearObject != nil {
		sessionState.LogEntry.Log(logger.WarningLevel, clearedObjects)
	}
	sessionState.DeRegisterEvents(clearedObjects)
}

func DebugPrintObjectSubscriptions(sessionState *session.State) {
	if !sessionState.LogEntry.ShouldLogDebug() {
		return
	}

	upLink := sessionState.Connection.Sense()
	objectsPointers := upLink.Objects.GetObjectsOfType(enigmahandlers.ObjTypeGenericObject)
	objects := make([]string, 0, len(objectsPointers))
	for _, object := range objectsPointers {
		if object == nil {
			continue
		}
		objects = append(objects, object.ID)
	}
	sessionState.LogEntry.LogDebug(fmt.Sprintf("current object subscriptions: %v", objects))
}

// Contains check whether any element in the supplied list matches (match func(s string) bool)
func Contains(list []string, match func(s string) bool) bool {
	for _, item := range list {
		if match(item) {
			return true
		}
	}
	return false
}
