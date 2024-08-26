## Example golang monorepo boilerplate

This repository serves as an example of using multiple modules/binaries under one git with multiple tags according to official [docs](https://go.dev/doc/modules/managing-source)

## Packages and binaries release process
1. Update module you'd like to release. It is preferable that you work on one module in one PR, but it is possible to release multiple packages from one commit.

2. Merge your PR. Then add module tags in format `mod/vX.X.X` on the `main` branch
    ```
    git tag one/v1.3.0
    git tag two/v1.2.0
    ```
   
   Annotated tags also work, but you need to pull them by commit like `go get ...@$SHA` then `go.mod` will be updated with a proper tag

   Preferred way is to use `lightweight tags` like `git tag mod/vX.X.X`

3. After your commit on `main` branch has corresponding tags you can consume it via `go get`
    ```
   go get github.com/skudasov/go-monorepo-boilerplate/one@v1.3.0
   go get github.com/skudasov/go-monorepo-boilerplate/two@v1.2.0
    ```
   In contradiction to official [docs] `require` won't work in you'll use `module/vX.X.X` tag so if you want to update `go.mod` manually use tag notation without module, ex.: `github.com/skudasov/go-monorepo-boilerplate/one v1.2.0`

4. Your module may have a binary release. Entrypoint must be under `cmd` and named as a package.
   
   After pushing a tag workflow will automatically create a release for `$package/vX.X.X` and publish binaries to [releases](https://github.com/skudasov/go-monorepo-boilerplate/releases) page