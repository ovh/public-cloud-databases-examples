<?php
require 'vendor/autoload.php';

if ($argc < 2) {
    fwrite(STDERR, "usage: ".$argv[0]." <URI>\n");
    exit(1); 
}
try {
    $pu = parse_url($argv[1]);
    $host = $pu["host"];
    $port = $pu["port"];
    $user = $pu["user"];
    $pass = $pu["pass"];

    $client = new \GuzzleHttp\Client();
    $url = "https://$host:$port/api/v1/labels";
    $response = $client->get($url, [
        'auth' => [
            $user,
            $pass
        ],
    ]);
    echo $response->getBody();
}
catch (Throwable $e) {
    echo "Captured Throwable for connection: " . $e->getMessage() . PHP_EOL;
}
?>
