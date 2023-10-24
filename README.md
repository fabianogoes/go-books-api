# PoC Go books API
> API publica para cadastro de livros sobre desenvolvimento de software

## Objetivo desta PoC
Criar uma API REST com [Golang][1] para praticar a linguagem. 
E usar a Plataform PaaS [Render][0] para fazer o deploy.

## Features para serem implementadas e testadas
- [x] Criar uma API Simples em [Go][1] que implemente um CRUD de cadastro de Livros, no primeiro momento usar um banco de dados InMemory.
- [x] Implementar o design de [Arquitetura Hexagonal][4].
- [x] Criar um serviço no [Render][0] do tipo `Web Service` e Configurar o Deploy usando a opção de `Build and deploy from git repository` com o Runtime esteja como `Docker`.
- [x] Implementar no projeto o Repository com [Posgres][2] e [GORM ORM][3].
- [x] Criar um novo Serviço no [Render][0] do tipo `Postgres`. 
- [x] Configurar no [Render][0] o link entre a API e o Posgres usando `Env Groups`.
- [ ] Implementar Testes unitários.
- [ ] Implementar Testes integrados.
- [ ] Implementar Documentação com Open API. 
- [x] Usar GitHub Actions para fazer o Build e Test.

## Stack
- [Golang][1]
- [GORM ORM][3]
- [PostgreSQL][2]
- [Render PassS][0]
- [Docker][5]

## Running

1. Up Postgres by Docker Compose
```shell
docker-compose up -d
```

2. Go app
```shell
go run main .
```

## Resultados
> Aproveitei essa PoC para documentar em um Post em meu LinkedIn: https://www.linkedin.com/pulse/deploying-aplica%C3%A7%C3%A3o-web-cloud-gratu%C3%ADto-renderuma-alternativa-g%C3%B3es

---

<img src="MIT.png" height="50" width="150">

[0]: https://render.com/
[1]: https://go.dev/
[2]: https://www.postgresql.org/
[3]: https://gorm.cn/
[4]: https://alistair.cockburn.us/hexagonal-architecture/
[5]: https://www.docker.com/


