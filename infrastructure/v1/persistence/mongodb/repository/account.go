package repository

import (
	"dealls/pkg/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"dealls/config"
	"dealls/core"
	"dealls/core/v1/entity"
	"dealls/infrastructure/v1/persistence/mongodb"
	"dealls/infrastructure/v1/persistence/mongodb/model"
)

type accountRepositoryImpl struct {
	collection mongodb.Collection
	cfg        *config.AppConfig
}

func NewAccountRepository(db mongodb.Database, cfg *config.AppConfig) *accountRepositoryImpl {
	return &accountRepositoryImpl{
		collection: db.Collection("account"),
		cfg:        cfg,
	}
}

func (r *accountRepositoryImpl) InsertAccount(ic *core.InternalContext, account *entity.Account) (*entity.Account, *core.CustomError) {
	doc := new(model.Account).Bind(account)
	doc.Created = time.Now()
	doc.Modified = time.Now()

	res, err := r.collection.InsertOne(ic.ToContext(), &doc)
	if err != nil {
		log.Error(ic.ToContext(), "failed to InsertAccount", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	doc.Id = res.InsertedID.(primitive.ObjectID)
	return doc.Entity(), nil
}

func (r *accountRepositoryImpl) FindAccountByUsername(ic *core.InternalContext, username string) (*entity.Account, *core.CustomError) {
	var data model.Account

	filter := bson.M{
		"username": username,
	}

	err := r.collection.FindOne(ic.ToContext(), filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error(ic.ToContext(), "failed to FindAccountByUsername", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}
	return data.Entity(), nil
}

func (r *accountRepositoryImpl) FindAccountByEmail(ic *core.InternalContext, email string) (*entity.Account, *core.CustomError) {
	var data model.Account

	filter := bson.M{
		"email": email,
	}

	err := r.collection.FindOne(ic.ToContext(), filter).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error(ic.ToContext(), "failed to FindAccountByEmail", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}
	return data.Entity(), nil
}

func (r *accountRepositoryImpl) GetAccountsExclude(ic *core.InternalContext, profileIds []string, meta map[string]any) ([]entity.Account, int32, *core.CustomError) {
	var _profileIds []primitive.ObjectID
	for _, id := range profileIds {
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return nil, 0, &core.CustomError{
				Code:    core.INTERNAL_SERVER_ERROR,
				Message: err.Error(),
			}
		}
		_profileIds = append(_profileIds, _id)
	}

	filter := bson.M{}
	if _profileIds != nil {
		filter["_id"] = bson.M{
			"$nin": _profileIds,
		}
	}

	match := bson.D{
		{Key: "$match", Value: filter},
	}

	page := helper.DataToInt(meta["page"])
	limit := helper.DataToInt(meta["limit"])
	offset := limit * (page - 1)

	facet := bson.D{
		{Key: "$facet", Value: bson.D{
			{Key: "metadata", Value: []bson.M{
				{"$count": "total"},
			}},
			{Key: "data", Value: []bson.M{
				{"$skip": offset},
				{"$limit": limit},
			}},
		}},
	}

	pipeline := mongo.Pipeline{match, facet}
	res, err := r.collection.Aggregate(ic.ToContext(), pipeline)
	if err != nil {
		log.Error(ic.ToContext(), "failed r.collection.Aggregate(ic.ToContext(), pipeline)", err)
		return nil, 0, &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}

	var doc []struct {
		Metadata []map[string]interface{} `bson:"metadata"`
		Data     model.Accounts           `bson:"data"`
	}

	err = res.All(ic.ToContext(), &doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, 0, nil
		}

		log.Error(ic.ToContext(), "failed res.All(ic.ToContext(), &doc)", err)
		return nil, 0, &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}

	if len(doc[0].Data) < 1 {
		return nil, 0, nil
	}

	total := doc[0].Metadata[0]["total"].(int32)
	return doc[0].Data.Entities(), total, nil
}

func (r *accountRepositoryImpl) FindAccountById(ic *core.InternalContext, id string) (*entity.Account, *core.CustomError) {
	var doc model.Account

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "invalid account id",
		}
	}

	filter := bson.M{
		"_id": objId,
	}

	err = r.collection.FindOne(ic.ToContext(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error(ic.ToContext(), "failed FindAccountById : %v", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	return doc.Entity(), nil
}

func (r *accountRepositoryImpl) FindAccountsActivation(ic *core.InternalContext) ([]entity.Account, *core.CustomError) {
	filter := bson.M{
		"metadata.status": core.XENDIT_STATUS_ACTIVE,
	}

	res, err := r.collection.Find(ic.ToContext(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error(ic.ToContext(), "failed FindAccountsActivation", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	var docs model.Accounts
	err = res.All(ic.ToContext(), &docs)
	if err != nil {
		log.Error(ic.ToContext(), "failed FindAccountsActivation", err)
		return nil, &core.CustomError{
			Code:    core.INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
	}

	return docs.Entities(), nil
}

func (r *accountRepositoryImpl) UpdateAccount(ic *core.InternalContext, account *entity.Account) (*entity.Account, *core.CustomError) {
	doc := new(model.Account).Bind(account)
	doc.Modified = time.Now()

	objId, err := primitive.ObjectIDFromHex(account.Id)
	if err != nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "invalid account id",
		}
	}

	filter := bson.M{"_id": objId}
	set := bson.M{"$set": doc}
	_, err = r.collection.UpdateOne(ctx(ic), filter, set)
	if err != nil {
		log.Error(ic.ToContext(), "failed UpdateAccount : %v", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	return doc.Entity(), nil
}
