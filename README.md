# jub0bs/iterutil

[![tag](https://img.shields.io/github/tag/jub0bs/iterutil.svg)](https://github.com/jub0bs/iterutil/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-%23007d9c)
[![Go Reference](https://pkg.go.dev/badge/github.com/jub0bs/iterutil.svg)](https://pkg.go.dev/github.com/jub0bs/iterutil)
[![license](https://img.shields.io/badge/License-MIT-yellow.svg?style=flat)](https://github.com/jub0bs/iterutil/raw/main/LICENSE)
[![build](https://github.com/jub0bs/iterutil/actions/workflows/iterutil.yml/badge.svg)](https://github.com/jub0bs/iterutil/actions/workflows/iterutil.yml)
[![codecov](https://codecov.io/gh/jub0bs/iterutil/branch/main/graph/badge.svg?token=N208BHWQTM)](https://app.codecov.io/gh/jub0bs/iterutil/tree/main)
[![goreport](https://goreportcard.com/badge/jub0bs/iterutil)](https://goreportcard.com/report/jub0bs/iterutil)

An experimental collection
(partly inspired by [Haskell][haskell]'s [prelude][prelude])
of utility functions for working with [Go][golang] [iterators].

## Installation

```shell
go get github.com/jub0bs/iterutil
```

jub0bs/iterutil requires Go 1.23 or above.

## Documentation

The documentation is available on [pkg.go.dev][pkgsite].

## Code coverage

![coverage](https://codecov.io/gh/jub0bs/iterutil/branch/main/graphs/sunburst.svg?token=N208BHWQTM)

## License

All source code is covered by the [MIT License][license].

## FAQ

### Can I depend on this library?

You can, but at your own peril. As stated above, this library is experimental;
more than anything else, it's an excuse for me to
flex my functional-programming muscles and
familiarize myself with [Go 1.23's iterators][iterators].
At this early stage, I reserve the right, upon new releases, to break the API:
some functions may see their names and/or signatures change,
and some functions may be removed altogether.

If you need a few functions from this library but do not want to depend on it,
feel free to copy their sources in your project;
[a little copying is better than a little dependency][copying].

### How should I use this library?

Above all, use it with parsimony.
The functional style displayed by this library is far from ideal in Go,
for several reasons: 

- code readability may suffer, in part
  because Go's idiosyncracies hinder a "[fluent interface][fluent]" style and
  because Go lacks a concise notation for anonymous functions;
- a more classic and imperative style is likely to prove much more performant;
- Go lacks the powerful [laziness][lazy] of [Haskell][haskell].

Bear in mind that the existence of this library is no license
to adopt a functional style all over your codebase!

[copying]: https://www.youtube.com/watch?v=PAAkCSZUG1c&t=568s
[fluent]: https://en.wikipedia.org/wiki/Fluent_interface
[golang]: https://go.dev/
[haskell]: https://www.haskell.org/
[iterators]: https://go.dev/blog/range-functions
[lazy]: https://en.wikipedia.org/wiki/Lazy_evaluation
[license]: https://github.com/jub0bs/iterutil/blob/main/LICENSE
[pkgsite]: https://pkg.go.dev/github.com/jub0bs/iterutil
[prelude]: https://downloads.haskell.org/ghc/9.8.2/docs/libraries/base-4.19.1.0-179c/Prelude.html