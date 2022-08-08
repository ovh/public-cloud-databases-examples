<?php

if ($argc < 2) {
    fwrite(STDERR, "URI is missing\n");
    exit(1); 
}
try {
    $manager = new MongoDB\Driver\Manager($argv[1]);
    $command = new MongoDB\Driver\Command(["buildInfo" => 1]);
    $cursor = $manager->executeCommand("admin", $command);
    $response = $cursor->toArray()[0];
    echo $response->version . PHP_EOL;
}
catch (Throwable $e) {
    echo "Captured Throwable for connection: " . $e->getMessage() . PHP_EOL;
}
?>
