Schema: Tumhari GraphQL definition (schema.graphqls) read hoti hai.

Exec: gqlgen ek generated.go banata hai jo GraphQL requests ko Go code se connect karta hai.

Model: GraphQL types ke liye Go structs banata hai (models_gen.go).

Resolver: Tumhare liye function stubs banata hai (internal/graph/resolvers/...), jahan tum database ya business logic likhte ho.


Run your project to call graphql hit this command
    go run cmd/api/main.go 

If any change in schema.graphql then hit this command
    gqlgen generate


Sabse pahle model create karo





========================= http://localhost:8080/ =========================

1)  Get All Exam ==========> 
    query ExamList($limit: Int = 10, $offset: Int = 0, $search: String, $sort: String) {
    examList(limit: $limit, offset: $offset, search: $search, sort: $sort) {
        total
        items {
        id
        title
        description
        durationMin   # or durationMinutes
        }
    }
    }


2) Create Exam ==========>
    mutation createExam($input: CreateExamInput!) {
    createExam(input: $input) {
        title
        description
        durationMin
        questions{
        statement
        questionType
        }
    }
    }

payload :
    {
    "input": {
        "title": "First Exam test",
        "description": "This is exam description",
        "durationMin": 100
    }
    }