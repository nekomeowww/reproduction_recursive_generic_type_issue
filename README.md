# recursive generic type reproduction

This repository showed how to reproduce the compiler unexpected behavior, or issue, bug when embedding a struct that has a field to reference the outer struct by using generic recursively: the compiler will complain `fatal error: all goroutines are asleep - deadlock!` or behave unexpectedly in all different ways.

Well I could understand the limitations of the current implementation for generics and the "should avoid" with recursive type embedding, but I still think this is a issue that the compiler or documentation should address when user attempts to do so. Especially it runs successfully and within the expectation when you change the package of `main_test.go` into `recursive_generic_type_issue_reproduction` without the `_test` prefix. You, as being a Golang developer, might only find out what has been wrong after you released the library without the `_test` test packages test included and caused the compiler issue and complains for all your users.

Even though such unexpected behavior could be resolved by `GOEXPERIMENT=nounified` flag[^1] and reported rarely, it's still a bit confusing for users to understand what's going on. At least it jammed me for a while.

## Behaviors

Compiler might complain the following panics message when running the test by `go test -v ./...`:

<details>
<summary>fatal error: all goroutines are asleep - deadlock!</summary>

```txt
# github.com/nekomeowww/recursive_generic_type_issue_reproduction_test [github.com/nekomeowww/recursive_generic_type_issue_reproduction.test]
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [sync.Mutex.Lock]:
sync.runtime_SemacquireMutex(0x140004e4200?, 0x30?, 0x140004e42e8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/runtime/sema.go:77 +0x28
sync.(*Mutex).lockSlow(0x140004e1448)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/sync/mutex.go:171 +0x178
sync.(*Mutex).Lock(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/sync/mutex.go:90
cmd/compile/internal/types2.(*Named).resolve(0x140004e1420)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:164 +0xa4
cmd/compile/internal/types2.(*Named).TypeParams(0x140004e43b8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:310 +0x1c
cmd/compile/internal/types2.(*subster).typ(0x140004e4a88, {0x105220da0?, 0x140004e1500?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x140004e47b8?, {0x105220dc8?, 0x1400006be60?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x140004ea148?, 0x140004e1570)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x140004e48f8?, {0x1400006be10, 0x2, 0x3?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x140004e4a98?, {0x105220e40?, 0x140004d9b30?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x0, {0x0?, 0x4d?, 0x0?}, {0x105220e40?, 0x140004d9b30}, 0x140004f00c0, 0x140004e18f0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x140004e18f0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x140004e18f0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x1?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x140004e18f0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x105220da0?, 0x140004e18f0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.computeUnionTypeSet(0x0, 0x1400000dfe0?, {0x0, 0x0, 0x0}, 0x1400000dfe0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeset.go:407 +0x12c
cmd/compile/internal/types2.computeInterfaceTypeSet(0x0, {0x140004d97d0?, 0x0?, 0x0?}, 0x140004ce3c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeset.go:295 +0x618
cmd/compile/internal/types2.(*TypeParam).iface(0x140004d9e30)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:138 +0x1e8
cmd/compile/internal/types2.(*TypeParam).SetConstraint(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:86
cmd/compile/internal/importer.(*reader).typeParamNames(0x14000489680)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:510 +0x254
cmd/compile/internal/importer.(*pkgReader).objIdx.func1.1(0x10?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:430 +0x24
cmd/compile/internal/types2.(*Named).resolve(0x140004e1420)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:203 +0x148
cmd/compile/internal/types2.(*Named).TypeParams(0x140004e54e8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:310 +0x1c
cmd/compile/internal/types2.(*subster).typ(0x140004e5be8, {0x105220da0?, 0x140004e1500?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x140004e5918?, {0x105220dc8?, 0x1400006be60?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x1400013df48?, 0x140004e1570)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x140004e5a58?, {0x1400006be10, 0x2, 0x3?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x140004e5bf8?, {0x105220e40?, 0x140004d9b30?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x140001483c0, {0x1400039f860?, 0x4c?, 0x0?}, {0x105220e40?, 0x140004d9b30}, 0x140004d9dd0, 0x140004e1810, 0x1400040c260)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x140004e1810)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x140004e1810)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x140004e1810)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x105220da0?, 0x140004e1810?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.lookupFieldOrMethod({0x105220dc8?, 0x1400006bed0?}, 0x0, 0x1400040b030?, {0x14000021900, 0x9}, 0x0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/lookup.go:100 +0xf0
cmd/compile/internal/types2.LookupFieldOrMethod({0x105220dc8?, 0x1400006bed0?}, 0x20?, 0x1400040b030?, {0x14000021900?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/lookup.go:66 +0xb0
cmd/compile/internal/types2.(*Checker).selector(0x140001483c0, 0x140004de640, 0x140004320a0, 0x0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:581 +0x1ec
cmd/compile/internal/types2.(*Checker).exprInternal(0x140001483c0, 0x140004de640, {0x105222ba0?, 0x140004320a0}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1590 +0xb7c
cmd/compile/internal/types2.(*Checker).rawExpr(0x140001483c0, 0x140004de640, {0x105222ba0?, 0x140004320a0?}, {0x0?, 0x0?}, 0x1)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).exprOrType(0x1400012ac98?, 0x1048039c4?, {0x105222ba0?, 0x140004320a0?}, 0x98?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1836 +0x38
cmd/compile/internal/types2.(*Checker).callExpr(0x140001483c0, 0x140004de640, 0x1400040b0a0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:110 +0xe4
cmd/compile/internal/types2.(*Checker).exprInternal(0x140001483c0, 0x140004de640, {0x105222820?, 0x1400040b0a0}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1639 +0xee8
cmd/compile/internal/types2.(*Checker).rawExpr(0x140001483c0, 0x140004de640, {0x105222820?, 0x1400040b0a0?}, {0x0?, 0x0?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).exprOrType(0x140004e0fc0?, 0x1?, {0x105222820?, 0x1400040b0a0?}, 0xa8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1836 +0x38
cmd/compile/internal/types2.(*Checker).selector(0x140001483c0, 0x140004de640, 0x14000432190, 0x0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:547 +0xc4
cmd/compile/internal/types2.(*Checker).exprInternal(0x140001483c0, 0x140004de640, {0x105222ba0?, 0x14000432190}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1590 +0xb7c
cmd/compile/internal/types2.(*Checker).rawExpr(0x140001483c0, 0x140004de640, {0x105222ba0?, 0x14000432190?}, {0x0?, 0x0?}, 0x1)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).exprOrType(0x13?, 0x140004e7fe8?, {0x105222ba0?, 0x14000432190?}, 0x10?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1836 +0x38
cmd/compile/internal/types2.(*Checker).callExpr(0x140001483c0, 0x140004de640, 0x1400040b110)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:110 +0xe4
cmd/compile/internal/types2.(*Checker).exprInternal(0x140001483c0, 0x140004de640, {0x105222820?, 0x1400040b110}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1639 +0xee8
cmd/compile/internal/types2.(*Checker).rawExpr(0x140001483c0, 0x140004de640, {0x105222820?, 0x1400040b110?}, {0x0?, 0x0?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).multiExpr(0x140004e0f50?, 0x0?, {0x105222820?, 0x1400040b110?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1817 +0x38
cmd/compile/internal/types2.(*Checker).exprList(0x140004e0540?, {0x140004e8ac8?, 0x0?, 0x3?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:273 +0x68
cmd/compile/internal/types2.(*Checker).initVars(0x10519f560?, {0x1400000e630, 0x1, 0x0?}, {0x140004e8ac8, 0x1, 0x1}, {0x0, 0x0})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/assignments.go:317 +0x74
cmd/compile/internal/types2.(*Checker).shortVarDecl(0x140001483c0, {0x1400039f860, 0x13, 0x11}, {0x140004e8ad8, 0x1, 0x0?}, {0x140004e8ac8?, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/assignments.go:505 +0x87c
cmd/compile/internal/types2.(*Checker).stmt(0x140001483c0, 0x0, {0x105221748?, 0x1400042c280?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:475 +0x5fc
cmd/compile/internal/types2.(*Checker).stmtList(0x105220620?, 0x0, {0x14000143300?, 0x104eed3c0?, 0x140001483c0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:123 +0x80
cmd/compile/internal/types2.(*Checker).funcBody(0x140001483c0, 0x14000488480, {0x14000022558?, 0x140001483c0?}, 0x140004de200, 0x1400042c240, {0x0, 0x0})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:43 +0x2bc
cmd/compile/internal/types2.(*Checker).funcDecl.func1()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/decl.go:760 +0x44
cmd/compile/internal/types2.(*Checker).processDelayed(0x140001483c0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:383 +0x14c
cmd/compile/internal/types2.(*Checker).checkFiles(0x140001483c0, {0x1400000e508, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:328 +0x144
cmd/compile/internal/types2.(*Checker).Files(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:300
cmd/compile/internal/types2.(*Config).Check(0x1400000da70?, {0x16b64a0f9?, 0x104ff1ddf?}, {0x1400000e508, 0x1, 0x1}, 0x1056e7fc0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/api.go:434 +0x68
cmd/compile/internal/noder.checkFiles({0x1400000e4e0, 0x1, 0x104fe4509?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/irgen.go:73 +0x3b4
cmd/compile/internal/noder.writePkgStub({0x1400000e4e0, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:210 +0x30
cmd/compile/internal/noder.unified({0x1400000e4e0?, 0x1400039f8f0?, 0x2?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:75 +0x94
cmd/compile/internal/noder.LoadPackage({0x140000003b0, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/noder.go:77 +0x400
cmd/compile/internal/gc.Main(0x10521bad0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/gc/main.go:196 +0xcd8
main.main()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/main.go:57 +0xf4
FAIL    github.com/nekomeowww/recursive_generic_type_issue_reproduction [build failed]
FAIL
```

</details>

Compiler might complain about the `invalid recursive type` if you try to move the `TypeB[T any]` definition to another file.

Compiler might complain about the `redecalred` if you try to reload the IDE.

[^1]: [cmd/compile: failed to compile some recursive generic type · Issue #54535 · golang/go](https://github.com/golang/go/issues/54535)
