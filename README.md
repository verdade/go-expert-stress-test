
# Go Expert Stress

## Instalação


Construa a imagem do projeto com:

```bash
make docker-build-image 
```

Este comando compilará o código fonte e criará um executável dentro de um container Docker.

## Uso

Para executar a ferramenta `go-expert-stress-test`, utilize o seguinte comando Docker:

```bash
docker run --rm ricanalista/go-expert-stress-test --url=https://google.com --requests=100 --concurrency=20
```

### Parâmetros

- `--url`: Especifica o URL alvo para o teste de carga. Substitua `https://google.com` pelo URL que você deseja testar.
- `--requests`: Define o número total de requisições a serem feitas ao URL alvo.
- `--concurrency`: Determina o número de requisições a serem feitas em paralelo, ou seja, o número de usuários simulados acessando o URL simultaneamente.


