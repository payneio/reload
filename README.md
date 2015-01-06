Reload
======

```bash
reload program "restart command"
```

`reload` will watch the given program file for modification. If it is changed, it will execute the supplied restart command. 

This is nice to automatically kick off restart scripts when an executable file changes. Much lighter weight than monit and requires no graceful restart code within your executable.


