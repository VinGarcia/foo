This is a repository created only for testing how dependencies work on Golang.

This package depends on: github.com/vingarcia/bar and bar depends on this package in an attempt to create a circular dependency.

The result of this experiment is that modules can depend on other modules circularly as long as there isn't a cycle of packages depending on one another.

E.g. if the package `bar/cmd` depends on the package `foo` and foo depends on the package `bar` (not `bar/cmd`) everything works.

But if `foo` also depends on `bar/cmd` then we got a cyclic dependency and Go won't accept that.
