This represents a wag module ait-a that imports a "library" ait-b that then goes ahead and imports what used to be a package provided by ait-c but was later (in version v0.0.3) converted to a module. ait-a also imports ait-c but the module version (simulating trying to import a wagv9 thing)

Attempting to `go get github.com/ChrisScotMartin/ait-c/pkgmodc@v0.0.3` results in an error:

```
github.com/ChrisScotMartin/ait-c/pkgmodc: ambiguous import: found package github.com/ChrisScotMartin/ait-c/pkgmodc in multiple modules:
        github.com/ChrisScotMartin/ait-c v0.0.1 (/Users/chris.martin/.gvm/pkgsets/go1.16.15/global/pkg/mod/github.com/!chris!scot!martin/ait-c@v0.0.1/pkgmodc)
        github.com/ChrisScotMartin/ait-c/pkgmodc v0.0.3 (/Users/chris.martin/.gvm/pkgsets/go1.16.15/global/pkg/mod/github.com/!chris!scot!martin/ait-c/pkgmodc@v0.0.3)
```

This is because pkgmodc is currently provided by two different modules `github.com/ChrisScotMartin/ait-c` and `github.com/ChrisScotMartin/ait-c/pkgmodc`

This can be "fixed" with two replaces in the go.mod.

```
replace github.com/ChrisScotMartin/ait-c/pkgmodc => github.com/ChrisScotMartin/ait-c/pkgmodc v0.0.3
replace github.com/ChrisScotMartin/ait-c => github.com/ChrisScotMartin/ait-c v0.0.3
```

However this becomes exponentially harder to maintain the more we convert packages to modules. Anything up the chain would need these replace statements for every module with ambiguous import issues.

And since replace directives are not respected on import if ait-d were to import ait-a and ait-c but NOT ait-b (the offending library) it would still need these same replaces.

<img width="701" alt="image" src="https://user-images.githubusercontent.com/101210272/202280014-2aace7a1-4e41-4ab3-9740-02c509831f33.png">
