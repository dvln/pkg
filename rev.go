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

// The rev.go sub-part of the 'pkg' package is related to working with
// package revs.  Keep in mind packages might be from a VCS or from a
// package management system or from a tar/patch mechanism potentially (or
// from all 3 different formats if a given package supports that).  Regardless,
// dvln is focused on pacakge management.  All package types share common
// traits as covered in the package definition in pkg.go, and every format
// of package used in 'dvln' needs to be referencable by a rev.
//
// The rev may be dynamic (branch/latest, or v.1.2.x if semver active)
// or may be static (v1.0.0, <sha1>, <datetime>, etc) and various package
// representations (VCS's, pkg system, etc) should be able to manipulate
// a package, switch between formats (source, binary, buildmeta, targz,
// rpm, deb, x, y, z) for any valid revision available in the given VCS
// system or other format.  The main point is these revs may be of
// varying formats for all these systems.  For VCS's like hg, git, svn
// are all valid revs potentially:
//   232eb232, labelx, 1.2.x, v0.1.0-alpha, branch, 320, branch/44, TAG1
// Some are sha1's (core VCS revisions) and some are tags which can
// take many formats, some are branches.
//
// Given a revision (VCS version) is static or dynamic always clear?
// Not without some indication, eg: "projx".  In git this might be
// the name of a changing branch tip reference or the name of a
// static unchanging tag (or even a tag that a team moves now and
// then, not really recommended but there ya go).  The 'dvln' tool
// will try and do it's best to identify static or not, eg for git:
//   "projx/latest" would be how one rides branch latest
//   "projx" would be seen as static by default (auto-introspection
//      could perhaps be added over time to determine it but that
//      would slow down the tool if many revs required that)
//   "<sha1>" would be seen as static
//   v10.1.0-alpha would be seen as static, if semver was on it
//      would also be seen as a valid semantic version (ver 2)
//   1.2.x could be seen as a static label unless semver was active
//      then if valid semver's regex matched it could be seen instead
//      as dynamic, matching the latest semver (eg: 1.2.9 now)
//
// Package revs will be (some of these over time):
// - for instantiating a pkg ver into a wkspc via 'dvln get'
// - within a development line (dvln) to identify package revs
// - have new revs pulled/merged into a wkspc via 'dvln pull'
// - be described via 'dvln describe'
//   - aside: could be run in or out of a workspace potentially
// - be used as src and target via 'dvln diff' and 'dvln log'
//   - aside: could be run in or out of a workspace potentially
// - group packages are package that contain a dvln version (w/pkg rev's)
//   - groups use the parent codebase namespace w/all pkg's used defined in
//     parent codebase, differeing from nested codebases as they have their
//     own namespace/governance (and only the codebase reference must exist
//     in the parent codebase definition file, not the comps in the nested
//     codebase since they are defined within that nested codebase)
// - the codebase definition file will be stored in a pkg and ref'd via pkg vers
// - etc

package pkg

import (
	"github.com/dvln/vcs"
)

// RevType is an indicator if a revision is static or dynamic (or unknown)
type RevType string

// dvln 'pkg' revision status, is it dynamic or static (or unknown)?
const (
	UnknownRevType RevType = ""
	Dynamic        RevType = "dynamic"
	Static         RevType = "static"
)

// Revision is a package revision (ie: VCS version), anything one might need
// to do with a revision will be part of the interface (and hence can be
// mock'd to ease testing and such)
type Revision interface {
	Exists(pkgRevSel string) (string, error)
	Get(pkgRevSel string) error
	Pull(pkgRevSel string) error
}

// Rev indicates the basic data needed for a given VCS pkg revision.  The
// name is useful for messaging and such whereas the ID is used for trivial
// lookup of the package meta-data in the codebase/pkg definition.
type Rev struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Rev     string         `json:"rev"`
	RevType string         `json:"rev_type"`
	Deps    map[int]string `json:"deps"`
	VCS     vcs.Reader
}

