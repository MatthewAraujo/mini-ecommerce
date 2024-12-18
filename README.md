# API REST com Go, PostgreSQL e Redis

Este é um projeto de API REST desenvolvido com Golang, PostgreSQL e Redis. Ele foi criado para demonstrar a integração de uma aplicação simples de e-commerce, com funcionalidade para gerenciar clientes, produtos, pedidos e autenticação, além de otimizações de performance com cache utilizando Redis.

## **Estrutura do Projeto**

### **1. Setup do Ambiente**
- **Go**: Ambiente de desenvolvimento configurado para trabalhar com Go.
- **PostgreSQL**: Banco de dados configurado para armazenar as informações da aplicação, como clientes, produtos, pedidos e estoque.
- **Redis**: Configuração para caching, melhorando a performance da aplicação.
- **Ferramentas de Testes**: A configuração para testes unitários e integração está em andamento.

### **2. Banco de Dados (PostgreSQL)**

A modelagem do banco de dados segue um esquema normalizado para evitar redundâncias e garantir consistência dos dados.

#### **Tabelas Criadas**
- `customers`: Armazena informações dos clientes (id, nome, email).
- `products`: Armazena informações dos produtos (id, nome, descrição, preço).
- `orders`: Armazena informações dos pedidos (id, customer_id, data_order, status).
- `itens_order`: Relaciona os produtos aos pedidos (id, order_id, product_id, quantidade).
- `stock`: Controla a quantidade de produtos disponíveis (id, product_id, quantidade_disponivel).

#### **Migrations**
- Migrations foram configuradas para versionar o esquema do banco de dados, facilitando alterações futuras e garantindo a consistência do modelo.
- Índices foram criados nas colunas mais consultadas, como `customer_id`, `product_id` e `order_id`, para melhorar a performance de consulta.

#### **Transações**
- Transações são utilizadas para garantir que a criação de pedidos seja atômica, incluindo a inserção de itens e a atualização do estoque.

### **3. API REST (Golang)**

A API foi desenvolvida com endpoints para gerenciar clientes, produtos, pedidos e autenticação de usuários.

#### **Endpoints Implementados**
- `POST /customers`: Cadastro de novos clientes.
- `POST /products`: Adicionar novos produtos.
- `POST /orders`: Criar um pedido.
- `POST /auth/login`: Autenticação de usuários com JWT.

#### **Endpoints Pendentes**
- `GET /customers/{id}`: Visualizar detalhes de um cliente.
- `GET /products`: Listar todos os produtos (com paginação).
- `GET /products/{id}`: Visualizar detalhes de um produto.
- `GET /orders/{id}`: Visualizar detalhes de um pedido.
- `PUT /orders/{id}`: Atualizar o status de um pedido.

#### **Autenticação com JWT**
- A autenticação foi implementada utilizando JWT para proteger as rotas privadas.
- Middleware foi adicionado para verificar o token em rotas protegidas.
- Sistema de roles para controlar o acesso de usuários.

### **4. Redis para Caching**

A implementação de cache foi planejada para melhorar a performance da aplicação, especialmente para produtos mais vendidos e detalhes de pedidos e clientes.

#### **Pendências**
- Implementação de cache para produtos mais vendidos.
- Implementação de cache para detalhes de pedidos e clientes.

#### **Configuração do Redis**
- O Redis foi configurado e integrado ao projeto para armazenar informações em cache, reduzindo o tempo de resposta de requisições frequentes.

## **Como Rodar o Projeto**

1. **Instalar dependências**:
   - Go
   - PostgreSQL
   - Redis

2. **Configurar Banco de Dados**:
   - Execute as migrations para criar as tabelas no PostgreSQL.

3. **Configurar Redis**:
   - Configure o Redis conforme necessário.

4. **Iniciar o Servidor**:
   - Execute o comando para rodar a aplicação em Go.

## **Melhorias Futuras**
- Implementação completa de caching para melhorar ainda mais a performance.
- Finalização dos endpoints pendentes.
- Adição de mais funcionalidades de gerenciamento de estoque e pedidos.

---

Esse projeto tem como objetivo proporcionar uma base sólida para criar APIs REST escaláveis com Go, PostgreSQL e Redis, além de implementar boas práticas como autenticação JWT e otimização de performance com caching.
