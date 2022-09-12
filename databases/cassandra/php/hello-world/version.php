<?php

if ($argc < 4) {
    fwrite(STDERR, "missing parameters, should be: Host, Port, Username, Password\n");
    exit(1); 
}

$ssl = \Cassandra::ssl()->withVerifyFlags(\Cassandra::VERIFY_PEER_CERT);
$ssl = $ssl->withTrustedCerts("/certificates/ca.certificate.pem");

$ssl = $ssl->build();


$cluster   = Cassandra::cluster()->withCredentials($argv[3], $argv[4])->withContactPoints($argv[1])->withPort((int)$argv[2])              
    ->withSSL($ssl)->build();
$session   = $cluster->connect();        // create session, optionally scoped to a keyspace
$statement = new Cassandra\SimpleStatement(       // also supports prepared and batch statements
    'SELECT cql_version FROM system.local'
);
$future    = $session->executeAsync($statement);  // fully asynchronous and easy parallel execution
$result    = $future->get();                      // wait for the result, with an optional timeout

foreach ($result as $row) {                       // results and rows implement Iterator, Countable and ArrayAccess
    printf("Version: %s\n", $row['cql_version']);
}
