## Value ID

A race-safe way to assign a string value an unchanging ID using an SQL or embedded database.

Sometimes I find I have strings I would rather not store copies of everywhere.
This library turns them into 64bit uint64 ids in an ACID-safe way.

I have the same basic query pattern at least once in every project:

```
CreateObjectID -> id
GetObjectByID -> object
GetOrCreateObjectID -> id (get id or create if missing, then get id)
```

Basically, I need to store some object and use the identifier everywhere else in the system. I’m tempted to use reflect or a code generator to make something so I don’t keep typing out these same functions eveywhere. However, type-safety is something I can't find a solution for.