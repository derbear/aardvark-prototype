// Copyright (C) 2019 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package protocol

// Tag represents a message type identifier.  Messages have a Tag field. Handlers can register to a given Tag.
// e.g., the agreement service can register to handle agreements with the Agreement tag.
type Tag string

// Tags, in lexicographic sort order to avoid duplicates.
const (
	UnknownMsgTag      Tag = "??"
	AgreementVoteTag   Tag = "AV"
	MsgSkipTag         Tag = "MS"
	NetPrioResponseTag Tag = "NP"
	PingTag            Tag = "pi"
	PingReplyTag       Tag = "pj"
	ProposalPayloadTag Tag = "PP"
	TxnTag             Tag = "TX"
	UniAccountReqTag   Tag = "UA"
	UniCatchupReqTag   Tag = "UC"
	UniDecommitReqTag  Tag = "UD"
	UniEnsBlockReqTag  Tag = "UE"
	UniDecommitResTag  Tag = "UM"
	UniAccountResTag   Tag = "UP"
	UniEnsBlockResTag  Tag = "US"
	UniCatchupResTag   Tag = "UT"
	VoteBundleTag      Tag = "VB"
)

// Complement is a convenience function for returning a corresponding response/request tag
func (t Tag) Complement() Tag {
	switch t {
	case UniCatchupResTag:
		return UniCatchupReqTag
	case UniCatchupReqTag:
		return UniCatchupResTag
	case UniEnsBlockResTag:
		return UniEnsBlockReqTag
	case UniEnsBlockReqTag:
		return UniEnsBlockResTag
	case UniAccountResTag:
		return UniAccountReqTag
	case UniAccountReqTag:
		return UniAccountResTag
	case UniDecommitResTag:
		return UniDecommitReqTag
	case UniDecommitReqTag:
		return UniDecommitResTag
	default:
		return UnknownMsgTag
	}
}
