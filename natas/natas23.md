
# General info

Username: natas23
Password: dIUQcI3uSus1JEOSSWRAEXBG8KbR8tRs
URL:      http://natas23.natas.labs.overthewire.org

# Special code section

Am facut totul din browser.
Codul vulnerabil era acesta:
```php
<?php
    if(array_key_exists("passwd",$_REQUEST)){
        if(strstr($_REQUEST["passwd"],"iloveyou") && ($_REQUEST["passwd"] > 10 )){
            echo "<br>The credentials for the next level are:<br>";
            echo "<pre>Username: natas24 Password: <censored></pre>";
        }
        else{
            echo "<br>Wrong!<br>";
        }
    }
    // morla / 10111
?>  
```
`strstr` m-am gandit ca verifica exista unui string intr-ul alt string, iar `$_REQUEST["passwd"] > 10` verifica ca numarul care exista in request sa fie mai mare decat 10.
Totusi, cum poate un text sa fie si numar, si string cu `iloveyou` in acelasi timp?
Am zis sa incerc `100 iloveyou` pe ideea ca s-ar face conversia doar pana la primul whitespace, iar aceasta idee a functionat.