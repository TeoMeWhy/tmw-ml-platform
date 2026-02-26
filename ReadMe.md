# Palantir

Um hub para consultar modelos de ML utilizando Feature Store e MLFlow.

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

## Setup e requisitos

### MLFlow

Um dos requisitos para o funcionamento da Palantir, é ter o MLFlow em seu ambiente. Com isso, defina a variável ambiente do seu sistema:

```bash
MLFLOW_URI = "http://host:port"
```

Além disso, o modelo de Machine Learning ou LLM deve estar registrado no MLFlow com as seguintes tags obrigatórias:

```bash
feature_store = "feature_store_table"
uri = "http://host:port/route"
```

Em que, `feature_store` se refere à tabela que contém as features/variáveis para o modelo registrado realizar inferência. E `uri` é o endereço de onde a API do modelo está disponível.

### MySQL

No momento, o driver utilizado para o uso da Palantir é o MySQL. É necessário adicionar a seguinte variável ambiente:

```bash
MYSQL_DSN = "user:pass@tcp(host:port)/dbname"
```

### API funcional

O modelo a ser utilizado na inferência deve estar disponível para ser acessado, respeitando os seguintes schema de dados:

#### Body da Chamada

```json
{
    "values": [
        {
            "id_1": "id1",
            "feature_1": "v1_1",
            ...
            "feature_n": "v1_n"
        },
        ...
        {
            "id_n": "idn",
            "feature_1": "vn_1",
            ...
            "feature_n": "vn_n"
        },
    ]
}
```

Em que `values` é uma lista de entidades, e cada entidade é representada por um par chave/valor (nome da feature / valor da feature).


#### Body da Resposta

```json
{
    "predictions": {
        "id_1": {"label_1":"v1", ..., "label_n":"vn"}
        ...
        "id_n": {"label_1":"v1",..., "label_n":"vn"}

    }
}

```

## Utilização

Para realizar a predição de uma entidade, basta fazer uma requisição `POST` na rota `/predict` dos serviço da Palantir, passando em seu Body a seguinte estrutura:

```json
{
    "model_name": "nome_modelo",
    "id": "id_entidade"
}
```

Com isso, a Palantir será responsável por identificar o endereço do Modelo e as Features necessárias. É importante que o ID da entidade solicitada esteja presenta na Feature Store correspodente.

Como retorno dessa requisição, temos a seguinte assinatura:

```json
{
    "predictions": {
        "id_1": {"label_1":"v1", ..., "label_n":"vn"}
        ...
        "id_n": {"label_1":"v1",..., "label_n":"vn"}
    }
}
```