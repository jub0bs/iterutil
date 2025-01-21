# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.5.1] (2025-01-21)

### Changed

- **Performance**: Further optimize function `SortedFromMap`.

## [0.5.0] (2025-01-20)

## Fixed

- **Bug**: Functions `Take` and `Drop` would previously produce incorrect
  results when used with unsigned integers.

## Added

- **API**: function `SortedFromMapFunc` (renamed from `SortedFuncFromMap`)
- **Tests**: Add tests for cases where `uint` is used as type argument for
  type parameters constrained by
  [`golang.org/x/exp/constraints.Integer`][constraints.Integer].
- **Tests**: Add benchmarks for `SortedFromMap`.

### Changed

- **Dependencies**: Go 1.23.2 (or above) is now required.
- **Dependencies**: Update golang.org/x/exp to latest.
- **API** (breaking change): The parameters of `Between` are now constrained by
  [`golang.org/x/exp/constraints.Signed`][constraints.Signed] rather than by
  [`golang.org/x/exp/constraints.Integer`][constraints.Integer].
- **Performance**: Functions `SortedFromMap` and `SortedFromMapFunc` are now
  optimized for cases where iteration is cut short.
- **Documentation**: minor fixes

### Removed

- **API** (breaking change): function `SortedFuncFromMap`

## [0.4.0] (2024-10-03)

### Added

- **API**: function `Reduce` (renamed from `Foldl`)

### Changed

- **API** (breaking change): The second argument of functions `At`, `Drop`,
  `Repeat`, and `Take` no longer needs be an `int`; for more flexibility,
  it can now be any value whose type is constrained by
  [`golang.org/x/exp/constraints.Integer`][constraints.Integer].
- **Documentation**: minor fixes

### Removed

- **API** (breaking change): function `Foldl`
- **API** (breaking change): function `AllErrors`

## [0.3.0] (2024-09-21)

### Added

- **API**: function `Flatten` (renamed from `Concat`)
- **API**: function `Between`
- **API**: function `Enumerate`
- **API**: functions `Equal`, `EqualFunc`
- **API**: functions `Compare`, `CompareFunc`, `IsSorted`, and `IsSortedFunc`
- **API**: functions `Min`, `MinFunc`, `Max`, and `MaxFunc`
- **API**: functions `SortedFromMap` and `SortedFuncFromMap`

### Changed

- **API** (breaking change): the type parameter in function `ContainsFunc`
  is now (correctly) unconstrained.
- **API** (breaking change): `Concat` previously took an iterator over iterators;
  it now takes a slice of iterators.
- **Behavior**: functions `Take` and `Drop` now tolerate a negative `count`
  argument.
- **Documentation**: sinks that may not terminate are now documented as such.
- **Documentation**: various other improvements

### Removed

- **API** (breaking change): function `AllLeafErrors`
- **API** (breaking change): function `Append`
- **API** (breaking change): function `FlatMap`
- **API** (breaking change): functions `Cons`, `Head`, `Tail`, and `Uncons`

## [0.2.0] (2024-09-16)

### Added

- **Tests**: augment test suite to reach 100% code coverage.

### Changed

- **API** (breaking changes): function `Repeat` now has an `count` parameter
  that specifies the number of repetitions of the desired value in the
  resulting iterator; if `count` is negative, the resulting iterator
  is infinite.
- **Behavior** (breaking change): function `At` now panics if its `count`
  argument is negative.
- **Behavior** (breaking change): function `Drop` now panics if its `count`
  argument is negative.
- **Behavior** (breaking change): function `Take` now panics if its `count`
  argument is negative.
- **Documentation**: various improvements
- **Tests**: improve examples

### Removed

- **API** (breaking change): function `Replicate`

## [0.1.0] (2024-09-14)

[0.5.1]: https://github.com/jub0bs/iterutil/compare/v0.5.0...v0.5.1
[0.5.0]: https://github.com/jub0bs/iterutil/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/jub0bs/iterutil/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/jub0bs/iterutil/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/jub0bs/iterutil/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/jub0bs/iterutil/releases/tag/v0.1.0

[constraints.Integer]: https://pkg.go.dev/golang.org/x/exp/constraints#Integer
[constraints.Signed]: https://pkg.go.dev/golang.org/x/exp/constraints#Signed
