# %%

import pandas as pd
import mlflow

from sklearn import tree
import json

mlflow.set_tracking_uri("http://localhost:5000")
mlflow.set_experiment(experiment_name="churn_exp")

# 
# %%

df = pd.read_csv("../data/abt_churn.csv")
df["id"] = df["idUsuario"]
del df['idUsuario']

# %%
clf = tree.DecisionTreeClassifier(min_samples_leaf=20)

target  = "flagChurn"
features = df.columns[2:-1]
X = df[features]
y = df[target]

features_metadata = {"feature_metadata":{"source":"fs_tmw_users", "features": X.dtypes.astype(str).to_dict()}}

features_metadata

# %%
with mlflow.start_run() as run:
    mlflow.sklearn.autolog()
    clf.fit(X, y)
    mlflow.register_model(model_uri=f"runs:/{run.info.run_id}/model",
                          name="Churn TMW")

    mlflow.log_params(features_metadata)


# %%

df = df.sort_values("dtRef", ascending=False).drop_duplicates(subset=["id"])
import sqlalchemy
con = sqlalchemy.create_engine("mysql+pymysql://root:rootpassword@localhost:3306/feature_store")
df.to_sql("tmw_users", con, if_exists="replace", index=False)
# %%
