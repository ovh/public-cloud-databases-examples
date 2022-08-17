<?php

if ($argc < 2) {
    fwrite(STDERR, "URI is missing\n");
    exit(1); 
}
try {
    $pu = parse_url($argv[1]);
    $host = $pu["host"];
    $port = $pu["port"];
    $user = $pu["user"];
    $pass = $pu["pass"];
    $conn = new PDO("mysql:host=$host; port=$port", $user, $pass);
    $rows = $conn->query("select version();")->fetchAll();
    echo $rows[0][0] . PHP_EOL;
    $conn = null;
}
catch (Throwable $e) {
    echo "Captured Throwable for connection: " . $e->getMessage() . PHP_EOL;
}
?>
