
# General info

Username: natas22
Password: d8rwGBl0Xslg3b76uh3fEbSlnOUBlozz
URL:      http://natas22.natas.labs.overthewire.org

# Special code section

```py
import requests

response = requests.get(
    url="http://natas22.natas.labs.overthewire.org/index.php",
    auth=("natas22", "d8rwGBl0Xslg3b76uh3fEbSlnOUBlozz"),
    params={
        "revelio": ""
    },
    allow_redirects=False
)

print(response.text)
print(*response.headers.items(), sep="\n")
print(response.status_code)
```

Am vorbit cu ChatGPT si am inteles ca, dupa acel apel de `header()`, ar trebui sa exista un `exit()`/`die()`, altfel codul se executa in continuare.
Fiind `header("Location: /");`, un browser simplu va urmari redirect-ul, ceea ce nu ne ajuta.
Totusi, daca folosesti o librarie / utilitar care te lasa sa NU urmaresti redicturile automat, poti vedea cum se executa codul in continuare.
De asta, daca folosim `allow_redirects=False`, putem vedea rezultatul.