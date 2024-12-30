# TODO

1. Improve and expand tests.
    + Use example tests or cmp to test usage output.
    + Add tests for defaults and isZero check.
    + Check coverage?
1. Change usage display to suit my preferences.  Do not special case short
   boolean flags; put usage on the same line as the flag. Use a tabwriter and
   wrap text to some good default for terminals?
1. Use a generic isZero check rather than reflect?
1. Add `Subcommand` and `Subcommands` (similar to `Alias` and `Aliases`).  These
   would support display of subcommands in usage, but I do not want to handle
   the overall parsing or running of subcommands (for now?).  Users will
   manually parse and run their subcommands.
1. Read both `flag` and `ff`.  Decide how I want to implement both long and
   short.  Do I want to use something like `Alias` and the stdlib, or should
   I do it from scratch.  Also, think about how I want to implement help
   displays and how much of the stdlib is worth carrying over.  E.g., `Visit`
   and `VisitAll`.
1. Demand that all flags implement an interface with two methods: `IsZero` and
   `Default`?  (Maybe call this a `Defaulter` interface?)  This could get rid of
   brittle, incomplete, `panic`-inducing, or `reflect`-requiring tests for
   whether something is a zero value.
