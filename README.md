### Requisitos:

- [Jenkins](https://www.jenkins.io/download/)
- [Packer](https://www.packer.io/downloads)
- [Docker](https://www.docker.com/get-started)

### Preparação:

1. Abra o terminal e utilize o comando `chmod 666 /var/run/docker.sock`
2. Na sequência abra o dashboard do Jenkins
3. Iremos criar o primeiro job:
   1. Vá em `New Item`, de o nome de `bake`, selecione `Pipeline` e clique em ok
   2. Selecione a aba `Pipeline`, em `Definition` altere para a opção `Pipeline script from SCM`
   3. Em `SCM` altere para `Git`, no campo `Repository URL` coloque `https://github.com/gustavoboliveira/devops-ilegra-jenkins` e em `Credentials` insira seu login e senha do Github
   4. Desça a página até encontrar o campo `Script Path` e insira `bake/Jenkinsfile`, depois clique em `Save`
4. Iremos criar o segundo job:
   1. Vá em `New Item`, de o nome de `launch`, selecione `Pipeline` e clique em ok 
   2. Selecione a aba `Pipeline`, em `Definition` altere para a opção `Pipeline script from SCM`
   3. Em `SCM` altere para `Git`, no campo `Repository URL` coloque `https://github.com/gustavoboliveira/devops-ilegra-jenkins` e em `Credentials` insira seu login e senha do Github
   4. Desça a página até encontrar o campo `Script Path` e insira `launch/Jenkinsfile`, depois clique em `Save`

### Execução:

1. No dashboard do Jenkins, localize o job `bake`, clique nele e em `Build Now`, espere concluir o processo
2. Voltando ao dashboard do Jenkins, localize o job `launch`, clique nele e em `Build Now`, espere concluir o processo
3. Aplicação deverá estar executando no endereço `localhost:8888/`
4. Para encerrar a aplicação, abra um terminal e use o comando `docker ps` e copie o `CONTAINER ID` referente a
   calculadora
5. Então utilize o comando `docker stop CONTAINER_ID`

### Endpoints:

- `localhost:8888/calc/sum/{x}/{y}` - Retorna a soma de `x` e `y`, caso sejam informadas letras ao invés de números,
  será apresentada mensagem de erro
- `localhost:8888/calc/sub/{x}/{y}` - Retorna a subtração de `x` e `y`, caso sejam informadas letras ao invés de
  números, será apresentada mensagem de erro
- `localhost:8888/calc/mul/{x}/{y}` - Retorna a multiplicação de `x` por `y`, caso sejam informadas letras ao invés de
  números, será apresentada mensagem de erro
- `localhost:8888/calc/div/{x}/{y}` - Retorna a divisão de `x` por `y`, caso sejam informadas letras ao invés de números
  ou `{y}` seja igual à 0, será apresentada mensagem de erro
- `localhost:8888/calc/history` - Retorna o histórico das operações já realizadas. Caso nenhuma tenha sido efetuada
  ainda, será mostrada uma mensagem erro