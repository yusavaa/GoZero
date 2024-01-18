# GoZero from Team Mudah Diingat

## Requirements
- Go 1.18 or higher. We aim to support the 3 latest versions of Go.
- MySQL (5.6+), MariaDB, Percona Server, Google CloudSQL or Sphinx (2.2.3+)

## How To Run
1. Clone Project to your destination directory.
```bash
git clone https://github.com/yusavaa/GoZero
```
2. Import `gozero.sql` file to MySQL server, can use tools like phpmyadmid or so on.
3. Customize `username`, `password`, and `database name` on file `config/database.go`.
   
    ![image](https://github.com/yusavaa/GoZero/assets/120007453/a95b4018-c4cb-48b3-82d9-f016a73d22f2)
4. Open your terminal or command prompt, navigate to the directory where your Go file is located, and use the `go run` command to execute.
```go
go run main.go
```
> When you run `main.go` file for the first time, you will download all the dependencies needed in the `go.mod` file.
5. Open `localhost:8080` in your browser, or adjust the port by changing `Addr` in the `main.go` file.

    ![image](https://github.com/yusavaa/GoZero/assets/120007453/4b056e26-32ed-4f3f-87ee-6c398e2c0cd9)
