# log
This package extends the current go logging package to support log levels.  It is a non-invasive in-place replacement of go's
standard logging package.

## Log Levels

- DEBUG
- INFO
- ERROR

Log levels can be set as follows:
```
log.SetLevel("INFO")
```

Aside from the one function above, all remainder functionality is taken directly from the go log library.
