// Copyright (c) 2022 Arthur Skowronek <0x5a17ed@tuta.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// <https://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package itertools

import (
	"github.com/0x5a17ed/iterkit"
)

type FilterFunc[T any] func(element T) bool

type FilterIter[T any] struct {
	it   iterkit.Iterator[T]
	fn   FilterFunc[T]
	next *T
}

func (f *FilterIter[T]) Next() bool {
	for f.it.Next() {
		if next := f.it.Value(); f.fn(next) {
			f.next = &next
			return true
		}
	}
	f.next = nil
	return false
}

func (f FilterIter[T]) Value() T { return *f.next }

// Filter returns an Iterator yielding items from the given iterator
// for which the given FilterFunc function returns true.
func Filter[T any](it iterkit.Iterator[T], cb FilterFunc[T]) iterkit.Iterator[T] {
	return &FilterIter[T]{it: it, fn: cb}
}
