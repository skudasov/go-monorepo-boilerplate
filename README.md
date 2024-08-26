## Example golang monorepo boilerplate

This repository serves as an example of using multiple modules under one git with multiple tags according to official [docs](https://go.dev/doc/modules/managing-source)

## Release process
1. Update modules
2. Tag each module with `mod/vX.X.X`, annotated tag
    ```
    git tag -a one/v1.1.0 -m "go mod package tag for one"
    git tag -a two/v1.0.0 -m "go mod package tag for two"
    ```
3. Update the deps in consumer, **use commit that has the tag**
    ```
   export UPDATE_SHA=$(git rev-parse HEAD)
   go get github.com/skudasov/go-monorepo-boilerplate/one@${UPDATE_SHA}
   go get github.com/skudasov/go-monorepo-boilerplate/two@${UPDATE_SHA}
    ```