## **Estrutura do Projeto**

### **1. Setup do Ambiente**

- [x] Configurar ambiente de desenvolvimento com Go.
- [x] Configurar PostgreSQL para o banco de dados.
- [x] Configurar Redis para cache.
- [??] Configurar ferramentas de testes unitários e integração.
- [x] Criar o projeto com a estrutura de diretórios para API REST.

### **2. Banco de Dados (PostgreSQL)**

- **Modelagem das Tabelas**:
  - [x] Criar a tabela `customers` com campos: `id`, `nome`, `email`
  - [x] Criar a tabela `products` com campos: `id`, `nome`, `descricao`, `preco`.
  - [x] Criar a tabela `orders` com campos: `id`, `customer_id`, `data_order`, `status`.
  - [x] Criar a tabela `itens_order` com campos: `id`, `order_id`, `product_id`, `quantidade`.
  - [x] Criar a tabela `stock` com campos: `id`, `product_id`, `quantidade_disponivel`.
- **Normalização**:
  - [x] Garantir que as tabelas estejam normalizadas para evitar redundâncias.
- **Migrations**:

  - [x] Configurar migrations para versionar o esquema do banco de dados.
  - [x] Criar migrations para adicionar e modificar tabelas e colunas.

- **Índices de Banco de Dados**:

  - [x] Criar índices nas colunas mais consultadas, como `customer_id`, `product_id` e `order_id`.

- **Transações (ACID)**:
  - [x] Implementar transações para garantir que as operações de criação de orders sejam atômicas (incluir itens do order e atualizar stock).

### **3. API REST (Golang)**

- **Estrutura de Endpoints**:

  - [x] Criar endpoint `POST /customers` para cadastro de novos customers.
  - [ ] Criar endpoint `GET /customers/{id}` para visualizar detalhes de um customer.
  - [x] Criar endpoint `POST /products` para adicionar novos products.
  - [x] Criar endpoint `GET /products` para pegar todos os produtos
    - [x] Criar um sistema de paginação
  - [ ] Criar endpoint `GET /products/{id}` para visualizar detalhes de um product.
  - [x] Criar endpoint `POST /orders` para criar um order.
  - [ ] Criar endpoint `GET /orders/{id}` para visualizar detalhes de um order.
  - [ ] Criar endpoint `PUT /orders/{id}` para atualizar o status do order.
  - [x] Criar endpoint `POST /auth/login` para autenticação de usuários com JWT.

- **JWT para Autenticaço**:
  - [x] Implementar autenticação com JWT nas rotas protegidas.
  - [x] Adicionar middleware para verificar o token em rotas privadas.
  - [x] Sistemas de ROLE para usuarios

### **4. Redis para Caching**

- **Implementação de Cache**:

  - [x] Implementar caching para products mais vendidos.
  - [ ] Implementar cache para detalhes de orders e customers para melhorar a performance.

- **Configuração do Redis**:
  - [x] Configurar o Redis no projeto e integrá-lo com a aplicação.
