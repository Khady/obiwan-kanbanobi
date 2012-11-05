createuser -P kanban
createdb -E UTF-8 -O kanban kanban
psql -U kanban -d kanban -f createdb.sql
