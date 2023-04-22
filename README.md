## FCLX - Full Cycle Learning Experience
Abril 2023

#### Agenda
- Entender o Projeto Prático
- Tecnologias que serão utilizadas
- Docker
- Inicio do Microsserviço

<br /><br />

# Projeto Prático
Desenvolveremos duas formas / interfaces para utilizarmos o ChatGPT
- Interface Web;
- WhatsApp.

ver: ./Dinâmica-Projeto.png


# TOOL

SQLC -> como orm
golang-migrate -> para gerar as migrations
migrate create -ext=mysql -dir=sql/migrations -seq init 