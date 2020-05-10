package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path/filepath"
)

var writers = map[string]chan interface{}{}
var conn *amqp.Connection
var alteryxChannel *amqp.Channel
var toAlteryxQueue amqp.Queue
var fromAlteryxQueue amqp.Queue

func main() {
	address := `localhost:35014`
	rabbitMqAddress := `amqp://guest:guest@localhost:5672/`
	var err error
	conn, err = amqp.Dial(rabbitMqAddress)
	if err != nil {
		print(err.Error())
		return
	}

	alteryxChannel, err = conn.Channel()
	if err != nil {
		print(err.Error())
		return
	}
	toAlteryxQueue, err = alteryxChannel.QueueDeclare(
		`to_alteryx`,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		print(err.Error())
		return
	}
	fromAlteryxQueue, err = alteryxChannel.QueueDeclare(
		`from_alteryx`,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		print(err.Error())
		return
	}

	msgs, msgErr := alteryxChannel.Consume(
		fromAlteryxQueue.Name,
		``,
		true,
		false,
		false,
		false,
		nil,
	)
	if msgErr != nil {
		print(msgErr.Error())
		return
	}

	go func() {
		for msg := range msgs {
			var i interface{}
			err = json.Unmarshal(msg.Body, &i)
			if err != nil {
				continue
			}
			iMap := i.(map[string]interface{})
			id, ok := iMap[`Id`]
			if !ok {
				continue
			}
			idStr, ok := id.(string)
			if !ok {
				continue
			}
			channel, ok := writers[idStr]
			if !ok {
				continue
			}
			iMap[`Error`] = ``
			channel <- iMap
		}
	}()

	http.HandleFunc(`/`, handleMain)
	http.HandleFunc("/main.dart.js", handleFile)
	http.HandleFunc("/main.dart.js.map", handleFile)
	http.HandleFunc("/main.dart.js.deps", handleFile)
	http.HandleFunc("/assets/", handleFile)
	println(`listening on ` + address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		print(err.Error())
	}
}

func generateRandomString() string {
	min := 65
	max := 90
	randomBytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		randomBytes[i] = byte(min + (rand.Intn(max - min)))
	}
	return string(randomBytes)
}

func handleMain(writer http.ResponseWriter, r *http.Request) {
	if r.Method == `GET` {
		http.ServeFile(writer, r, filepath.Join(`html`, `index.html`))
		return
	}

	id := generateRandomString()
	toSend := id
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErrorResponse(writer, err.Error())
	}
	bodyString := string(body)
	if bodyString != `` {
		toSend = fmt.Sprintf(`%v|%v`, id, bodyString)
	}
	println(`sending ` + toSend)
	channel := make(chan interface{})
	writers[id] = channel
	err = alteryxChannel.Publish(
		``,
		toAlteryxQueue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: `text/plain`,
			Body:        []byte(toSend),
		},
	)
	if err != nil {
		sendErrorResponse(writer, err.Error())
		delete(writers, id)
	}
	response := <-channel
	close(channel)
	delete(writers, id)
	println(fmt.Sprintf(`received back: %v`, response))
	setHeaders(writer, "application/json")
	responseBytes, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		buffer := bytes.NewBufferString(marshalErr.Error())
		_, _ = writer.Write(buffer.Bytes())
	}
	_, _ = writer.Write(responseBytes)
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Path
	ext := filepath.Ext(file)
	if ext == `.js` || ext == `.json` || ext == `.map` {
		setHeaders(w, "application/javascript")
	}
	if ext == `.deps` {
		setHeaders(w, "text/plain")
	}
	http.ServeFile(w, r, filepath.Join(`html`, file))
}

func sendErrorResponse(w http.ResponseWriter, err string) {
	setHeaders(w, "application/json")
	response := map[string]interface{}{`Error`: err}
	responseBytes, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		var buffer = bytes.NewBufferString(marshalErr.Error())
		_, _ = w.Write(buffer.Bytes())
		return
	}
	_, _ = w.Write(responseBytes)
}

func setHeaders(w http.ResponseWriter, contentType string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", contentType)
}
