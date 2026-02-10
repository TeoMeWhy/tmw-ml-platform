# %%
print("Script do modelo sendo executado...")

import pandas as pd
import mlflow

from sklearn import tree

import sqlalchemy

mlflow.set_tracking_uri("http://localhost:5000")
mlflow.set_experiment(experiment_name="churn_exp")

print("Lendo os dados...")
df = pd.read_csv("abt_churn.csv")
df["id"] = df["idUsuario"]
del df['idUsuario']

# %%
print("Preparando os dados e o modelo...")

clf = tree.DecisionTreeClassifier(min_samples_leaf=20, random_state=42)
target  = "flagChurn"
features = df.columns[2:-1]
X = df[features]
y = df[target]

features_metadata = {"feature_metadata":{"source":"fs_tmw_users", "features": X.dtypes.astype(str).to_dict()}}

print("Iniciando o experimento no MLflow...")
with mlflow.start_run() as run:
    mlflow.sklearn.autolog()

    print("    Treinando o modelo...")
    clf.fit(X, y)

    print("    Registrando os modelos...")
    mlflow.register_model(model_uri=f"runs:/{run.info.run_id}/model",
                          name="Churn TMW")


    mlflow.log_params(features_metadata)



# %%
client = mlflow.client.MlflowClient()

client.set_registered_model_tag(name="Churn TMW", key="feature_store", value="tmw_users")
client.set_registered_model_tag(name="Churn TMW", key="uri", value="http://ml_churn_app:5001/predict")

client.create_registered_model(name="Teste 01")


client.create_registered_model(name="Teste 02", tags={"uri": "http://ml_churn_app:5001/predict"})

client.create_registered_model(name="Teste 03", tags ={"uri":"http://ml_churn_app:5001/predict", "feature_store":"not_exists"})

client.create_registered_model(name="Teste 04", tags={"feature_store": "tmw_users", "uri": "http://ml_churn_app:5002/predict"})

print("Gravando a feature store...")
df = df.sort_values("dtRef", ascending=False).drop_duplicates(subset=["id"])
con = sqlalchemy.create_engine("mysql+pymysql://root:rootpassword@localhost:3306/feature_store")
df.to_sql("tmw_users", con, if_exists="replace", index=False)
# %%
