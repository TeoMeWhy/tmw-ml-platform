

## Objetivo

Ter um serviço que centraliza todas as chamadas de modelos.

- Conhecer os metadados dos modelos:
    - url
    - payload / features

- Acesso à feature store
- Acesso ao MLFlow


## Fluxo

1. Recebe requisição do cliente;
2. Identifica qual o modelo e entidade a ser consultada;
3. Busca Features na Feature Store;
4. Realiza requisição no URI do modelo;
5. Devolve para ao cliente o payload das predições;

## Funcionalidades adicionais

- Autodeploy de modelos