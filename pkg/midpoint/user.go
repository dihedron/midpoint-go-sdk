package midpoint

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/dihedron/rawdata"
)

type UserService struct {
	Service
}

type User struct {
	Metadata *struct {
		Ns      *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
		Storage *struct {
			CreateTimestamp *time.Time `json:"createTimestamp,omitempty" yaml:"createTimestamp,omitempty"`
			CreateChannel   *string    `json:"createChannel,omitempty" yaml:"createChannel,omitempty"`
		} `json:"storage,omitempty" yaml:"storage,omitempty"`
		Process *struct {
			RequestTimestamp *time.Time `json:"requestTimestamp,omitempty" yaml:"requestTimestamp,omitempty"`
		} `json:"process,omitempty" yaml:"process,omitempty"`
		ID int `json:"@id,omitempty" yaml:"@id,omitempty"`
	} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
	Oid            *string `json:"oid,omitempty" yaml:"oid,omitempty"`
	Version        *string `json:"version,omitempty" yaml:"version,omitempty"`
	Name           *string `json:"name,omitempty" yaml:"name,omitempty"`
	Indestructible bool    `json:"indestructible,omitempty" yaml:"indestructible,omitempty"`
	Assignment     *[]struct {
		Ns       *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
		Metadata *struct {
			Storage *struct {
				CreateTimestamp *time.Time `json:"createTimestamp,omitempty" yaml:"createTimestamp,omitempty"`
				CreateChannel   *string    `json:"createChannel,omitempty" yaml:"createChannel,omitempty"`
			} `json:"storage,omitempty" yaml:"storage,omitempty"`
			Process *struct {
				RequestTimestamp *time.Time `json:"requestTimestamp,omitempty" yaml:"requestTimestamp,omitempty"`
			} `json:"process,omitempty" yaml:"process,omitempty"`
			Provenance *struct {
				Acquisition *struct {
					Channel   *string    `json:"channel,omitempty" yaml:"channel,omitempty"`
					Timestamp *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
					ID        *string    `json:"@id,omitempty" yaml:"@id,omitempty"`
				} `json:"acquisition,omitempty" yaml:"acquisition,omitempty"`
			} `json:"provenance,omitempty" yaml:"provenance,omitempty"`
			ID int `json:"@id,omitempty" yaml:"@id,omitempty"`
		} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
		ID         int     `json:"@id,omitempty" yaml:"@id,omitempty"`
		Identifier *string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
		TargetRef  *struct {
			Oid      *string `json:"oid,omitempty" yaml:"oid,omitempty"`
			Relation *string `json:"relation,omitempty" yaml:"relation,omitempty"`
			Type     *string `json:"type,omitempty" yaml:"type,omitempty"`
		} `json:"targetRef,omitempty" yaml:"targetRef,omitempty"`
		Activation *struct {
			EffectiveStatus *string `json:"effectiveStatus,omitempty" yaml:"effectiveStatus,omitempty"`
		} `json:"activation,omitempty" yaml:"activation,omitempty"`
	} `json:"assignment,omitempty" yaml:"assignment,omitempty"`
	Iteration      int     `json:"iteration,omitempty" yaml:"iteration,omitempty"`
	IterationToken *string `json:"iterationToken,omitempty" yaml:"iterationToken,omitempty"`
	ArchetypeRef   *struct {
		Oid      *string `json:"oid,omitempty" yaml:"oid,omitempty"`
		Relation *string `json:"relation,omitempty" yaml:"relation,omitempty"`
		Type     *string `json:"type,omitempty" yaml:"type,omitempty"`
	} `json:"archetypeRef,omitempty" yaml:"archetypeRef,omitempty"`
	RoleMembershipRef []struct {
		Metadata *struct {
			Ns      *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
			ID      int     `json:"@id,omitempty" yaml:"@id,omitempty"`
			Storage *struct {
				CreateTimestamp *time.Time `json:"createTimestamp,omitempty" yaml:"createTimestamp,omitempty"`
			} `json:"storage,omitempty" yaml:"storage,omitempty"`
			Provenance *struct {
				AssignmentPath *struct {
					SourceRef *struct {
						Oid      *string `json:"oid,omitempty" yaml:"oid,omitempty"`
						Relation *string `json:"relation,omitempty" yaml:"relation,omitempty"`
						Type     *string `json:"type,omitempty" yaml:"type,omitempty"`
					} `json:"sourceRef,omitempty" yaml:"sourceRef,omitempty"`
					Segment *struct {
						ID           int `json:"@id,omitempty" yaml:"@id,omitempty"`
						SegmentOrder int `json:"segmentOrder,omitempty" yaml:"segmentOrder,omitempty"`
						AssignmentID int `json:"assignmentId,omitempty" yaml:"assignmentId,omitempty"`
						TargetRef    *struct {
							Oid      *string `json:"oid,omitempty" yaml:"oid,omitempty"`
							Relation *string `json:"relation,omitempty" yaml:"relation,omitempty"`
							Type     *string `json:"type,omitempty" yaml:"type,omitempty"`
						} `json:"targetRef,omitempty" yaml:"targetRef,omitempty"`
						MatchingOrder bool `json:"matchingOrder,omitempty" yaml:"matchingOrder,omitempty"`
					} `json:"segment,omitempty" yaml:"segment,omitempty"`
				} `json:"assignmentPath,omitempty" yaml:"assignmentPath,omitempty"`
			} `json:"provenance,omitempty" yaml:"provenance,omitempty"`
		} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
		Oid      *string `json:"oid,omitempty" yaml:"oid,omitempty"`
		Relation *string `json:"relation,omitempty" yaml:"relation,omitempty"`
		Type     *string `json:"type,omitempty" yaml:"type,omitempty"`
	} `json:"roleMembershipRef,omitempty" yaml:"roleMembershipRef,omitempty"`
	Activation *struct {
		AdministrativeStatus *string    `json:"administrativeStatus,omitempty" yaml:"administrativeStatus,omitempty"`
		EffectiveStatus      *string    `json:"effectiveStatus,omitempty" yaml:"effectiveStatus,omitempty"`
		EnableTimestamp      *time.Time `json:"enableTimestamp,omitempty" yaml:"enableTimestamp,omitempty"`
		LockoutStatus        *string    `json:"lockoutStatus,omitempty" yaml:"lockoutStatus,omitempty"`
	} `json:"activation,omitempty" yaml:"activation,omitempty"`
	Credentials *struct {
		Password *struct {
			Metadata *struct {
				Ns      *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
				Storage *struct {
					CreateTimestamp *time.Time `json:"createTimestamp,omitempty" yaml:"createTimestamp,omitempty"`
					CreateChannel   *string    `json:"createChannel,omitempty" yaml:"createChannel,omitempty"`
				} `json:"storage,omitempty" yaml:"storage,omitempty"`
				ID int `json:"@id,omitempty" yaml:"@id,omitempty"`
			} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
			Value *struct {
				EncryptedData *struct {
					EncryptionMethod *struct {
						Algorithm *string `json:"algorithm,omitempty" yaml:"algorithm,omitempty"`
					} `json:"encryptionMethod,omitempty" yaml:"encryptionMethod,omitempty"`
					KeyInfo *struct {
						KeyName *string `json:"keyName,omitempty" yaml:"keyName,omitempty"`
					} `json:"keyInfo,omitempty" yaml:"keyInfo,omitempty"`
					CipherData *struct {
						CipherValue *string `json:"cipherValue,omitempty" yaml:"cipherValue,omitempty"`
					} `json:"cipherData,omitempty" yaml:"cipherData,omitempty"`
				} `json:"encryptedData,omitempty" yaml:"encryptedData,omitempty"`
			} `json:"value,omitempty" yaml:"value,omitempty"`
		} `json:"password,omitempty" yaml:"password,omitempty"`
	} `json:"credentials,omitempty" yaml:"credentials,omitempty"`
	Behavior *struct {
		Authentication []struct {
			Ns                  *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
			ID                  int     `json:"@id,omitempty" yaml:"@id,omitempty"`
			LastSuccessfulLogin *struct {
				Timestamp *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
				From      *string    `json:"from,omitempty" yaml:"from,omitempty"`
			} `json:"lastSuccessfulLogin,omitempty" yaml:"lastSuccessfulLogin,omitempty"`
			PreviousSuccessfulLogin *struct {
				Timestamp *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
				From      *string    `json:"from,omitempty" yaml:"from,omitempty"`
			} `json:"previousSuccessfulLogin,omitempty" yaml:"previousSuccessfulLogin,omitempty,omitempty"`
			AuthenticationAttempt *struct {
				SequenceIdentifier           *string `json:"sequenceIdentifier,omitempty" yaml:"sequenceIdentifier,omitempty"`
				ModuleIdentifier             *string `json:"moduleIdentifier,omitempty" yaml:"moduleIdentifier,omitempty"`
				LastSuccessfulAuthentication *struct {
					Timestamp *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
					From      *string    `json:"from,omitempty" yaml:"from,omitempty"`
				} `json:"lastSuccessfulAuthentication,omitempty" yaml:"lastSuccessfulAuthentication,omitempty"`
			} `json:"authenticationAttempt,omitempty" yaml:"authenticationAttempt,omitempty"`
			SequenceIdentifier *string `json:"sequenceIdentifier,omitempty" yaml:"sequenceIdentifier,omitempty"`
		} `json:"authentication,omitempty" yaml:"authentication,omitempty"`
	} `json:"behavior,omitempty" yaml:"behavior,omitempty"`
	FullName   *string `json:"fullName,omitempty" yaml:"fullName,omitempty"`
	GivenName  *string `json:"givenName,omitempty" yaml:"givenName,omitempty"`
	FamilyName *string `json:"familyName,omitempty" yaml:"familyName,omitempty"`
}

// UnmarshalFlag provides support for loading user's data at the
// command line from a file on disk.
func (u *User) UnmarshalFlag(value string) error {
	return rawdata.UnmarshalInto(value, u)
}

type userWrapper struct {
	User *User `json:"user,omitempty" yaml:"user,omitempty"`
}

func (s *UserService) Read(ctx context.Context, id string) (*User, error) {
	entity, result, err := s.client.Get[userWrapper](ctx, fmt.Sprintf("users/%s?options=raw", id))
	if err != nil {
		slog.Error("error reading user", "id", id, "result", result, "error", err)
		return nil, err
	}
	return entity.User, nil
}

func (s *UserService) Create(ctx context.Context, user *User) error {
	result, err := s.client.Post(ctx, "users?options=raw", &userWrapper{User: user})
	if err != nil {
		slog.Error("error creating user", "result", result, "error", err)
		return err
	}
	return nil
}
