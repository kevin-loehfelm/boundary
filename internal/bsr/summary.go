// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bsr

import (
	"context"
	"fmt"
)

// ChannelSummaryAllocFunc is a function that returns a channel summary type
type ChannelSummaryAllocFunc func(ctx context.Context) ChannelSummary

type channelSummaryAllocFuncRegistry map[Protocol]ChannelSummaryAllocFunc

func (r channelSummaryAllocFuncRegistry) Get(p Protocol) (ChannelSummaryAllocFunc, bool) {
	af, ok := r[p]
	return af, ok
}

var ChannelSummaryAllocFuncs channelSummaryAllocFuncRegistry

// RegisterChannelSummaryAllocFunc registers a ChannelSummaryAllocFunc for the given Protocol.
// A given Protocol can only have one ChannelSummaryAllocFunc function registered.
func RegisterChannelSummaryAllocFunc(p Protocol, af ChannelSummaryAllocFunc) error {
	const op = "bsr.RegisterChannelSummaryAllocFunc"

	if ChannelSummaryAllocFuncs == nil {
		ChannelSummaryAllocFuncs = make(map[Protocol]ChannelSummaryAllocFunc)
	}

	_, ok := ChannelSummaryAllocFuncs[p]
	if ok {
		return fmt.Errorf("%s: %s: %w", op, p, ErrAlreadyRegistered)
	}
	ChannelSummaryAllocFuncs[p] = af
	return nil
}
