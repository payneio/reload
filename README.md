Reload
======

```bash
reload program "restart command"
```

`reload` will watch the given program file for modification. If it is changed, it will execute the supplied restart command.

You can create the program.pid like so:

```bash
/usr/local/bin/program & echo $! > /var/run/program.pid
```
