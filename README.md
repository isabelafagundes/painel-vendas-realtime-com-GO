# Painel de Vendas em Tempo Real (Go + Echo)

Aplicação backend em Go que simula um **painel de vendas em tempo real**, expondo:

- API REST (JSON) para:
  - Listar lojas
  - Obter o estado atual do painel de vendas
- WebSocket para:
  - Enviar o **estado atual** do painel ao conectar
  - Enviar **novas vendas** em tempo real

Arquitetura baseada em **MVC**, usando o framework HTTP **Echo**.

---

## 📦 Tecnologias

- **Linguagem:** Go
- **Framework HTTP:** [Echo](https://echo.labstack.com/)
- **WebSocket:** `gorilla/websocket` (ou nativo, a depender da implementação)
- **Arquitetura:** MVC (Model, View/Controller, Service)
- **Estado interno:** Em memória, com simulador de vendas (`SimuladorVendas`)

---

## 🗂 Estrutura de pastas (proposta)

```txt
.
├── cmd
│   └── server
│       └── main.go           # Inicialização do servidor e injeção de dependências
├── internal
│   ├── config
│   │   └── config.go         # Carregamento de variáveis de ambiente / configuração
│   ├── infra
│   │   └── log
│       └── log.go            # Logger estruturado (Info, Error, etc.)
│   ├── model                 # M (Model) – domínio
│   │   ├── loja.go           # struct Loja
│   │   ├── terminal.go       # struct Terminal
│   │   ├── venda.go          # struct Venda
│   │   ├── painel.go         # struct EstadoPainelVendas / ResumoLoja
│   │   └── ws.go             # struct MensagemWS (protocolo WebSocket)
│   ├── service               # Regras de negócio
│   │   ├── simulador_vendas.go  # SimuladorVendas – gera vendas em memória
│   │   ├── painel_service.go    # PainelService – monta EstadoPainelVendas
│   │   └── ws_hub.go            # HubPainel / Client – gestão de clientes WebSocket
│   └── controller            # C (Controller) – Echo handlers
│       ├── saude_controller.go   # /saude
│       ├── loja_controller.go    # /v1/lojas
│       ├── painel_controller.go  # /v1/painel/snapshot
│       └── painel_ws_controller.go # /v1/ws/painel
└── go.mod
