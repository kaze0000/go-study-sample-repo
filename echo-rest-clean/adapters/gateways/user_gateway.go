package gateways

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"echo-rest-clean/entities"

	"cloud.google.com/go/firestore"
)

type FirestoreClientFactory interface {
	NewClient(ctx context.Context) (*firestore.Client, error)
}

type UserGateway struct {
	//[役割]Firestoreクライアントの生成を抽象化する
	clientFactory FirestoreClientFactory
}

func NewUserGateway(clientFactory FirestoreClientFactory) *UserGateway {
	// [役割]ports.UserReositoryインターフェイスを実装したオブジェクトを返す
	return &UserGateway{clientFactory: clientFactory}
}

func (g *UserGateway) AddUser(ctx context.Context, user *entities.User) ([]*entities.User, error) {
	// [役割] Firestoreクライアントを生成し、引数として渡されたユーザーをFirestoreのusersコレクションに追加する
	if user == nil {
		return nil, errors.New("user is nil")
	}

	client, err := g.clientFactory.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}
	defer client.Close()

	_, err = client.Collection("users").Doc(user.Name).Set(ctx, map[string]interface{} {
		"age": user.Age,
		"address": user.Address,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return getUsers(ctx, client)
}

func (gateway *UserGateway) GetUsers(ctx context.Context) ([]*entities.User, error) {
	// [役割] Firestoreクライアントを生成し、usersコレクションからすべてのユーザーを取得してリストで返す
	client, err := gateway.clientFactory.NewClient(ctx)
 	if err != nil {
	return nil, fmt.Errorf("failed GetUsers NewClient: %v", err)
	}
	defer client.Close()

	return getUsers(ctx, client)
}

func getUsers(ctx context.Context, client *firestore.Client) ([]*entities.User, error) {
	allData := client.Collection("users").Documents(ctx)

	docs, err := allData.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed GetUsers GetAll: %v", err)
	}
	users := make([]*entities.User, 0)
	for _, doc := range docs {
		u := new(entities.User)
		err = mapToStruct(doc.Data(), &u)
		if err != nil {
			return nil, fmt.Errorf("failed GetUsers mapToStruct: %v", err)
		}
		u.Name = doc.Ref.ID
		users = append(users, u)
	}

 	return users, nil
}

// map -> 構造体の変換
func mapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}
