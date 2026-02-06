FROM python:3.13

RUN pip install mlflow[extras]==3.9.0

EXPOSE 5000

WORKDIR /app

CMD ["mlflow", "server", "--host", "0.0.0.0", "--port", "5000", "--backend-store-uri", "sqlite:///mlflow.db", "--default-artifact-root", "file:/app/artifacts", "--allowed-hosts", "*"]