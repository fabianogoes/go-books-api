curl http://localhost:3000/books \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "DDD", "description": "test DDD", "author": "Uncle Bob"}'