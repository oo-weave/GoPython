# Xeno

## Getting Started - Go

Install dependencies
```shell
go install
```

#### Running Extraction
```shell
go run main.go
```

#### Running Tests
```shell
go run main.go
```

## Getting Started - Python

Setup Python virtual env (linux)
```shell
mkdir -p env
python3 -m venv env/venv
source env/venv/bin/activate
pip install --upgrade pip
pip install requests
```

#### Running Extraction
```shell
python main.py
```


#### Running Tests
```shell
python -m unittest -v xeno_py/api_test.py
```




