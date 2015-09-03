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

// Package pkg contains structures and methods related to pkg definitions
// from a 'dvln' codebase.  These packages might be of any of these types:
// - a "leaf" pkg: identifying a simple "single" repo (git, hg, etc)
// - a "dvln" pkg: same as leaf but also contains a "sub" dvln (pkg manifest)
// - a "codebase" pkg: this pkg contains a codebase defn inside it (see codebase pkg)
//
// A package may be arbitrarily defined (eg: within a codebase definition file)
// or may be instantiated at a given version inside a users workspace (or by
// querying some version of a pkg on the server).  This is generic info this
// run of the tool has about the pkg (in it's current "use", eg: if querying
// a diff from the server with no workspace we may indicate what the workspace
// path would be if it were instantiated but that path won't exist whereas if
// we have a workspace we could find the package there in the workspace.
package pkg

import (
	"github.com/dvln/vcs"
)

type Class string

// dvln 'pkg' classes, ie "single pkg", "group pkg", "codebase pkg"
const (
	UnknownClass Class = ""
	Single       Class = "single"
	Group        Class = "group"
	Codebase     Class = "codebase"
)

// Defn identifies the data needed to "define" a dvln "package"... *not*
// including any specific version information.  This is all about basic
// package details... VCS info for the package, where it will exist in the
// workspace (ie: sandbox or work tree), what a pkg might have been called
// in the past or aliases that it might still need to be referenced by, etc
type Defn struct {
	ID       int                          `json:"id"`
	Name     string                       `json:"name"`
	Aliases  map[string]string            `json:"aliases,omitempty"`
	Desc     string                       `json:"desc,omitempty"`
	Ws       string                       `json:"ws,omitempty"`
	Contacts []string                     `json:"contacts,omitempty"`
	Class    Class                        `json:"class,omitempty"`
	VCS      vcs.Type                     `json:"vcs,omitempty"`
	Repo     map[string]string            `json:"repo,omitempty"`
	Arch     []string                     `json:"arch,omitempty"`
	OS       []string                     `json:"os,omitempty"`
	Stage    []string                     `json:"stage,omitempty"`
	Attrs    map[string]string            `json:"attrs,omitempty"`
	Remotes  map[string]map[string]string `json:"remotes,omitempty"`
	Access   map[string]string            `json:"access,omitempty"`
	Status   string                       `json:"status,omitempty"`
}
