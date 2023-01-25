package main

import (
	"context"
	"log"
	"time"

	ampq "github.com/rabbitmq/amqp091-go"
)

func main(){
  conn, err := ampq.Dial("ampq://guest:guest@localhost:5672/")
  failOnError(err,"Failed to connect to RabbitMQ")
  defer conn.Close()

  ch, err := conn.Channel()
  failOnError(err, "Failed To open Channel")
  defer ch.Close()

  q, err := ch.QueueDeclare(
    "Hello", // name
    false,  // dueable 
    false, 
    false, 
    false,
    nil,
    )
  failOnError(err, "Failed to declare a queue")
  ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
  defer cancel()

  body := "Hello There!"
  err = ch.PublishWithContext(
    ctx,
    "", // exchange 
    q.Name, // routing k
    false, // mandatory
    false, // immediate 
    ampq.Publishing{
      ContentType: "text/plain",
      Body: []byte(body),
    })
  failOnError(err, "failed to publish the message")

}

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s",msg,err)
  }
}

