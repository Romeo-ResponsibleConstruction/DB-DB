## Setup instructions
to install fiber and gorm:

run
> ```go get github.com/gofiber/fiber/v2```
> ```go get -u gorm.io/gorm```
> ```go get -u gorm.io/driver/mysql```
> ```go get github.com/google/uuid```
from the DB-DB directory

## Usage instructions
Ensure you have mysql installed and running on the same machine, with an existing database.

The file `dsn.txt` should exist in the DB-DB directory and contain exactly one line:
> ```root:{rootPassword}@/{databaseName}```,

where `{rootPassword}` and `{databaseName}` are the root password of your mysql instance and the name of the database you intend to use.

The file `dashboardaddress.txt` should exist in the DB-DB directory and contain exactly one line:
> ```{IPAddress}:{PortNumber}```,

where `{IPAddress}` is the IP address where you want the online dashboard to have and `{PortNumber}` it the port number the dashboard can be accessed from (we recommend using port 8080 as it is commonly used for web servers). In our case we used the IP address 127.0.0.1, not the public IP of the virtual machine our dashboard is hosted on, as the routing is handled by a virtual cloud network above the VM. Leaving `{IPAddress}` empty and only writing `:{PortNumber}` results in the same behaviour.

Also ensure that in the DB-DB directory, you have a subdirectory ./data/images.
