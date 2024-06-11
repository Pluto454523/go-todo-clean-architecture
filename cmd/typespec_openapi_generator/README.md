# typespec_openapi_generator
This is a tool to generate OpenAPI 3.0 specification from Typespec.

## Generate OpenAPI 3.0 using Docker
```bash

docker compose up -d
```

## Generate OpenAPI 3.0 using Node.js
```bash
npm install -g typespec
npm install

tsp compile . --emit "@typespec/openapi3"
```