# portchecker

portchecker is a simple application to check the port of host is connectable or not

## Usage

```
portchecker [-constr connection string] [-env environment-variable-name] [-attempt maximum-attempt-times] [-wait connection-failed-waiting-time]
    -constr connection string under test
        example:
            tcp://127.0.0.0.1:1234
            postgres://localhost/test
            mysql://testuser:password@mysql-host.local:1234
            redis://localhost

    -env environment variable name
        example:
            DB_URL
            REDIS_URL

    -attempt how many times to try, default: 5 times
    -wait wait how many time to try next time, default: 3s
```
