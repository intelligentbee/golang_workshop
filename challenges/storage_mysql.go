package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
)

//var storage map[string]string
type Storage struct {
	dbCon *sql.DB
}

func (storage *Storage) Set(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// read json from request body
	body := r.Body
	dec := json.NewDecoder(body)
	reqBody := map[string]string{}
	if err := dec.Decode(&reqBody); err != nil {
		if err.Error() == "EOF" {
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	value, ok := reqBody["value"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	//set key-value in storage
	err := storage.addToStorage(key, value)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// respond back
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (storage *Storage) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := storage.deleteFromStorage(key)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (storage *Storage) Get(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// read key from url params
	key := ps.ByName("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := storage.getFromStorage(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(result); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	myStorage := Storage{dbCon: db}

	router := httprouter.New()
	router.DELETE("/storage/:key", myStorage.Delete)
	router.POST("/storage/:key", myStorage.Set)
	router.GET("/storage/:key", myStorage.Get)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(:3306)/test")
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS storage(id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY, keyx varchar(50), valx varchar(50))")
	if err != nil {
		log.Fatal(err)

		return nil, err
	}

	return db, nil
}

func (storage *Storage) addToStorage(key, value string) error {
	query := fmt.Sprintf("INSERT INTO test.storage(keyx, valx) VALUES('%s', '%s')", key, value)
	_, err := storage.dbCon.Exec(query)
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func (storage *Storage) deleteFromStorage(key string) error {
	query := fmt.Sprintf("DELETE FROM test.storage WHERE keyx = '%s'", key)
	_, err := storage.dbCon.Exec(query)
	if err != nil {
		log.Fatal(err)

		return err
	}

	return nil
}

func (storage *Storage) getFromStorage(key string) (map[string]string, error) {
	result := map[string]string{}

	query := fmt.Sprintf("select keyx, valx FROM test.storage WHERE keyx = '%s' limit 1", key)
	rows, err := storage.dbCon.Query(query)
	if err != nil {
		log.Fatal(err)

		return result, err
	}

	if rows.Next() {
		var key, value string
		if err := rows.Scan(&key, &value); err != nil {
			log.Fatal(err)

			return result, err
		}
		log.Printf("found row containing %d, %q", key, value)
		result[key] = value
	}
	rows.Close()

	return result, nil
}
