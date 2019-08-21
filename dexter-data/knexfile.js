// Update with your config settings.
var dotenv = require("dotenv");

dotenv.config();

const common = {
  client: "postgresql",
  connection: {
    database: process.env.DB_NAME || "dexter",
    user: process.env.DB_USER,
    password: process.env.DB_PASSWORD
  },
  migrations: {
    tableName: "knex_migrations"
  },
  directory: "./migrations"
};

module.exports = {
  development: common,
  production: common
};
