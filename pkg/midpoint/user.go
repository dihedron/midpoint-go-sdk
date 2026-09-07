package midpoint

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/dihedron/rawdata"
)

type UserService struct {
	Service
}

type User struct {
	Ns             *string `json:"@ns,omitempty"`
	Type           *string `json:"@type,omitempty"`
	Oid            *string `json:"oid,omitempty" yaml:"oid,omitempty"`
	Version        *string `json:"version,omitempty" yaml:"version,omitempty"`
	Name           *string `json:"name,omitempty" yaml:"name,omitempty"`
	Indestructible bool    `json:"indestructible,omitempty" yaml:"indestructible,omitempty"`
	Description    *string `json:"description,omitempty" yaml:"description,omitempty"`
	Metadata       *struct {
		Ns      *string  `json:"@ns,omitempty" yaml:"@ns,omitempty"`
		ID      int      `json:"@id,omitempty" yaml:"@id,omitempty"`
		Storage *Storage `json:"storage,omitempty" yaml:"storage,omitempty"`
		Process *struct {
			RequestTimestamp *time.Time `json:"requestTimestamp,omitempty" yaml:"requestTimestamp,omitempty"`
			RequestorRef     *Reference `json:"requestorRef"`
		} `json:"process,omitempty" yaml:"process,omitempty"`
		Provisioning *struct {
			LastProvisioningTimestamp *time.Time `json:"lastProvisioningTimestamp,omitempty" yaml:"lastProvisioningTimestamp,omitempty"`
		} `json:"provisioning,omitempty" yaml:"provisioning,omitempty"`
	} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
	Assignment *[]struct {
		ID         int     `json:"@id,omitempty" yaml:"@id,omitempty"`
		Identifier *string `json:"identifier,omitempty" yaml:"identifier,omitempty"`
		Ns         *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
		Metadata   *struct {
			ID      int      `json:"@id,omitempty" yaml:"@id,omitempty"`
			Ns      *string  `json:"@ns,omitempty" yaml:"@ns,omitempty"`
			Storage *Storage `json:"storage,omitempty" yaml:"storage,omitempty"`
			Process *struct {
				RequestTimestamp *time.Time `json:"requestTimestamp,omitempty" yaml:"requestTimestamp,omitempty"`
				RequestorRef     *Reference `json:"requestorRef,omitempty" yaml:"requestorRef,omitempty"`
			} `json:"process,omitempty" yaml:"process,omitempty"`
			Provenance *struct {
				Acquisition *struct {
					Channel   *string    `json:"channel,omitempty" yaml:"channel,omitempty"`
					Timestamp *time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
					ID        *string    `json:"@id,omitempty" yaml:"@id,omitempty"`
				} `json:"acquisition,omitempty" yaml:"acquisition,omitempty"`
			} `json:"provenance,omitempty" yaml:"provenance,omitempty"`
		} `json:"@metadata,omitempty" yaml:"@metadata,omitempty"`
		TargetRef  *Reference `json:"targetRef,omitempty" yaml:"targetRef,omitempty"`
		Activation *struct {
			EffectiveStatus *string `json:"effectiveStatus,omitempty" yaml:"effectiveStatus,omitempty"`
		} `json:"activation,omitempty" yaml:"activation,omitempty"`
	} `json:"assignment,omitempty" yaml:"assignment,omitempty"`
	Iteration         int        `json:"iteration,omitempty" yaml:"iteration,omitempty"`
	IterationToken    *string    `json:"iterationToken,omitempty" yaml:"iterationToken,omitempty"`
	ArchetypeRef      *Reference `json:"archetypeRef,omitempty" yaml:"archetypeRef,omitempty"`
	RoleMembershipRef []struct {
		Ns       *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
		Metadata *struct {
			Ns         *string  `json:"@ns,omitempty" yaml:"@ns,omitempty"`
			ID         int      `json:"@id,omitempty" yaml:"@id,omitempty"`
			Storage    *Storage `json:"storage,omitempty" yaml:"storage,omitempty"`
			Provenance *struct {
				AssignmentPath *struct {
					SourceRef *Reference `json:"sourceRef,omitempty" yaml:"sourceRef,omitempty"`
					Segment   *struct {
						ID            int        `json:"@id,omitempty" yaml:"@id,omitempty"`
						SegmentOrder  int        `json:"segmentOrder,omitempty" yaml:"segmentOrder,omitempty"`
						AssignmentID  int        `json:"assignmentId,omitempty" yaml:"assignmentId,omitempty"`
						TargetRef     *Reference `json:"targetRef,omitempty" yaml:"targetRef,omitempty"`
						MatchingOrder bool       `json:"matchingOrder,omitempty" yaml:"matchingOrder,omitempty"`
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
				Ns      *string  `json:"@ns,omitempty" yaml:"@ns,omitempty"`
				Storage *Storage `json:"storage,omitempty" yaml:"storage,omitempty"`
				ID      int      `json:"@id,omitempty" yaml:"@id,omitempty"`
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
	FullName   *string    `json:"fullName,omitempty" yaml:"fullName,omitempty"`
	GivenName  *string    `json:"givenName,omitempty" yaml:"givenName,omitempty"`
	FamilyName *string    `json:"familyName,omitempty" yaml:"familyName,omitempty"`
	Title      string     `json:"title,omitempty" yaml:"title,omitempty"`
	LinkRef    *Reference `json:"linkRef,omitempty" yaml:"linkRef,omitempty"`
}

// UnmarshalFlag provides support for loading user's data at the
// command line from a file on disk.
func (u *User) UnmarshalFlag(value string) error {
	return rawdata.UnmarshalInto(value, u)
}

type Reference struct {
	Ns         *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
	Oid        *string `json:"oid,omitempty" yaml:"oid,omitempty"`
	Relation   *string `json:"relation,omitempty" yaml:"relation,omitempty"`
	Type       *string `json:"type,omitempty" yaml:"type,omitempty"`
	TargetName *string `json:"targetName,omitempty" yaml:"targetName,omitempty"`
}

type Password struct {
	MinOccurs                     *string    `json:"minOccurs,omitempty" yaml:"minOccurs,omitempty"`
	LockoutMaxFailedAttempts      int        `json:"lockoutMaxFailedAttempts,omitempty" yaml:"lockoutMaxFailedAttempts,omitempty"`
	LockoutFailedAttemptsDuration *string    `json:"lockoutFailedAttemptsDuration,omitempty" yaml:"lockoutFailedAttemptsDuration,omitempty"`
	LockoutDuration               *string    `json:"lockoutDuration,omitempty" yaml:"lockoutDuration,omitempty"`
	ValuePolicyRef                *Reference `json:"valuePolicyRef,omitempty" yaml:"valuePolicyRef,omitempty"`
}

type Storage struct {
	CreateTimestamp *time.Time `json:"createTimestamp,omitempty" yaml:"createTimestamp,omitempty"`
	CreateChannel   *string    `json:"createChannel,omitempty" yaml:"createChannel,omitempty"`
	CreatorRef      *Reference `json:"creatorRef"`
	CreateTaskRef   *Reference `json:"createTaskRef"`
	ModifyTimestamp *time.Time `json:"modifyTimestamp,omitempty" yaml:"modifyTimestamp,omitempty"`
	ModifierRef     *Reference `json:"modifierRef"`
	ModifyChannel   *string    `json:"modifyChannel,omitempty" yaml:"modifyChannel,omitempty"`
	ModifyTaskRef   *Reference `json:"modifyTaskRef"`
}

type userWrapper struct {
	User *User `json:"user,omitempty" yaml:"user,omitempty"`
}

type userListWrapper struct {
	Ns     string `json:"@ns"`
	Object struct {
		Ns     string `json:"@ns"`
		Type   string `json:"@type"`
		Object []User `json:"object"`
	} `json:"object"`
}

type userPasswordPolicyWrapper struct {
	Ns     *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
	Object struct {
		Type     *string   `json:"@type,omitempty" yaml:"@type,omitempty"`
		Password *Password `json:"password,omitempty" yaml:"password,omitempty"`
	} `json:"object,omitempty" yaml:"object,omitempty"`
}

type PasswordResetRequest struct {
	ResetMethod *string `json:"resetMethod,omitempty" yaml:"resetMethod,omitempty"`
	UserEntry   *string `json:"userEntry,omitempty" yaml:"userEntry,omitempty"`
}

type PasswordResetResponse struct {
	Ns     *string `json:"@ns,omitempty" yaml:"@ns,omitempty"`
	Object struct {
		Type    string `json:"@type"`
		Message struct {
			Type            string `json:"@type"`
			Key             string `json:"key"`
			FallbackMessage string `json:"fallbackMessage"`
		} `json:"message"`
	} `json:"object"`
}

type userPasswordResetWrapper struct {
	ExecuteCredentialResetRequest *PasswordResetRequest `json:"executeCredentialResetRequest,omitempty" yaml:"executeCredentialResetRequest,omitempty"`
}

// Read reads information about a user.
func (s *UserService) Read(ctx context.Context, id string) (*User, error) {
	error := &Error{}
	response, err := s.client.
		R().
		SetContext(ctx).
		SetPathParam("id", id).
		SetQueryParam("options", "raw").
		SetResult(&userWrapper{}).
		SetResultError(error).
		Get("/users/{id}")
	if err != nil {
		slog.Error("error reading user", "id", id, "error", err, "result", response)
		return nil, err
	}
	return response.Result().(*userWrapper).User, nil
}

func (s *UserService) Create(ctx context.Context, user *User) (string, error) {
	error := &Error{}
	response, err := s.client.
		R().
		SetContext(ctx).
		SetQueryParam("options", "raw").
		SetBody(&userWrapper{User: user}).
		SetResultError(error).
		Post("/users")
	if err != nil {
		slog.Error("error creating user", "user", *user, "error", err, "result", response)
		return "", err
	}
	if response.StatusCode() == http.StatusCreated {
		location := response.Header().Get("Location")
		id := location[strings.LastIndex(location, "users/")+len("users/"):]
		slog.Debug("user created", "location", location, "id", id)
		return id, nil
	} else if response.StatusCode() == http.StatusConflict {
		slog.Warn("user already exists", "user", *user)
		return "", fmt.Errorf("user already exists")
	}
	return "", fmt.Errorf("unexpected status code: %d", response.StatusCode())
}

func (s *UserService) Search(ctx context.Context, query string) ([]User, error) {
	error := &Error{}
	slog.Debug("running search", "query", query)
	response, err := s.client.
		R().
		SetContext(ctx).
		SetQueryParam("options", "raw").
		SetContentType("application/json").
		SetBody(Query(query)). // NOTE: raw body
		SetResult(userListWrapper{}).
		SetResultError(error).
		Post("/users/search")
	if err != nil && response.StatusCode() != http.StatusOK {
		slog.Error("error searching for users", "query", query, "error", err, "result", response)
		return nil, err
	}

	if response.StatusCode() == StatusHandledError {
		slog.Debug("the operation incurred an error which qwas automatically fixed", "error", *error.Object.Message)
	} else if response.StatusCode() == StatusPartialError {
		slog.Debug("the operation was only partially successful", "error", *error.Object.Message)
	}

	return response.Result().(*userListWrapper).Object.Object, nil
	// func() {
	// 	f, _ := os.Create("output.json")
	// 	defer f.Close()
	// 	fmt.Fprintf(f, "------------------------ RESPONSE ------------------------ %s\n----------------------------------------------------------", response.String())
	// }()
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	error := &Error{}
	response, err := s.client.
		R().
		SetContext(ctx).
		SetPathParam("id", id).
		SetQueryParam("options", "raw").
		SetResultError(error).
		Delete("/users/{id}")
	if err != nil {
		slog.Error("error deleting user", "id", id, "error", err, "result", response)
		return err
	}
	if response.StatusCode() != http.StatusNoContent {
		slog.Error("invalid response status code in user deletion", "id", id, "error", err, "result", response)
	}
	return nil
}

func (s *UserService) ReadPasswordPolicy(ctx context.Context, id string) (*Password, error) {
	error := &Error{}
	response, err := s.client.
		R().
		SetContext(ctx).
		SetPathParam("id", id).
		SetQueryParam("options", "raw").
		SetResult(&userPasswordPolicyWrapper{}).
		SetResultError(error).
		Get("/users/{id}/policy")
	if err != nil {
		slog.Error("error reading user password policy", "id", id, "error", err, "result", response)
		return nil, err
	}

	// func() {
	// 	f, _ := os.Create("output.json")
	// 	defer f.Close()
	// 	fmt.Fprintf(f, "------------------------ RESPONSE ------------------------ %s\n----------------------------------------------------------", response.String())
	// }()

	return response.Result().(*userPasswordPolicyWrapper).Object.Password, nil
}

func (s *UserService) ResetPassword(ctx context.Context, id string, password string) error {
	error := &Error{}
	response, err := s.client.
		R().
		SetContext(ctx).
		SetQueryParam("options", "raw").
		SetPathParam("id", id).
		SetBody(&userPasswordResetWrapper{ExecuteCredentialResetRequest: &PasswordResetRequest{
			ResetMethod: new("passwordReset"),
			UserEntry:   new(password),
		}}).
		SetResult(&PasswordResetResponse{}).
		SetResultError(error).
		Post("/users/{id}/credential")
	if err != nil {
		slog.Error("error creating user password reset", "user", id, "password", password, "error", err, "result", response)
		return err
	}
	if response.StatusCode() != http.StatusOK {
		slog.Error("error resetting user password", "id", id, "password", password, "message", error.Object.Message)
		return fmt.Errorf("%s", error.Object.Message)
	}
	slog.Debug("user's password reset", "id", id, "password", password, "message", response.Result().(*PasswordResetResponse).Object.Message)
	return nil
}
