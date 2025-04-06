# 🧰 Vault CLI

A lightweight CLI tool to interact with the [Vault Security API](https://github.com/your-username/vault-security). This CLI supports login, secret storage, and retrieval via JWT authentication and Role-Based Access Control (RBAC).

---

## ⚙️ Features

- 🔑 Login via `user_id`
- 🔐 Store secrets (if allowed by role)
- 📖 Read secrets
- 🧪 Works with HashiCorp Vault + Vault Security API

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/vaultcli.git
cd vaultcli
```

---

## ▶️ Build CLI

```bash
go build -o vaultcli main.go
```

This generates a binary called `vaultcli`.

---

## 🛠️ Usage

```bash
./vaultcli -cmd=<command> [flags]
```

### 🔐 Login (Receive JWT Token)

```bash
./vaultcli -cmd=login -user=admin
```

### 💾 Store a Secret (Admin / Dev Only)

```bash
./vaultcli -cmd=store -key=<your-key> -value=<your-value> -token=<your-jwt-token>
```

### 📖 Read a Secret (All Roles)

```bash
./vaultcli -cmd=read -key=<your-key> -token=<your-jwt-token>
```

---

## 🧪 Example

```bash
./vaultcli -cmd=login -user=readonly
# returns JWT token

./vaultcli -cmd=read -key=project -token=<jwt-token>
# returns secret or error
```

---

## 🧩 Notes

- This CLI assumes your Vault Security API is running at `http://localhost:8080`
- Tokens expire based on the TTL defined in Vault
- If you want to run this from any directory, move the binary:
  ```bash
  sudo mv vaultcli /usr/local/bin/
  ```

---
