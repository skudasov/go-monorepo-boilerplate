## Example golang monorepo boilerplate

This repository serves as an example of using multiple modules under one git with multiple tags according to official [docs](https://go.dev/doc/modules/managing-source)

## Release process
1. Update modules
2. Tag each module with `mod/vX.X.X`, annotated tag
    ```
    git tag one/v1.3.0
    git tag two/v1.2.0
    ```
   Annotated tags also work, but you need to pull them by commit like `go get ...@$SHA` then `go.mod` will be updated with a proper tag
   Preferred way is to use `lightweight tags` like `git tag mod/vX.X.X`

3. Update the deps in consumer, **use commit that has the tag**
    ```
   go get github.com/skudasov/go-monorepo-boilerplate/one@v1.3.0
   go get github.com/skudasov/go-monorepo-boilerplate/two@v1.2.0
    ```
   In contradiction to official [docs] `require` won't work in you'll use `module/vX.X.X` tag so if you want to update `go.mod` manually use tag notation without module, ex.: `github.com/skudasov/go-monorepo-boilerplate/one v1.2.0`