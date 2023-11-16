# Versão:Python 3.12.0
# Autor:Juliano Alessandro dos Santos
# Dia: 03/11/2023
# Biblioteca: Requests utilizada para fazer a requisição de um serviço na Web
# refinamento do Algoritimo criando classes e metodos
# FrameWork FastAPI, responsável por criar uma URL,localmente no meu computador, estou exibindo na
# web os meus dados

from fastapi import FastAPI, HTTPException
from Classes import CadastroUsuario as c

app = FastAPI()


@app.get("/weather/{duracao},{latitude},{longitude}")  # decoration
def Previsao(duracao: int, latitude: float, longitude: float):
    try:

        usuarios = c.CadastroUsuario()
        valores = usuarios.realizarPrevisao(duracao, latitude, longitude)
        return valores
    except Exception as e:
        return HTTPException(status_code=500, detail="Error %s" % str(e))


@app.get("/weather/{latitude},{longitude}")
def Indice(latitude: float, longitude: float):
    try:
        usuarios = c.CadastroUsuario()
        valores = usuarios.BuscarPrecipitacao(latitude, longitude)

        return valores

    except Exception as e:
        return HTTPException(status_code=500, detail="Error %s" % str(e))
