// Code generated by jrpc. DO NOT EDIT.

package scenedef

import (
	"encoding/json"

	"github.com/jakewright/home-automation/libraries/go/errors"
	"github.com/jakewright/home-automation/libraries/go/firehose"
)

// Publish publishes the event to the Firehose
func (m *SetSceneEvent) Publish() error {
	if err := m.Validate(); err != nil {
		return err
	}

	return firehose.Publish("set-scene", m)
}

// SetSceneEventHandler implements the necessary functions to be a Firehose handler
type SetSceneEventHandler func(*SetSceneEvent) firehose.Result

// EventName returns the Firehose channel name
func (h SetSceneEventHandler) EventName() string {
	return "set-scene"
}

// HandleEvent handles the Firehose event
func (h SetSceneEventHandler) HandleEvent(e *firehose.Event) firehose.Result {
	var body SetSceneEvent
	if err := json.Unmarshal(e.Payload, &body); err != nil {
		return firehose.Discard(errors.WithMessage(err, "failed to unmarshal payload"))
	}
	return h(&body)
}