# Local Blockchain Network

A **fully decentralized blockchain network** developed in **Go (Golang)**. This project enables secure and tamper-resistant transactions between nodes while implementing core blockchain functionalities such as block creation, transaction management, consensus mechanisms, and peer-to-peer networking.

## Features

- **Block Creation & Validation** - Ensures each block contains valid transactions before being added to the blockchain.
- **Transaction Management** - Allows users to create, sign, and broadcast transactions securely.
- **Consensus Algorithm** - Implements a proof-based mechanism to ensure agreement across nodes.
- **Networking Layer** - Establishes a peer-to-peer communication protocol for efficient blockchain synchronization.
- **Immutable Ledger** - Uses cryptographic hashing to maintain data integrity and prevent tampering.
- **Node Synchronization** - Ensures all participating nodes have the latest valid blockchain state.

---

## Tech Stack

- **Programming Language**: Go (Golang)
- **Networking**: HTTP, WebSockets
- **Consensus Mechanism**: Custom implementation
- **Cryptographic Security**: SHA256 Hashing
- **Data Structure**: Blockchain (Linked Blocks)

---

## Setup & Installation

### **Prerequisites**
- Install **Go**: [Download Golang](https://go.dev/doc/install)
- Set up **Go environment**: `export GOPATH=$HOME/go`

### **Clone the Repository**
```bash
git clone https://github.com/syedibtisam/local-blockchain-network.git
cd local-blockchain-network
```

### **Install Dependencies**
```bash
go mod tidy
```

---

## Running the Blockchain Node

To start a blockchain node, use:
```bash
go run main.go
```

This will initialize a blockchain node that:
- Creates an initial **Genesis Block**.
- Listens for **incoming transactions**.
- Mines new blocks and syncs with connected nodes.

---

## Interacting with the Network

1. **Create a Transaction**
   ```bash
   curl -X POST http://localhost:8080/transaction -d '{"sender":"Alice", "receiver":"Bob", "amount":10}'
   ```

2. **Mine a Block**
   ```bash
   curl -X POST http://localhost:8080/mine
   ```

3. **View the Blockchain**
   ```bash
   curl http://localhost:8080/blockchain
   ```

---

## Future Enhancements
- Implement Proof-of-Work (PoW) or Proof-of-Stake (PoS)
- Enhance networking with full **peer discovery**
- Add smart contract execution support

---

## Contributing

Contributions are welcome! Feel free to **fork** this repository and submit **pull requests**.

---

## License

This project is licensed under the **MIT License**. See [LICENSE](LICENSE) for details.
