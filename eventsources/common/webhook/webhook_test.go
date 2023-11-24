/*
Copyright 2018 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestValidateWebhook(t *testing.T) {
	convey.Convey("Given a webhook, validate it", t, func() {
		convey.So(ValidateWebhookContext(Hook), convey.ShouldBeNil)
	})
}

func TestNewWebhookHelper(t *testing.T) {
	convey.Convey("Make sure webhook helper is not empty", t, func() {
		controller := NewController()
		convey.So(controller, convey.ShouldNotBeNil)
	})
}
