# Examples of use of OVHcloud Cloud Databases

![githubbanner](https://user-images.githubusercontent.com/3379410/27423240-3f944bc4-5731-11e7-87bb-3ff603aff8a7.png)

[![Maintenance](https://img.shields.io/maintenance/yes/2022.svg)]()

This repository centralize all resources and examples that could be use with the [OVHcloud Cloud Databases product](https://www.ovhcloud.com/en-gb/public-cloud/databases/)

# Installation

## Get the sources

```bash
git clone https://github.com/ovh/public-cloud-databases-examples.git

cd public-cloud-databases-examples
```

## Read and experiment the tutorials

The tutorials are categorised by **database engine**, then **language** and finally **example** type.
They all embedded a Dockerfile so you can try it with a simple build and run of it, by setting the environment variable URI to connect to the database.

```bash
databases/
└── postgresql
    └── bash
        └── hello-world
            ├── Dockerfile
            └── example
```

## Available examples

Engine    | Example   |       bash       |       python     |        php       |       golang     |     terraform
----------|-----------|:----------------:|:----------------:|:----------------:|:----------------:|:----------------:
Mongodb   |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:
Mysql     |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:
Postgresql|hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:
Redis™    |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|   :wrench: 
Kafka     |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:
Opensearch|hello-world|     :wrench:     |    :wrench:      |     :wrench:     |     :wrench:     |   :wrench:
Cassandra |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:
M3DB      |hello-world|:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|   :wrench: 

# Related links

 * Documentation: https://docs.ovh.com/gb/en/publiccloud/databases/
 * Contribute: https://github.com/ovh/public-cloud-databases-examples/blob/master/CONTRIBUTING.md
 * Report bugs: https://github.com/ovh/public-cloud-databases-examples/issues
 * Discord: https://discord.gg/ZnF86Crdbf

# License

Copyright 2021 OVH SAS

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
