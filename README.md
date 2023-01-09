## Installation & Usage

* Clone this repo 

```bash
git clone https://github.com/mfturkcanoglu/echo-todo.git
```

* Run Docker Postgres Container 

```bash
make create_postgres
```

* Copy sql scripts to Docker Container

```bash
make copy_sql_files
```

* Insert sql scripts

```bash
make run_sql_files
```

To run this application, execute:

```bash
go run ./cmd/server
```

You should be able to access this application at `http://127.0.0.1:8080`
