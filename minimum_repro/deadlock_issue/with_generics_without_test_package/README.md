# Minimum reproduction of **disappeared** deadlock error when using generics **without** test package (`_test` suffix)

Unlike the [`minimum_repro/with_generics`](../with_generics) example I used the package with a suffix of `_test` in the tests file (`_test.go`), I
removed the `_test` suffix from the package name and the unit tests file instead.

The compiler no longer complain about the
`fatal error: all goroutines are asleep - deadlock!` issue, and the unit tests passed.

The most strange thing is this, by removing the `_test` suffix from the [`test_without_test_package_test.go`](../with_generics/main_without_test_package_test.go) file that belongs to a test package, the compile error will disappear, and the test will pass with no error. Imagine that you maintain a package that is widely used by other users, you never test your code with a test package that suffix with `_test` (it's not enforced or not part of the regulation), and you would have no ways to understand how and why the compiler would complain and throw a panic with `fatal error: all goroutines are asleep - deadlock!` message when users imported your package after you released a version.

## How to reproduce the issue?

Run the following command in the project root directory:

```shell
go test -v ./minimum_repro/deadlock_issue/with_generics_without_test_package
```

<details>
<summary>Output</summary>

```shell
$ go test -v ./minimum_repro/deadlock_issue/with_generics_without_test_package
=== RUN   TestT1
--- PASS: TestT1 (0.00s)
PASS
ok      github.com/nekomeowww/recursive_generic_type_issue_reproduction/minimum_repro/deadlock_issue/with_generics_without_test_package 0.004s
```

</details>
