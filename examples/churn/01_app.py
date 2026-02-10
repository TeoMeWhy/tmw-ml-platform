import time

time.sleep(20)  # Aguarda o MLflow Server iniciar

import json
import flask
app = flask.Flask(__name__)

import dotenv
import os

dotenv.load_dotenv()

MLFLOW_URI = os.getenv("MLFLOW_URI")

# %%
import mlflow
import pandas as pd

mlflow.set_tracking_uri(MLFLOW_URI)
MODEL = mlflow.sklearn.load_model("models:/Churn TMW/1")


@app.route("/")
def hello():
    return "Hello, World!"


@app.route("/predict", methods=["POST"])
def predict():

    data = flask.request.json
    values = pd.DataFrame(data["values"])

    X = values[MODEL.feature_names_in_]
    y_pred = MODEL.predict_proba(X)

    df_pred = pd.DataFrame(y_pred, columns=MODEL.classes_)
    df_pred["id"] = values["id"].copy()
    df_pred = (df_pred.set_index("id")
                      .to_dict(orient="index"))

    payload = json.loads(json.dumps({"predictions": df_pred}))
    return payload


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)