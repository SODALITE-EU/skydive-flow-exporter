/*
 * Copyright (C) 2019 IBM, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy ofthe License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specificlanguage governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"github.com/skydive-project/skydive-flow-exporter/core"
	vpc "github.com/skydive-project/skydive-flow-exporter/vpclogs/pkg"
)

func main() {
	core.Main("/etc/skydive/vpclogs.yml")
}

func init() {
	core.TransformerHandlers.Register("vpclogs", vpc.NewTransform, false)
	core.StoreHeaderHandlers.Register("vpclogs", vpc.NewStoreHeaderVpc, false)
	core.ManglerHandlers.Register("action", vpc.NewAction, false)
}
