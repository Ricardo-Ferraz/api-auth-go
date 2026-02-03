# 🔐 API Auth Go

API de autenticação e autorização desenvolvida em **Go**, utilizando o framework **Gin**, com autenticação baseada em **JWT**, controle de usuários e roles, persistência em **AWS RDS** e pipeline de **CI/CD com GitHub Actions**.

A aplicação é **cloud-native**, empacotada em Docker e preparada para execução em **AWS EKS**, podendo rodar tanto em **Node Groups (EC2)** quanto em **Fargate**.

---

## Visão Geral

Esta API fornece um serviço centralizado de autenticação para aplicações backend, frontend ou microserviços, oferecendo:

- Registro de usuários
- Autenticação com geração de JWT
- Validação de token
- Controle de acesso por roles
- Persistência de dados em RDS
- Deploy automatizado via GitHub Actions
- Execução containerizada Kubernetes (EKS)

---

## 🏗 Arquitetura

- **Linguagem:** Go
- **Framework HTTP:** Gin
- **Autenticação:** JWT
- **Autorização:** Roles
- **Banco de Dados:** AWS RDS (MySQL ou compatível)
- **Containerização:** Docker
- **CI/CD:** GitHub Actions
- **Orquestração:** Kubernetes
- **Cloud Provider:** AWS (EKS)

A mesma imagem Docker pode ser executada:
- Em **EKS com Node Groups (EC2)** → maior controle da infraestrutura
- Em **EKS com Fargate** → modelo serverless, sem gerenciamento de nós

---

## Variáveis de Ambiente

A aplicação depende das seguintes variáveis de ambiente:

```env
DB_STRING=user:password@tcp(host:3306)/database_name (RDS)
JWT_SECRET=chave_secreta_jwt
```

---

## Modelo de Banco de Dados (RDS)

### Tabela users

```sql
CREATE TABLE users (
id INT NOT NULL AUTO_INCREMENT,
username VARCHAR(255) DEFAULT NULL,
pass_hash VARCHAR(255) DEFAULT NULL,
PRIMARY KEY (id),
UNIQUE KEY username (username)
);
```

### Tabela roles
```sql
CREATE TABLE roles (
id INT AUTO_INCREMENT PRIMARY KEY,
name VARCHAR(255),
user_id INT,
FOREIGN KEY (user_id) REFERENCES users(id)
);
```

---

## Inicialização da Aplicação

    go run main.go

    A aplicação escuta, por padrão, na porta 8080.

---

##  Endpoints da API

### Autenticação

#### Registra um novo usuário

`POST /auth/register`

**Corpo da requisição:**
```json
{
    "username": "string",
    "password": "string"
}
```

**Resposta** – 201 Created
```json
{
    "id": 1,
    "username": "string"
}
```
#### Possíveis erros
| Código | Motivo               |
| ------ | -------------------- |
| 400    | Dados inválidos      |
| 409    | Usuário já existente |

#### Geração de Token/Login

`POST /auth/login`

**Corpo da requisição:**
```json
{
    "username": "string",
    "password": "string"
}
```

**Resposta** – 200 OK
```json
{
  "token": "string",
  "exp": 0
}
```
#### Possíveis erros

| Código | Motivo                |
| ------ | --------------------- |
| 401    | Credenciais inválidas |

#### Verifica Token

`GET /auth/verify`

**Headers:**

    Authorization: Bearer <Token>

**Resposta** – 200 OK
```json
{
  "valid": true
}
```

### Usuários (Protegida)

`GET /users/:id`

Obs.:

- JWT válido
- Role obrigatória associada ao usuário (ROLE_0)
- Busca pelo ID

**Headers:**

    Authorization: Bearer <Token>

**Resposta** – 200 OK
```json
{
    "id": 0,
    "username": "string",
    "roles": [
        {
            "id": 0,
            "nameRole": "string"
        }
    ]
}
```
#### Possíveis erros
| Código | Motivo                    |
| ------ | ------------------------- |
| 401    | Token inválido ou ausente |
| 403    | Acesso negado             |
| 404    | Usuário não encontrado    |

### Roles (Protegida)

`POST /role/register`

Obs.:

- JWT válido
- Role obrigatória associada ao usuário (ROLE_0)
- Adiciona uma Role a um usuário

**Headers:**

    Authorization: Bearer <Token>

**Resposta** – 200 OK
```json
{
    "idUser": 0,
    "nameRole": "string"
}
```

**Resposta** – 200 OK
```json
{
    "id": 0,
    "name": "string"
}
```

### Versionamento

#### Versão atual

`GET /version`

**Resposta** – 200 OK
```json
{
    "version": "string",
    "commit": "string",
    "buildTime": "string"
}
```


### Tratamento de Exceções
A aplicação utiliza um padrão centralizado para tratamento de erros, garantindo consistência nas respostas HTTP.

#### Formato padrão
```json
{
    "code": "ERROR_CODE",
    "message": "Descrição do erro"
}
```

#### Tipos de erros tratados
- Senhas armazenadas como hash (pass_hash)

- JWT assinado com JWT_SECRET

- Middleware de autenticação

- Middleware de autorização por roles

- Rotas protegidas por token


### 🐳 Docker

A aplicação é empacotada em uma imagem Docker, utilizada tanto localmente quanto em produção.

A mesma imagem é reutilizada para (AWS):

- EKS Node Group

- EKS Fargate

### 🔁 CI/CD – GitHub Actions
Workflow localizado em: `.github/workflows/deploy.yml`

#### Trigger

- Push na branch main
- Pull request para main
- Execução manual (workflow_dispatch)

#### Etapas

- Checkout do código
- Geração automática da tag da imagem
- Login no Docker Hub
- Build da imagem Docker
- Push da imagem para o registry
- Assunção de role AWS via OIDC
- Configuração do kubeconfig
- Deploy automático no AWS EKS

### ☁️ Deploy no AWS EKS
A aplicação é implantada em dois ambientes Kubernetes distintos:

- EKS com Node Group (EC2)
- EKS com Fargate

O deploy utiliza kubectl set image, garantindo atualização da aplicação sem downtime.