# CaleoX-SpaceForum 🌌

A **cross-platform command-line forum** built in **Go**, allowing users to sign up, post threads, reply, delete posts, and manage accounts — all stored in **JSONBin.io**. Fully portable CLI with embedded secrets — no `.env` required.

---

## 🔹 Features

- Signup / Login system  
- Create and view threads  
- Reply to threads  
- Delete your posts  
- Delete your account and all your posts  
- Pagination for threads with multiple posts  
- Fully cross-platform: macOS, Linux, Windows  
- Single executable — no setup required  

---

## 💻 Installation

### **1. Download the binary**

- macOS: `caleox-spaceforum-mac`  
- Linux: `caleox-spaceforum-linux`  
- Windows: `caleox-spaceforum.exe`

> Binaries are pre-built and portable — no `.env` needed.

### **2. Run the CLI**

```bash
# macOS / Linux
./caleox-spaceforum-mac

# Windows
caleox-spaceforum.exe
````

---

## 🛠️ Commands

```text
help                 Show all commands
signup               Create a new account
login                Login to your account
threads              List all threads
post                 Create a new thread
view [thread] [page] View a thread and its posts (optional page number)
reply [thread]       Reply to a thread
deletepost [thread] [post]  Delete one of your posts
deleteaccount        Delete your account and all your posts
exit                 Exit the CLI
```

### Example:

```bash
> signup
Username: jogo
Password: root
Signup successful!

> login
Username: jogo
Password: root
Logged in as jogo

> post
Thread title: Test
Thread created successfully!

> threads
1. Test (by jogo)

> view 1
Title: Test | Author: jogo
No posts yet. Be the first to reply!

> reply 1
Your reply: Hello world!
Reply posted successfully!

> deletepost 1 1
Are you sure? (y/n): y
Post deleted successfully!
```

---

## ⚙️ Build from Source

Make sure you have Go installed (>=1.20).

```bash
# Clone repository
git clone https://github.com/jagath-sajjan/caleox-spaceforum.git
cd caleox-spaceforum

# Build for your OS
go build -o caleox-spaceforum main.go

# Run
./caleox-spaceforum
```

### Cross-platform builds

```bash
# macOS Intel
GOOS=darwin GOARCH=amd64 go build -o caleox-spaceforum-mac main.go

# macOS ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o caleox-spaceforum-mac-arm main.go

# Linux x64
GOOS=linux GOARCH=amd64 go build -o caleox-spaceforum-linux main.go

# Linux ARM64 (Raspberry Pi)
GOOS=linux GOARCH=arm64 go build -o caleox-spaceforum-linux-arm main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o caleox-spaceforum.exe main.go
```

---

## 📂 Project Structure

```
caleox-spaceforum/
├── commands/      CLI commands (threads, auth)
├── utils/         JSONBin & secret handling
├── models/        Thread & Post models
├── main.go        Entry point
├── .gitignore
└── README.md
```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Commit your changes: `git commit -am 'Add new feature'`
4. Push to the branch: `git push origin feature/my-feature`
5. Open a Pull Request

---

## ⚠️ License

This project is **open-source**. Feel free to use, modify, and contribute.

---

Enjoy building your universe in the terminal! 🌌
