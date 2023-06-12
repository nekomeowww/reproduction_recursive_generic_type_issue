# Minimum reproduction of deadlock error when using generics with pointer reference to the generic type

Unlike the [`minimum_repro/with_generics`](../with_generics) example I used `type innerT1[T any, R T1[T]] struct` where the `R T1[T]` is not the a pointer, I
wrapped with pointer to ask compiler to generate a dynamic memory allocation for the type `innerT1[T any, R *T1[T]] struct` this time instead.

The compiler would immediately complain about the
`fatal error: all goroutines are asleep - deadlock!` issue when executing `go test -v ./minimum_repro/with_generics`, and the unit tests never bot ran. Just the same as [`minimum_repro/with_generics`](../with_generics).

## How to reproduce the issue?

Run the following command in the project root directory:

```shell
$ go test -v ./minimum_repro/deadlock_issue/with_generics_with_pointer
```

<details>
<summary>Output</summary>

```shell
$ go test -v ./minimum_repro/deadlock_issue/with_generics_with_pointer
# github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics_with_pointer_test [github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics_with_pointer.test]
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [sync.Mutex.Lock]:
sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/runtime/sema.go:77 +0x28
sync.(*Mutex).lockSlow(0x140004566b8)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/sync/mutex.go:171 +0x178
sync.(*Mutex).Lock(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/sync/mutex.go:90
cmd/compile/internal/types2.(*Named).resolve(0x14000456690)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:164 +0xa4
cmd/compile/internal/types2.(*Named).TypeParams(0x0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:310 +0x1c
cmd/compile/internal/types2.(*subster).typ(0x1400045cc48, {0x104e44da0?, 0x14000456770?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x1400045c978?, {0x104e44dc8?, 0x1400009f7b0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x140000ebd68?, 0x140004567e0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x1400045cab8?, {0x140000b65c8, 0x1, 0x3?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x1400045cc58?, {0x104e44e40?, 0x140004524e0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x0, {0x0?, 0x75?, 0x0?}, {0x104e44e40?, 0x140004524e0}, 0x14000452900, 0x140004568c0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x1400045cec8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x104e44da0?, 0x140004568c0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.computeInterfaceTypeSet(0x0, {0x14000452330?, 0x0?, 0x0?}, 0x140003ff810)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeset.go:275 +0x40c
cmd/compile/internal/types2.(*TypeParam).iface(0x14000452720)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:138 +0x1e8
cmd/compile/internal/types2.(*TypeParam).SetConstraint(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:86
cmd/compile/internal/importer.(*reader).typeParamNames(0x140004511a0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:510 +0x254
cmd/compile/internal/importer.(*pkgReader).objIdx.func1.1(0x140003c83c0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:430 +0x24
cmd/compile/internal/types2.(*Named).resolve(0x14000456690)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:203 +0x148
cmd/compile/internal/types2.(*Named).TypeParams(0x1400045d658?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:310 +0x1c
cmd/compile/internal/types2.(*subster).typ(0x1400045dcd8, {0x104e44da0?, 0x14000456770?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x1400045da08?, {0x104e44dc8?, 0x1400009f7b0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x140000ebc28?, 0x140004567e0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x1400045db48?, {0x140000b65c8, 0x1, 0xa?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x1400045dce8?, {0x104e44e40?, 0x140004524e0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x140000f83c0, {0x1400035b860?, 0x7c?, 0x0?}, {0x104e44e40?, 0x140004524e0}, 0x140004526c0, 0x14000456850, 0x140003c81a0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x3?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x104e44da0?, 0x14000456850?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.coreType({0x104e44da0?, 0x14000456850?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:33 +0xa0
cmd/compile/internal/types2.(*Checker).exprInternal(0x140000f83c0, 0x1400044cb80, {0x104e468a0?, 0x140000f1280}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1403 +0x720
cmd/compile/internal/types2.(*Checker).rawExpr(0x140000f83c0, 0x1400044cb80, {0x104e468a0?, 0x140000f1280?}, {0x0?, 0x0?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).multiExpr(0x104e44e68?, 0x0?, {0x104e468a0?, 0x140000f1280?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1817 +0x38
cmd/compile/internal/types2.(*Checker).exprList(0x140000f83c0?, {0x1400045eac8?, 0x104e46820?, 0x140003c6e70?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:273 +0x68
cmd/compile/internal/types2.(*Checker).assignVars(0x140000f83c0, {0x1400045ead8?, 0x1, 0x1}, {0x1400045eac8?, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/assignments.go:384 +0x5c
cmd/compile/internal/types2.(*Checker).stmt(0x140000f83c0, 0x0, {0x104e45748?, 0x140003e8280?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:472 +0x62c
cmd/compile/internal/types2.(*Checker).stmtList(0x104e44620?, 0x0, {0x1400009eea0?, 0x104b113c0?, 0x140000f83c0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:123 +0x80
cmd/compile/internal/types2.(*Checker).funcBody(0x140000f83c0, 0x140004500c0, {0x140000bd760?, 0x140000f83c0?}, 0x1400044c7c0, 0x140003e8240, {0x0, 0x0})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:43 +0x2bc
cmd/compile/internal/types2.(*Checker).funcDecl.func1()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/decl.go:760 +0x44
cmd/compile/internal/types2.(*Checker).processDelayed(0x140000f83c0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:383 +0x14c
cmd/compile/internal/types2.(*Checker).checkFiles(0x140000f83c0, {0x140000b64c8, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:328 +0x144
cmd/compile/internal/types2.(*Checker).Files(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:300
cmd/compile/internal/types2.(*Config).Check(0x140000a99e0?, {0x16ba26001?, 0x104c15ddf?}, {0x140000b64c8, 0x1, 0x1}, 0x10530bfc0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/api.go:434 +0x68
cmd/compile/internal/noder.checkFiles({0x140000b64a0, 0x1, 0x104c08509?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/irgen.go:73 +0x3b4
cmd/compile/internal/noder.writePkgStub({0x140000b64a0, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:210 +0x30
cmd/compile/internal/noder.unified({0x140000b64a0?, 0x1400035b8f0?, 0x2?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:75 +0x94
cmd/compile/internal/noder.LoadPackage({0x140000c6130, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/noder.go:77 +0x400
cmd/compile/internal/gc.Main(0x104e3fad0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/gc/main.go:196 +0xcd8
main.main()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/main.go:57 +0xf4
FAIL    github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics_with_pointer [build failed]
FAIL
```

</details>
