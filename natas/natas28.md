
# General info

Username: natas28

Password: 1JNwQM1Oi6J6j1k49Xyw7ZN6pXMQInVj

URL:      http://natas28.natas.labs.overthewire.org

# Special code section

Pentru acealasi input, avem acelasi query.
Totusi, nu e garantat ca vom avea aceleasi "glume" (nu cred ca e relevant).

Query-ul este _URL encoded_, nu _base64URL encoded_. 
Am invatat sa fac diferenta intre ele.

Query-ul selecteaza glumele care contin un string dat ca input.

Primele doua linii vazute cu `xxd` sunt mereu la fel.
Ultimele doua linii se schimba doar daca dimensiunea inputului se schimba.

---

De aici am inceput sa citesc un write-up si am vazut ceva legat de AES.
Ma voi uita peste asta in continuare.

---

Am incercat sa citesc doua writeups, am inteles cat de cat ce se intampla, dar nu total.
Voi mai incerca cu [acest videoclip](https://www.youtube.com/watch?v=qpC2sNcRj5o).

---

```py
import requests
import urllib.parse
import base64


for i in range(64):
    q = "a" * i
    response = requests.post(
        url="http://natas28.natas.labs.overthewire.org/",
        auth=("natas28", "1JNwQM1Oi6J6j1k49Xyw7ZN6pXMQInVj"),
        data={
            "query" : q
        }
    )
    link = response.url.split("query=")[1]
    link2 = urllib.parse.unquote(link)
    link3 = base64.b64decode(link2)

    print(f"{i:<5} : {len(link):<5} {len(link2):<5} {len(link3):<5} ")
```

Asa aflam ce block size avem.

LE: ma dau batut, nu inteleg cum trebuie facuta aceasta problema, asa ca voi trece la urmatoarea si voi reveni la ea.