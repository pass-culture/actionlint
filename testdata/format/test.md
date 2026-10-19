### Error at line 3, col 5 of `testdata/format/test.yaml`

unexpected key "branch" for "push" section. expected one of "branches", "branches-ignore", "paths", "paths-ignore", "tags", "tags-ignore", "types", "workflows"

```
    branch: main
    ^~~~~~~
```

### Error at line 13, col 20 of `testdata/format/test.yaml`

property "msg" is not defined in object type {}

```
          MSG: ${{ matrix.msg }}
                   ^~~~~~~~~~
```

### Error at line 14, col 9 of `testdata/format/test.yaml`

unexpected key "with" for step to run shell command. expected one of "continue-on-error", "env", "id", "if", "name", "run", "shell", "timeout-minutes", "working-directory"

```
        with:
        ^~~~~
```

