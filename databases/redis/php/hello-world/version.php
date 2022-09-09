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
    $redis = new Redis() or die("Cannot load Redis module in PHP.");
    //Connecting to Redis
    $redis->connect("tls://$host", $port);
    $redis->auth(["$user", "$pass"]);

    //Get the server version
    $info = $redis->info();
    $version = $info["redis_version"];
    echo "Redis server version: $version";
}

catch (Throwable $e) {
    echo "Captured Throwable for connection: " . $e->getMessage() . PHP_EOL;
}

?>
