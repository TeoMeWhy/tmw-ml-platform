import flask
app = flask.Flask(__name__)

# %%
import mlflow
import pandas as pd

mlflow.set_tracking_uri("http://localhost:5000")
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

    payload = (pd.DataFrame(y_pred, columns=MODEL.classes_)
                 .to_dict(orient="records"))

    return payload


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5001)