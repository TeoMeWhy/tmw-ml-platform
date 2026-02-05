# %%

import pandas as pd
import requests
import datetime

df = pd.read_csv("abt_churn.csv")

# %%

values = df.sample(2).to_dict(orient="records")
payload = {"values": values}
payload

# %%

res = requests.post("http://localhost:5001/predict", json=payload)
res.json()
# %%

ids = df['idUsuario'].unique().tolist()

for i in ids:

        data = {"id": i,
                "model_name": "Churn TMW"}

        start = datetime.datetime.now().isoformat()

        res = requests.post("http://localhost:3000/predict", json=data)

        stop  = datetime.datetime.now().isoformat()
        time_duration = datetime.datetime.fromisoformat(stop) - datetime.datetime.fromisoformat(start)
        print(f"{res.status_code} - {time_duration} - {res.json()}")

# %%
df.columns