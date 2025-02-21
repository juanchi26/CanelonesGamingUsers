package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/juanchi26/CanelonesGamingUsers/awsgo"
	"github.com/juanchi26/CanelonesGamingUsers/bd"
	"github.com/juanchi26/CanelonesGamingUsers/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	// Codigo de la lambda
	awsgo.InicializoAWS() // inicializo AWS

	if !ValidoParametros() { //devuelvo el mismo evento y el error generado
		fmt.Println("No se recibieron los parametros necesarios : SecretName")
		err := errors.New("error en los parametros necesarios : SecretName")
		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("sub = " + datos.UserUUID)
		}

	}

	err := bd.ReadSecret()

	if err != nil {
		fmt.Println("error al leer el secret")
		return event, err
	}

	err = bd.SignUp(datos)
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("Secretname")
	return traeParametro
}
