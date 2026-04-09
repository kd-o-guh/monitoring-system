# Monitoring System

Sistema de monitoramento composto por um agente coletor e uma API backend para armazenamento e consulta de métricas de sistema.

## 📌 Descrição

Este projeto consiste em um sistema distribuído para coleta e exposição de métricas de infraestrutura.

O sistema é composto por dois componentes principais:

- **Agente coletor** responsável por capturar métricas da máquina
- **API backend** responsável por receber, armazenar e disponibilizar as métricas coletadas

O agente coleta métricas como:

- Uso de CPU
- Uso de memória
- Uso de disco

A arquitetura permite a adição de novas métricas futuramente, como:

- GPU Usage
- Network Usage
- Temperature Metrics

Os dados coletados são armazenados em um banco de dados PostgreSQL e podem ser consultados via API.

---

## 🧱 Arquitetura

Agente → API → PostgreSQL → Consulta via Endpoint

O agente envia periodicamente os dados coletados para a API, que armazena as informações no banco de dados PostgreSQL.

---

## 🚀 Tecnologias utilizadas

- Go (Golang)
- REST API
- PostgreSQL
- Docker
- Docker Compose
  
---

## 📡 Endpoints disponíveis

### `GET /metrics`

Retorna as métricas coletadas pelo sistema.

Exemplo de resposta:

```json
{
  "cpu": "35%",
  "memory": "62%",
  "disk": "48%"
}
