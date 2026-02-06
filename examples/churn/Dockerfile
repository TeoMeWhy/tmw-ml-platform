FROM python:3.13

WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY 01_app.py .
COPY model.pkl .

CMD ["python", "01_app.py"]