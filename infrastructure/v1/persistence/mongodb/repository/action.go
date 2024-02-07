package repository

import (
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

type actionRepositoryImpl struct {
	collection mongodb.Collection
	cfg        *config.AppConfig
}

func NewActionRepository(db mongodb.Database, cfg *config.AppConfig) *actionRepositoryImpl {
	return &actionRepositoryImpl{
		collection: db.Collection("action"),
		cfg:        cfg,
	}
}

func (r *actionRepositoryImpl) InsertAction(ic *core.InternalContext, action *entity.Action) (*entity.Action, *core.CustomError) {
	doc := new(model.Action).Bind(action)
	doc.Created = time.Now()
	doc.Modified = time.Now()

	res, err := r.collection.InsertOne(ic.ToContext(), doc)
	if err != nil {
		log.Error(ic.ToContext(), "failed to InsertAction", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	doc.Id = res.InsertedID.(primitive.ObjectID)
	return doc.Entity(), nil
}

func (r *actionRepositoryImpl) FindActionByTargetId(ic *core.InternalContext, targetId string) (*entity.Action, *core.CustomError) {
	var doc model.Action

	objId, err := primitive.ObjectIDFromHex(targetId)
	if err != nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "invalid target id",
		}
	}

	filter := bson.M{
		"target_id": objId,
	}

	err = r.collection.FindOne(ic.ToContext(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		log.Error(ic.ToContext(), "failed FindActionByTargetId : %v", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	return doc.Entity(), nil
}

func (r *actionRepositoryImpl) UpdateAction(ic *core.InternalContext, action *entity.Action) (*entity.Action, *core.CustomError) {
	doc := new(model.Action).Bind(action)
	doc.Modified = time.Now()

	objId, err := primitive.ObjectIDFromHex(action.Id)
	if err != nil {
		return nil, &core.CustomError{
			Code:    core.UNPROCESSABLE_ENTITY,
			Message: "invalid action id",
		}
	}

	filter := bson.M{"_id": objId}
	set := bson.M{"$set": doc}
	_, err = r.collection.UpdateOne(ctx(ic), filter, set)
	if err != nil {
		log.Error(ic.ToContext(), "failed UpdateAction : %v", err)
		return nil, &core.CustomError{
			Code: core.INTERNAL_SERVER_ERROR,
		}
	}

	return doc.Entity(), nil
}
