//+build djavul

package main

import "C"

import (
	"unsafe"

	"github.com/sanctuary/djavul/pfile"
)

// --- [ pfile ] ---------------------------------------------------------------

//export UICreateSave
func UICreateSave(heroInfo unsafe.Pointer) bool {
	return pfile.UICreateSave((*pfile.HeroInfo)(heroInfo))
}
