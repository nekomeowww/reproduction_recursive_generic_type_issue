# Minimum reproduction of **disappeared** `invalid recursive type` error when defining one of the type parameter of type `innerT` with union type.

Unlike the separated file reproduction case in [`minimum_repro/unstable_invalid_recursive_issue/seperated_files`](../seperated_files) where it is harder to reproduce and the definition of type `innerT`, `T1` and `T2` are separated into two files and compiler would complain about the `invalid recursive type` error. It is more stable and easier to reproduce when the definition of type `innerT`, `T1` and `T2` are in the same file while we can observe another strange behavior of the compiler: it wouldn't complain about the `invalid recursive type` error but it also wouldn't generate the needed types and runtime assembly codes.

## How to reproduce the issue?

1. You have to create a file that named `file_1.go` and containing the following definition (the same as the separated file reproduction case):

```go
package seperated_files

// It doesn't matter if the innerT struct unexported or exported, the result is the same.
// It also doesn't matter if the R type parameter is infer to a pointer or not, the result is the same.
type innerT[T any, R T1[T] | T2[T]] struct {
    reference *R
}

type T1[T any] struct {
    e *innerT[T, T1[T]]
}

type T2[T any] struct {
    e *innerT[T, T2[T]]
}
```

2. You may noticed that the `invalid recursive type innerT compilerInvalidDeclCycle` error will no longer appear in both editor, IDE, and compiler if you try to build the project no matter how you rename the source code file just like before. The case is more strange and unlike than the [`minimum_repro/unstable_invalid_recursive_issue/seperated_files`](../seperated_files) case where the compiler will print the `invalid recursive type innerT compilerInvalidDeclCycle` error. However, this doesn't mean the issue is fixed or faded, the actual issue is behind the scene and relate to the compiled Assembly code from Golang compiler itself.

3. Let's run the following command to ask the compiler to build (`-gcflags=-S` means print assembly listing, you can check the description by running `go tool compile --help`).

```shell
go build -gcflags=-S ./minimum_repro/unstable_invalid_recursive_issue/single_file
# github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/unstable_invalid_recursive_issue/single_file
go:cuinfo.producer.github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/unstable_invalid_recursive_issue/single_file SDWARFCUINFO dupok size=0
        0x0000 2d 73 68 61 72 65 64 20 72 65 67 61 62 69        -shared regabi
go:cuinfo.packagename.github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/unstable_invalid_recursive_issue/single_file SDWARFCUINFO dupok size=0
        0x0000 73 69 6e 67 6c 65 5f 66 69 6c 65                 single_file
```

4. You may have noticed that the corresponding, runtime related assembly code for type `innerT`, `T1` and `T2` are missing in the output. This means the compiler optimized the invalid recursive type away. The normal code would be like this I you would try.
