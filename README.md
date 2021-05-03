To run the server on your system:

1. Make sure you have pgAdmin4 installed
2. Run `git clone https://github.com/nico1029/ticketManagement.git` in your destination folder
3. Run `go install github.com/nico1029/ticketManagement`
4. Run `go build` to create the binary (`github.com/nico1029/ticketManagement`)
6. Launch the SQL Shell (psql)
7. Create the database called `ticket_database` in pgAdmin. 
8. Run the binary :`./ticketManagement`

Create the followings tables in your database before running the application :

```sql
CREATE TABLE registered_tickets (
    numberid SERIAL PRIMARY KEY, 
    id VARCHAR(256), 
    usuario VARCHAR(256), 
    fechacreacion VARCHAR(256), 
    fechaactualizacion VARCHAR(256), 
    estatus VARCHAR(256)
);

CREATE TABLE deleted_tickets (
    numberid SERIAL PRIMARY KEY, 
    id VARCHAR(256), 
    usuario VARCHAR(256), 
    fechacreacion VARCHAR(256), 
    fechaactualizacion VARCHAR(256), 
    estatus VARCHAR(256)
);
```

Before running the application, edit the `connString` variable inside the `main` function to specify your postgres database connection. Replace $YOUR_PASSWORD$ with your password used on database. 

When a Query is pulled from the server, it is necesarry to select the .html file to continue. The registered.html file contains lines necessary to display the interface to register a ticket and the deleted.html to delete it. On the other hand, the tickets are filtered acorrding how exactly you input the information. 

TODO
- Deploy the API with Docker containers 
