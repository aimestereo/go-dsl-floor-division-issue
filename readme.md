# DSL Parsing Issue

## Reproduce

```bash
git clone git@github.com:aimestereo/go-dsl-floor-division-issue.git
cd go-dsl-floor-division-issue
go mod tidy -v

go test -v -race -buildvcs ./...
```

## Output

```
=== RUN   TestOperationParser
2024/11/03 18:38:55 Testing: +
main.Expr{
  Op: "+",
}
2024/11/03 18:38:55 Testing: !=
main.Expr{
  Op: "!=",
}
2024/11/03 18:38:55 Testing: /
main.Expr{
  Op: "/",
}
2024/11/03 18:38:55 Testing: //
main.Expr{
}
    parser_test.go:28: Did not expect an error but got:
        1:3: unexpected token "<EOF>"
--- FAIL: TestOperationParser (0.00s)
FAIL
FAIL    dsl     0.300s
FAIL


```

## Redraw diagram

Expecting `railroad` pre-built as `/tmp/bin/railroad`

```bash
go run .
open open diagram/dsl.html
```
