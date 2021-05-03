To run the server on your system:

1. Make sure you have pgAdmin4 installed
2. Run `git clone https://github.com/nico1029/ticketManagement.git` in your destination folder
3. Run `go install github.com/nico1029/ticketManagement`
4. Run `go install` to create the binary (`github.com/nico1029/ticketManagement`)
5. Run the binary : `./blog_example__go_web_db`
6. Launch the SQL Shell (psql)
7. Create the database called `ticket_database` in pgAdmin. 

Create the followings tables in your database before running the application :

```sql
CREATE TABLE registered_tickets (
    numberid SERIAL PRIMARY KEY, 
    id VARCHAR(256), 
    usuario VARCHAR(256), 
    fechacreacion VARCHAR(256), 
    fechaactualizacion VARCHAR(256), 
    estatus VARCHAR(256)
)

CREATE TABLE deleted_tickets (
    numberid SERIAL PRIMARY KEY, 
    id VARCHAR(256), 
    usuario VARCHAR(256), 
    fechacreacion VARCHAR(256), 
    fechaactualizacion VARCHAR(256), 
    estatus VARCHAR(256)
)
```

Before running the application, edit the `connString` variable inside the `main` function to specify your postgres database connection. Replace $YOUR_PASSWORD$ with your password used on database
