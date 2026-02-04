# %%

import pandas as pd
import requests

df = pd.read_csv("../data/abt_churn.csv")


values = df.sample(2).to_dict(orient="records")
payload = {"values": values}
payload

res = requests.post("http://localhost:5001/predict", json=payload)
res.json()
# %%


df.head(2).to_dict(orient="records")