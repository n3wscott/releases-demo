# releases-demo
Trying to make multi-step releases with GitHub Actions

In the repo are 2 releasable components:

- `./subcomponent`, v1 of the api/implementation, produces `module github.com/n3wscott/releases-demo/subcomponent`.
- `./v2`, v1 of the api/implementation, produces `module github.com/n3wscott/releases-demo/v2`.
- `./subcomponent`, using both v1 and v2 apis, produces `github.com/n3wscott/releases-demo/subcomponent`

Note:

- `v2` depends on `subcomponent`
- `subcomponent` depends on `v1` and `v2`.

## Release process

1. Create a `release-v<major>.<minor>` release branch.
2. Release `subcomponent`.
   1. Tag the release branch repo with the correct semver tag in the form `v1.<minor>.<patch>`
3. Release `v2`.
   1. Confirm `v1` has a new release.
   2. Update the version of `v1` in `v2`'s go mod.
   3. Drop the `replace` directive in `v2`'s go mod file.
   4. Tag the release branch repo with the correct semver tag in the form `v2.<minor>.<patch>`
4. Release `subcomponent`
   1. Confirm `v1` has a new release.
   2. Update the version of `v1` in `subcomponent`'s go mod.
   3. Confirm `v2` has a new release.
   4. Update the version of `v2` in `subcomponent`'s go mod.
   5. Drop the `replace` directive in `v2`'s go mod file.
   6. Tag the release branch repo with the correct semver tag in the form `subcomponent/v2.<minor>.<patch>`

