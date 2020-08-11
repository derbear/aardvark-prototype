// Code generated by "stringer -type=actionType"; DO NOT EDIT.

package agreement

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[noop-0]
	_ = x[ignore-1]
	_ = x[broadcast-2]
	_ = x[relay-3]
	_ = x[disconnect-4]
	_ = x[broadcastVotes-5]
	_ = x[verifyVote-6]
	_ = x[verifyPayload-7]
	_ = x[verifyBundle-8]
	_ = x[ensure-9]
	_ = x[rezero-10]
	_ = x[attest-11]
	_ = x[assemble-12]
	_ = x[repropose-13]
	_ = x[checkpoint-14]
}

const _actionType_name = "noopignorebroadcastrelaydisconnectbroadcastVotesverifyVoteverifyPayloadverifyBundleensurerezeroattestassemblereproposecheckpoint"

var _actionType_index = [...]uint8{0, 4, 10, 19, 24, 34, 48, 58, 71, 83, 89, 95, 101, 109, 118, 128}

func (i actionType) String() string {
	if i < 0 || i >= actionType(len(_actionType_index)-1) {
		return "actionType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _actionType_name[_actionType_index[i]:_actionType_index[i+1]]
}
