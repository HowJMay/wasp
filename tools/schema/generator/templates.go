// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import "github.com/iotaledger/wasp/tools/schema/model"

var commonTemplates = model.StringMap{
	// *******************************
	"else": `
`,
	// *******************************
	"newline": `

`,
	// *******************************
	"copyright": `
$#emit initGlobals
$#if copyrightMessage userCopyrightMessage defaultCopyrightMessage
`,
	// *******************************
	"userCopyrightMessage": `
$copyrightMessage
`,
	// *******************************
	"defaultCopyrightMessage": `
// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

`,
	// *******************************
	"warning": `
// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the schema definition file instead

`,
	// *******************************
	"../README.md Lib": `
## $package

Interface library for: $scDesc
`,
	// *******************************
	"../README.md Impl": `
## $package$+impl

Implementation library for: $scDesc
`,
	// *******************************
	"../README.md Wasm": `
## $package$+wasm

Wasm VM host stub for: $scDesc
`,
	// *******************************
	"test.go": `
$#emit copyright
package test

import (
	"testing"

	"$module/go/$package"
	"github.com/iotaledger/wasp/packages/wasmvm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, $package.ScName, $package.OnDispatch)
	require.NoError(t, ctx.ContractExists($package.ScName))
}
`,
	// *******************************
	"setupInitFunc": `
$#set initFunc $nil
$#if init setInitFunc
`,
	// *******************************
	"setInitFunc": `
$#set initFunc Init
`,
	// *******************************
	"alignCalculate": `
$#set align $nil
$#if events align1space
$#if param align1space
$#if result align2spaces
$#set salign $align
$#set align $nil
$#if result align1space
$#set falign $nil
$#if salign alignSetFunc
`,
	// *******************************
	"align1space": `
$#set align $space
`,
	// *******************************
	"align2spaces": `
$#set align $space$space
`,
	// *******************************
	"alignSetFunc": `
$#set falign $salign$space
`,
}
