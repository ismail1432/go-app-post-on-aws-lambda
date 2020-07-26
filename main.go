    package main

    import (
        "net/http"
        "html/template"
    )

    var tpl *template.Template

    func init() {
        tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
    }

    func main() {
        http.HandleFunc("/", index)
        http.ListenAndServe(":8080", nil)
    }

    func index(w http.ResponseWriter, r *http.Request) {
        tpl.ExecuteTemplate(w, "index.gohtml", nil)
    }

    func process_message(w http.ResponseWriter, r *http.Request) {
        if r.Method != "POST" {
            http.Redirect(w, r, "/", http.StatusSeeOther)
            return
        }

        //fcontent := r.FormValue("linkedin_content")

        //values := map[string]string{"username": username, "password": password}

        //jsonValue, _ := json.Marshal(values)

        //resp, err := http.Post(authAuthenticatorUrl, "application/json", bytes.NewBuffer(jsonValue))

    }
