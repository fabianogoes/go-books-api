curl http://localhost:3000/books/38df470b-4605-48ad-980e-2642c340439f \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"title": "DDD", "description": "DDD updated", "author": "Eric Evans"}'