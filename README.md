# Projeto

## Descrição
Uma api com clean arch, destinada a pegar commits e repos do github, baseado no desenvolvimento utilizando os 12 fatores.

## Tecnologias
 - Golang
 - Docker
 - Open Telemetry
 - Jaeger
 - Prometheus
 - Swagger

## Solução adotada
Visto que o teste era pra mostrar minhas habilidades, utilizei da arquitetura limpa com alguns padrões de desenvolvimento do GO.

## Executando o projeto em ambiente de desenvolvimento com o Docker
### É necessário possuir o Docker instalado na sua máquina
Se ainda não possui clique aqui: https://www.docker.com

Para usuários Linux também é necessário instalar o Compose, veja a documentação aqui: https://docs.docker.com/compose/install/

### O projeto já está dockerizado, sendo assim basta utilizar os próximos comandos para inicializar a aplicação, junto com o open telemtry, Jaeger e o Prometheus

Para executar com o docker rode os seguintes comandos:
```
docker-compose up -d --build
```

**Após a execução o projeto poderá ser acessado por http://localhost:3000 (máquina local)**

**A documentação da API poderá ser acessada por http://localhost:3000/git-api/documentation/index.html (máquina local)**

**O jeager poderá ser acessado por http://localhost:16686/trace**

**O Prometheus poderá ser acessado por http://localhost:9090**

