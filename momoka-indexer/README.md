# Momoka Indexer

## Installation

```bash
yarn install
```

## Usage

```bash
yarn start
```

## Docker

```bash
docker build -t momoka-indexer .
docker run -p 3000:3000 -e DATABASE_URL=postgres://postgres:postgres@localhost:5432/lenscan momoka-indexer
```
