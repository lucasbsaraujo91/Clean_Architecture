/* Para rodar a criação do banco de dados, tabela e inserção dos dados */

cd Clean_Architecture/docker-compose
E rodar o comando docker-compose up -d

/* Para colocar o servidor no ar */
Entrar na pasta Clean_Architecture/cmd/ordersystem
go run main.go wire_gen.go

/* Descrição do projeto */
Desafio: 
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL

Resultado: 
Criei os três retornos baseando-se em outros exercícios que realizamos
realizando páginação e dando a opção de ordenação (asc ou desc)
Nesse caso que fará a requisição deverá passar 3 parâmetros:
Page
Limit
Sort (ordenação da listagem)
 
/* Para realizar uma requisição Endpoint REST */
Entrar na pasta Clean_Architecture/api/
E rodar o arquivo findall_order.http

/* Para rodar o GraphQL */
No browser entrar na seguinte endereço: http://localhost:8082/

E aplicar o comando abaixo

query ListOrders {
  ListOrders(input: {Page: "1", Limit: 10, Sort: "desc"}){
    id
    Price, 
    Tax,
    FinalPrice
  }
}


/* Para gerar a estrutura do GRPC OBS: esse comando não precisa ser rodado */
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

/* Aplicar o Comando abaixo: obs ferramenta de linha de comando para interagir com servidores gRPC */
evans -r repl

/* Para simular a chamada ListOrders GRPC*/
show package
package pb
service OrderService
call ListOrders

Digite esses parametros por exemplo
page (TYPE_STRING) => 1
limit (TYPE_INT32) => 10
sort (TYPE_STRING) => desc