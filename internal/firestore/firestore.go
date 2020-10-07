package firestore

import (
	"context"
	"log"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/firestore"
	"github.com/gookit/color"
	"google.golang.org/api/option"

	"dataclouder/internal/logger"
)

var client, ctx = GetClient()

type User struct {
	Apellido   string `firestore:"apellido,omitempty"`
	Nacimiento string `firestore:"nacimiento,omitempty"`
	Nombre     string `firestore:"nombre,omitempty"`
}

func GetClient() (client *firestore.Client, ctx context.Context) {
	projectID := "adamo-test-1"

	ctx = context.Background()

	sa := option.WithCredentialsFile("resources/credentials.json")

	client, err := firestore.NewClient(ctx, projectID, sa)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, ctx
}

func AddDocWithoutID(collection string, data map[string]interface{}) map[string]interface{} {
	logger.Debug("Se recibió el siguiente objeto ", data)
	// TODO! : deje de lado lo de cerrar la conexión
	//defer client.Close()

	// Overwrite the document with the given data. Any other fields currently
	// in the document will be removed.
	doc, wr, err := client.Collection(collection).Add(ctx, data)
	if err != nil {
		// TODO: Handle error.
		logger.Error("Ocurrió un error", err)
	}

	data["id"] = doc.ID

	color.Success.Println("Se guardó el objeto", doc.ID, wr)
	return data
}

// GetAllDocsInCollection obtiene las collection de firestore
// regresa un arreglo de map,
func GetAllDocsInCollection(name string) ([]map[string]interface{}, error) {
	// [START fs_get_all_docs]
	color.Debug.Println("buscando documentos de ", name)

	iter := client.Collection(name).Documents(ctx)

	var collection []map[string]interface{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			color.Danger.Println("Ocurrió un problema", err)
			return nil, err
		}

		var data map[string]interface{}
		doc.DataTo(&data)
		// guardó sobre el mismo mapa el id
		data["id"] = doc.Ref.ID
		collection = append(collection, data)
	}
	color.Warn.Println("Los objetos a mostrar son", collection)
	// [END fs_get_all_docs]
	return collection, nil
}

func DeleteDoc(name string, id string) error {
	// [START fs_delete_doc]

	color.Danger.Println("Aqui tengo que eliminar.")
	result, err := client.Collection(name).Doc(id).Delete(ctx)

	color.Danger.Println(result)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
	color.Debug.Println("Se eliminó el documento con id", id)
	// [END fs_delete_doc]
	return err
}

func ValidateEntity() {
	logger.Info("Validando la Entidad")

}
