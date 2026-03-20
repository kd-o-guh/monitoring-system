# 🚀 Monitoring System

Sistema de monitoramento distribuído desenvolvido em Go, capaz de coletar métricas reais da máquina (CPU, memória e disco) através de um agente e enviá-las para uma API REST com persistência em PostgreSQL.

---

## 📌 Sobre o projeto

Este projeto simula o funcionamento de ferramentas como sistemas de monitoramento (ex: Zabbix, Datadog), onde um agente coleta informações do sistema operacional e envia para um servidor central.

O backend processa e armazena essas métricas, permitindo consulta posterior via API.

---

## 🏗️ Arquitetura

O sistema é dividido em duas partes principais:

### 🔹 API (Backend)

* Desenvolvida em Go
* Arquitetura em camadas:

  * Handler
  * Service
  * Repository
* Integração com PostgreSQL
* Endpoints REST:

  * `POST /metrics` → recebe métricas
  * `GET /metrics` → lista métricas
  * `GET /health` → status da API

---

### 🔹 Agent (Coletor)

* Desenvolvido em Go
* Coleta métricas reais da máquina:

  * CPU
  * Memória
  * Disco
* Envio automático para a API
* Configurável via `.env`
* Implementa:

  * Retry
  * Timeout HTTP
  * Execução contínua

---

## 🐳 Como rodar com Docker

### 1. Subir o ambiente

```bash
docker-compose up --build
```

---

### 2. Acessar a API

```bash
http://localhost:8080/health
```

---

### 3. Testar envio de métricas

```bash
curl -X POST http://localhost:8080/metrics \
-H "Content-Type: application/json" \
-d '{"hostname":"server-01","cpu":50,"memory":70,"disk":80}'
```

---

## ⚙️ Configuração do Agent

Crie um arquivo `.env`:

```env
API_URL=http://localhost:8080/metrics
INTERVAL=5
```

Rodar o agent:

```bash
go run scripts/agent/main.go
```

---

## 🧰 Tecnologias utilizadas

* Go (Golang)
* PostgreSQL
* Docker & Docker Compose
* gopsutil (coleta de métricas do sistema)

---

## 💡 Diferenciais

* Coleta de métricas reais do sistema operacional
* Arquitetura limpa (separação de responsabilidades)
* Sistema distribuído (agent + API)
* Resiliência com retry e timeout
* Configuração via variáveis de ambiente
* Ambiente totalmente dockerizado

---

## 📈 Possíveis melhorias

* Dashboard (Grafana ou frontend próprio)
* Autenticação na API
* Paginação e filtros nas métricas
* Sistema de alertas
* Migrations automatizadas

---

## 👨‍💻 Autor

Desenvolvido por Gustavo Barbosa
