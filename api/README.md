# API

## Create a database user

```sql
create role web_anon nologin;

grant usage on schema public to web_anon;
grant select on all tables in schema public to web_anon;
```

## Run PostgRest

```bash
postgrest lenscan.conf
```
