curl http://localhost:3000/books/29a30c47-f534-4bd3-bbe4-6d8113b2f2cc \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"title": "DDD updated", "description": "test DDD updated", "author": "Uncle Bob"}'