import os
import requests as r

TOKEN = os.environ.get("WEATHER_TOKEN")

class CadastroUsuario:
    def BuscarPrecipitacao(self, latitude, longitude):
        url = f"https://atlas.microsoft.com/weather/currentConditions/json?api-version=1.0&query={latitude},{longitude}&subscription-key={TOKEN}"
        valores = r.get(url)
        formatoDicionario = valores.json()
        resumoPrecipitação = formatoDicionario["results"][0]["precipitationSummary"]
        return resumoPrecipitação

    def realizarPrevisao(self, duracao: int, latitude: float, longitude: float):
        url = f"https://atlas.microsoft.com/weather/forecast/daily/json?api-version=1.0&query={latitude},{longitude}&duration={duracao}&subscription-key={TOKEN}"
        valores = r.get(url)
        formatoDicionario = valores.json()
        valores = range(0, duracao)
        for valor in valores:
            ResumoPrecipitacao = formatoDicionario["forecasts"][valor]["day"][
                "totalLiquid"
            ]
            DiaPrecipitacao = formatoDicionario["forecasts"][valor]["date"]
        return ResumoPrecipitacao, DiaPrecipitacao
