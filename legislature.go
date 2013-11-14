package main

import (
	"io/ioutil"
	"strings"
	"net/http"
	"html/template"
	"regexp"
	"errors"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

type Bill struct {
	Name string
	Tags []string
}

func (b* Bill) TagString() string {
	return strings.Join(b.Tags, ",")
}

func (b* Bill) save() error {
	filename := b.Name + ".txt"
	return ioutil.WriteFile(filename, []byte(b.TagString()), 0600)
}

func loadBill(title string) (*Bill, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	tags := strings.Split(string(body), ",")

	return &Bill{Name: title, Tags: tags}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, b *Bill) {
	err := templates.ExecuteTemplate(w, tmpl, b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	b, err := loadBill(title)
	if err != nil {
		http.Redirect(w, r, "http://localhost:8080/edit/" + title, http.StatusFound)
		return
	}

	renderTemplate(w, "view.html", b)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	b, err := loadBill(title)
	if err != nil {
		b = &Bill{Name: title}
	}

	renderTemplate(w, "edit.html", b)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	tags_string := r.FormValue("tags")
	tags := strings.Split(tags_string, ",")
	b := &Bill{Name: title, Tags: tags}
	err = b.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "http://localhost:8080/view/" + title, http.StatusFound)
}

func main() {
	b1 := &Bill{Name: "TestBill", Tags: []string{"sample", "new", "test", "bill"}}
	b1.save()
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
