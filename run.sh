####
 # Copyright (c) 2023 by jimyag, All Rights Reserved. 
 # Licensed under the MIT License. See LICENSE file in the project root for license information.
### 
#!/bin/bash
ps ax | grep ./bin/mactools | grep -v grep | awk '{print $1}' | xargs kill -9
./bin/mactools > /dev/null 2>&1 &
