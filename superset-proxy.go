package main

import (
	"log"
	"os"
	"syscall"
	"text/template"
)

func env(s, d string) string {
	r := os.Getenv(s)
	if r == "" {
		return d
	}
	return r
}

func substring(start, length int, s string) string {
	if start < 0 {
		return s[:length]
	}
	if length < 0 {
		return s[start:]
	}
	return s[start:length]
}

func main() {
	funcMap := template.FuncMap{
		"env":       env,
		"substring": substring,
	}

	tmpl, err := template.New("nginx").Funcs(funcMap).ParseFiles("/etc/nginx/nginx.conf.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	out, err := os.Create("/etc/nginx/nginx.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	err = tmpl.ExecuteTemplate(out, "nginx.conf.tmpl", "")
	if err != nil {
		log.Fatal(err)
	}

	env := os.Environ()
	err = syscall.Exec(os.Args[1], os.Args[1:], env)
	if err != nil {
		log.Fatal(err)
	}
}
