# Lenscan

A open-source explorer for exploring the Lens Protocol, built with Next.js.

## Installation

Clone the project to your local machine.

```bash
git clone https://github.com/daoleno/lenscan
cd lenscan/dashboard
bun install
bun dev
```

## Build Docker Image

```bash
docker build --build-arg DUCKDB_PATH=/path/to/v2_polygon.db -t lenscan .
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
