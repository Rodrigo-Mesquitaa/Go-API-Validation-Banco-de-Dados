# Go-API-Validation
API Crud realizada em Golang para estudo.


Blockcoin - Go Api Restful - Para Estudo
=========================================

Endpoints:

    GET/  http://localhost:3000/
    |
    |_ Home page

    POST/ http://localhost:3000/api/users
    |
    |_ Recebe um json com nickname, email e senha.

    POST/ http://localhost:3000/ap�/login
    |
    |_ Recebe um json com email e senha para autentica��o
    |_ Retorna um JWT

    GET/ http://localhost:3000/api/users
    |
    |_ Retorna um json com uma lista de usu�rios

    POST/ http://localhost:3000/api/users
    |
    |_ Recebe um json com nickname, email e senha
    |_ Precisa de Autentica��o

    GET/ http://localhost:3000/api/users/1
    |
    |_ Retorna um usu�rio buscando pelo Id
    |_ Precisa de Autentica��o

    PUT/ http://localhost:3000/api/users/1
    |
    |_ Recebe um json com os dados que ser�o atualizados do usu�rio que possui o Id informado na Url
    |_ Retorna a quantidade de linhas atualizadas
    |_ Precisa de Autentica��o

    DELETE/ http://localhost:3000/api/users/1
    |
    |_ Deleta um usu�rio informando Id pel  'a Url
    |_ Retorna quantidade de linhas afetadas
    |_ Precisa de Autentica��o

    GET/ http://localhost:3000/api/wallets
    |
    |_ Retorna uma lista de catteries

    GET/ http://localhost:3000/api/wallets/public_key
    |
    |_ Retorna uma carteira informando a chave publica na Url

    PUT/ http://localhost:3000/api/wallets/public_key
    |
    |_ Atualiza o saldo de uma carteira informando a chave publica na Url

    PUT/ http://localhost:3000/api/wallets/add/public_key
    |
    |_ Adiciona um valor via json ao saldo da carteira informando a chave publica na Url
    |_ Precisa de autentica��o

    GET/ http://localhost:3000/api/transactions
    |
    |_ Retorna um json com uma lista de transa��es

    POST/ http://localhost:3000/api/transactions/public_key
    |
    |_ Recebe um json com a chave p�blica e um valor de uma carteira que ser�
    |  trasferido para a carteira cuja chave p�blica e passada pela url
    |_ Precisa de autentica��o