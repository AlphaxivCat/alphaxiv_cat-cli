# Changelog

## 0.5.0 (2026-04-23)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([bbe07d4](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/bbe07d4411b2a7ab3502cac47e241cab36eba284))


### Chores

* **internal:** more robust bootstrap script ([12543dd](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/12543dd0175a9f6bd4483c564c49185c67df473a))

## 0.4.0 (2026-04-18)

Full Changelog: [v0.3.1...v0.4.0](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.3.1...v0.4.0)

### Features

* **cli:** add `--raw-output`/`-r` option to print raw (non-JSON) strings ([2eb9846](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/2eb9846df90eeeab8a6c02b94def0cdd1efe15e5))
* **cli:** send filename and content type when reading input from files ([1275415](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/12754156504dd13f763d3d59295c789bd5f4d5f6))


### Chores

* **ci:** support manually triggering release workflow ([7a1344c](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/7a1344c0a844bb79e03a0840e6e5601eecc09cfe))
* **cli:** use `ShowJSONOpts` as argument to `formatJSON` instead of many positionals ([9cb55ed](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/9cb55ed0d102d2ea5035fad1c030776f9077b72e))

## 0.3.1 (2026-04-16)

Full Changelog: [v0.3.0...v0.3.1](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.3.0...v0.3.1)

### Chores

* **cli:** switch long lists of positional args over to param structs ([99acc54](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/99acc540d457526be101913dfaa908a79fc035a7))

## 0.3.0 (2026-04-15)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.2.0...v0.3.0)

### Features

* **api:** api update ([2fab5d8](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/2fab5d8fe069f2df56fb60e1ec1cbf6df4f7ebf1))
* **cli:** alias parameters in data with `x-stainless-cli-data-alias` ([d1f22c2](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/d1f22c29f33b1d05922a6f5ae9de0facc47de772))


### Chores

* **cli:** fall back to JSON when using default "explore" with non-TTY ([2927694](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/2927694dfbbf6b130c262906f6441b45a06824e6))

## 0.2.0 (2026-04-14)

Full Changelog: [v0.1.2...v0.2.0](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.1.2...v0.2.0)

### Features

* add default description for enum CLI flags without an explicit description ([10a03fc](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/10a03fc708e7e5e81be41d8c623581a2294fc995))
* allow `-` as value representing stdin to binary-only file parameters in CLIs ([da39f06](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/da39f06663765ba233b632bfff766b525445de8e))
* **api:** api update ([4d8e42b](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/4d8e42b7ead59344b3c4a85ad266771c6dcbb339))
* better error message if scheme forgotten in CLI `*_BASE_URL`/`--base-url` ([6fd2f08](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/6fd2f084ba36cca0cf5adc2ff039e8e4ebc5463c))
* binary-only parameters become CLI flags that take filenames only ([9098313](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/9098313044507cba62bc3b759a7691ae550c5a3b))
* set CLI flag constant values automatically where `x-stainless-const` is set ([f2b90b6](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/f2b90b691d16fa41c65c52c400b26a338e1ae207))


### Bug Fixes

* cli no longer hangs when stdin is attached to a pipe with empty input ([ed98c7e](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/ed98c7e7980f2cfe9edb8da5a1c861214be1b1ae))
* **cli:** fix incompatible Go types for flag generated as array of maps ([a43f326](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/a43f326e2e739853fc00f01e1436322d74cfe8c9))
* fall back to main branch if linking fails in CI ([f2c2949](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/f2c2949da51aa8ebc1fea89691273ed32a735755))
* fix for failing to drop invalid module replace in link script ([682547f](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/682547fa3b8e9060b628e09d19f37b633dabca02))
* fix for off-by-one error in pagination logic ([042ff72](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/042ff7286d7aec63c31a4fc708320aefa0b84482))
* fix quoting typo ([f7ecd06](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/f7ecd06677265a0ea37e2e9a01834a8adb19e41e))
* handle empty data set using `--format explore` ([9004ee3](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/9004ee31e373d8b38fb90e3069aced5799e25fe0))
* use `RawJSON` when iterating items with `--format explore` in the CLI ([db150da](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/db150da3472e3c3c804e2518f6d43a1a7bd0b003))


### Chores

* add documentation for ./scripts/link ([1bf9315](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/1bf93159267efded479aa25f3049a56f8bec7a9f))
* **ci:** skip lint on metadata-only changes ([d877e60](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/d877e60b4df7d413f5f5b389a5c44c42063d5886))
* **cli:** additional test cases for `ShowJSONIterator` ([3b1f62a](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/3b1f62a470cb3ed5c7bcc07629762873dc80d733))
* **cli:** let `--format raw` be used in conjunction with `--transform` ([1c94002](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/1c94002d3fddb3ab9a49aba2d4df5ae6a7b72c37))
* **internal:** update gitignore ([c83f699](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/c83f699524f3a0a9ee45d4ff180721457f17a06e))
* mark all CLI-related tests in Go with `t.Parallel()` ([359dc17](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/359dc1756dde46e4328f87ae9a2fee32769793ac))
* modify CLI tests to inject stdout so mutating `os.Stdout` isn't necessary ([47aa309](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/47aa30941c09aaa8f808f7a6d4284390717b2814))
* omit full usage information when missing required CLI parameters ([52ebf4f](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/52ebf4ffe6e47e4a9d42d0ae9a390e7a59510826))
* switch some CLI Go tests from `os.Chdir` to `t.Chdir` ([5406d0e](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/5406d0e765bf8f6e2b65db5700ae5b8866640543))
* **tests:** bump steady to v0.19.4 ([9f4da7a](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/9f4da7ada3b22decd701e841b836714814733e8b))

## 0.1.2 (2026-03-19)

Full Changelog: [v0.1.1...v0.1.2](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.1.1...v0.1.2)

### Bug Fixes

* avoid reading from stdin unless request body is form encoded or json ([1674f47](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/1674f474425c12eb583dd0993e002d624a12347b))
* improve linking behavior when developing on a branch not in the Go SDK ([162ab83](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/162ab83ac7027220da99954c617f0c85b25d6f56))

## 0.1.1 (2026-03-18)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.1.0...v0.1.1)

### Bug Fixes

* better support passing client args in any position ([713b181](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/713b181d6e253790bb95f4c2f85cf27aeabd7a70))
* improved workflow for developing on branches ([e89d658](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/e89d65869a7ec13e553dcfb349205333037dbb36))
* no longer require an API key when building on production repos ([29f1c7c](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/29f1c7c536b84ba6a7d2f39030d694eff8522813))
* only set client options when the corresponding CLI flag or env var is explicitly set ([6b464df](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/6b464dfa3bee4fa786b4cbc3a5d6665548854b88))


### Chores

* **internal:** tweak CI branches ([b6cac50](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/b6cac50dc72c35c8705a5cae471ef3358b8dd34d))

## 0.1.0 (2026-03-12)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/AlphaxivCat/alphaxiv_cat-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([b67ae58](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/b67ae58bba57758e625e5e8735e3341f0a9d7d6c))


### Chores

* configure new SDK language ([c226b56](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/c226b5642ab8f6990e092843a701017900ab56fa))
* update SDK settings ([b78ba62](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/b78ba62b59bb64ea299aeff2c9cd875f510b727e))
* update SDK settings ([3104e53](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/3104e5305ea50f6b201dbb055d30ba4220ba5dcb))
* update SDK settings ([8734e3a](https://github.com/AlphaxivCat/alphaxiv_cat-cli/commit/8734e3a196d672f3265fa193a7ad771b998b7ce6))
