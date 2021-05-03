package main

// The sql go library is needed to interact with the database
import (
	"database/sql"
	"math/rand"
	"strconv"
	"time"
)

type Store interface {
	CreateTicket(ticket *Ticket) error
	GetTickets() ([]*Ticket, error)
	GetDeletedTickets() ([]*Ticket, error)
	DeleteTicket(ticket *Ticket) error
	UpdateTicket(ticket *Ticket, newId string) error
	InsertDeleteTicket(ticket *Ticket) error
	RecoverTicket(ticket *Ticket) error
	InsertRecoverTicket(ticket *Ticket) error
	FilterTicket(filter string, searchFilter string) error
	DeleteFilterTicket() error
}

// The `dbStore` struct will implement the `Store` interface
// It also takes the sql DB connection object, which represents
// the database connection.
var SearchFilter string
var Filter string

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) CreateTicket(ticket *Ticket) error {

	dt := time.Now()
	rand.Seed(time.Now().UnixNano())
	_, err := store.db.Query("INSERT INTO registered_tickets(id, usuario, fechacreacion, fechaactualizacion, estatus) VALUES ($1,$2,$3,$4,$5)",
		strconv.Itoa(rand.Intn(10000000)), ticket.Usuario, dt.Format("01-02-2006"), "Sin Actualizaciones", "Abierto")
	return err
}

func (store *dbStore) GetTickets() ([]*Ticket, error) {

	tickets := []*Ticket{}

	if SearchFilter != "" {

		rows, err := store.db.Query("SELECT id, usuario, fechacreacion, fechaactualizacion, estatus from registered_tickets")

		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {

			ticket := &Ticket{}
			if err := rows.Scan(&ticket.Id, &ticket.Usuario, &ticket.FechaCreacion, &ticket.FechaActualizacion, &ticket.Estatus); err != nil {
				return nil, err
			}

			if Filter == "id" {
				if ticket.Id == SearchFilter {
					tickets = append(tickets, ticket)
				}

			} else if Filter == "usuario" {
				if ticket.Usuario == SearchFilter {
					tickets = append(tickets, ticket)
				}

			} else if Filter == "fechacreacion" {
				if ticket.FechaCreacion == SearchFilter {
					tickets = append(tickets, ticket)
				}

			} else if Filter == "fechaactualizacion" {
				if ticket.FechaActualizacion == SearchFilter {
					tickets = append(tickets, ticket)
				}

			} else {
				if ticket.Estatus == SearchFilter {
					tickets = append(tickets, ticket)
				}
			}

		}

	} else {
		rows, err := store.db.Query("SELECT id, usuario, fechacreacion, fechaactualizacion, estatus from registered_tickets")

		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {

			ticket := &Ticket{}
			if err := rows.Scan(&ticket.Id, &ticket.Usuario, &ticket.FechaCreacion, &ticket.FechaActualizacion, &ticket.Estatus); err != nil {
				return nil, err
			}
			tickets = append(tickets, ticket)

		}

	}

	return tickets, nil
}

func (store *dbStore) GetDeletedTickets() ([]*Ticket, error) {

	rows, err := store.db.Query("SELECT id, usuario, fechacreacion, fechaactualizacion, estatus from deleted_tickets")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	deleted_tickets := []*Ticket{}
	for rows.Next() {

		ticket := &Ticket{}

		if err := rows.Scan(&ticket.Id, &ticket.Usuario, &ticket.FechaCreacion, &ticket.FechaActualizacion, &ticket.Estatus); err != nil {
			return nil, err
		}

		deleted_tickets = append(deleted_tickets, ticket)
	}
	return deleted_tickets, nil
}

func (store *dbStore) DeleteTicket(ticket *Ticket) error {
	sqlStatement := `
	DELETE FROM registered_tickets
	WHERE id = $1;`

	_, err := store.db.Exec(sqlStatement, ticket.Id)

	return err

}

func (store *dbStore) InsertDeleteTicket(ticket *Ticket) error {
	dt := time.Now()

	_, err := store.db.Query("INSERT INTO deleted_tickets(id, usuario, fechacreacion, fechaactualizacion, estatus) VALUES ($1,$2,$3,$4,$5)",
		ticket.Id, ticket.Usuario, ticket.FechaCreacion, dt.Format("01-02-2006"), "Cerrado")

	return err

}

func (store *dbStore) UpdateTicket(ticket *Ticket, newId string) error {
	dt := time.Now()

	sqlStatement := `
	UPDATE registered_tickets
	SET id = $2, usuario = $3, fechaactualizacion = $4, estatus = $5
	WHERE id = $1;`

	_, err := store.db.Exec(sqlStatement, ticket.Id, newId, ticket.Usuario, dt.Format("01-02-2006"), ticket.Estatus)
	return err
}

func (store *dbStore) RecoverTicket(ticket *Ticket) error {
	sqlStatement := `
	DELETE FROM deleted_tickets
	WHERE id = $1;`

	_, err := store.db.Exec(sqlStatement, ticket.Id)

	return err

}

func (store *dbStore) InsertRecoverTicket(ticket *Ticket) error {
	dt := time.Now()

	_, err := store.db.Query("INSERT INTO registered_tickets(id, usuario, fechacreacion, fechaactualizacion, estatus) VALUES ($1,$2,$3,$4,$5)",
		ticket.Id, ticket.Usuario, ticket.FechaCreacion, dt.Format("01-02-2006"), "Abierto")

	return err

}

func (store *dbStore) FilterTicket(filter string, searchfilter string) error {
	SearchFilter = searchfilter
	Filter = filter

	return nil
}

func (store *dbStore) DeleteFilterTicket() error {
	SearchFilter = ""
	Filter = ""

	return nil
}

// The store variable is a package level variable that will be available for
// use throughout our application code
var store Store

/*
We will need to call the InitStore method to initialize the store. This will
typically be done at the beginning of our application (in this case, when the server starts up)
This can also be used to set up the store as a mock, which we will be observing
later on
*/
func InitStore(s Store) {
	store = s
}
