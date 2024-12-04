## **Estrutura do Projeto**

### **1. Setup do Ambiente**
- [X] Configurar ambiente de desenvolvimento com Go.
- [X] Configurar PostgreSQL para o banco de dados.
- [X] Configurar Redis para cache.
- [??] Configurar ferramentas de testes unitários e integração.
- [X] Criar o projeto com a estrutura de diretórios para API REST.

### **2. Banco de Dados (PostgreSQL)**

- **Modelagem das Tabelas**:
  - [X] Criar a tabela `customers` com campos: `id`, `nome`, `email`  
  - [X] Criar a tabela `products` com campos: `id`, `nome`, `descricao`, `preco`.
  - [X] Criar a tabela `orders` com campos: `id`, `customer_id`, `data_order`, `status`.
  - [X] Criar a tabela `itens_order` com campos: `id`, `order_id`, `product_id`, `quantidade`.
  - [X] Criar a tabela `stock` com campos: `id`, `product_id`, `quantidade_disponivel`.
  
- **Normalização**:
  - [X] Garantir que as tabelas estejam normalizadas para evitar redundâncias.
  
- **Migrations**:
  - [X] Configurar migrations para versionar o esquema do banco de dados.
  - [X] Criar migrations para adicionar e modificar tabelas e colunas.

- **Índices de Banco de Dados**:
  - [X] Criar índices nas colunas mais consultadas, como `customer_id`, `product_id` e `order_id`.

- **Transações (ACID)**:
  - [X] Implementar transações para garantir que as operações de criação de orders sejam atômicas (incluir itens do order e atualizar stock).

  
### **3. API REST (Golang)**

- **Estrutura de Endpoints**:
  - [X] Criar endpoint `POST /customers` para cadastro de novos customers.
  - [ ] Criar endpoint `GET /customers/{id}` para visualizar detalhes de um customer.
  - [X] Criar endpoint `POST /products` para adicionar novos products.
  - [ ] Criar endpoint `GET /products/{id}` para visualizar detalhes de um product.
  - [X] Criar endpoint `POST /orders` para criar um order.
  - [ ] Criar endpoint `GET /orders/{id}` para visualizar detalhes de um order.
  - [ ] Criar endpoint `PUT /orders/{id}` para atualizar o status do order.
  - [X] Criar endpoint `POST /auth/login` para autenticação de usuários com JWT.

- **JWT para Autenticaço**:
  - [X] Implementar autenticação com JWT nas rotas protegidas.
  - [X] Adicionar middleware para verificar o token em rotas privadas.
  - [X] Sistemas de ROLE para usuarios

### **4. Redis para Caching**

- **Implementação de Cache**:
  - [ ] Implementar caching para products mais vendidos.
  - [ ] Implementar cache para detalhes de orders e customers para melhorar a performance.

- **Configuração do Redis**:
  - [X] Configurar o Redis no projeto e integrá-lo com a aplicação.
