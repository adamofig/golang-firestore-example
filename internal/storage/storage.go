package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/api/iterator"

	"dataclouder/configs"

	"cloud.google.com/go/storage"
)

func TestStorage() {
	println("Las configuraciones son :", configs.ProjectId, configs.Bucket)

	println("Iniciando una prubea de storage")

	uploadSimpleFile()
	// configs.MiFuncion()
	//initBucket()
	//ListStorageItems()
	// var r io.Reader

	// dat, err := ioutil.ReadFile("../googleCloud/hola.txt")

	// if err != nil {
	// 	log.Fatal("No pudo leer el archivo hola.txt")
	// }
	// log.Println("Se tiene el archivo ", string(dat))

	// f, err := os.Open("../googleCloud/hola.txt")

	// if err != nil {
	// 	log.Fatal("Ocurió un error", err)
	// }

	// r = f
	// _ = r

}

func uploadSimpleFile() {

	_, ctx := getClient()

	f, err := os.Open("../googleCloud/hola.txt")

	if err != nil {
		log.Fatal("ocurrió un error al abrir el archivo")
	}

	_, atrr, err := upload(ctx, f, configs.ProjectId, configs.Bucket, "archivo.txt", true)
	println("Subiendo archivo")
	_ = atrr
	if err != nil {
		log.Fatal("ocurrió un error al subir el archivo", err)
	}

}

func ListStorageItems() {
	storageClient, ctx := getClient()

	println(storageClient)

	var projectId string = "gnp-datapullsinautos-qa"

	buckets := storageClient.Buckets(ctx, projectId)

	for {
		attrs, err := buckets.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(attrs)

	}
	// Aqui vamos  a listar los archivos
	println("______")
	//var bucket string = "gnp-datapullsinautos-qa.appspot.com"
	var bucket string = "nubedatos1.appspot.com"

	bucketHandle := storageClient.Bucket(bucket)

	query := &storage.Query{Prefix: ""}

	var names []string
	it := bucketHandle.Objects(ctx, query)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)

		// descomntar para ver los buckets
		//fmt.Println(names)
	}
}

func getClient() (*storage.Client, context.Context) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalln("No pudo obtener el cliente", err)
	}
	return client, ctx
}

func list(w io.Writer, client *storage.Client, bucket string) error {
	ctx := context.Background()
	// [START storage_list_files]
	it := client.Bucket(bucket).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintln(w, attrs.Name)
	}
	// [END storage_list_files]
	return nil
}

func upload(ctx context.Context, r io.Reader, projectID, bucket, name string, public bool) (*storage.ObjectHandle, *storage.ObjectAttrs, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}

	bh := client.Bucket(bucket)
	// Next check if the bucket exists
	if _, err = bh.Attrs(ctx); err != nil {
		return nil, nil, err
	}

	obj := bh.Object(name)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, r); err != nil {
		return nil, nil, err
	}
	if err := w.Close(); err != nil {
		return nil, nil, err
	}

	if public {
		if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
			return nil, nil, err
		}
	}

	attrs, err := obj.Attrs(ctx)
	return obj, attrs, err
}
