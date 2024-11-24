# TODO.md - Gerenciamento de Pedidos para E-commerce

## Fase 1: Estrutura do Banco de Dados, Modelagem e CRUD Básico

### Tarefa 1: Modelagem do Banco de Dados
- **Descrição**: Criar o esquema básico do banco de dados com as tabelas `Clientes`, `Produtos`, `Pedidos`, `Itens_Pedido` e `Estoque`.
  - Definir os campos e tipos para cada tabela.
  - Garantir relacionamentos entre as tabelas: `Clientes` para `Pedidos`, `Pedidos` para `Itens_Pedido`, `Produtos` para `Itens_Pedido`.
  - Normalizar o banco de dados (1NF, 2NF, 3NF).
- **Prazo**: 3 dias.
  
### Tarefa 2: Implementar o CRUD para Produtos e Clientes
- **Descrição**: Criar as funcionalidades para cadastro, leitura, atualização e exclusão (CRUD) de `Produtos` e `Clientes` no sistema.
  - Criar APIs RESTful para manipulação de dados.
  - Validar entradas com regras simples (ex: nome do produto, e-mail do cliente).
- **Prazo**: 3 dias.

### Tarefa 3: Implementar o CRUD para Pedidos e Itens de Pedido
- **Descrição**: Criar as funcionalidades para criação de pedidos e seus itens. O pedido pode incluir múltiplos produtos.
  - Criar endpoints para adicionar itens a um pedido.
  - Validar a quantidade disponível em estoque para cada item.
- **Prazo**: 4 dias.

## Fase 2: Transações, Normalização e Resolução de Problemas de Performance

### Tarefa 4: Implementar Transações (ACID) em Pedidos
- **Descrição**: Implementar transações para garantir que a criação de um pedido e seus itens sejam atômicas. 
  - Usar transações para garantir que se um item falhar, o pedido inteiro seja revertido.
  - Testar a integridade dos dados ao falhar em algum ponto da transação.
- **Prazo**: 3 dias.

### Tarefa 5: Resolver Problema N+1 no Carregamento de Pedidos e Itens
- **Descrição**: Otimizar o carregamento de pedidos e seus itens para evitar o problema N+1.
  - Implementar **eager loading** ou **joins** para buscar os pedidos com seus itens em uma única consulta.
  - Validar que o número de consultas SQL foi reduzido.
- **Prazo**: 4 dias.

### Tarefa 6: Normalização do Banco de Dados
- **Descrição**: Garantir que o banco de dados esteja corretamente normalizado.
  - Verificar redundâncias e aplicar a normalização (1NF, 2NF, 3NF).
  - Criar novas tabelas se necessário para remover dados duplicados.
- **Prazo**: 3 dias.

## Fase 3: Migrations, Indexação e Profiling

### Tarefa 7: Criar Migrations e Gerenciar Alterações no Banco de Dados
- **Descrição**: Criar migrations para versionar o banco de dados e gerenciar alterações futuras.
  - Criar uma migration inicial para definir as tabelas.
  - Adicionar novas migrations conforme a evolução do sistema (ex: nova tabela ou coluna).
- **Prazo**: 3 dias.

### Tarefa 8: Implementar Índices para Otimização de Consultas
- **Descrição**: Adicionar índices nas colunas mais utilizadas nas consultas (ex: `cliente_id`, `produto_id`, `pedido_id`).
  - Analisar as consultas mais frequentes e criar índices para melhorar a performance.
  - Validar se as consultas estão mais rápidas após a adição dos índices.
- **Prazo**: 4 dias.

### Tarefa 9: Profiling de Performance e Análise de Consultas Lentas
- **Descrição**: Realizar profiling das consultas SQL para encontrar e otimizar consultas lentas.
  - Usar ferramentas de profiling como EXPLAIN ou pg_stat_statements.
  - Analisar o tempo de execução e ajustar as consultas para melhorar a performance.
- **Prazo**: 3 dias.

## Fase 4: Testes e Ajustes Finais

### Tarefa 10: Criar Testes de Unidade para CRUD e Lógica de Pedidos
- **Descrição**: Escrever testes de unidade para as funcionalidades de CRUD e lógica de criação de pedidos.
  - Garantir que todas as operações estão funcionando como esperado.
  - Criar testes para transações e para verificar a integridade dos dados.
- **Prazo**: 4 dias.

### Tarefa 11: Testar Integração com Banco de Dados (Testes de Performance)
- **Descrição**: Testar a performance do sistema integrando as consultas com o banco de dados real.
  - Validar se o tempo de resposta das consultas está dentro de um padrão aceitável.
  - Realizar testes de carga para verificar como o sistema se comporta com grandes volumes de dados.
- **Prazo**: 4 dias.

### Tarefa 12: Revisão de Código e Ajustes Finais
- **Descrição**: Revisar o código, corrigir bugs e melhorar a legibilidade.
  - Refatorar o código se necessário para garantir aderência aos princípios de SOLID e boas práticas.
- **Prazo**: 2 dias.
