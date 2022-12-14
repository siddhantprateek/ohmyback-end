# GraphQL + Mongodb + Go

Integration of Mongodb and GraphQL


## Getting started

![Reference](https://www.apollographql.com/blog/graphql/golang/using-graphql-with-golang/)

Let’s get started by installing gqlgen and initializing our project. We can fetch the library using the following command.

```shell
go get github.com/99designs/gqlgen
```

Next, add gqlgen to your project’s tools.go.

```shell
printf '// +build tools\npackage tools\nimport _ "github.com/99designs/gqlgen"' | gofmt > tools.go
```

Then initialize gqlgen config and generate the models.

```shell
go run github.com/99designs/gqlgen init
```

Finally, start the GraphQL server.
```shell
go run server.go
```
