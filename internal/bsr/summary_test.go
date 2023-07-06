// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bsr_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/hashicorp/boundary/internal/bsr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegisterChannelSummaryAllocFunc(t *testing.T) {
	ctx := context.Background()
	startTime := time.Now()
	endTime := time.Now()

	cases := []struct {
		name            string
		p               bsr.Protocol
		getP            bsr.Protocol
		cf              bsr.ChannelSummaryAllocFunc
		want            *bsr.BaseChannelSummary
		wantRegisterErr error
		wantGetErr      bool
	}{
		{
			"valid",
			bsr.Protocol("TEST_PROTOCOL"),
			bsr.Protocol("TEST_PROTOCOL"),
			func(ctx context.Context) bsr.ChannelSummary {
				return &bsr.BaseChannelSummary{
					Id:                    "TEST_ID",
					ConnectionRecordingId: "TEST_CONNECTION_RECORDING_ID",
					StartTime:             startTime,
					EndTime:               endTime,
				}
			},
			&bsr.BaseChannelSummary{
				Id:                    "TEST_ID",
				ConnectionRecordingId: "TEST_CONNECTION_RECORDING_ID",
				StartTime:             startTime,
				EndTime:               endTime,
			},
			nil,
			false,
		},
		{
			"already-registered-protocol",
			bsr.Protocol("TEST_PROTOCOL"),
			bsr.Protocol("TEST_PROTOCOL"),
			nil,
			&bsr.BaseChannelSummary{},
			errors.New("bsr.RegisterChannelSummaryAllocFunc: TEST_PROTOCOL: type already registered"),
			false,
		},
		{
			"invalid-protocol",
			bsr.Protocol("TEST_PROTOCOL_2"),
			bsr.Protocol("TEST_INVALID_PROTOCOL"),
			nil,
			&bsr.BaseChannelSummary{},
			nil,
			true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := bsr.RegisterChannelSummaryAllocFunc(tc.p, tc.cf)
			if tc.wantRegisterErr != nil {
				assert.EqualError(t, tc.wantRegisterErr, err.Error())
				return
			}
			require.NoError(t, err)

			af, ok := bsr.ChannelSummaryAllocFuncs.Get(tc.getP)
			if tc.wantGetErr {
				require.False(t, ok, "found invalid channel summary")
				return
			}
			require.True(t, ok, "could not get channel summary")

			got := af(ctx)

			assert.Equal(t, tc.want.GetId(), got.GetId())
			assert.Equal(t, tc.want.GetConnectionRecordingId(), got.GetConnectionRecordingId())
			assert.Equal(t, tc.want.GetChannelType(), got.GetChannelType())
			assert.Equal(t, tc.want.GetStartTime(), got.GetStartTime())
			assert.Equal(t, tc.want.GetEndTime(), got.GetEndTime())
			assert.Equal(t, tc.want.GetBytesUp(), got.GetBytesUp())
			assert.Equal(t, tc.want.GetBytesDown(), got.GetBytesDown())
		})
	}
}
