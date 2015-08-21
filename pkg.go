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

// Defn identifies the data needed to "define" a 'dvln' "package"... *not*
// including any specific version information.  This is all about basic
// package details... where the packages "live", where they will go in the
// sandbox or working tree (ie: "workspace"), what they might have been called
// in the past or aliases that they might still need to be referenced by, what
// VCS/SCM system is storing the package and data about where to find that, ...
type Defn struct {
	ID       int                          `json:"id"`
	Name     string                       `json:"name"`
	Aliases  map[string]string            `json:"aliases"`
	Desc     string                       `json:"desc"`
	Codebase string                       `json:"codebase"`
	Contacts []string                     `json:"contacts"`
	VCS      string                       `json:"vcs"`
	Repo     map[string]string            `json:"repo"`
	WsPath   string                       `json:"ws_path"`
	Arch     []string                     `json:"arch"`
	OS       []string                     `json:"os"`
	DevStage []string                     `json:"dev_stage"`
	Attrs    map[string]string            `json:"attrs"`
	Remotes  map[string]map[string]string `json:"remotes"`
	Access   map[string]string            `json:"access"`
	Status   string                       `json:"status"`
}

