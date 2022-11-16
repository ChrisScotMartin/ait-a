Notice that in this example I get the output:

```
❯ go run main.go
Hello World
From pkgmoda
Hello, I'm either a package or a module...depends on version.
From pkgmodc
Hello, I'm now a module (pkgmodc) at this version.
I'm a func from the lib function in pkgbodb
From pkgmodc
Hello, I'm either a package or a module...depends on version.

```

Both versions of pkgmodc are being used in this example. The module version called by ait-a produced:

```
From pkgmodc
Hello, I'm now a module (pkgmodc) at this version.
```

while the package version imported from ait-b produced

```
From pkgmodc
Hello, I'm either a package or a module...depends on version
```
