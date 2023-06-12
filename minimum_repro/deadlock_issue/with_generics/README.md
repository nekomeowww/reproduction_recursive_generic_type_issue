# Minimum reproduction of deadlock error when using generics with test package (`_test` suffix)

Unlike the [`minimum_repro/with_generics_without_test_package`](../with_generics_without_test_package) example, I **didn't** used the package with a suffix of `_test` in the tests file (`_test.go`), I
appended a `_test` suffix to the package name instead.

The compiler would immediately complain about the
`fatal error: all goroutines are asleep - deadlock!` issue when executing `go test -v ./minimum_repro/with_generics`, and the unit tests never bot ran.

The most strange thing is that after you removed the `_test` suffix from the [`test_without_test_package_test.go`](./main_without_test_package_test.go) file that belongs to a test package, the compile error will disappear, and the test will pass with no error, you can check the minimum reproduction I created [`minimum_repro/with_generics_without_test_package`](../with_generics_without_test_package).

## How to reproduce the issue?

Run the following command in the project root directory:

```shell
go test -v ./minimum_repro/deadlock_issue/with_generics
```

<details>
<summary>Output</summary>

```shell
$ go test -v ./minimum_repro/deadlock_issue/with_generics
# github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics_test [github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics.test]
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
cmd/compile/internal/types2.(*subster).typ(0x1400045cc48, {0x10113cda0?, 0x14000456770?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x1400045c978?, {0x10113cdc8?, 0x1400009f790?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x140000d5cc8?, 0x140004567e0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x1400045cab8?, {0x140000b65d8, 0x1, 0x3?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x1400045cc58?, {0x10113ce40?, 0x140004524e0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x0, {0x0?, 0x75?, 0x0?}, {0x10113ce40?, 0x140004524e0}, 0x14000452900, 0x140004568c0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x1400045cec8?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x140004568c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x10113cda0?, 0x140004568c0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.computeInterfaceTypeSet(0x0, {0x14000452330?, 0x0?, 0x0?}, 0x140003ff810)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeset.go:275 +0x40c
cmd/compile/internal/types2.(*TypeParam).iface(0x14000452720)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:138 +0x1e8
cmd/compile/internal/types2.(*TypeParam).SetConstraint(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/typeparam.go:86
cmd/compile/internal/importer.(*reader).typeParamNames(0x140004511a0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:510 +0x254
cmd/compile/internal/importer.(*pkgReader).objIdx.func1.1(0x140003c83e0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/importer/ureader.go:430 +0x24
cmd/compile/internal/types2.(*Named).resolve(0x14000456690)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:203 +0x148
cmd/compile/internal/types2.(*Named).TypeParams(0x1400045d658?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:310 +0x1c
cmd/compile/internal/types2.(*subster).typ(0x1400045dcd8, {0x10113cda0?, 0x14000456770?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:223 +0xb08
cmd/compile/internal/types2.(*subster).typ(0x1400045da08?, {0x10113cdc8?, 0x1400009f790?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:118 +0x308
cmd/compile/internal/types2.(*subster).var_(0x140000d5b88?, 0x140004567e0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:285 +0x30
cmd/compile/internal/types2.(*subster).varList(0x1400045dae8?, {0x140000b65d8, 0x1, 0x1006e606c?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:311 +0x84
cmd/compile/internal/types2.(*subster).typ(0x1400045dce8?, {0x10113ce40?, 0x140004524e0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:111 +0x1e0
cmd/compile/internal/types2.(*Checker).subst(0x140000f83c0, {0x1400035b860?, 0x7c?, 0x0?}, {0x10113ce40?, 0x140004524e0}, 0x140004526c0, 0x14000456850, 0x140003c81c0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/subst.go:78 +0x160
cmd/compile/internal/types2.(*Named).expandUnderlying(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:623 +0x45c
cmd/compile/internal/types2.(*Named).resolve(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:177 +0x1a8
cmd/compile/internal/types2.(*Named).Underlying(0x3?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:456 +0x1c
cmd/compile/internal/types2.(*Named).under(0x14000456850)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/named.go:484 +0x28
cmd/compile/internal/types2.under({0x10113cda0?, 0x14000456850?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:19 +0x4c
cmd/compile/internal/types2.coreType({0x10113cda0?, 0x14000456850?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/type.go:33 +0xa0
cmd/compile/internal/types2.(*Checker).exprInternal(0x140000f83c0, 0x1400044cb80, {0x10113e8a0?, 0x140000f1280}, {0x0?, 0x0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1403 +0x720
cmd/compile/internal/types2.(*Checker).rawExpr(0x140000f83c0, 0x1400044cb80, {0x10113e8a0?, 0x140000f1280?}, {0x0?, 0x0?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1252 +0x128
cmd/compile/internal/types2.(*Checker).multiExpr(0x10113ce68?, 0x0?, {0x10113e8a0?, 0x140000f1280?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/expr.go:1817 +0x38
cmd/compile/internal/types2.(*Checker).exprList(0x140000f83c0?, {0x1400045eac8?, 0x10113e820?, 0x140003c6e70?}, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/call.go:273 +0x68
cmd/compile/internal/types2.(*Checker).assignVars(0x140000f83c0, {0x1400045ead8?, 0x1, 0x1}, {0x1400045eac8?, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/assignments.go:384 +0x5c
cmd/compile/internal/types2.(*Checker).stmt(0x140000f83c0, 0x0, {0x10113d748?, 0x140003e8280?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:472 +0x62c
cmd/compile/internal/types2.(*Checker).stmtList(0x10113c620?, 0x0, {0x1400009ee80?, 0x100e093c0?, 0x140000f83c0?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:123 +0x80
cmd/compile/internal/types2.(*Checker).funcBody(0x140000f83c0, 0x140004500c0, {0x140000bd760?, 0x140000f83c0?}, 0x1400044c7c0, 0x140003e8240, {0x0, 0x0})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/stmt.go:43 +0x2bc
cmd/compile/internal/types2.(*Checker).funcDecl.func1()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/decl.go:760 +0x44
cmd/compile/internal/types2.(*Checker).processDelayed(0x140000f83c0, 0x0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:383 +0x14c
cmd/compile/internal/types2.(*Checker).checkFiles(0x140000f83c0, {0x140000b64d8, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:328 +0x144
cmd/compile/internal/types2.(*Checker).Files(...)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/check.go:300
cmd/compile/internal/types2.(*Config).Check(0x140000ab9f8?, {0x16f72e039?, 0x100f0dddf?}, {0x140000b64d8, 0x1, 0x1}, 0x101603fc0?)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/types2/api.go:434 +0x68
cmd/compile/internal/noder.checkFiles({0x140000b64b0, 0x1, 0x100f00509?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/irgen.go:73 +0x3b4
cmd/compile/internal/noder.writePkgStub({0x140000b64b0, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:210 +0x30
cmd/compile/internal/noder.unified({0x140000b64b0?, 0x1400035b8f0?, 0x2?})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/unified.go:75 +0x94
cmd/compile/internal/noder.LoadPackage({0x140000c6130, 0x1, 0x1})
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/noder/noder.go:77 +0x400
cmd/compile/internal/gc.Main(0x101137ad0)
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/internal/gc/main.go:196 +0xcd8
main.main()
        /opt/homebrew/Cellar/go/1.20.3/libexec/src/cmd/compile/main.go:57 +0xf4
FAIL    github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics [build failed]
FAIL
```

</details>
