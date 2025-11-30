# TON Exporter

A lightweight **Prometheus exporter** that exposes the **latest TON masterchain block height** using TON Center’s **v2** or **v3** API.

---

## 🚀 Features

- Supports **TON Center API v2** and **v3**
- Periodically scrapes the latest masterchain block height

## ⚙️ Command-line Flags

| Flag     | Description                              | Default |
|----------|------------------------------------------|---------|
| `--port` | Port for the HTTP metrics server         | `4000`  |
| `--api`  | TON Center API version (v2 or v3)        | `none`  |
| `--rpc`  | TON RPC endpoint URL                     | `none`  |

## 🧪 Running the Exporter

### API v2
```bash
./ton-exporter --port=4000 \
  --api=v2 \
  --rpc=https://toncenter.com/api/v2/getMasterchainInfo
```

### API v3
```bash
./ton-exporter --port=4000 \
  --api=v3 \
  --rpc=https://toncenter.com/api/v2/getMasterchainInfo
```