package rout

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"short/database"
	"short/models"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

const Pattern string = "http://127.1.1.0:2000/"

type HTTPHandler struct {
	Pool *pgxpool.Pool
}

func HomeGet(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("http/home.html")
	if err != nil {
		log.Println("не загрузил шаблон главной страницы")
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("http/errorPage.html")
	if err != nil {
		log.Println("не загрузил шаблон  страницы")
		http.Error(w, err.Error(), 500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

	return
}

func (h HTTPHandler) ErrorPageDuplicateData(w http.ResponseWriter, r *http.Request) {
	log.Println("not exec data to table Form rom DB")
	//defer h.Pool.Close()
	tmpl, err := template.ParseFiles("http/err2.html")
	if err != nil {
		log.Println("не загрузил шаблон главной страницы")
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

}

func (h HTTPHandler) ResultPost(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	userURL := r.FormValue("userURL")
	newURL := r.FormValue("newURL")

	newUserURL, err := ValidateData(userURL, newURL)
	if err != nil {
		ErrorPage(w, r)
	}

	mf := &models.Form{UserURL: newUserURL,
		NewURL:       Pattern + "s/" + newURL,
		StatisticURL: Pattern + "stat/" + newURL}

	err = database.InsertData(ctx, h.Pool, mf.UserURL, mf.NewURL)
	if err != nil {
		h.ErrorPageDuplicateData(w, nil)
		return
	}

	tmpl, err := template.ParseFiles("http/resultPage.html")
	if err != nil {
		log.Println("не загрузил шаблон главной страницы")
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, mf)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

}

func (h HTTPHandler) AllResults(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	form, _ := database.ShowAllResult(ctx, h.Pool)
	tmpl, err := template.ParseFiles("http/allresults.html")
	if err != nil {
		log.Println("не загрузил шаблон страницы")
		http.Error(w, err.Error(), 500)
		return
	}
	err = tmpl.Execute(w, form)
	if err != nil {
		log.Println("не могу вывести шаблон страницы")
		http.Error(w, err.Error(), 500)
		return
	}
	log.Println("Запрос на отоброжение всех результатов успешно выполнен")
}

func (h HTTPHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	vars := mux.Vars(r)

	name := vars["name"]

	redirect, err := database.Search(ctx, h.Pool, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("data base not response ")

	}

	log.Println(redirect.UserURL)

	type RedirectAdres struct {
		UserURL string
		NewURL  string
	}

	mf := RedirectAdres{
		UserURL: redirect.UserURL,
		NewURL:  redirect.NewURL,
	}

	tmpl, err := template.ParseFiles("http/redirect.html")
	if err != nil {
		log.Println("не загрузил шаблон главной страницы")
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, mf)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

	time := time.Now()
	ip := r.RemoteAddr

	err = database.InsertStat(ctx, h.Pool, ip, time, mf.NewURL)
	if err != nil {
		log.Println("not exec data to table statistic from DB")
	}

}

func (h HTTPHandler) GetStatistic(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	vars := mux.Vars(r)

	name := vars["name"]

	hints, err := database.SearchRowStat(ctx, h.Pool, name)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("data base not response ")

	}

	tmpl, err := template.ParseFiles("http/statistic.html")
	if err != nil {
		log.Println("не загрузил шаблон главной страницы")
		http.Error(w, err.Error(), 500)
		return
	}

	err = tmpl.Execute(w, hints)
	if err != nil {
		log.Println("не удалось вывести шаблон на экран")
		http.Error(w, err.Error(), 500)
		return
	}

}
