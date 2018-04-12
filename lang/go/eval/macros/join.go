package macros

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
	. "github.com/kocircuit/kocircuit/lang/go/eval"
	. "github.com/kocircuit/kocircuit/lang/go/eval/symbol"
	. "github.com/kocircuit/kocircuit/lang/go/kit/util"
)

func init() {
	RegisterEvalMacro("", new(EvalJoinMacro))
}

type EvalJoinMacro struct{}

func (m EvalJoinMacro) MacroID() string { return m.Help() }

func (m EvalJoinMacro) Label() string { return "join" }

func (m EvalJoinMacro) MacroSheathString() *string { return PtrString("Join") }

func (m EvalJoinMacro) Help() string { return "Join" }

func (EvalJoinMacro) Invoke(span *Span, arg Arg) (returns Return, effect Effect, err error) {
	return Construct(arg.(*StructSymbol)), nil, nil
}

func Construct(a *StructSymbol) Symbol {
	a = FilterEmptyStructFields(a)
	if a.IsEmpty() {
		return EmptySymbol{}
	} else if monadic := a.SelectMonadic(); !IsEmptySymbol(monadic) {
		return monadic
	} else {
		return a
	}
}
