package weave

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/compile"
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
)

var GoIdiomRepo = compileIdiomRepo()

func compileIdiomRepo() Repo {
	repo := MustCompileString(IdiomRootPkg, "idiom.ko", IdiomSrc) // package "idiom"
	return repo
}

const IdiomSrc = `
Return(pass?) { return: pass }

Branch(if, then, else) {
	return: Yield(if: if, then: then, else: else)()
}

import "github.com/kocircuit/kocircuit/lib/strings"
import "tree"

Equal(m?) {
	hashes: Range(
		over: m
		with: equalProjectionHash(elem) {
			return: (
				emit: tree.Hash(tree.Projection(elem))
			)
		}
	)
	return: strings.GoEqualStrings(hashes.image)
}
`
