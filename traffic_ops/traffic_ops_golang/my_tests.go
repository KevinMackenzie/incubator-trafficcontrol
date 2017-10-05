package main;

import (
        "fmt"
        "net/http"

        "github.com/jmoiron/sqlx"
)

const TestExtPrivLevel = PrivLevelReadOnly;

func testExtHandler(db *sqlx.DB) RegexHandlerFunc {
        return func(w http.ResponseWriter, r *http.Request, p PathParams) {
                /*handleErr := func(err error, status int) {
                        log.Errorf("%v %v\n", r.RemoteAddr, err)
                        w.WriteHeader(status)
                        fmt.Fprintf(w, http.StatusText(status))
                }

                cdnName := p["cdn"]

                resp, err := getMonitoringJson(cdnName, db)
                if err != nil {
                        handleErr(err, http.StatusInternalServerError)
                        return
                }

                respBts, err := json.Marshal(resp)
                if err != nil {
                        handleErr(err, http.StatusInternalServerError)
                        return
                }*/

                w.Header().Set("Content-Type", "text/plain")
                fmt.Fprintf(w, "Go Extension!")
        }
}

