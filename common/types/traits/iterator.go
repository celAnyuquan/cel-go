// Copyright 2018 Google LLC
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

package traits

import (
	"github.com/celAnyuquan/cel-go/common/types/ref"
)

// Iterable aggregate types permit traversal over their elements.
type Iterable interface {
	// Iterator returns a new iterator view of the struct.
	Iterator() Iterator
}

// Iterator permits safe traversal over the contents of an aggregate type.
type Iterator interface {
	ref.Val

	// HasNext returns true if there are unvisited elements in the Iterator.
	HasNext() ref.Val

	// Next returns the next element.
	Next() ref.Val
}
