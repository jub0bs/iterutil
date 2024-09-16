# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] (2024-09-16)

## Added

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

[0.2.0]: https://github.com/jub0bs/iterutil/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/jub0bs/iterutil/releases/tag/v0.1.0
