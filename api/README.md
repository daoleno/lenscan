# API

## Create a database user

```sql
create role web_anon nologin;

grant usage on schema public to web_anon;
grant select on all tables in schema public to web_anon;
```

## Run PostgRest

```sh
postgrest lenscan.conf
```

```sh
# dump db
pg_dump -Fc postgres://postgres:postgres@localhost:5432/lenscan > lenscan.dump
```

```sh
# restore db
pg_restore -d postgres://postgres:postgres@localhost:5432/lenscan lenscan-conf/lenscan.dump
```
