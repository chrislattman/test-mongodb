# MongoDB examples

> These examples require Docker to be installed.

One time only: run `docker pull mongo` to download the MongoDB Docker image.

Use the provided Makefile to run the examples.

- While MongoDB is written in C++, there is a C driver available
- To view the contents of this database from the command line, run `docker exec -it mongodb mongosh`
    - In the mongosh shell run `use mydb` and then run your commands as `db.customers.<command>`
- To view the contents of this database in a GUI, use MongoDB Compass
- To view the contents of this database in a web browser, use [mongo-express](https://github.com/mongo-express/mongo-express)

## Differences with SQL

While SQL RDBMSes and MongoDB both have databases, RDBMS databases store _tables_ that contains _rows_ of _column_ data, whereas MongoDB databases store _collections_ that contains _documents_ of _field_ data. RDBMS rows have a primary key that uniquely identifies them, whereas MongoDB documents have a `_id` field that uniquely identifies them.

- MongoDB is great if your data is unstructured (technically, it follows the [BSON](https://www.mongodb.com/docs/manual/reference/bson-types/) format)
- However, if you know that your data is going to follow a certain structure, use SQL as it's faster and more optimized

> The insertMany function calls in these examples are _not_ atomic, read more [here](https://www.mongodb.com/docs/manual/core/write-operations-atomicity/)
