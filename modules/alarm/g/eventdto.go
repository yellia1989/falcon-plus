// Copyright 2017 Xiaomi, Inc.
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

package g

import (
	"fmt"
	cmodel "github.com/yellia1989/falcon-plus/common/model"
)

func Link(event *cmodel.Event) string {
	tplId := event.TplId()
	if tplId != 0 {
		return fmt.Sprintf("%s/portal/template/view/%d", Config().Api.Dashboard, tplId)
	}

	eid := event.ExpressionId()
	if eid != 0 {
		return fmt.Sprintf("%s/portal/expression/view/%d", Config().Api.Dashboard, eid)
	}

	return ""
}
