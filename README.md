# boa

`github.com/schleppy/boa` is a small CLI command for parsing flags

Inspiration taken from Cobra

- **Commands and nested subcommands**: `Command`, `AddCommand`, aliases, `Find`/`Traverse`, command groups
- **Flags** via [pflag](https://github.com/spf13/pflag): local, persistent, required, and flag groups
  (`MarkFlagsRequiredTogether`, `MarkFlagsOneRequired`, `MarkFlagsMutuallyExclusive`)
- **Positional args**: `NoArgs`, `ExactArgs`, `MinimumNArgs`, `RangeArgs`, `OnlyValidArgs`, ...
- **Hooks**: `PreRun` / `Run` / `PostRun` (and persistent and `E` variants)
- **Help and usage**: auto `help` command, `--help`/`--version`, usage on bad input, "did you mean" suggestions

Removed: shell completions, active help, and doc generation.

```go
root := &boa.Command{Use: "app", Short: "demo"}
root.AddCommand(&boa.Command{
    Use:  "greet NAME",
    Args: boa.ExactArgs(1),
    RunE: func(c *boa.Command, args []string) error { c.Println("hi", args[0]); return nil },
})
root.Execute()
```
