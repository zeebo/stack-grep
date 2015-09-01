# stack-grep

Grep go stacks

# Usage

```
cat goroutine.dump | stack-grep [-v] <go regex>
```

`-v` will invert. Only reads on stdin because of reasons.
