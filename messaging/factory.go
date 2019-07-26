//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package messaging

import (
	"fmt"
	"strings"

	"github.com/antoniomtz/go-mod-messaging/pkg/types"

	"github.com/antoniomtz/go-mod-messaging/internal/pkg/zeromq"
)

const (
	// ZeroMQ messaging implementation
	ZeroMQ = "zero"
)

// NewMessageClient is a factory function to instantiate different message client depending on
// the "Type" from the configuration
func NewMessageClient(msgConfig types.MessageBusConfig) (MessageClient, error) {

	if msgConfig.PublishHost.IsHostInfoEmpty() && msgConfig.SubscribeHost.IsHostInfoEmpty() {
		return nil, fmt.Errorf("unable to create messageClient: host info not set")
	}

	switch lowerMsgType := strings.ToLower(msgConfig.Type); lowerMsgType {
	case ZeroMQ:
		return zeromq.NewZeroMqClient(msgConfig)
	default:
		return nil, fmt.Errorf("unknown message type '%s' requested", msgConfig.Type)
	}
}
