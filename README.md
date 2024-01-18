# GoZero from Team Mudah Diingat

## Team
- `Hustler` Novera Prestiana Putri
- `Hipster` Diena Mukafasyadiah
- `Hipster` Kasyiful Kurobi Alqorrosyai'
- `Hacker` Fath Yusava Arden

## Product Preview
GoZero is a website-based application that focuses on `Zero Waste` from SDG 11, namely Sustainable Cities and Communities. GoZero will not only empower individuals but also contribute to a more sustainable and inclusive urban environment, reducing the environmental impact of cities and promoting responsible consumption and production patterns.<br><br>
GoZero is develop by using:
- Go for back-end
- HTML, CSS, and JS for front-end
- MySQL for database

### Feature
- Home Page
   ![image](https://github.com/yusavaa/GoZero/assets/120007453/3e3847d2-4051-4412-9843-3683d7b403b1)
- Community Gallery
   ![image](https://github.com/yusavaa/GoZero/assets/120007453/4447b69f-45ee-4c06-9d16-681b2d0362e6)
- Peta Interaktif
   ![image](https://github.com/yusavaa/GoZero/assets/120007453/81bff3eb-ca95-4de3-99e7-017abab0c619)
- Quiz
   ![image](https://github.com/yusavaa/GoZero/assets/120007453/301487f5-7bb3-420a-8338-738c50733cf0)
- Etc

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
