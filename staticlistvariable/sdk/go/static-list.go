// Copyright 2024 The Perses Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package staticlist

import (
	listvariable "github.com/perses/perses/go-sdk/variable/list-variable"
)

type PluginSpec struct {
	Values []string `json:"values" yaml:"values"`
}

type Option func(plugin *Builder) error

type Builder struct {
	PluginSpec
}

func create(options ...Option) (Builder, error) {
	var builder = &Builder{
		PluginSpec: PluginSpec{},
	}

	var defaults []Option

	for _, opt := range append(defaults, options...) {
		if err := opt(builder); err != nil {
			return *builder, err
		}
	}

	return *builder, nil
}

func StaticList(options ...Option) listvariable.Option {
	return func(builder *listvariable.Builder) error {
		t, err := create(options...)
		if err != nil {
			return err
		}
		builder.ListVariableSpec.Plugin.Kind = "StaticListVariable"
		builder.ListVariableSpec.Plugin.Spec = t
		return nil
	}
}
