package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var Id string

type Ticket struct {
	Id                 string `json:"id"`
	Usuario            string `json:"usuario"`
	FechaCreacion      string `json:"fechacreacion"`
	FechaActualizacion string `json:"fechaactualizacion"`
	Estatus            string `json:"estatus"`
}

func getTicketsHandler(w http.ResponseWriter, r *http.Request) {

	tickets, err := store.GetTickets()

	ticketListBytes, err := json.Marshal(tickets)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(ticketListBytes)
}

func getDeletedTicketsHandler(w http.ResponseWriter, r *http.Request) {

	tickets, err := store.GetDeletedTickets()

	ticketListBytes, err := json.Marshal(tickets)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(ticketListBytes)
}

func createTicketHandler(w http.ResponseWriter, r *http.Request) {
	ticket := Ticket{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ticket.Id = r.Form.Get("id")
	ticket.Usuario = r.Form.Get("usuario")
	ticket.FechaCreacion = r.Form.Get("fechacreacion")
	ticket.FechaActualizacion = r.Form.Get("fechaactualizacion")
	ticket.Estatus = r.Form.Get("estatus")

	err = store.CreateTicket(&ticket)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func deleteTicketHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tickets, err := store.GetTickets()

	for _, ticket := range tickets {
		if ticket.Id == r.Form.Get("id") {
			err = store.DeleteTicket(ticket)

			if err != nil {
				fmt.Println(err)
			}

			err = store.InsertDeleteTicket(ticket)

			if err != nil {
				fmt.Println(err)
			}

			break
		}
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func updateTicketHandler(w http.ResponseWriter, r *http.Request) {
	ticket := Ticket{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ticket.Id = r.Form.Get("id")
	ticket.Usuario = r.Form.Get("usuario")
	ticket.Estatus = r.Form.Get("estatus")

	err = store.UpdateTicket(&ticket, r.Form.Get("newid"))
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)

}

func recoverTicketHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tickets, err := store.GetDeletedTickets()

	for _, ticket := range tickets {
		if ticket.Id == r.Form.Get("id") {
			err = store.InsertRecoverTicket(ticket)

			if err != nil {
				fmt.Println(err)
			}

			err = store.RecoverTicket(ticket)

			if err != nil {
				fmt.Println(err)
			}

			break
		}
	}

	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func recoverAllTicketsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tickets, err := store.GetDeletedTickets()

	for _, ticket := range tickets {
		err = store.InsertRecoverTicket(ticket)

		if err != nil {
			fmt.Println(err)
		}

		err = store.RecoverTicket(ticket)
		http.Redirect(w, r, "/assets/", http.StatusFound)
	}

}

func filterTicketsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	store.FilterTicket(r.Form.Get("filter"), r.Form.Get("filter-s"))
	http.Redirect(w, r, "/assets/", http.StatusFound)

}

func deleteFilterTicketsHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	store.DeleteFilterTicket()
	http.Redirect(w, r, "/assets/", http.StatusFound)

}
