// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bsr

import (
	"context"
	"time"
)

type (
	// SessionSummary encapsulates data for a session, including its session id, connection count,
	// and start/end time using a monotonic clock
	SessionSummary struct {
		Id              string
		ConnectionCount uint64
		StartTime       time.Time
		EndTime         time.Time
		Errors          error
	}

	// ConnectionSummary encapsulates data for a connection, including its connection id, channel count,
	// start/end time using a monotonic clock, and the aggregate bytes up/ down of its channels
	ConnectionSummary struct {
		Id           string
		ChannelCount uint64
		StartTime    time.Time
		EndTime      time.Time
		BytesUp      uint64
		BytesDown    uint64
		Errors       error
	}

	// BaseChannelSummary encapsulates data for a channel, including its id, channel type,
	// start/end time using a monotonic clock, and the bytes up/ down seen on this channel
	BaseChannelSummary struct {
		Id                    string
		ConnectionRecordingId string
		StartTime             time.Time
		EndTime               time.Time
		BytesUp               uint64
		BytesDown             uint64
		ChannelType           string
	}

	ChannelSummary interface {
		// GetId returns the Id of the summary file.
		GetId() string
		// GetConnectionRecordingId returns the id of the connection
		GetConnectionRecordingId() string
		// GetStartTime returns the start time using a monotonic clock of the summary.
		GetStartTime() time.Time
		// GetEndTime returns the end time using a monotonic clock of the summary.
		GetEndTime() time.Time
		// GetBytesUp returns upload bytes.
		GetBytesUp() uint64
		// BytesDown returns download bytes.
		GetBytesDown() uint64
		// GetChannelType the type of summary channel.
		GetChannelType() string
	}
)

func AllocSessionSummary(_ context.Context) any {
	return &SessionSummary{}
}

func AllocConnectionSummary(_ context.Context) any {
	return &ConnectionSummary{}
}

func AllocChannelSummary(_ context.Context) any {
	return &BaseChannelSummary{}
}

// GetId returns the Id of the summary file.
func (b *BaseChannelSummary) GetId() string {
	return b.Id
}

// GetId returns the Id of the summary file.
func (b *BaseChannelSummary) GetConnectionRecordingId() string {
	return b.ConnectionRecordingId
}

// GetStartTime returns the start time using a monotonic clock of the summary.
func (b *BaseChannelSummary) GetStartTime() time.Time {
	return b.StartTime
}

// GetEndTime returns the end time using a monotonic clock of the summary.
func (b *BaseChannelSummary) GetEndTime() time.Time {
	return b.EndTime
}

// GetBytesUp returns upload bytes.
func (b *BaseChannelSummary) GetBytesUp() uint64 {
	return b.BytesUp
}

// GetBytesDown returns download bytes.
func (b *BaseChannelSummary) GetBytesDown() uint64 {
	return b.BytesDown
}

// GetChannelType the type of summary channel.
func (b *BaseChannelSummary) GetChannelType() string {
	return b.ChannelType
}
