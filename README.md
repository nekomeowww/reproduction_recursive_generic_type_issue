# recursive generic type issue reproduction

This repository showed how to reproduce the compiler unexpected behavior, or issue, bug when a struct that has a field to reference the outer struct that has a field to reference the initial struct by using generic in the way recursively: the compiler will complain `fatal error: all goroutines are asleep - deadlock!` or other behave unexpectedly in all different ways.

## Expected

- Compiler or documentation should address the issue when user attempts to do so.
- Compiler should have a consistent behavior when user attempts to do so.
- Compiler should produce the needed assembly codes for the generic recursive types or produces error or warning messages when user attempts to do so.

## Behaviors

- Compiler might complain the `fatal error: all goroutines are asleep - deadlock!` panics message when running the test by `go test -v ./...`.
- Compiler might be able to compile the needed assembly codes and run the test successfully when you write the test codes without the `_test` suffix (exported test).
- Compiler might complain about the `invalid recursive type` if you try to move the `T2[T any]` definition to another file.
- Compiler might not be able to generate the needed assembly codes for the generic recursive types and produces no error or warning messages silently.
- Compiler might complain about the `redecalred` if you try to reload the IDE somehow.

## How to reproduce the issue from this repository?

This is the most needed directory structure for reproducing the issue, each cases are well documented with the case, summary, and how-to-reproduce, you may navigate into the directory and run the command I provided in the `README.md` file to reproduce the issue by yourself on-demand. Or you could use the next section to click the link to the issue and case you want to reproduce by following the instructions.

```shell
.
├── minimum_repro
│   ├── deadlock_issue # issue 1: fatal error: all goroutines are asleep - deadlock!
│   │   ├── with_generics # at the case of using generics
│   │   ├── with_generics_with_pointer # at the case of using generics with pointer wrapped
│   │   └── with_generics_without_test_package # at the case of using generics without test package (no _test suffix)
│   ├── unstable_invalid_recursive_issue # issue 2: invalid recursive type
│   │   ├── seperated_files # at the case of seperated files
│   │   └── single_file # at the case of single file
│   └── unstable_redeclare_issue # issue 3: redeclared in this block
│       └── seperated_files # at the case of seperated files
└── original_scenario # the original scenario that I encountered the issue
```

### Issue 1: fatal error: all goroutines are asleep - deadlock!

- [when using generics with test package (`_test` suffix)](./minimum_repro/deadlock_issue/with_generics)
- [when using generics with pointer reference to the generic type](./minimum_repro/deadlock_issue/with_generics_with_pointer)
- [**disappeared** deadlock error when using generics **without** test package (`_test` suffix)](./minimum_repro/deadlock_issue/with_generics_without_test_package)

### Issue 2: invalid recursive type

- [when defining one of the type parameter of type `innerT` with union type](./minimum_repro/unstable_invalid_recursive_issue/seperated_files)
- [**disappeared** `invalid recursive type` error when defining one of the type parameter of type `innerT` with union type.](./minimum_repro/unstable_invalid_recursive_issue/single_file)

### Issue 3: redeclared in this block

- [when defining one of the type parameter of type `innerT` with union type](./minimum_repro/unstable_redeclare_issue/seperated_files)

## How did I encounter?

I was working on an experimental channel wrapper utilities library in [nekomeowww/xo/exp](https://github.com/nekomeowww/xo/tree/main/exp), and I tried to discover the possibility of using generics to make the common options assignment easier and more readable by using chained calls pattern in [`exp/channelx/puller.go`](https://github.com/nekomeowww/xo/commit/cf486f5ea50b84c6df530b1203f5d82a5a4dbc0c#diff-d45899a6268a5dbaf2b914a1e382cfa7d495f9079dfd161624e123dde032c05e). You could see I commented out all the examples in the [`exp/channelx/example_test.go`](https://github.com/nekomeowww/xo/commit/cf486f5ea50b84c6df530b1203f5d82a5a4dbc0c#diff-c489abd579609224cb8d2b19e42f630cfe8fb1420a5a630b93d760c3b02b5787) in commit [`cf486f5`](https://github.com/nekomeowww/xo/commit/cf486f5ea50b84c6df530b1203f5d82a5a4dbc0c) just because I got stucked by the compiler error `fatal error: all goroutines are asleep - deadlock!` with the test package (`_test` suffix) included for examples. I didn't realized that the problem was came from by Golang compiler it self, just because the tests in [`exp/channelx/puller_test.go`](https://github.com/nekomeowww/xo/commit/cf486f5ea50b84c6df530b1203f5d82a5a4dbc0c#diff-518be0525b26d5ac85efe6f40ba39e23b6d0f92540123409d98d7ab7d50c73de) can be run successfully and without any problem until now.

## How do I think?

TL;DR: It behaves confusingly for users to understand what's going on. At least it jammed me for a while. And I think it should be fixed or addressed by the compiler or documentation.

Even though I could understand the limitations of the current implementation for generics and the "should avoid" with recursive types, but still, I think this is a issue that either Golang team needs to decide whether the compiler or documentation should address when user attempts to do so or Golang team fixes such issue and turns it into the expected way.

Especially it runs successfully without generating the needed assembly codes, and ran successfully when you change the package of `main_test.go` into `recursive_generic_type_issue_reproduction` without the `_test` suffix while developing. Imaging you as being a Golang developer, might only find out what has been wrong after you released the library without the `_test` test packages test included and caused the compiler issue and complains for all your users, as I stated out in [**disappeared** deadlock error when using generics **without** test package (`_test` suffix)](./minimum_repro/deadlock_issue/with_generics_without_test_package) case.

## Workaround?

Such unexpected behavior could be resolved by `GOEXPERIMENT=nounified` flag[^1] and it is rarely reported (or even nobody have mentioned so far and so detailed).

[^1]: [cmd/compile: failed to compile some recursive generic type · Issue #54535 · golang/go](https://github.com/golang/go/issues/54535)
