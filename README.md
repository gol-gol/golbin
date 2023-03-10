
## golbin

> facade for regular used Go calls to manage Shell Evals or Shell style function calls
>
> extracted from an older pandora box of such packages at [abhishekkr/gol](https://github.com/abhishekkr/gol)

### Public Functions

* `IsSystemCmd(cmd string) bool`

* `(konsole *Console) Run() (err error)`
* `ExecOutput(cmdline string) string`
* `RunWithAssignedApp(runThis string) string`
* `Exec(cmd string) (string, error)`
* `ExecWithEnv(cmd string, env map[string]string) (string, error)`

---
