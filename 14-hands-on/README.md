## Commands

POST: curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user

GET: curl http://localhost:8080/user/:id

DELETE: curl -X DELETE http://localhost:8080/user/:id