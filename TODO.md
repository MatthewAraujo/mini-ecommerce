## **Estrutura do Projeto**

### **1. Setup do Ambiente**
- [X] Configurar ambiente de desenvolvimento com Go.
- [X] Configurar PostgreSQL para o banco de dados.
- [X] Configurar Redis para cache.
- [ ] Configurar ferramentas de testes unitários e integração.
- [ ] Criar o projeto com a estrutura de diretórios para API REST.

### **2. Banco de Dados (PostgreSQL)**

- **Modelagem das Tabelas**:
  - [ ] Criar a tabela `clientes` com campos: `id`, `nome`, `email`, `endereco`.
  - [ ] Criar a tabela `produtos` com campos: `id`, `nome`, `descricao`, `preco`.
  - [ ] Criar a tabela `pedidos` com campos: `id`, `cliente_id`, `data_pedido`, `status`.
  - [ ] Criar a tabela `itens_pedido` com campos: `id`, `pedido_id`, `produto_id`, `quantidade`.
  - [ ] Criar a tabela `estoque` com campos: `id`, `produto_id`, `quantidade_disponivel`.
  
- **Normalização**:
  - [ ] Garantir que as tabelas estejam normalizadas para evitar redundâncias.
  
- **Migrations**:
  - [ ] Configurar migrations para versionar o esquema do banco de dados.
  - [ ] Criar migrations para adicionar e modificar tabelas e colunas.

- **Índices de Banco de Dados**:
  - [ ] Criar índices nas colunas mais consultadas, como `cliente_id`, `produto_id` e `pedido_id`.

- **Transações (ACID)**:
  - [ ] Implementar transações para garantir que as operações de criação de pedidos sejam atômicas (incluir itens do pedido e atualizar estoque).

- **Performance de Consultas**:
  - [ ] Realizar profiling das consultas e otimizar com índices ou ajustes.
  
### **3. API REST (Golang)**

- **Estrutura de Endpoints**:
  - [ ] Criar endpoint `POST /clientes` para cadastro de novos clientes.
  - [ ] Criar endpoint `GET /clientes/{id}` para visualizar detalhes de um cliente.
  - [ ] Criar endpoint `POST /produtos` para adicionar novos produtos.
  - [ ] Criar endpoint `GET /produtos/{id}` para visualizar detalhes de um produto.
  - [ ] Criar endpoint `POST /pedidos` para criar um pedido.
  - [ ] Criar endpoint `GET /pedidos/{id}` para visualizar detalhes de um pedido.
  - [ ] Criar endpoint `PUT /pedidos/{id}` para atualizar o status do pedido.
  - [ ] Criar endpoint `POST /auth/login` para autenticação de usuários com JWT.

- **JWT para Autenticação**:
  - [ ] Implementar autenticação com JWT nas rotas protegidas.
  - [ ] Adicionar middleware para verificar o token em rotas privadas.

### **4. Redis para Caching**

- **Implementação de Cache**:
  - [ ] Implementar caching para produtos mais vendidos.
  - [ ] Implementar cache para detalhes de pedidos e clientes para melhorar a performance.

- **Configuração do Redis**:
  - [ ] Configurar o Redis no projeto e integrá-lo com a aplicação.

### **5. Testes (Unitários e de Integração)**

- **Testes Unitários**:
  - [ ] Escrever testes unitários para funções de lógica de negócio.
  - [ ] Escrever testes para a validação de dados (ex: validação de entrada de pedidos, dados de clientes).

- **Testes de Integração**:
  - [ ] Escrever testes de integração para os endpoints da API REST (testar criação de pedidos, relacionamento entre produtos e pedidos).
  
- **Testes de Banco de Dados**:
  - [ ] Escrever testes para garantir que as migrations e transações de banco de dados estão funcionando corretamente.
  
- **Cobertura de Testes**:
  - [ ] Garantir cobertura de testes de 80% ou mais.

### **6. CI/CD (GitHub Actions)**

- **Configuração de CI/CD**:
  - [ ] Criar pipeline de CI/CD usando GitHub Actions.
  - [ ] Configurar o pipeline para rodar testes unitários e de integração a cada commit.
  - [ ] Configurar o pipeline para realizar o deploy automático da aplicação (inicialmente em um ambiente de testes, depois produção).

### **7. Performance e Otimização**

- **Profiling de Performance**:
  - [ ] Usar ferramentas como `pprof` para análise de performance da aplicação.
  - [ ] Monitorar e otimizar endpoints de maior carga.

- **Cache e Indices**:
  - [ ] Testar e otimizar as estratégias de cache no Redis.
  - [ ] Analisar e melhorar as consultas SQL usando `EXPLAIN`.
 
- **README**:
  - [ ] Escrever o README do projeto com detalhes sobre a configuração, execução e funcionamento do sistema.
  
