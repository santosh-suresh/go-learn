package main

import (
	"io/ioutil"
  "net/http"
  "html/template"
  "regexp"
  "errors"
  "log"
  "flag"
  "net"
  "os"
)

var templates = template.Must(template.ParseFiles("tmpl/edit.html","tmpl/view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var (
    		addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
  os.Mkdir("data",0755)
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body,err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }

  return &Page{Title: title, Body: body}, nil
}


func getTitle(w http.ResponseWriter, r *http.Request) (string,error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
      http.NotFound(w,r)
      return "", errors.New("Invalid Page Title")
    }
    return m[2], nil
}

func makeHandler(fn func(w http.ResponseWriter, r *http.Request, title string)) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
      http.NotFound(w,r)
      return
    }
    fn(w, r, m[2])
  }
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
  p,err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/"+title, http.StatusFound)
    return
  }
  renderTemplate(w, "view", p)
}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w,"edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  err := templates.ExecuteTemplate(w, tmpl+".html",p)
  if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
   }
}




func main() {
  flag.Parse()
   http.HandleFunc("/view/", makeHandler(viewHandler))
   http.HandleFunc("/edit/", makeHandler(editHandler))
   http.HandleFunc("/save/", makeHandler(saveHandler))

   if *addr {
       l, err := net.Listen("tcp", "127.0.0.1:0")
       if err != nil {
           log.Fatal(err)
       }
       err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 0644)
       if err != nil {
           log.Fatal(err)
       }
       s := &http.Server{}
       s.Serve(l)
       return
   }

   http.ListenAndServe(":8080", nil)}
