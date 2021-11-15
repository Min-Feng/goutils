// Code generated by "stringer -type=TriggerKind -linecomment -output=trigger_kind_string.go"; DO NOT EDIT.

package xLog

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TriggerKindUnknown-0]
	_ = x[TriggerKindHttp-1]
	_ = x[TriggerKindWebsocket-2]
	_ = x[TriggerKindGrpc-3]
	_ = x[TriggerKindMessageBroker-4]
}

const _TriggerKind_name = "unknownhttpwebsocketgrpcmq"

var _TriggerKind_index = [...]uint8{0, 7, 11, 20, 24, 26}

func (i TriggerKind) String() string {
	if i < 0 || i >= TriggerKind(len(_TriggerKind_index)-1) {
		return "TriggerKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TriggerKind_name[_TriggerKind_index[i]:_TriggerKind_index[i+1]]
}