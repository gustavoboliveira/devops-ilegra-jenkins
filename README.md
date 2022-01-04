### Requisitos:

- [Packer](https://www.packer.io/downloads)
- [Docker](https://www.docker.com/get-started)

### Execução:

1. Após clonar o repositório, abrir o terminal na pasta raiz do projeto
2. Utilize o comando `packer build build.json`
3. Em seguida, ainda no terminal, utilize o comando `docker run -it --rm -p 8080:8080 calculator:1.0`
4. Aplicação deverá estar executando no endereço `localhost:8080/`
5. Para encerrar a aplicação, abra um novo terminal e use o comando `docker ps` e copie o `CONTAINER ID` referente a calculadora
6. Então utilize o comando `docker stop CONTAINER_ID`

### Endpoints:

- `localhost:8080/calc/sum/{x}/{y}` - Retorna a soma de `x` e `y`, caso sejam informadas letras ao invés de números, será apresentada mensagem de erro
- `localhost:8080/calc/sub/{x}/{y}` - Retorna a subtração de `x` e `y`, caso sejam informadas letras ao invés de números, será apresentada mensagem de erro
- `localhost:8080/calc/mul/{x}/{y}` - Retorna a multiplicação de `x` por `y`, caso sejam informadas letras ao invés de números, será apresentada mensagem de erro
- `localhost:8080/calc/div/{x}/{y}` - Retorna a divisão de `x` por `y`, caso sejam informadas letras ao invés de números ou `{y}` seja igual à 0, será apresentada mensagem de erro
- `localhost:8080/calc/history` - Retorna o histórico das operações já realizadas. Caso nenhuma tenha sido efetuada ainda, será mostrada uma mensagem erro