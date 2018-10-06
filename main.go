package main
 
import (
    "fmt"
    "net/http"
 
    helper "github.com/reg1/helper"
)
 
func main() {
 
    uName, email, pwd, pwdConfirm := "", "", "", ""
 
    mux := http.NewServeMux()
    // Signup
    mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        /*
            // Available for testing.
            for key, value := range r.Form {
                fmt.Printf("%s = %s\n", key, value)
            }
        */
 
        uName = r.FormValue("username")     // Data from the form
        email = r.FormValue("email")        // Data from the form
        pwd = r.FormValue("password")       // Data from the form
        pwdConfirm = r.FormValue("confirm") // Data from the form
    
 
        // Empty data checking
        uNameCheck := helper.IsEmpty(uName)
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
        pwdConfirmCheck := helper.IsEmpty(pwdConfirm)

 
        if uNameCheck || emailCheck || pwdCheck || pwdConfirmCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }
 
        if pwd == pwdConfirm {
            // Save to database (username, email and password)
            fmt.Fprintln(w, "Registration successful.")
        } else {
            fmt.Fprintln(w, "Password information must be the same.")
        }
    })
 
    // Login
    mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        r.ParseForm()
 
        email = r.FormValue("email")  // Data from the form
        pwd = r.FormValue("password") // Data from the form
         
        // Empty data checking
        emailCheck := helper.IsEmpty(email)
        pwdCheck := helper.IsEmpty(pwd)
 
        if emailCheck || pwdCheck {
            fmt.Fprintf(w, "ErrorCode is -10 : There is empty data.")
            return
        }
 
        dbPwd := "1234!*."                   // DB simulation
        dbEmail := "abc@hotmail.com" // DB simulation
 
        if email == dbEmail && pwd == dbPwd {
            fmt.Fprintln(w, "Login succesful!")
        } else {
            fmt.Fprintln(w, "Login failed!")
        }
    })
     println("starting server")
    http.ListenAndServe(":8080", mux)
}