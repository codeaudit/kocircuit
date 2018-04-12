package eval

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/compile"
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
)

var EvalIdiomRepo = compileIdiomRepo()

func compileIdiomRepo() Repo {
	repo := MustCompileString(IdiomRootPkg, "idiom.ko", EvalIdiomSrc) // package "idiom"
	return repo
}

const EvalIdiomSrc = `
// Nothing returns an empty value.
Nothing() {
	return: ()
}

// Return returns its default argument.
Return(pass?) {
	return: pass
}

// Expect panics if its default argument is empty.
Expect(value?) {
	return: When(
		have: value
		then: Return
		else: Panic[expect: "expecting a value"]
	)
}

Pick(either, or) {
	return: When(
		have: either
		then: Return
		else: Return[or]
	)
}

Invoke(variety?) {
	return: variety()
}

Branch(if, then, else) {
	return: When(
		have: Yield(if: if, then: then, else: else)
		then: Invoke
		else: Nothing
	)
}
`
