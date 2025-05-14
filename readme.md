# Gerenciador de Sites - Backend em Go

Uma API RESTful para gerenciamento de sites, módulos e componentes, desenvolvida em Go, seguindo Clean Architecture e Domain-Driven Design (DDD).

## 📋 Descrição

Este projeto permite:
- Cadastro e autenticação de usuários
- Cadastro e organização de sites
- Gerenciamento de módulos de cada site (ex: Home, Sobre, Contato)
- Gerenciamento de componentes dinâmicos de cada módulo
- Configuração de settings para cada componente

Ideal para servir como backend de um construtor de sites dinâmico, onde o frontend pode montar páginas a partir dos dados estruturados da API.

## 🛠️ Tecnologias
- **Go 1.24+**
- **Gin** (framework web)
- **GORM** (ORM)
- **MySQL** (ou PostgreSQL, adaptável)
- **JWT** (autenticação)
- **Docker** (opcional)

## 🏗️ Estrutura do Projeto

```
backend/
├── application/
│   ├── dtos/                   # Data Transfer Objects
│   ├── interfaces/
│   │   └── repositories/       # Interfaces dos repositórios
│   └── services/               # Lógica de negócio
│       └── test/               # Testes unitários dos serviços
├── config/                     # Configuração da aplicação
├── domain/
│   ├── entities/               # Entidades do domínio (Site, Module, Component, etc)
│   └── errors/                 # Erros de domínio
├── docs/                       # Documentação e guias
├── infrastructure/
│   ├── database/               # Configuração do banco de dados
│   └── repositories/           # Implementação dos repositórios
├── presentation/
│   ├── handlers/               # Handlers das rotas HTTP
│   ├── middlewares/            # Middlewares (ex: autenticação)
│   └── routes/                 # Definição das rotas
├── docker-compose.yml          # Orquestração Docker
├── dockerfile                  # Build da aplicação
├── go.mod / go.sum             # Dependências Go
├── main.go                     # Ponto de entrada da aplicação
├── readme.md                   # Este arquivo
└── routes.json                 # (Opcional) Rotas extras
```

## ⚙️ Como rodar

1. Instale Go 1.24+ e MySQL (ou use Docker)
2. Configure o banco de dados e variáveis de ambiente (`.env`)
3. Instale dependências:
   ```sh
   go mod tidy
   ```
4. Rode a aplicação:
   ```sh
   go run main.go
   ```

## 🔀 Endpoints principais

- `POST   /api/user/register` — Cadastro de usuário
- `POST   /api/site` — Cadastro de site
- `GET    /api/site/:slug` — Buscar site completo por slug (com módulos e componentes)
- `POST   /api/site/module` — Cadastro de módulo
- `POST   /api/site/component` — Cadastro de componente

## 🔐 Autenticação
A API utiliza JWT. Para acessar rotas protegidas, envie o token no header:
```
Authorization: Bearer SEU_TOKEN
```

## 📝 Exemplo de resposta do endpoint de site
```json
{
  "site_id": 10000,
  "site_name": "Site Isaac",
  "site_slug": "isaacmachado",
  "modules": [
    {
      "module_id": 30000,
      "module_name": "Home",
      "module_slug": "home",
      "components": [
        {
          "component_id": 70000,
          "component_type": "BannerTop",
          "component_name": "Banner Principal",
          "component_settings": {
            "title": "Home do site",
            "subtitle": "Essa é a home...",
          }
        }
      ]
    }
  ]
}
```

## 📄 Licença
MIT

Desenvolvido para estudos e como base para construtores de sites dinâmicos.

