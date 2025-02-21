package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/juanchi26/CanelonesGamingUsers/awsgo"
	"github.com/juanchi26/CanelonesGamingUsers/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSjson, error) { //funcion de SECRET MANAGER
	var datosSecret models.SecretRDSjson
	fmt.Println("> Pudio secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error()) //si hay algun error retorna datos secret sin cambios y el error
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret) //estructura para procesar lo que me devuelta secret value, primero es lo que recibe y lo segundo donde lo graba
	fmt.Sprintln(" > lectura secret OK" + nombreSecret)
	return datosSecret, nil
}
