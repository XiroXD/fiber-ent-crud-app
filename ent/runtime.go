// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/XiroXD/fiber-ent-crud-app/ent/pet"
	"github.com/XiroXD/fiber-ent-crud-app/ent/schema"
	"github.com/XiroXD/fiber-ent-crud-app/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	petFields := schema.Pet{}.Fields()
	_ = petFields
	// petDescName is the schema descriptor for name field.
	petDescName := petFields[0].Descriptor()
	// pet.NameValidator is a validator for the "name" field. It is called by the builders before save.
	pet.NameValidator = petDescName.Validators[0].(func(string) error)
	// petDescBreed is the schema descriptor for breed field.
	petDescBreed := petFields[1].Descriptor()
	// pet.BreedValidator is a validator for the "breed" field. It is called by the builders before save.
	pet.BreedValidator = petDescBreed.Validators[0].(func(string) error)
	// petDescAge is the schema descriptor for age field.
	petDescAge := petFields[2].Descriptor()
	// pet.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	pet.AgeValidator = petDescAge.Validators[0].(func(int) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
