<?php
$f = fopen("log.txt", "a+");
fwrite($f, date("Y-m-d H:i:s") . "_test2\r\n");
fclose($f);
sleep(5);
exit;