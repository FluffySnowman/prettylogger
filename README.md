# PL - Simple Pretty Logger

PL is a simple, easy to use pretty logger made in go.


# Usage

Initialise the pretty logger with a config (can be SIMPLE, TIMEBASED) which
is automatically used everywhere once set.

Example:

    pl.InitPrettyLogger("SIMPLE")       // basic 
    pl.InitPrettyLogger("TIMEBASED")    // shows timestamps
    pl.LogInfo("Hello World")

Multiple arguments:

    pl.LogDebug("this is a debug log %v", "foo bar")

Force a timestamp log:

    pl.LogDebug("this is a debug log %v", "foo bar").Timestamp().Print()

Changing SIMPLE to TIMEBASED & vice versa will cause problems hence use the
.Timestamp() and .Print() when timestamps are needed.

