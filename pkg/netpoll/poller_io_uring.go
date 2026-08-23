// Copyright (c) 2026 The Gnet Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build linux && io_uring_opt

package netpoll

import (
	"errors"

	"github.com/panjf2000/gnet/v2/pkg/queue"
)

// errNotImplemented is returned by every io_uring operation until the
// ring setup/submission/completion machinery is actually implemented.
var errNotImplemented = errors.New("netpoll: io_uring poller is not implemented yet")

// Poller represents a poller backed by io_uring which is in charge of monitoring file-descriptors.
type Poller struct {
	asyncTaskQueue       queue.AsyncTaskQueue // queue with low priority
	urgentAsyncTaskQueue queue.AsyncTaskQueue // queue with high priority
}

// OpenPoller instantiates a poller.
func OpenPoller() (poller *Poller, err error) {
	return nil, errNotImplemented
}

// Close closes the poller.
func (p *Poller) Close() error {
	return errNotImplemented
}

// Trigger enqueues task and wakes up the poller to process pending tasks.
func (p *Poller) Trigger(priority queue.EventPriority, fn queue.Func, param any) (err error) {
	return errNotImplemented
}

// Polling blocks the current goroutine, monitoring the registered file descriptors and waiting for network I/O.
// When I/O occurs on any of the file descriptors, the provided callback function is invoked.
func (p *Poller) Polling(callback PollEventHandler) error {
	return errNotImplemented
}

// AddReadWrite registers the given file descriptor with readable and writable events to the poller.
func (p *Poller) AddReadWrite(pa *PollAttachment, edgeTriggered bool) error {
	return errNotImplemented
}

// AddRead registers the given file descriptor with readable event to the poller.
func (p *Poller) AddRead(pa *PollAttachment, edgeTriggered bool) error {
	return errNotImplemented
}

// AddWrite registers the given file descriptor with writable event to the poller.
func (p *Poller) AddWrite(pa *PollAttachment, edgeTriggered bool) error {
	return errNotImplemented
}

// ModRead modifies the given file descriptor with readable event in the poller.
func (p *Poller) ModRead(pa *PollAttachment, edgeTriggered bool) error {
	return errNotImplemented
}

// ModReadWrite modifies the given file descriptor with readable and writable events in the poller.
func (p *Poller) ModReadWrite(pa *PollAttachment, edgeTriggered bool) error {
	return errNotImplemented
}

// Delete removes the given file descriptor from the poller.
func (p *Poller) Delete(fd int) error {
	return errNotImplemented
}
