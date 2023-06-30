package secretmgo

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/ss-Contreras/GambitUcc/awsgo"
	"github.com/ss-Contreras/GambitUcc/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJon, error) {
	var datosSecret models.SecretRDSJon
	fmt.Println(" >  Pido Secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})
	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret Correcta " + nombreSecret)
	return datosSecret, nil
}
