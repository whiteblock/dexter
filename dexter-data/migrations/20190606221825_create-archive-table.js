exports.up = function(knex, Promise) {
  return knex.schema.createTable('archive', function(table) {
    table.increments('id');
    table.string('exchange', 255).notNullable();
    table.string('market', 255).notNullable();
    table.string('timeframe', 16).notNullable();
    table.timestamps();
  });
};

exports.down = function(knex, Promise) {
  return knex.schema.dropTable('archive');
};
