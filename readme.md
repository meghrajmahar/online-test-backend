Schema: Tumhari GraphQL definition (schema.graphqls) read hoti hai.

Exec: gqlgen ek generated.go banata hai jo GraphQL requests ko Go code se connect karta hai.

Model: GraphQL types ke liye Go structs banata hai (models_gen.go).

Resolver: Tumhare liye function stubs banata hai (internal/graph/resolvers/...), jahan tum database ya business logic likhte ho.


Run your project to call graphql hit this command
    go run cmd/api/main.go 

If any change in schema.graphql then hit this command
    gqlgen generate