// Copyright © 2015 Erik Brady <brady@dvln.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package test: pkg
//      Simple testing for the 'pkg' package ...

package pkg

import (
	"fmt"
	"testing"
)

func TestPkg(t *testing.T) {
	pkg := &Defn{}
	fmt.Printf("pkg: %+v", pkg)
}
