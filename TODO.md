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

### **5. Testes (Unitários e de Integração)**

- **Testes Unitários**:
  - [ ] Escrever testes unitários para funções de lógica de negócio.
  - [ ] Escrever testes para a validação de dados (ex: validação de entrada de orders, dados de customers).

- **Testes de Integração**:
  - [ ] Escrever testes de integração para os endpoints da API REST (testar criação de orders, relacionamento entre products e orders).
  
- **Testes de Banco de Dados**:
  - [ ] Escrever testes para garantir que as migrations e transações de banco de dados estão funcionando corretamente.
  
- **Cobertura de Testes**:
  - [ ] Garantir cobertura de testes de 80% ou mais.

### **6. CI/CD (GitHub Actions)**

- **Configuração de CI/CD**:
  - [ ] Criar pipeline de CI/CD usando GitHub Actions.
  - [ ] Configurar o pipeline para rodar testes unitários e de integração a cada commit.
  - [ ] Configurar o pipeline para realizar o deploy automático da aplicação (inicialmente em um ambiente de testes, depois produção).

### **7. Criar um Microsserviço para Gerenciar os Status dos Orders**

#### **Estrutura do Microsserviço**

1. **Definição do Propósito**:
   - O microsserviço será responsável por:
     - Gerenciar mudanças de status dos pedidos.
     - Emitir eventos baseados em alterações de status (ex.: enviar notificações).
     - Monitorar inconsistências ou falhas em transições de estado.

2. **Configuração Inicial**:
   - [ ] Criar um novo repositório para o microsserviço.
   - [ ] Configurar o ambiente do microsserviço com Go e dependências necessárias.
   - [ ] Configurar o banco de dados (PostgreSQL) para armazenar logs de transições de status.
   - [ ] Configurar um serviço de mensagens (ex.: RabbitMQ, Kafka) para comunicação assíncrona com outros microsserviços.

3. **Estrutura do Banco de Dados**:
   - [ ] Criar tabela `order_status_logs` com campos:
     - `id`: Chave primária.
     - `order_id`: Identificador do pedido.
     - `old_status`: Status anterior.
     - `new_status`: Novo status.
     - `changed_at`: Timestamp da mudança.
   - [ ] Adicionar triggers no banco de dados principal para enviar eventos em cada mudança de status.

4. **Endpoints do Microsserviço**:
   - [ ] `PUT /orders/{id}/status`:
     - Valida a transição de status com base em regras de negócio.
     - Atualiza o status do pedido.
     - Armazena logs de mudança no banco de dados.
   - [ ] `GET /orders/{id}/status`:
     - Retorna o status atual do pedido.
   - [ ] `GET /orders/logs`:
     - Retorna um histórico de transições de status para análise.

5. **Regras de Negócio**:
   - [ ] Validar as transições permitidas:
     - **Exemplo de transições válidas**:
       - `pending` → `send`.
       - `send` → `arrived`.
     - Transições inválidas devem retornar um erro apropriado.
   - [ ] Adicionar suporte para novos estados (ex.: `cancelled`).

6. **Eventos e Notificações**:
   - [ ] Emitir eventos assíncronos (ex.: `OrderStatusChanged`) via RabbitMQ ou Kafka.
   - [ ] Configurar consumidores para esses eventos em outros serviços (ex.: notificações por e-mail ou SMS).
   - [ ] Publicar mensagens para o Redis para integrar com serviços de cache.

7. **Testes**:
   - [ ] Testar as regras de validação de transições.
   - [ ] Testar a criação e recuperação de logs de status.
   - [ ] Testar a publicação de eventos em serviços de mensagens.

8. **Escalabilidade e Observabilidade**:
   - [ ] Configurar métricas e logs para monitorar o microsserviço (ex.: Prometheus, Grafana).
   - [ ] Configurar alertas para monitorar falhas em transições de status ou atrasos no processamento de eventos.

9. **CI/CD para o Microsserviço**:
   - [ ] Configurar pipelines independentes para o deploy e testes do microsserviço.
   - [ ] Implementar deploy automático em um cluster Kubernetes (se aplicável).
