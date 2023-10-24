curl http://localhost:3000/books \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"title": "Domain-driven design: atacando as complexidades no coração do software", "description": "A comunidade de desenvolvimento de softwares reconhece que a modelagem de domínios é fundamental para o design de softwares. Através de modelos de domínios, os desenvolvedores de software conseguem expressar valiosas funcionalidades e traduzi-las em uma implementação de software que realmente atenda às necessidades de seus usuários. Mas, apesar de sua óbvia importância, existem poucos recursos práticos que explicam como incorporar uma modelagem de domínios eficiente no processo de desenvolvimento de softwares", "author": "Eric Evans"}'