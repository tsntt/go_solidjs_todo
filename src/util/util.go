package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func SetEnvs(fileUrl string, variable *map[string]string) error {
	jsonFile, err := os.Open(fileUrl)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	json.Unmarshal(byteValue, &variable)

	return nil
}

func StringToTimeUnix(t string) time.Time {
	tInt, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		log.Printf("unable to parse time: %+v\n", err)
	}
	return time.Unix(tInt/1000, 0)
}

func addCorsHeader(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "*")
	headers.Add("Access-Control-Allow-Methods", "*")
}

func WriteJson(w http.ResponseWriter, status int, msg any) {
	w.Header().Set("Content-Type", "application/json")
	addCorsHeader(w)
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		log.Printf("%+v", err)
	}
}

func ReadJson(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if r.Header.Get("content-type") != "application/json" {
		return nil, errors.New("unsupported media format")
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	vars := make(map[string]string)

	err := dec.Decode(&vars)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return nil, errors.New(msg)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return nil, errors.New("request body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return nil, errors.New(msg)

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("request body contains unknown field %s", fieldName)
			return nil, errors.New(msg)

		case errors.Is(err, io.EOF):
			return nil, errors.New("request body must not be empty")

		case err.Error() == "http: request body too large":
			return nil, errors.New("request body must not be larger than 1MB")

		default:
			return nil, err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		return nil, errors.New("request body must only contain a single JSON object")
	}

	return vars, nil
}
